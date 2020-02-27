package component

import (
	"github.com/containers-ai/federatorai-operator/pkg/util"
)

type ComponentConfigOption func(*ComponentConfig)

func WithNamespace(namespace string) ComponentConfigOption {
	return func(cc *ComponentConfig) {
		cc.NameSpace = namespace
	}
}

func WithImageConfig(ic ImageConfig) ComponentConfigOption {
	return func(cc *ComponentConfig) {
		cc.Image = ic
	}
}

func WithPodSecurityPolicyGroup(podSecurityPolicyGroup string) ComponentConfigOption {
	return func(cc *ComponentConfig) {
		cc.PodSecurityPolicyGroup = podSecurityPolicyGroup
	}
}

func WithPodSecurityPolicyVersion(podSecurityPolicyVersion string) ComponentConfigOption {
	return func(cc *ComponentConfig) {
		cc.PodSecurityPolicyVersion = podSecurityPolicyVersion
	}
}

func WithFedermeterConfig(clusterType string) ComponentConfigOption {
	return func(cc *ComponentConfig) {
		if clusterType == "NKS" {
			cc.FedemeterConfig.FedemeterWorkerNodeLowerLimit = util.NKS_FEDEMETER_WORKER_NODE_LOWER_LIMIT
			cc.FedemeterConfig.FedemeterFilterTable = util.NKS_FEDEMETER_FILTER_TABLE
		} else {
			cc.FedemeterConfig.FedemeterWorkerNodeLowerLimit = util.Openshift_FEDEMETER_WORKER_NODE_LOWER_LIMIT
			cc.FedemeterConfig.FedemeterFilterTable = util.Openshift_FEDEMETER_FILTER_TABLE
		}
	}
}
