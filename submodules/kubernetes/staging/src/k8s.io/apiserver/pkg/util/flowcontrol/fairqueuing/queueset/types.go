/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package queueset

import (
	"context"
	"fmt"
	"math"
	"time"

	genericrequest "k8s.io/apiserver/pkg/endpoints/request"
	"k8s.io/apiserver/pkg/util/flowcontrol/debug"
	fq "k8s.io/apiserver/pkg/util/flowcontrol/fairqueuing"
	"k8s.io/apiserver/pkg/util/flowcontrol/fairqueuing/promise"
	fcrequest "k8s.io/apiserver/pkg/util/flowcontrol/request"
)

// request is a temporary container for "requests" with additional
// tracking fields required for the functionality FQScheduler
type request struct {
	ctx context.Context

	qs *queueSet

	flowDistinguisher string
	fsName            string

	// The relevant queue.  Is nil if this request did not go through
	// a queue.
	queue *queue

	// startTime is the real time when the request began executing
	startTime time.Time

	// estimated amount of work of the request
	workEstimate fcrequest.WorkEstimate

	// decision gets set to a `requestDecision` indicating what to do
	// with this request.  It gets set exactly once, when the request
	// is removed from its queue.  The value will be decisionReject,
	// decisionCancel, or decisionExecute.
	//
	// decision.Set is called with the queueSet locked.
	// decision.Get is called without the queueSet locked.
	decision promise.WriteOnce

	// arrivalTime is the real time when the request entered this system
	arrivalTime time.Time

	// arrivalR is R(arrivalTime).  R is, confusingly, also called "virtual time".
	arrivalR SeatSeconds

	// descr1 and descr2 are not used in any logic but they appear in
	// log messages
	descr1, descr2 interface{}

	// Indicates whether client has called Request::Wait()
	waitStarted bool

	queueNoteFn fq.QueueNoteFn

	// Removes this request from its queue. If the request is not put into a
	// a queue it will be nil.
	removeFromQueueFn removeFromFIFOFunc
}

// queue is an array of requests with additional metadata required for
// the FQScheduler
type queue struct {
	// The requests not yet executing in the real world are stored in a FIFO list.
	requests fifo

	// nextDispatchR is the R progress meter reading at
	// which the next request will be dispatched in the virtual world.
	nextDispatchR SeatSeconds

	// requestsExecuting is the count in the real world
	requestsExecuting int
	index             int

	// seatsInUse is the total number of "seats" currently occupied
	// by all the requests that are currently executing in this queue.
	seatsInUse int
}

// Enqueue enqueues a request into the queue and
// sets the removeFromQueueFn of the request appropriately.
func (q *queue) Enqueue(request *request) {
	request.removeFromQueueFn = q.requests.Enqueue(request)
}

// Dequeue dequeues a request from the queue
func (q *queue) Dequeue() (*request, bool) {
	request, ok := q.requests.Dequeue()
	return request, ok
}

func (q *queue) dump(includeDetails bool) debug.QueueDump {
	digest := make([]debug.RequestDump, q.requests.Length())
	i := 0
	q.requests.Walk(func(r *request) bool {
		// dump requests.
		digest[i].MatchedFlowSchema = r.fsName
		digest[i].FlowDistinguisher = r.flowDistinguisher
		digest[i].ArriveTime = r.arrivalTime
		digest[i].StartTime = r.startTime
		if includeDetails {
			userInfo, _ := genericrequest.UserFrom(r.ctx)
			digest[i].UserName = userInfo.GetName()
			requestInfo, ok := genericrequest.RequestInfoFrom(r.ctx)
			if ok {
				digest[i].RequestInfo = *requestInfo
			}
		}
		i++
		return true
	})
	return debug.QueueDump{
		VirtualStart:      q.nextDispatchR.ToFloat(), // TODO: change QueueDump to use SeatSeconds
		Requests:          digest,
		ExecutingRequests: q.requestsExecuting,
		SeatsInUse:        q.seatsInUse,
	}
}

// SeatSeconds is a measure of work, in units of seat-seconds, using a fixed-point representation.
// `SeatSeconds(n)` represents `n/ssScale` seat-seconds.
// The constants `ssScale` and `ssScaleDigits` are private to the implementation here,
// no other code should use them.
type SeatSeconds uint64

// MaxSeatsSeconds is the maximum representable value of SeatSeconds
const MaxSeatSeconds = SeatSeconds(math.MaxUint64)

// MinSeatSeconds is the lowest representable value of SeatSeconds
const MinSeatSeconds = SeatSeconds(0)

// SeatsTimeDuration produces the SeatSeconds value for the given factors.
// This is intended only to produce small values, increments in work
// rather than amount of work done since process start.
func SeatsTimesDuration(seats float64, duration time.Duration) SeatSeconds {
	return SeatSeconds(math.Round(seats * float64(duration/time.Nanosecond) / (1e9 / ssScale)))
}

// ToFloat converts to a floating-point representation.
// This conversion may lose precision.
func (ss SeatSeconds) ToFloat() float64 {
	return float64(ss) / ssScale
}

// String converts to a string.
// This is suitable for large as well as small values.
func (ss SeatSeconds) String() string {
	const div = SeatSeconds(ssScale)
	quo := ss / div
	rem := ss - quo*div
	return fmt.Sprintf("%d.%08dss", quo, rem)
}

const ssScale = 1e8
