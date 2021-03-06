// This file has messages related to metric data of containers, pods, and nodes

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v4.0.0
// source: alameda_api/v1alpha1/datahub/metrics/metrics.proto

package metrics

import (
	common "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/common"
	resources "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/resources"
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Represents metric data of a container
type ContainerMetric struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	MetricData []*common.MetricData `protobuf:"bytes,2,rep,name=metric_data,json=metricData,proto3" json:"metric_data,omitempty"`
}

func (x *ContainerMetric) Reset() {
	*x = ContainerMetric{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alameda_api_v1alpha1_datahub_metrics_metrics_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContainerMetric) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContainerMetric) ProtoMessage() {}

func (x *ContainerMetric) ProtoReflect() protoreflect.Message {
	mi := &file_alameda_api_v1alpha1_datahub_metrics_metrics_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContainerMetric.ProtoReflect.Descriptor instead.
func (*ContainerMetric) Descriptor() ([]byte, []int) {
	return file_alameda_api_v1alpha1_datahub_metrics_metrics_proto_rawDescGZIP(), []int{0}
}

func (x *ContainerMetric) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ContainerMetric) GetMetricData() []*common.MetricData {
	if x != nil {
		return x.MetricData
	}
	return nil
}

// Represents metric data of a pod
type PodMetric struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ObjectMeta       *resources.ObjectMeta `protobuf:"bytes,1,opt,name=object_meta,json=objectMeta,proto3" json:"object_meta,omitempty"`
	ContainerMetrics []*ContainerMetric    `protobuf:"bytes,2,rep,name=container_metrics,json=containerMetrics,proto3" json:"container_metrics,omitempty"`
}

func (x *PodMetric) Reset() {
	*x = PodMetric{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alameda_api_v1alpha1_datahub_metrics_metrics_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PodMetric) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PodMetric) ProtoMessage() {}

func (x *PodMetric) ProtoReflect() protoreflect.Message {
	mi := &file_alameda_api_v1alpha1_datahub_metrics_metrics_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PodMetric.ProtoReflect.Descriptor instead.
func (*PodMetric) Descriptor() ([]byte, []int) {
	return file_alameda_api_v1alpha1_datahub_metrics_metrics_proto_rawDescGZIP(), []int{1}
}

func (x *PodMetric) GetObjectMeta() *resources.ObjectMeta {
	if x != nil {
		return x.ObjectMeta
	}
	return nil
}

func (x *PodMetric) GetContainerMetrics() []*ContainerMetric {
	if x != nil {
		return x.ContainerMetrics
	}
	return nil
}

type ControllerMetric struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ObjectMeta *resources.ObjectMeta `protobuf:"bytes,1,opt,name=object_meta,json=objectMeta,proto3" json:"object_meta,omitempty"`
	Kind       resources.Kind        `protobuf:"varint,2,opt,name=kind,proto3,enum=containersai.alameda.v1alpha1.datahub.resources.Kind" json:"kind,omitempty"`
	MetricData []*common.MetricData  `protobuf:"bytes,3,rep,name=metric_data,json=metricData,proto3" json:"metric_data,omitempty"`
}

func (x *ControllerMetric) Reset() {
	*x = ControllerMetric{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alameda_api_v1alpha1_datahub_metrics_metrics_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ControllerMetric) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ControllerMetric) ProtoMessage() {}

func (x *ControllerMetric) ProtoReflect() protoreflect.Message {
	mi := &file_alameda_api_v1alpha1_datahub_metrics_metrics_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ControllerMetric.ProtoReflect.Descriptor instead.
func (*ControllerMetric) Descriptor() ([]byte, []int) {
	return file_alameda_api_v1alpha1_datahub_metrics_metrics_proto_rawDescGZIP(), []int{2}
}

func (x *ControllerMetric) GetObjectMeta() *resources.ObjectMeta {
	if x != nil {
		return x.ObjectMeta
	}
	return nil
}

func (x *ControllerMetric) GetKind() resources.Kind {
	if x != nil {
		return x.Kind
	}
	return resources.Kind_KIND_UNDEFINED
}

func (x *ControllerMetric) GetMetricData() []*common.MetricData {
	if x != nil {
		return x.MetricData
	}
	return nil
}

type ApplicationMetric struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ObjectMeta *resources.ObjectMeta `protobuf:"bytes,1,opt,name=object_meta,json=objectMeta,proto3" json:"object_meta,omitempty"`
	MetricData []*common.MetricData  `protobuf:"bytes,2,rep,name=metric_data,json=metricData,proto3" json:"metric_data,omitempty"`
}

