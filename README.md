# mini-kube-scheduler: Scheduler for learning Kubernetes Scheduler

Hello world. 

This repository is scenario system for kube-scheduler. You can write scenario like this and check the scheduler's behaviours.

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

## custom scenario

You can write scenario [here](/sched.go#L70) and check this scheduler's behaviour.

## How to start this scheduler and scenario

To run this scheduler and start scenario, you have to install Go and etcd.
You can install etcd with [kubernetes/kubernetes/hack/install-etcd.sh](https://github.com/kubernetes/kubernetes/blob/master/hack/install-etcd.sh).

And, `make start` starts your scenario.

## Note

This mini-kube-scheduler starts scheduler, etcd, api-server and pv-controller.

The whole mechanism is based on [kubernetes-sigs/kube-scheduler-simulator](https://github.com/kubernetes-sigs/kube-scheduler-simulator)

