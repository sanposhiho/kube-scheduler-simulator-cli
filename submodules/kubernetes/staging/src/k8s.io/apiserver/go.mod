// This is a generated file. Do not edit directly.

module k8s.io/apiserver

go 1.16

require (
	github.com/NYTimes/gziphandler v1.1.1 // indirect
	github.com/coreos/go-oidc v2.1.0+incompatible
	github.com/coreos/go-systemd/v22 v22.3.2
	github.com/davecgh/go-spew v1.1.1
	github.com/emicklei/go-restful v2.9.5+incompatible
	github.com/evanphx/json-patch v4.11.0+incompatible
	github.com/fsnotify/fsnotify v1.4.9
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/go-openapi/jsonreference v0.19.5 // indirect
	github.com/go-openapi/swag v0.19.14 // indirect
	github.com/gogo/protobuf v1.3.2
	github.com/google/go-cmp v0.5.5
	github.com/google/gofuzz v1.1.0
	github.com/google/uuid v1.1.2
	github.com/googleapis/gnostic v0.5.5
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822
	github.com/pquerna/cachecontrol v0.0.0-20171018203845-0dec1b30a021 // indirect
	github.com/sirupsen/logrus v1.8.1 // indirect
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.7.0
	go.etcd.io/etcd/api/v3 v3.5.0
	go.etcd.io/etcd/client/pkg/v3 v3.5.0
	go.etcd.io/etcd/client/v3 v3.5.0
	go.etcd.io/etcd/server/v3 v3.5.0
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.20.0
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.20.0
	go.opentelemetry.io/otel v0.20.0
	go.opentelemetry.io/otel/exporters/otlp v0.20.0
	go.opentelemetry.io/otel/sdk v0.20.0
	go.opentelemetry.io/otel/trace v0.20.0
	go.uber.org/zap v1.19.0
	golang.org/x/crypto v0.0.0-20210817164053-32db794688a5
	golang.org/x/net v0.0.0-20210813160813-60bc85c4be6d
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	golang.org/x/sys v0.0.0-20210820121016-41cdb8703e55
	google.golang.org/grpc v1.38.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	gopkg.in/square/go-jose.v2 v2.2.2
	k8s.io/api v0.0.0
	k8s.io/apimachinery v0.0.0
	k8s.io/client-go v0.0.0
	k8s.io/component-base v0.0.0
	k8s.io/klog/v2 v2.20.0
	k8s.io/kube-openapi v0.0.0-20210817084001-7fbd8d59e5b8
	k8s.io/utils v0.0.0-20210930125809-cb0fa318a74b
	sigs.k8s.io/apiserver-network-proxy/konnectivity-client v0.0.23
	sigs.k8s.io/structured-merge-diff/v4 v4.1.2
	sigs.k8s.io/yaml v1.2.0
)

replace (
	k8s.io/api => ../api
	k8s.io/apimachinery => ../apimachinery
	k8s.io/apiserver => ../apiserver
	k8s.io/client-go => ../client-go
	k8s.io/component-base => ../component-base
)
