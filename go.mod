module github.com/containers-ai/federatorai-operator

go 1.13

require (
	cloud.google.com/go v0.56.0 // indirect
	github.com/containers-ai/alameda v0.9.168-0.20200714074612-47b570672d31
	github.com/containers-ai/api v4.2.790-0.20200802143522-a83d2a3dbcb5+incompatible
	github.com/coreos/prometheus-operator v0.39.0
	github.com/ghodss/yaml v1.0.0
	github.com/go-logr/logr v0.1.0
	github.com/go-logr/zapr v0.1.1
	github.com/golang/protobuf v1.4.2
	github.com/googleapis/gnostic v0.4.0 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.0
	github.com/influxdata/influxdb1-client v0.0.0-20200515024757-02f0bf5dbca3
	github.com/jetstack/cert-manager v0.15.2
	github.com/kr/pretty v0.2.0 // indirect
	github.com/magiconair/properties v1.8.1 // indirect
	github.com/mitchellh/mapstructure v1.2.2 // indirect
	github.com/openshift/api v0.0.0-20200526144822-34f54f12813a
	github.com/pelletier/go-toml v1.4.0 // indirect
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.7.1 // indirect
	github.com/spf13/cast v1.3.1 // indirect
	github.com/spf13/cobra v1.0.0 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/viper v1.4.0
	github.com/stretchr/testify v1.5.1
	go.uber.org/zap v1.15.0
	golang.org/x/crypto v0.0.0-20200423211502-4bdfaf469ed5 // indirect
	golang.org/x/net v0.0.0-20200602114024-627f9648deb9 // indirect
	golang.org/x/sync v0.0.0-20200317015054-43a5402ce75a
	golang.org/x/time v0.0.0-20200416051211-89c76fbcd5d1 // indirect
	golang.org/x/tools v0.0.0-20200603131246-cc40288be839 // indirect
	google.golang.org/appengine v1.6.6 // indirect
	google.golang.org/genproto v0.0.0-20200603110839-e855014d5736
	google.golang.org/grpc v1.29.1
	gopkg.in/yaml.v3 v3.0.0-20200603094226-e3079894b1e8
	k8s.io/api v0.18.6
	k8s.io/apiextensions-apiserver v0.18.6
	k8s.io/apimachinery v0.18.6
	k8s.io/client-go v12.0.0+incompatible
	k8s.io/kube-aggregator v0.18.6
	sigs.k8s.io/controller-runtime v0.6.1
)

replace k8s.io/client-go => k8s.io/client-go v0.18.6
