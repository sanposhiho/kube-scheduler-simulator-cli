# kube-scheduler-simulator-cli: Kubernetes Scheduler simulator on CLI and scenario system. 

Hello world.

This repository is scenario system for kube-scheduler. You can write scenario and check the scheduler's behaviours.

And, also you can change the scheduler implementations on submodule/kubernetes. It helps changing implementations and debugging for scheeduler.

## How to write the scenario

see `scenario` function on [sched.go](/sched.go).

```go
func scenario(client clientset.Interface) error {
	ctx := context.Background()

	// create node0 ~ node9, but all nodes are unschedulable
	for i := 0; i < 9; i++ {
		suffix := strconv.Itoa(i)
		_, err := client.CoreV1().Nodes().Create(ctx, &v1.Node{
			ObjectMeta: metav1.ObjectMeta{
				Name: "node" + suffix,
			},
			Spec: v1.NodeSpec{
				Unschedulable: true,
			},
		}, metav1.CreateOptions{})
		if err != nil {
			return fmt.Errorf("create node: %w", err)
		}
	}

	klog.Info("scenario: all nodes created")

	_, err := client.CoreV1().Pods("default").Create(ctx, &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{Name: "pod1"},
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{
					Name:  "container1",
					Image: "k8s.gcr.io/pause:3.5",
				},
			},
		},
	}, metav1.CreateOptions{})
	if err != nil {
		return fmt.Errorf("create pod: %w", err)
	}

	klog.Info("scenario: pod1 created")

	// wait to schedule
	time.Sleep(3 * time.Second)

	pod1, err := client.CoreV1().Pods("default").Get(ctx, "pod1", metav1.GetOptions{})
	if err != nil {
		return fmt.Errorf("get pod: %w", err)
	}

    klog.Info("scenario: pod1 is bound to " + pod1.Spec.NodeName)

    return nil
}
```

## How to change the scheduler implementations 

We have kubernetes/kubernetes repo as git submodule on [submodules/kubernetes](./submodules/kubernetes)

On this simulator, we build the scheduler with the submodule.
So you can change scheduler version or even you can change implementations by changing the scheduler implementations on submodule/kubernetes.

## How to start this scheduler and scenario

### 0. install etcd

To run this scheduler and start scenario, you have to install Go and etcd.
You can install etcd with [submodules/kubernetes/kubernetes/kubernetes/hack/install-etcd.sh](https://github.com/kubernetes/kubernetes/blob/master/hack/install-etcd.sh).

### 1. edit go.mod on submodule

add `k8s.io/kubernetes => ./` on replace directive. like this:

```
replace (
...
...
(the other replacements...)
..
.

k8s.io/kubernetes => ./
```

### 2. let's start the scenario and scheduler.

And, `make start` starts the scheduler and your scenario.

## Note

This scheduler-playground starts scheduler, etcd, api-server and pv-controller.

The whole mechanism is based on [kubernetes-sigs/kube-scheduler-simulator](https://github.com/kubernetes-sigs/kube-scheduler-simulator)
