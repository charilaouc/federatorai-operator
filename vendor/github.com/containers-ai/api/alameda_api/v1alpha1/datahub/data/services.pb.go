// This file has messages related to read & write data

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v4.0.0
// source: alameda_api/v1alpha1/datahub/data/services.proto

package data

import (
	schemas "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/schemas"
	proto "github.com/golang/protobuf/proto"
	status "google.golang.org/genproto/googleapis/rpc/status"
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

type WriteDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SchemaMeta *schemas.SchemaMeta `protobuf:"bytes,1,opt,name=schema_meta,json=schemaMeta,proto3" json:"schema_meta,omitempty"`
	WriteData  []*WriteData        `protobuf:"bytes,2,rep,name=write_data,json=writeData,proto3" json:"write_data,omitempty"`
}

func (x *WriteDataRequest) Reset() {
	*x = WriteDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alameda_api_v1alpha1_datahub_data_services_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteDataRequest) ProtoMessage() {}

func (x *WriteDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_alameda_api_v1alpha1_datahub_data_services_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteDataRequest.ProtoReflect.Descriptor instead.
func (*WriteDataRequest) Descriptor() ([]byte, []int) {
	return file_alameda_api_v1alpha1_datahub_data_services_proto_rawDescGZIP(), []int{0}
}

func (x *WriteDataRequest) GetSchemaMeta() *schemas.SchemaMeta {
	if x != nil {
		return x.SchemaMeta
	}
	return nil
}

func (x *WriteDataRequest) GetWriteData() []*WriteData {
	if x != nil {
		return x.WriteData
	}
	return nil
}

type ReadDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SchemaMeta *schemas.SchemaMeta `protobuf:"bytes,1,opt,name=schema_meta,json=schemaMeta,proto3" json:"schema_meta,omitempty"`
	ReadData   []*ReadData         `protobuf:"bytes,2,rep,name=read_data,json=readData,proto3" json:"read_data,omitempty"`
}

func (x *ReadDataRequest) Reset() {
	*x = ReadDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alameda_api_v1alpha1_datahub_data_services_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadDataRequest) ProtoMessage() {}