func (x *ApplicationMetric) Reset() {
	*x = ApplicationMetric{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alameda_api_v1alpha1_datahub_metrics_metrics_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplicationMetric) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplicationMetric) ProtoMessage() {}

func (x *ApplicationMetric) ProtoReflect() protoreflect.Message {
	mi := &file_alameda_api_v1alpha1_datahub_metrics_metrics_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplicationMetric.ProtoReflect.Descriptor instead.
func (*ApplicationMetric) Descriptor() ([]byte, []int) {
	return file_alameda_api_v1alpha1_datahub_metrics_metrics_proto_rawDescGZIP(), []int{3}
}

func (x *ApplicationMetric) GetObjectMeta() *resources.ObjectMeta {
	if x != nil {
		return x.ObjectMeta
	}
	return nil
}

func (x *ApplicationMetric) GetMetricData() []*common.MetricData {
	if x != nil {
		return x.MetricData
	}
	return nil
}

type NamespaceMetric struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ObjectMeta *resources.ObjectMeta `protobuf:"bytes,1,opt,name=object_meta,json=objectMeta,proto3" json:"object_meta,omitempty"`
	MetricData []*common.MetricData  `protobuf:"bytes,2,rep,name=metric_data,json=metricData,proto3" json:"metric_data,omitempty"`
}

func (x *NamespaceMetric) Reset() {
	*x = NamespaceMetric{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alameda_api_v1alpha1_datahub_metrics_metrics_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NamespaceMetric) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NamespaceMetric) ProtoMessage() {}

func (x *NamespaceMetric) ProtoReflect() protoreflect.Message {
	mi := &file_alameda_api_v1alpha1_datahub_metrics_metrics_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NamespaceMetric.ProtoReflect.Descriptor instead.
func (*NamespaceMetric) Descriptor() ([]byte, []int) {
	return file_alameda_api_v1alpha1_datahub_metrics_metrics_proto_rawDescGZIP(), []int{4}
}

func (x *NamespaceMetric) GetObjectMeta() *resources.ObjectMeta {
	if x != nil {
		return x.ObjectMeta
	}
	return nil
}

func (x *NamespaceMetric) GetMetricData() []*common.MetricData {
	if x != nil {
		return x.MetricData
	}
	return nil
}

// Represents metric data of a node
type NodeMetric struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ObjectMeta *resources.ObjectMeta `protobuf:"bytes,1,opt,name=object_meta,json=objectMeta,proto3" json:"object_meta,omitempty"`
	MetricData []*common.MetricData  `protobuf:"bytes,2,rep,name=metric_data,json=metricData,proto3" json:"metric_data,omitempty"`
}

func (x *NodeMetric) Reset() {
	*x = NodeMetric{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alameda_api_v1alpha1_datahub_metrics_metrics_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeMetric) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeMetric) ProtoMessage() {}

func (x *NodeMetric) ProtoReflect() protoreflect.Message {
	mi := &file_alameda_api_v1alpha1_datahub_metrics_metrics_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeMetric.ProtoReflect.Descriptor instead.
func (*NodeMetric) Descriptor() ([]byte, []int) {
	return file_alameda_api_v1alpha1_datahub_metrics_metrics_proto_rawDescGZIP(), []int{5}
}

func (x *NodeMetric) GetObjectMeta() *resources.ObjectMeta {
	if x != nil {
		return x.ObjectMeta
	}
	return nil
}

func (x *NodeMetric) GetMetricData() []*common.MetricData {
	if x != nil {
		return x.MetricData
	}
	return nil
}

type ClusterMetric struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ObjectMeta *resources.ObjectMeta `protobuf:"bytes,1,opt,name=object_meta,json=objectMeta,proto3" json:"object_meta,omitempty"`
	MetricData []*common.MetricData  `protobuf:"bytes,2,rep,name=metric_data,json=metricData,proto3" json:"metric_data,omitempty"`
}

func (x *ClusterMetric) Reset() {
	*x = ClusterMetric{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alameda_api_v1alpha1_datahub_metrics_metrics_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterMetric) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterMetric) ProtoMessage() {}

func (x *ClusterMetric) ProtoReflect() protoreflect.Message {
	mi := &file_alameda_api_v1alpha1_datahub_metrics_metrics_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterMetric.ProtoReflect.Descriptor instead.
func (*ClusterMetric) Descriptor() ([]byte, []int) {
	return file_alameda_api_v1alpha1_datahub_metrics_metrics_proto_rawDescGZIP(), []int{6}
}

func (x *ClusterMetric) GetObjectMeta() *resources.ObjectMeta {
	if x != nil {
		return x.ObjectMeta
	}
	return nil
}

func (x *ClusterMetric) GetMetricData() []*common.MetricData {
	if x != nil {
		return x.MetricData
	}
	return nil
}

type WriteMetric struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MetricType common.MetricType `protobuf:"varint,1,opt,name=metric_type,json=metricType,proto3,enum=containersai.alameda.v1alpha1.datahub.common.MetricType" json:"metric_type,omitempty"`
	WriteData  *common.WriteData `protobuf:"bytes,2,opt,name=write_data,json=writeData,proto3" json:"write_data,omitempty"`
}

