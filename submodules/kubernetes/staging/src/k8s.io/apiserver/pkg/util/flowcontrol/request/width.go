/*
Copyright 2021 The Kubernetes Authors.

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

package request

import (
	"fmt"
	"net/http"
	"time"

	apirequest "k8s.io/apiserver/pkg/endpoints/request"
	"k8s.io/klog/v2"
)

const (
	// the minimum number of seats a request must occupy
	minimumSeats = 1

	// the maximum number of seats a request can occupy
	maximumSeats = 10
)

type WorkEstimate struct {
	// InitialSeats represents the number of initial seats associated with this request
	InitialSeats uint

	// AdditionalLatency specifies the additional duration the seats allocated
	// to this request must be reserved after the given request had finished.
	// AdditionalLatency should not have any impact on the user experience, the
	// caller must not experience this additional latency.
	AdditionalLatency time.Duration
}

// objectCountGetterFunc represents a function that gets the total
// number of objects for a given resource.
type objectCountGetterFunc func(string) (int64, error)

// NewWorkEstimator estimates the work that will be done by a given request,
// if no WorkEstimatorFunc matches the given request then the default
// work estimate of 1 seat is allocated to the request.
func NewWorkEstimator(countFn objectCountGetterFunc) WorkEstimatorFunc {
	estimator := &workEstimator{
		listWorkEstimator: newListWorkEstimator(countFn),
	}
	return estimator.estimate
}

// WorkEstimatorFunc returns the estimated work of a given request.
// This function will be used by the Priority & Fairness filter to
// estimate the work of of incoming requests.
type WorkEstimatorFunc func(*http.Request) WorkEstimate

func (e WorkEstimatorFunc) EstimateWork(r *http.Request) WorkEstimate {
	return e(r)
}

type workEstimator struct {
	// listWorkEstimator estimates work for list request(s)
	listWorkEstimator WorkEstimatorFunc
}

func (e *workEstimator) estimate(r *http.Request) WorkEstimate {
	requestInfo, ok := apirequest.RequestInfoFrom(r.Context())
	if !ok {
		klog.ErrorS(fmt.Errorf("no RequestInfo found in context"), "Failed to estimate work for the request", "URI", r.RequestURI)
		// no RequestInfo should never happen, but to be on the safe side let's return maximumSeats
		return WorkEstimate{InitialSeats: maximumSeats}
	}

	switch requestInfo.Verb {
	case "list":
		return e.listWorkEstimator.EstimateWork(r)
	}

	return WorkEstimate{InitialSeats: minimumSeats}
}