func (x *ReadDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_alameda_api_v1alpha1_datahub_data_services_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadDataRequest.ProtoReflect.Descriptor instead.
func (*ReadDataRequest) Descriptor() ([]byte, []int) {
	return file_alameda_api_v1alpha1_datahub_data_services_proto_rawDescGZIP(), []int{1}
}

func (x *ReadDataRequest) GetSchemaMeta() *schemas.SchemaMeta {
	if x != nil {
		return x.SchemaMeta
	}
	return nil
}

func (x *ReadDataRequest) GetReadData() []*ReadData {
	if x != nil {
		return x.ReadData
	}
	return nil
}

type ReadDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *status.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Data   *Data          `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ReadDataResponse) Reset() {
	*x = ReadDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alameda_api_v1alpha1_datahub_data_services_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadDataResponse) ProtoMessage() {}

func (x *ReadDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_alameda_api_v1alpha1_datahub_data_services_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadDataResponse.ProtoReflect.Descriptor instead.
func (*ReadDataResponse) Descriptor() ([]byte, []int) {
	return file_alameda_api_v1alpha1_datahub_data_services_proto_rawDescGZIP(), []int{2}
}

func (x *ReadDataResponse) GetStatus() *status.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *ReadDataResponse) GetData() *Data {
	if x != nil {
		return x.Data
	}
	return nil
}

type DeleteDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SchemaMeta *schemas.SchemaMeta `protobuf:"bytes,1,opt,name=schema_meta,json=schemaMeta,proto3" json:"schema_meta,omitempty"`
	DeleteData []*DeleteData       `protobuf:"bytes,2,rep,name=delete_data,json=deleteData,proto3" json:"delete_data,omitempty"`
}

func (x *DeleteDataRequest) Reset() {
	*x = DeleteDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alameda_api_v1alpha1_datahub_data_services_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDataRequest) ProtoMessage() {}

func (x *DeleteDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_alameda_api_v1alpha1_datahub_data_services_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDataRequest.ProtoReflect.Descriptor instead.
func (*DeleteDataRequest) Descriptor() ([]byte, []int) {
	return file_alameda_api_v1alpha1_datahub_data_services_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteDataRequest) GetSchemaMeta() *schemas.SchemaMeta {
	if x != nil {
		return x.SchemaMeta
	}
	return nil
}

func (x *DeleteDataRequest) GetDeleteData() []*DeleteData {
	if x != nil {
		return x.DeleteData
	}
	return nil
}

var File_alameda_api_v1alpha1_datahub_data_services_proto protoreflect.FileDescriptor

var file_alameda_api_v1alpha1_datahub_data_services_proto_rawDesc = []byte{
	0x0a, 0x30, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2f, 0x64,
	0x61, 0x74, 0x61, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x2a, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x61, 0x69,
	0x2e, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x2c,
	0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2f, 0x64, 0x61, 0x74,
	0x61, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2d, 0x61, 0x6c,
	0x61, 0x6d, 0x65, 0x64, 0x61, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x30, 0x61, 0x6c, 0x61,
	0x6d, 0x65, 0x64, 0x61, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc4, 0x01, 0x0a, 0x10, 0x57, 0x72, 0x69, 0x74, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x5a, 0x0a, 0x0b, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x39, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x61, 0x69, 0x2e,
	0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73,
	0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x0a, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x54, 0x0a, 0x0a, 0x77, 0x72, 0x69, 0x74, 0x65,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x61, 0x69, 0x2e, 0x61, 0x6c, 0x61, 0x6d, 0x65,
	0x64, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x68, 0x75, 0x62, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x09, 0x77, 0x72, 0x69, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x22, 0xc0, 0x01,
	0x0a, 0x0f, 0x52, 0x65, 0x61, 0x64, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x5a, 0x0a, 0x0b, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x6d, 0x65, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e,
	0x65, 0x72, 0x73, 0x61, 0x69, 0x2e, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x4d, 0x65, 0x74,
	0x61, 0x52, 0x0a, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x51, 0x0a,
	0x09, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x34, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x61, 0x69, 0x2e,
	0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x52, 0x65,
	0x61, 0x64, 0x44, 0x61, 0x74, 0x61, 0x52, 0x08, 0x72, 0x65, 0x61, 0x64, 0x44, 0x61, 0x74, 0x61,
	0x22, 0x84, 0x01, 0x0a, 0x10, 0x52, 0x65, 0x61, 0x64, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72,
	0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x44, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x30, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x61, 0x69, 0x2e, 0x61,
	0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xc8, 0x01, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x5a, 0x0a,
	0x0b, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x39, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x61,
	0x69, 0x2e, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x73, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x0a, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x57, 0x0a, 0x0b, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x36,
	0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x61, 0x69, 0x2e, 0x61, 0x6c,
	0x61, 0x6d, 0x65, 0x64, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x42, 0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x2d, 0x61, 0x69, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2f,
	0x64, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_alameda_api_v1alpha1_datahub_data_services_proto_rawDescOnce sync.Once
	file_alameda_api_v1alpha1_datahub_data_services_proto_rawDescData = file_alameda_api_v1alpha1_datahub_data_services_proto_rawDesc
)

func file_alameda_api_v1alpha1_datahub_data_services_proto_rawDescGZIP() []byte {
	file_alameda_api_v1alpha1_datahub_data_services_proto_rawDescOnce.Do(func() {
		file_alameda_api_v1alpha1_datahub_data_services_proto_rawDescData = protoimpl.X.CompressGZIP(file_alameda_api_v1alpha1_datahub_data_services_proto_rawDescData)
	})
	return file_alameda_api_v1alpha1_datahub_data_services_proto_rawDescData
}

var file_alameda_api_v1alpha1_datahub_data_services_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_alameda_api_v1alpha1_datahub_data_services_proto_goTypes = []interface{}{
	(*WriteDataRequest)(nil),   // 0: containersai.alameda.v1alpha1.datahub.data.WriteDataRequest
	(*ReadDataRequest)(nil),    // 1: containersai.alameda.v1alpha1.datahub.data.ReadDataRequest
	(*ReadDataResponse)(nil),   // 2: containersai.alameda.v1alpha1.datahub.data.ReadDataResponse
	(*DeleteDataRequest)(nil),  // 3: containersai.alameda.v1alpha1.datahub.data.DeleteDataRequest
	(*schemas.SchemaMeta)(nil), // 4: containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta
	(*WriteData)(nil),          // 5: containersai.alameda.v1alpha1.datahub.data.WriteData
	(*ReadData)(nil),           // 6: containersai.alameda.v1alpha1.datahub.data.ReadData
	(*status.Status)(nil),      // 7: google.rpc.Status
	(*Data)(nil),               // 8: containersai.alameda.v1alpha1.datahub.data.Data
	(*DeleteData)(nil),         // 9: containersai.alameda.v1alpha1.datahub.data.DeleteData
}
var file_alameda_api_v1alpha1_datahub_data_services_proto_depIdxs = []int32{
	4, // 0: containersai.alameda.v1alpha1.datahub.data.WriteDataRequest.schema_meta:type_name -> containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta
	5, // 1: containersai.alameda.v1alpha1.datahub.data.WriteDataRequest.write_data:type_name -> containersai.alameda.v1alpha1.datahub.data.WriteData
	4, // 2: containersai.alameda.v1alpha1.datahub.data.ReadDataRequest.schema_meta:type_name -> containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta
	6, // 3: containersai.alameda.v1alpha1.datahub.data.ReadDataRequest.read_data:type_name -> containersai.alameda.v1alpha1.datahub.data.ReadData
	7, // 4: containersai.alameda.v1alpha1.datahub.data.ReadDataResponse.status:type_name -> google.rpc.Status
	8, // 5: containersai.alameda.v1alpha1.datahub.data.ReadDataResponse.data:type_name -> containersai.alameda.v1alpha1.datahub.data.Data
	4, // 6: containersai.alameda.v1alpha1.datahub.data.DeleteDataRequest.schema_meta:type_name -> containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta
	9, // 7: containersai.alameda.v1alpha1.datahub.data.DeleteDataRequest.delete_data:type_name -> containersai.alameda.v1alpha1.datahub.data.DeleteData
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_alameda_api_v1alpha1_datahub_data_services_proto_init() }
func file_alameda_api_v1alpha1_datahub_data_services_proto_init() {
	if File_alameda_api_v1alpha1_datahub_data_services_proto != nil {
		return
	}
	file_alameda_api_v1alpha1_datahub_data_data_proto_init()
	file_alameda_api_v1alpha1_datahub_data_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_alameda_api_v1alpha1_datahub_data_services_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteDataRequest); i {
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
		file_alameda_api_v1alpha1_datahub_data_services_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadDataRequest); i {
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
		file_alameda_api_v1alpha1_datahub_data_services_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadDataResponse); i {
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
		file_alameda_api_v1alpha1_datahub_data_services_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDataRequest); i {
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
			RawDescriptor: file_alameda_api_v1alpha1_datahub_data_services_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_alameda_api_v1alpha1_datahub_data_services_proto_goTypes,
		DependencyIndexes: file_alameda_api_v1alpha1_datahub_data_services_proto_depIdxs,
		MessageInfos:      file_alameda_api_v1alpha1_datahub_data_services_proto_msgTypes,
	}.Build()
	File_alameda_api_v1alpha1_datahub_data_services_proto = out.File
	file_alameda_api_v1alpha1_datahub_data_services_proto_rawDesc = nil
	file_alameda_api_v1alpha1_datahub_data_services_proto_goTypes = nil
	file_alameda_api_v1alpha1_datahub_data_services_proto_depIdxs = nil
}
