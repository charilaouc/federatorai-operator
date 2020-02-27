package component

import (
	"github.com/containers-ai/federatorai-operator/pkg/util"
)

type InfluxDBConfig struct {
	Address  string
	Username string
	Password string
}

type FedemeterConfig struct {
	FedemeterWorkerNodeLowerLimit string
	FedemeterFilterTable          string
}

func NewDefaultFedemeterConfig() FedemeterConfig {
	return FedemeterConfig{
		FedemeterWorkerNodeLowerLimit: util.Openshift_FEDEMETER_WORKER_NODE_LOWER_LIMIT,
		FedemeterFilterTable:          util.Openshift_FEDEMETER_FILTER_TABLE,
	}
}

type PrometheusConfig struct {
	Address  string
	Username string
	Password string
}

type FederatoraiAgentGPUDatasourceConfig struct {
	InfluxDB   InfluxDBConfig
	Prometheus PrometheusConfig
}

type FederatoraiAgentGPUConfig struct {
	Datasource FederatoraiAgentGPUDatasourceConfig
}

func NewDefaultFederatoraiAgentGPUConfig() FederatoraiAgentGPUConfig {
	return FederatoraiAgentGPUConfig{
		Datasource: FederatoraiAgentGPUDatasourceConfig{
			InfluxDB: InfluxDBConfig{
				Address:  "",
				Username: "",
				Password: "",
			},
			Prometheus: PrometheusConfig{
				Address:  "",
				Username: "",
				Password: "",
			},
		},
	}
}