func (x *WriteMetric) Reset() {
	*x = WriteMetric{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alameda_api_v1alpha1_datahub_metrics_metrics_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteMetric) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteMetric) ProtoMessage() {}

func (x *WriteMetric) ProtoReflect() protoreflect.Message {
	mi := &file_alameda_api_v1alpha1_datahub_metrics_metrics_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteMetric.ProtoReflect.Descriptor instead.
func (*WriteMetric) Descriptor() ([]byte, []int) {
	return file_alameda_api_v1alpha1_datahub_metrics_metrics_proto_rawDescGZIP(), []int{7}
}

func (x *WriteMetric) GetMetricType() common.MetricType {
	if x != nil {
		return x.MetricType
	}
	return common.MetricType_METRICS_TYPE_UNDEFINED
}

func (x *WriteMetric) GetWriteData() *common.WriteData {
	if x != nil {
		return x.WriteData
	}
	return nil
}

var File_alameda_api_v1alpha1_datahub_metrics_metrics_proto protoreflect.FileDescriptor

var file_alameda_api_v1alpha1_datahub_metrics_metrics_proto_rawDesc = []byte{
	0x0a, 0x32, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2f, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73,
	0x61, 0x69, 0x2e, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x6d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x73, 0x1a, 0x31, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x5f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75,
	0x62, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x31, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x5f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x64, 0x61, 0x74,
	0x61, 0x68, 0x75, 0x62, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x72, 0x61, 0x77, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x35, 0x61, 0x6c, 0x61, 0x6d, 0x65,
	0x64, 0x61, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f,
	0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x80, 0x01, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x4d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x59, 0x0a, 0x0b, 0x6d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x38, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x61, 0x69, 0x2e, 0x61, 0x6c, 0x61,
	0x6d, 0x65, 0x64, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0a, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x44,
	0x61, 0x74, 0x61, 0x22, 0xd6, 0x01, 0x0a, 0x09, 0x50, 0x6f, 0x64, 0x4d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x12, 0x5c, 0x0a, 0x0b, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e,
	0x65, 0x72, 0x73, 0x61, 0x69, 0x2e, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4d,
	0x65, 0x74, 0x61, 0x52, 0x0a, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x12,
	0x6b, 0x0a, 0x11, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x5f, 0x6d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3e, 0x2e, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x61, 0x69, 0x2e, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64,
	0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68,
	0x75, 0x62, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x52, 0x10, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x22, 0x96, 0x02, 0x0a,
	0x10, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x12, 0x5c, 0x0a, 0x0b, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e,
	0x65, 0x72, 0x73, 0x61, 0x69, 0x2e, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4d,
	0x65, 0x74, 0x61, 0x52, 0x0a, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x12,
	0x49, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x35, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x61, 0x69, 0x2e, 0x61, 0x6c, 0x61,
	0x6d, 0x65, 0x64, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e,
	0x4b, 0x69, 0x6e, 0x64, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x59, 0x0a, 0x0b, 0x6d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x38, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x61, 0x69, 0x2e, 0x61,
	0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0a, 0x6d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x44, 0x61, 0x74, 0x61, 0x22, 0xcc, 0x01, 0x0a, 0x11, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x12, 0x5c, 0x0a, 0x0b, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x3b, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x61, 0x69, 0x2e,
	0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x0a, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x59, 0x0a, 0x0b, 0x6d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x38,
	0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x61, 0x69, 0x2e, 0x61, 0x6c,
	0x61, 0x6d, 0x65, 0x64, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0a, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x44, 0x61, 0x74, 0x61, 0x22, 0xca, 0x01, 0x0a, 0x0f, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x12, 0x5c, 0x0a, 0x0b, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3b, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x61, 0x69, 0x2e, 0x61, 0x6c, 0x61,
	0x6d, 0x65, 0x64, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x0a, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x59, 0x0a, 0x0b, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x61, 0x69, 0x2e, 0x61, 0x6c, 0x61, 0x6d, 0x65,
	0x64, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0a, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x44, 0x61, 0x74,
	0x61, 0x22, 0xc5, 0x01, 0x0a, 0x0a, 0x4e, 0x6f, 0x64, 0x65, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x12, 0x5c, 0x0a, 0x0b, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x73, 0x61, 0x69, 0x2e, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4d, 0x65,
	0x74, 0x61, 0x52, 0x0a, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x59,
	0x0a, 0x0b, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73,
	0x61, 0x69, 0x2e, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0a, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x44, 0x61, 0x74, 0x61, 0x22, 0xc8, 0x01, 0x0a, 0x0d, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x12, 0x5c, 0x0a, 0x0b, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x3b, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x61, 0x69, 0x2e,
	0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x0a, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x59, 0x0a, 0x0b, 0x6d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x38,
	0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x61, 0x69, 0x2e, 0x61, 0x6c,
	0x61, 0x6d, 0x65, 0x64, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0a, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x44, 0x61, 0x74, 0x61, 0x22, 0xc0, 0x01, 0x0a, 0x0b, 0x57, 0x72, 0x69, 0x74, 0x65, 0x4d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x12, 0x59, 0x0a, 0x0b, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x38, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x61, 0x69, 0x2e, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x0a, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x56, 0x0a, 0x0a, 0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73,
	0x61, 0x69, 0x2e, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x09, 0x77, 0x72,
	0x69, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x42, 0x43, 0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73,
	0x2d, 0x61, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x5f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x64, 0x61, 0x74,
	0x61, 0x68, 0x75, 0x62, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_alameda_api_v1alpha1_datahub_metrics_metrics_proto_rawDescOnce sync.Once
	file_alameda_api_v1alpha1_datahub_metrics_metrics_proto_rawDescData = file_alameda_api_v1alpha1_datahub_metrics_metrics_proto_rawDesc
)

func file_alameda_api_v1alpha1_datahub_metrics_metrics_proto_rawDescGZIP() []byte {
	file_alameda_api_v1alpha1_datahub_metrics_metrics_proto_rawDescOnce.Do(func() {
		file_alameda_api_v1alpha1_datahub_metrics_metrics_proto_rawDescData = protoimpl.X.CompressGZIP(file_alameda_api_v1alpha1_datahub_metrics_metrics_proto_rawDescData)
	})
	return file_alameda_api_v1alpha1_datahub_metrics_metrics_proto_rawDescData
}

var file_alameda_api_v1alpha1_datahub_metrics_metrics_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_alameda_api_v1alpha1_datahub_metrics_metrics_proto_goTypes = []interface{}{
	(*ContainerMetric)(nil),      // 0: containersai.alameda.v1alpha1.datahub.metrics.ContainerMetric
	(*PodMetric)(nil),            // 1: containersai.alameda.v1alpha1.datahub.metrics.PodMetric
	(*ControllerMetric)(nil),     // 2: containersai.alameda.v1alpha1.datahub.metrics.ControllerMetric
	(*ApplicationMetric)(nil),    // 3: containersai.alameda.v1alpha1.datahub.metrics.ApplicationMetric
	(*NamespaceMetric)(nil),      // 4: containersai.alameda.v1alpha1.datahub.metrics.NamespaceMetric
	(*NodeMetric)(nil),           // 5: containersai.alameda.v1alpha1.datahub.metrics.NodeMetric
	(*ClusterMetric)(nil),        // 6: containersai.alameda.v1alpha1.datahub.metrics.ClusterMetric
	(*WriteMetric)(nil),          // 7: containersai.alameda.v1alpha1.datahub.metrics.WriteMetric
	(*common.MetricData)(nil),    // 8: containersai.alameda.v1alpha1.datahub.common.MetricData
	(*resources.ObjectMeta)(nil), // 9: containersai.alameda.v1alpha1.datahub.resources.ObjectMeta
	(resources.Kind)(0),          // 10: containersai.alameda.v1alpha1.datahub.resources.Kind
	(common.MetricType)(0),       // 11: containersai.alameda.v1alpha1.datahub.common.MetricType
	(*common.WriteData)(nil),     // 12: containersai.alameda.v1alpha1.datahub.common.WriteData
}
var file_alameda_api_v1alpha1_datahub_metrics_metrics_proto_depIdxs = []int32{
	8,  // 0: containersai.alameda.v1alpha1.datahub.metrics.ContainerMetric.metric_data:type_name -> containersai.alameda.v1alpha1.datahub.common.MetricData
	9,  // 1: containersai.alameda.v1alpha1.datahub.metrics.PodMetric.object_meta:type_name -> containersai.alameda.v1alpha1.datahub.resources.ObjectMeta
	0,  // 2: containersai.alameda.v1alpha1.datahub.metrics.PodMetric.container_metrics:type_name -> containersai.alameda.v1alpha1.datahub.metrics.ContainerMetric
	9,  // 3: containersai.alameda.v1alpha1.datahub.metrics.ControllerMetric.object_meta:type_name -> containersai.alameda.v1alpha1.datahub.resources.ObjectMeta
	10, // 4: containersai.alameda.v1alpha1.datahub.metrics.ControllerMetric.kind:type_name -> containersai.alameda.v1alpha1.datahub.resources.Kind
	8,  // 5: containersai.alameda.v1alpha1.datahub.metrics.ControllerMetric.metric_data:type_name -> containersai.alameda.v1alpha1.datahub.common.MetricData
	9,  // 6: containersai.alameda.v1alpha1.datahub.metrics.ApplicationMetric.object_meta:type_name -> containersai.alameda.v1alpha1.datahub.resources.ObjectMeta
	8,  // 7: containersai.alameda.v1alpha1.datahub.metrics.ApplicationMetric.metric_data:type_name -> containersai.alameda.v1alpha1.datahub.common.MetricData
	9,  // 8: containersai.alameda.v1alpha1.datahub.metrics.NamespaceMetric.object_meta:type_name -> containersai.alameda.v1alpha1.datahub.resources.ObjectMeta
	8,  // 9: containersai.alameda.v1alpha1.datahub.metrics.NamespaceMetric.metric_data:type_name -> containersai.alameda.v1alpha1.datahub.common.MetricData
	9,  // 10: containersai.alameda.v1alpha1.datahub.metrics.NodeMetric.object_meta:type_name -> containersai.alameda.v1alpha1.datahub.resources.ObjectMeta
	8,  // 11: containersai.alameda.v1alpha1.datahub.metrics.NodeMetric.metric_data:type_name -> containersai.alameda.v1alpha1.datahub.common.MetricData
	9,  // 12: containersai.alameda.v1alpha1.datahub.metrics.ClusterMetric.object_meta:type_name -> containersai.alameda.v1alpha1.datahub.resources.ObjectMeta
	8,  // 13: containersai.alameda.v1alpha1.datahub.metrics.ClusterMetric.metric_data:type_name -> containersai.alameda.v1alpha1.datahub.common.MetricData
	11, // 14: containersai.alameda.v1alpha1.datahub.metrics.WriteMetric.metric_type:type_name -> containersai.alameda.v1alpha1.datahub.common.MetricType
	12, // 15: containersai.alameda.v1alpha1.datahub.metrics.WriteMetric.write_data:type_name -> containersai.alameda.v1alpha1.datahub.common.WriteData
	16, // [16:16] is the sub-list for method output_type
	16, // [16:16] is the sub-list for method input_type
	16, // [16:16] is the sub-list for extension type_name
	16, // [16:16] is the sub-list for extension extendee
	0,  // [0:16] is the sub-list for field type_name
}

func init() { file_alameda_api_v1alpha1_datahub_metrics_metrics_proto_init() }
func file_alameda_api_v1alpha1_datahub_metrics_metrics_proto_init() {
	if File_alameda_api_v1alpha1_datahub_metrics_metrics_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_alameda_api_v1alpha1_datahub_metrics_metrics_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContainerMetric); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_alameda_api_v1alpha1_datahub_metrics_metrics_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PodMetric); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_alameda_api_v1alpha1_datahub_metrics_metrics_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ControllerMetric); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_alameda_api_v1alpha1_datahub_metrics_metrics_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplicationMetric); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_alameda_api_v1alpha1_datahub_metrics_metrics_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NamespaceMetric); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_alameda_api_v1alpha1_datahub_metrics_metrics_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeMetric); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_alameda_api_v1alpha1_datahub_metrics_metrics_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterMetric); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_alameda_api_v1alpha1_datahub_metrics_metrics_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteMetric); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_alameda_api_v1alpha1_datahub_metrics_metrics_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_alameda_api_v1alpha1_datahub_metrics_metrics_proto_goTypes,
		DependencyIndexes: file_alameda_api_v1alpha1_datahub_metrics_metrics_proto_depIdxs,
		MessageInfos:      file_alameda_api_v1alpha1_datahub_metrics_metrics_proto_msgTypes,
	}.Build()
	File_alameda_api_v1alpha1_datahub_metrics_metrics_proto = out.File
	file_alameda_api_v1alpha1_datahub_metrics_metrics_proto_rawDesc = nil
	file_alameda_api_v1alpha1_datahub_metrics_metrics_proto_goTypes = nil
	file_alameda_api_v1alpha1_datahub_metrics_metrics_proto_depIdxs = nil
}
