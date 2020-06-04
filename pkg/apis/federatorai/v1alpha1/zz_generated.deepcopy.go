// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlamedaComponentSpec) DeepCopyInto(out *AlamedaComponentSpec) {
	*out = *in
	if in.Storages != nil {
		in, out := &in.Storages, &out.Storages
		*out = make([]StorageSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.BootStrapContainer = in.BootStrapContainer
	if in.EnvVars != nil {
		in, out := &in.EnvVars, &out.EnvVars
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Resources.DeepCopyInto(&out.Resources)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlamedaComponentSpec.
func (in *AlamedaComponentSpec) DeepCopy() *AlamedaComponentSpec {
	if in == nil {
		return nil
	}
	out := new(AlamedaComponentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlamedaService) DeepCopyInto(out *AlamedaService) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlamedaService.
func (in *AlamedaService) DeepCopy() *AlamedaService {
	if in == nil {
		return nil
	}
	out := new(AlamedaService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AlamedaService) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlamedaServiceList) DeepCopyInto(out *AlamedaServiceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AlamedaService, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlamedaServiceList.
func (in *AlamedaServiceList) DeepCopy() *AlamedaServiceList {
	if in == nil {
		return nil
	}
	out := new(AlamedaServiceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AlamedaServiceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlamedaServiceSpec) DeepCopyInto(out *AlamedaServiceSpec) {
	*out = *in
	if in.EnableDispatcher != nil {
		in, out := &in.EnableDispatcher, &out.EnableDispatcher
		*out = new(bool)
		**out = **in
	}
	if in.Storages != nil {
		in, out := &in.Storages, &out.Storages
		*out = make([]StorageSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Resources.DeepCopyInto(&out.Resources)
	if in.ServiceExposures != nil {
		in, out := &in.ServiceExposures, &out.ServiceExposures
		*out = make([]ServiceExposureSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.Keycode = in.Keycode
	in.InfluxdbSectionSet.DeepCopyInto(&out.InfluxdbSectionSet)
	in.GrafanaSectionSet.DeepCopyInto(&out.GrafanaSectionSet)
	in.AlamedaAISectionSet.DeepCopyInto(&out.AlamedaAISectionSet)
	in.AlamedaOperatorSectionSet.DeepCopyInto(&out.AlamedaOperatorSectionSet)
	in.AlamedaDatahubSectionSet.DeepCopyInto(&out.AlamedaDatahubSectionSet)
	in.AlamedaEvictionerSectionSet.DeepCopyInto(&out.AlamedaEvictionerSectionSet)
	in.AdmissionControllerSectionSet.DeepCopyInto(&out.AdmissionControllerSectionSet)
	in.AlamedaRecommenderSectionSet.DeepCopyInto(&out.AlamedaRecommenderSectionSet)
	in.AlamedaExecutorSectionSet.DeepCopyInto(&out.AlamedaExecutorSectionSet)
	in.AlamedaFedemeterSectionSet.DeepCopyInto(&out.AlamedaFedemeterSectionSet)
	in.AlamedaFedemeterInfluxDBSectionSet.DeepCopyInto(&out.AlamedaFedemeterInfluxDBSectionSet)
	in.AlamedaWeavescopeSectionSet.DeepCopyInto(&out.AlamedaWeavescopeSectionSet)
	in.AlamedaDispatcherSectionSet.DeepCopyInto(&out.AlamedaDispatcherSectionSet)
	in.AlamedaRabbitMQSectionSet.DeepCopyInto(&out.AlamedaRabbitMQSectionSet)
	in.AlamedaAnalyzerSectionSet.DeepCopyInto(&out.AlamedaAnalyzerSectionSet)
	in.AlamedaNotifierSectionSet.DeepCopyInto(&out.AlamedaNotifierSectionSet)
	in.FederatoraiAgentSectionSet.DeepCopyInto(&out.FederatoraiAgentSectionSet)
	in.FederatoraiAgentGPUSectionSet.DeepCopyInto(&out.FederatoraiAgentGPUSectionSet)
	in.FederatoraiRestSectionSet.DeepCopyInto(&out.FederatoraiRestSectionSet)
	in.FederatoraiAgentPreloaderSectionSet.DeepCopyInto(&out.FederatoraiAgentPreloaderSectionSet)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlamedaServiceSpec.
func (in *AlamedaServiceSpec) DeepCopy() *AlamedaServiceSpec {
	if in == nil {
		return nil
	}
	out := new(AlamedaServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlamedaServiceStatus) DeepCopyInto(out *AlamedaServiceStatus) {
	*out = *in
	out.CRDVersion = in.CRDVersion
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]AlamedaServiceStatusCondition, len(*in))
		copy(*out, *in)
	}
	out.KeycodeStatus = in.KeycodeStatus
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlamedaServiceStatus.
func (in *AlamedaServiceStatus) DeepCopy() *AlamedaServiceStatus {
	if in == nil {
		return nil
	}
	out := new(AlamedaServiceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlamedaServiceStatusCRDVersion) DeepCopyInto(out *AlamedaServiceStatusCRDVersion) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlamedaServiceStatusCRDVersion.
func (in *AlamedaServiceStatusCRDVersion) DeepCopy() *AlamedaServiceStatusCRDVersion {
	if in == nil {
		return nil
	}
	out := new(AlamedaServiceStatusCRDVersion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlamedaServiceStatusCondition) DeepCopyInto(out *AlamedaServiceStatusCondition) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlamedaServiceStatusCondition.
func (in *AlamedaServiceStatusCondition) DeepCopy() *AlamedaServiceStatusCondition {
	if in == nil {
		return nil
	}
	out := new(AlamedaServiceStatusCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FederatoraiAgentGPUSpec) DeepCopyInto(out *FederatoraiAgentGPUSpec) {
	*out = *in
	in.AlamedaComponentSpec.DeepCopyInto(&out.AlamedaComponentSpec)
	if in.Prometheus != nil {
		in, out := &in.Prometheus, &out.Prometheus
		*out = new(PrometheusConfig)
		**out = **in
	}
	if in.InfluxDB != nil {
		in, out := &in.InfluxDB, &out.InfluxDB
		*out = new(InfluxDBConfig)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FederatoraiAgentGPUSpec.
func (in *FederatoraiAgentGPUSpec) DeepCopy() *FederatoraiAgentGPUSpec {
	if in == nil {
		return nil
	}
	out := new(FederatoraiAgentGPUSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Imagestruct) DeepCopyInto(out *Imagestruct) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Imagestruct.
func (in *Imagestruct) DeepCopy() *Imagestruct {
	if in == nil {
		return nil
	}
	out := new(Imagestruct)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfluxDBConfig) DeepCopyInto(out *InfluxDBConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfluxDBConfig.
func (in *InfluxDBConfig) DeepCopy() *InfluxDBConfig {
	if in == nil {
		return nil
	}
	out := new(InfluxDBConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycodeSpec) DeepCopyInto(out *KeycodeSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycodeSpec.
func (in *KeycodeSpec) DeepCopy() *KeycodeSpec {
	if in == nil {
		return nil
	}
	out := new(KeycodeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeycodeStatus) DeepCopyInto(out *KeycodeStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeycodeStatus.
func (in *KeycodeStatus) DeepCopy() *KeycodeStatus {
	if in == nil {
		return nil
	}
	out := new(KeycodeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodePortSpec) DeepCopyInto(out *NodePortSpec) {
	*out = *in
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]PortSpec, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodePortSpec.
func (in *NodePortSpec) DeepCopy() *NodePortSpec {
	if in == nil {
		return nil
	}
	out := new(NodePortSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PortSpec) DeepCopyInto(out *PortSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PortSpec.
func (in *PortSpec) DeepCopy() *PortSpec {
	if in == nil {
		return nil
	}
	out := new(PortSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusConfig) DeepCopyInto(out *PrometheusConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusConfig.
func (in *PrometheusConfig) DeepCopy() *PrometheusConfig {
	if in == nil {
		return nil
	}
	out := new(PrometheusConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceExposureSpec) DeepCopyInto(out *ServiceExposureSpec) {
	*out = *in
	if in.NodePort != nil {
		in, out := &in.NodePort, &out.NodePort
		*out = new(NodePortSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceExposureSpec.
func (in *ServiceExposureSpec) DeepCopy() *ServiceExposureSpec {
	if in == nil {
		return nil
	}
	out := new(ServiceExposureSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageSpec) DeepCopyInto(out *StorageSpec) {
	*out = *in
	if in.Class != nil {
		in, out := &in.Class, &out.Class
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageSpec.
func (in *StorageSpec) DeepCopy() *StorageSpec {
	if in == nil {
		return nil
	}
	out := new(StorageSpec)
	in.DeepCopyInto(out)
	return out
}
