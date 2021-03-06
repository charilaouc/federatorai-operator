package component

import (
	"github.com/containers-ai/federatorai-operator/pkg/util"
)

type InfluxDBConfig struct {
	Address   string
	BasicAuth BasicAuth
}

type PrometheusConfig struct {
	Address         string
	Host            string
	Port            string
	Protocol        string
	BasicAuth       BasicAuth
	BearerTokenFile string
	TLS             TLSConfig
}

type FedemeterConfig struct {
	FedemeterWorkerNodeLowerLimit string
	FedemeterFilterTable          string
}

type KafkaConfig struct {
	Enabled         bool
	BrokerAddresses []string
	Version         string

	SASL SASLConfig
	TLS  TLSConfig
}

type NginxConfig struct {
	Enabled bool
}

type ClusterAutoScalerConfig struct {
	EnableExecution bool
}

type VolumeNameSuffixes struct {
	Data string
	Log  string
}

func NewDefaultNginxConfig() NginxConfig {
	return NginxConfig{Enabled: false}
}

func NewDefaultFedemeterConfig() FedemeterConfig {
	return FedemeterConfig{
		FedemeterWorkerNodeLowerLimit: util.Openshift_FEDEMETER_WORKER_NODE_LOWER_LIMIT,
		FedemeterFilterTable:          util.Openshift_FEDEMETER_FILTER_TABLE,
	}
}

type BasicAuth struct {
	Username string
	Password string
}

type SASLConfig struct {
	Enabled   bool
	BasicAuth BasicAuth
}

type TLSConfig struct {
	Enabled            bool
	InsecureSkipVerify bool
}

type FederatoraiAgentGPUDatasourceConfig struct {
	InfluxDB   InfluxDBConfig
	Prometheus PrometheusConfig
}

type FederatoraiAgentGPUConfig struct {
	Enabled    bool
	Datasource FederatoraiAgentGPUDatasourceConfig
}

func NewDefaultFederatoraiAgentGPUConfig() FederatoraiAgentGPUConfig {
	return FederatoraiAgentGPUConfig{
		Enabled: false,
		Datasource: FederatoraiAgentGPUDatasourceConfig{
			InfluxDB: InfluxDBConfig{
				Address: "",
				BasicAuth: BasicAuth{
					Username: "",
					Password: "",
				},
			},
			Prometheus: PrometheusConfig{
				Address: "",
				BasicAuth: BasicAuth{
					Username: "",
					Password: "",
				},
			},
		},
	}
}

type AIDispatcherConfig struct {
	Enabled bool
}

func NewDefaultAIDispatcherConfig() AIDispatcherConfig {
	return AIDispatcherConfig{
		Enabled: true,
	}
}

type ExecutionConfig struct {
	EnabledVPA bool
}

func NewDefaultExecutionConfig() ExecutionConfig {
	return ExecutionConfig{
		EnabledVPA: true,
	}
}
