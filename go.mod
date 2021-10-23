module github.com/sanposhiho/scheduler-playground

go 1.14

replace (
	k8s.io/api => ./submodules/kubernetes/staging/src/k8s.io/api
	k8s.io/apiextensions-apiserver => ./submodules/kubernetes/staging/src/k8s.io/apiextensions-apiserver
	k8s.io/apimachinery => ./submodules/kubernetes/staging/src/k8s.io/apimachinery
	k8s.io/apiserver => ./submodules/kubernetes/staging/src/k8s.io/apiserver
	k8s.io/cli-runtime => ./submodules/kubernetes/staging/src/k8s.io/cli-runtime
	k8s.io/client-go => ./submodules/kubernetes/staging/src/k8s.io/client-go
	k8s.io/cloud-provider => ./submodules/kubernetes/staging/src/k8s.io/cloud-provider
	k8s.io/cluster-bootstrap => ./submodules/kubernetes/staging/src/k8s.io/cluster-bootstrap
	k8s.io/code-generator => ./submodules/kubernetes/staging/src/k8s.io/code-generator
	k8s.io/component-base => ./submodules/kubernetes/staging/src/k8s.io/component-base
	k8s.io/component-helpers => ./submodules/kubernetes/staging/src/k8s.io/component-helpers
	k8s.io/controller-manager => ./submodules/kubernetes/staging/src/k8s.io/controller-manager
	k8s.io/cri-api => ./submodules/kubernetes/staging/src/k8s.io/cri-api
	k8s.io/csi-translation-lib => ./submodules/kubernetes/staging/src/k8s.io/csi-translation-lib
	k8s.io/gengo => k8s.io/gengo v0.0.0-20210813121822-485abfe95c7c
	k8s.io/kube-aggregator => ./submodules/kubernetes/staging/src/k8s.io/kube-aggregator
	k8s.io/kube-controller-manager => ./submodules/kubernetes/staging/src/k8s.io/kube-controller-manager
	k8s.io/kube-openapi => k8s.io/kube-openapi v0.0.0-20210817084001-7fbd8d59e5b8
	k8s.io/kubernetes => ./submodules/kubernetes
	k8s.io/kube-proxy => ./submodules/kubernetes/staging/src/k8s.io/kube-proxy
	k8s.io/kube-scheduler => ./submodules/kubernetes/staging/src/k8s.io/kube-scheduler
	k8s.io/kubectl => ./submodules/kubernetes/staging/src/k8s.io/kubectl
	k8s.io/kubelet => ./submodules/kubernetes/staging/src/k8s.io/kubelet
	k8s.io/legacy-cloud-providers => ./submodules/kubernetes/staging/src/k8s.io/legacy-cloud-providers
	k8s.io/metrics => ./submodules/kubernetes/staging/src/k8s.io/metrics
	k8s.io/mount-utils => ./submodules/kubernetes/staging/src/k8s.io/mount-utils
	k8s.io/pod-security-admission => ./submodules/kubernetes/staging/src/k8s.io/pod-security-admission
	k8s.io/sample-apiserver => ./submodules/kubernetes/staging/src/k8s.io/sample-apiserver
	k8s.io/sample-cli-plugin => ./submodules/kubernetes/staging/src/k8s.io/sample-cli-plugin
	k8s.io/sample-controller => ./submodules/kubernetes/staging/src/k8s.io/sample-controller
)

require (
	github.com/golang/mock v1.6.0
	github.com/google/uuid v1.3.0
	github.com/stretchr/testify v1.7.0
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1
	k8s.io/api v0.0.0
	k8s.io/apiextensions-apiserver v0.0.0
	k8s.io/apimachinery v0.0.0
	k8s.io/apiserver v0.0.0
	k8s.io/client-go v0.0.0
	k8s.io/component-base v0.0.0
	k8s.io/klog/v2 v2.9.0
	k8s.io/kube-openapi v0.0.0-20210421082810-95288971da7e
	k8s.io/kube-scheduler v0.0.0
	k8s.io/kubernetes v1.22.2
)
