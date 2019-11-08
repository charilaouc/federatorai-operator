// Code generated by protoc-gen-go. DO NOT EDIT.
// source: alameda_api/v1alpha1/datahub/recommendations/services.proto

package recommendations

import (
	fmt "fmt"
	common "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/common"
	resources "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/resources"
	proto "github.com/golang/protobuf/proto"
	status "google.golang.org/genproto/googleapis/rpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Represents a request for creating a pod's recommendation
type CreatePodRecommendationsRequest struct {
	PodRecommendations   []*PodRecommendation `protobuf:"bytes,1,rep,name=pod_recommendations,json=podRecommendations,proto3" json:"pod_recommendations,omitempty"`
	Granularity          int64                `protobuf:"varint,2,opt,name=granularity,proto3" json:"granularity,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *CreatePodRecommendationsRequest) Reset()         { *m = CreatePodRecommendationsRequest{} }
func (m *CreatePodRecommendationsRequest) String() string { return proto.CompactTextString(m) }
func (*CreatePodRecommendationsRequest) ProtoMessage()    {}
func (*CreatePodRecommendationsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_603987ba119fba8b, []int{0}
}

func (m *CreatePodRecommendationsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreatePodRecommendationsRequest.Unmarshal(m, b)
}
func (m *CreatePodRecommendationsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreatePodRecommendationsRequest.Marshal(b, m, deterministic)
}
func (m *CreatePodRecommendationsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePodRecommendationsRequest.Merge(m, src)
}
func (m *CreatePodRecommendationsRequest) XXX_Size() int {
	return xxx_messageInfo_CreatePodRecommendationsRequest.Size(m)
}
func (m *CreatePodRecommendationsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePodRecommendationsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePodRecommendationsRequest proto.InternalMessageInfo

func (m *CreatePodRecommendationsRequest) GetPodRecommendations() []*PodRecommendation {
	if m != nil {
		return m.PodRecommendations
	}
	return nil
}

func (m *CreatePodRecommendationsRequest) GetGranularity() int64 {
	if m != nil {
		return m.Granularity
	}
	return 0
}

// Represents a request for creating a controller's recommendation
type CreateControllerRecommendationsRequest struct {
	ControllerRecommendations []*ControllerRecommendation `protobuf:"bytes,1,rep,name=controller_recommendations,json=controllerRecommendations,proto3" json:"controller_recommendations,omitempty"`
	XXX_NoUnkeyedLiteral      struct{}                    `json:"-"`
	XXX_unrecognized          []byte                      `json:"-"`
	XXX_sizecache             int32                       `json:"-"`
}

func (m *CreateControllerRecommendationsRequest) Reset() {
	*m = CreateControllerRecommendationsRequest{}
}
func (m *CreateControllerRecommendationsRequest) String() string { return proto.CompactTextString(m) }
func (*CreateControllerRecommendationsRequest) ProtoMessage()    {}
func (*CreateControllerRecommendationsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_603987ba119fba8b, []int{1}
}

func (m *CreateControllerRecommendationsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateControllerRecommendationsRequest.Unmarshal(m, b)
}
func (m *CreateControllerRecommendationsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateControllerRecommendationsRequest.Marshal(b, m, deterministic)
}
func (m *CreateControllerRecommendationsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateControllerRecommendationsRequest.Merge(m, src)
}
func (m *CreateControllerRecommendationsRequest) XXX_Size() int {
	return xxx_messageInfo_CreateControllerRecommendationsRequest.Size(m)
}
func (m *CreateControllerRecommendationsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateControllerRecommendationsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateControllerRecommendationsRequest proto.InternalMessageInfo

func (m *CreateControllerRecommendationsRequest) GetControllerRecommendations() []*ControllerRecommendation {
	if m != nil {
		return m.ControllerRecommendations
	}
	return nil
}

// Represents a request for listing recommendations of pods
type ListPodRecommendationsRequest struct {
	QueryCondition       *common.QueryCondition    `protobuf:"bytes,1,opt,name=query_condition,json=queryCondition,proto3" json:"query_condition,omitempty"`
	NamespacedName       *resources.NamespacedName `protobuf:"bytes,2,opt,name=namespaced_name,json=namespacedName,proto3" json:"namespaced_name,omitempty"`
	Kind                 resources.Kind            `protobuf:"varint,3,opt,name=kind,proto3,enum=containersai.alameda.v1alpha1.datahub.resources.Kind" json:"kind,omitempty"`
	Granularity          int64                     `protobuf:"varint,4,opt,name=granularity,proto3" json:"granularity,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *ListPodRecommendationsRequest) Reset()         { *m = ListPodRecommendationsRequest{} }
func (m *ListPodRecommendationsRequest) String() string { return proto.CompactTextString(m) }
func (*ListPodRecommendationsRequest) ProtoMessage()    {}
func (*ListPodRecommendationsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_603987ba119fba8b, []int{2}
}

func (m *ListPodRecommendationsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPodRecommendationsRequest.Unmarshal(m, b)
}
func (m *ListPodRecommendationsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPodRecommendationsRequest.Marshal(b, m, deterministic)
}
func (m *ListPodRecommendationsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPodRecommendationsRequest.Merge(m, src)
}
func (m *ListPodRecommendationsRequest) XXX_Size() int {
	return xxx_messageInfo_ListPodRecommendationsRequest.Size(m)
}
func (m *ListPodRecommendationsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPodRecommendationsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListPodRecommendationsRequest proto.InternalMessageInfo

func (m *ListPodRecommendationsRequest) GetQueryCondition() *common.QueryCondition {
	if m != nil {
		return m.QueryCondition
	}
	return nil
}

func (m *ListPodRecommendationsRequest) GetNamespacedName() *resources.NamespacedName {
	if m != nil {
		return m.NamespacedName
	}
	return nil
}

func (m *ListPodRecommendationsRequest) GetKind() resources.Kind {
	if m != nil {
		return m.Kind
	}
	return resources.Kind_POD
}

func (m *ListPodRecommendationsRequest) GetGranularity() int64 {
	if m != nil {
		return m.Granularity
	}
	return 0
}

// Represents a response for listing pod recommendations request
type ListPodRecommendationsResponse struct {
	Status               *status.Status       `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	PodRecommendations   []*PodRecommendation `protobuf:"bytes,2,rep,name=pod_recommendations,json=podRecommendations,proto3" json:"pod_recommendations,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ListPodRecommendationsResponse) Reset()         { *m = ListPodRecommendationsResponse{} }
func (m *ListPodRecommendationsResponse) String() string { return proto.CompactTextString(m) }
func (*ListPodRecommendationsResponse) ProtoMessage()    {}
func (*ListPodRecommendationsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_603987ba119fba8b, []int{3}
}

func (m *ListPodRecommendationsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPodRecommendationsResponse.Unmarshal(m, b)
}
func (m *ListPodRecommendationsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPodRecommendationsResponse.Marshal(b, m, deterministic)
}
func (m *ListPodRecommendationsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPodRecommendationsResponse.Merge(m, src)
}
func (m *ListPodRecommendationsResponse) XXX_Size() int {
	return xxx_messageInfo_ListPodRecommendationsResponse.Size(m)
}
func (m *ListPodRecommendationsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPodRecommendationsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListPodRecommendationsResponse proto.InternalMessageInfo

func (m *ListPodRecommendationsResponse) GetStatus() *status.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ListPodRecommendationsResponse) GetPodRecommendations() []*PodRecommendation {
	if m != nil {
		return m.PodRecommendations
	}
	return nil
}

// Represents a request for listing recommendations of controllers
type ListControllerRecommendationsRequest struct {
	QueryCondition       *common.QueryCondition    `protobuf:"bytes,1,opt,name=query_condition,json=queryCondition,proto3" json:"query_condition,omitempty"`
	NamespacedName       *resources.NamespacedName `protobuf:"bytes,2,opt,name=namespaced_name,json=namespacedName,proto3" json:"namespaced_name,omitempty"`
	RecommendedType      ControllerRecommendedType `protobuf:"varint,3,opt,name=recommended_type,json=recommendedType,proto3,enum=containersai.alameda.v1alpha1.datahub.recommendations.ControllerRecommendedType" json:"recommended_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *ListControllerRecommendationsRequest) Reset()         { *m = ListControllerRecommendationsRequest{} }
func (m *ListControllerRecommendationsRequest) String() string { return proto.CompactTextString(m) }
func (*ListControllerRecommendationsRequest) ProtoMessage()    {}
func (*ListControllerRecommendationsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_603987ba119fba8b, []int{4}
}

func (m *ListControllerRecommendationsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListControllerRecommendationsRequest.Unmarshal(m, b)
}
func (m *ListControllerRecommendationsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListControllerRecommendationsRequest.Marshal(b, m, deterministic)
}
func (m *ListControllerRecommendationsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListControllerRecommendationsRequest.Merge(m, src)
}
func (m *ListControllerRecommendationsRequest) XXX_Size() int {
	return xxx_messageInfo_ListControllerRecommendationsRequest.Size(m)
}
func (m *ListControllerRecommendationsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListControllerRecommendationsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListControllerRecommendationsRequest proto.InternalMessageInfo

func (m *ListControllerRecommendationsRequest) GetQueryCondition() *common.QueryCondition {
	if m != nil {
		return m.QueryCondition
	}
	return nil
}

func (m *ListControllerRecommendationsRequest) GetNamespacedName() *resources.NamespacedName {
	if m != nil {
		return m.NamespacedName
	}
	return nil
}

func (m *ListControllerRecommendationsRequest) GetRecommendedType() ControllerRecommendedType {
	if m != nil {
		return m.RecommendedType
	}
	return ControllerRecommendedType_CRT_Undefined
}

// Represents a response for listing controller recommendations request
type ListControllerRecommendationsResponse struct {
	Status                    *status.Status              `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	ControllerRecommendations []*ControllerRecommendation `protobuf:"bytes,2,rep,name=controller_recommendations,json=controllerRecommendations,proto3" json:"controller_recommendations,omitempty"`
	XXX_NoUnkeyedLiteral      struct{}                    `json:"-"`
	XXX_unrecognized          []byte                      `json:"-"`
	XXX_sizecache             int32                       `json:"-"`
}

func (m *ListControllerRecommendationsResponse) Reset()         { *m = ListControllerRecommendationsResponse{} }
func (m *ListControllerRecommendationsResponse) String() string { return proto.CompactTextString(m) }
func (*ListControllerRecommendationsResponse) ProtoMessage()    {}
func (*ListControllerRecommendationsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_603987ba119fba8b, []int{5}
}

func (m *ListControllerRecommendationsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListControllerRecommendationsResponse.Unmarshal(m, b)
}
func (m *ListControllerRecommendationsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListControllerRecommendationsResponse.Marshal(b, m, deterministic)
}
func (m *ListControllerRecommendationsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListControllerRecommendationsResponse.Merge(m, src)
}
func (m *ListControllerRecommendationsResponse) XXX_Size() int {
	return xxx_messageInfo_ListControllerRecommendationsResponse.Size(m)
}
func (m *ListControllerRecommendationsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListControllerRecommendationsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListControllerRecommendationsResponse proto.InternalMessageInfo

func (m *ListControllerRecommendationsResponse) GetStatus() *status.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ListControllerRecommendationsResponse) GetControllerRecommendations() []*ControllerRecommendation {
	if m != nil {
		return m.ControllerRecommendations
	}
	return nil
}

func init() {
	proto.RegisterType((*CreatePodRecommendationsRequest)(nil), "containersai.alameda.v1alpha1.datahub.recommendations.CreatePodRecommendationsRequest")
	proto.RegisterType((*CreateControllerRecommendationsRequest)(nil), "containersai.alameda.v1alpha1.datahub.recommendations.CreateControllerRecommendationsRequest")
	proto.RegisterType((*ListPodRecommendationsRequest)(nil), "containersai.alameda.v1alpha1.datahub.recommendations.ListPodRecommendationsRequest")
	proto.RegisterType((*ListPodRecommendationsResponse)(nil), "containersai.alameda.v1alpha1.datahub.recommendations.ListPodRecommendationsResponse")
	proto.RegisterType((*ListControllerRecommendationsRequest)(nil), "containersai.alameda.v1alpha1.datahub.recommendations.ListControllerRecommendationsRequest")
	proto.RegisterType((*ListControllerRecommendationsResponse)(nil), "containersai.alameda.v1alpha1.datahub.recommendations.ListControllerRecommendationsResponse")
}

func init() {
	proto.RegisterFile("alameda_api/v1alpha1/datahub/recommendations/services.proto", fileDescriptor_603987ba119fba8b)
}

var fileDescriptor_603987ba119fba8b = []byte{
	// 530 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x55, 0xcd, 0x8a, 0x53, 0x31,
	0x14, 0xe6, 0xb6, 0xc3, 0x2c, 0x52, 0x68, 0xe5, 0xba, 0xb0, 0x16, 0xd4, 0x52, 0x54, 0x8a, 0x60,
	0x42, 0x2b, 0x05, 0x41, 0x41, 0x98, 0x6e, 0x1c, 0x46, 0x74, 0x8c, 0xae, 0xdc, 0x94, 0x4c, 0x72,
	0x68, 0x83, 0xb7, 0x49, 0x9a, 0xe4, 0x0e, 0x14, 0xdf, 0xc1, 0xf7, 0x71, 0xe3, 0xd6, 0xa7, 0x70,
	0xe7, 0x23, 0xf8, 0x00, 0x92, 0x36, 0xad, 0xd3, 0x96, 0x6b, 0x7f, 0x50, 0x17, 0xb3, 0xcb, 0xef,
	0xf7, 0x9d, 0xf3, 0xe5, 0x3b, 0x39, 0xe8, 0x19, 0xcb, 0xd8, 0x18, 0x04, 0x1b, 0x30, 0x23, 0xc9,
	0x65, 0x87, 0x65, 0x66, 0xc4, 0x3a, 0x44, 0x30, 0xcf, 0x46, 0xf9, 0x05, 0xb1, 0xc0, 0xf5, 0x78,
	0x0c, 0x4a, 0x30, 0x2f, 0xb5, 0x72, 0xc4, 0x81, 0xbd, 0x94, 0x1c, 0x1c, 0x36, 0x56, 0x7b, 0x9d,
	0xf6, 0xb8, 0x56, 0x9e, 0x49, 0x05, 0xd6, 0x31, 0x89, 0x23, 0x12, 0x5e, 0xa0, 0xe0, 0x88, 0x82,
	0xd7, 0x50, 0x1a, 0x9d, 0x3f, 0x72, 0x86, 0xb3, 0x5a, 0x91, 0x49, 0x0e, 0x56, 0x2e, 0x98, 0x1a,
	0x27, 0x7b, 0x85, 0xb9, 0x36, 0x8f, 0x18, 0x4f, 0xf7, 0xc2, 0xf0, 0x53, 0xb3, 0x64, 0xef, 0x6d,
	0xb9, 0xe9, 0x74, 0x6e, 0x39, 0x38, 0x32, 0x06, 0xcf, 0xc2, 0x6a, 0xbc, 0xd6, 0xdd, 0xf1, 0xda,
	0x55, 0xaa, 0x5b, 0x43, 0xad, 0x87, 0x19, 0x10, 0x6b, 0x38, 0x71, 0x9e, 0xf9, 0x3c, 0x6e, 0xb4,
	0xbe, 0x26, 0xe8, 0x5e, 0xdf, 0x02, 0xf3, 0x70, 0xae, 0x05, 0x5d, 0x0d, 0x96, 0xc2, 0x24, 0x07,
	0xe7, 0xd3, 0x29, 0xba, 0x69, 0xb4, 0x18, 0xac, 0xa5, 0x52, 0x4f, 0x9a, 0xe5, 0x76, 0xa5, 0xfb,
	0x12, 0x1f, 0xf4, 0x5a, 0x78, 0x83, 0x8e, 0xa6, 0x66, 0x23, 0x82, 0xb4, 0x89, 0x2a, 0x43, 0xcb,
	0x54, 0x9e, 0x31, 0x2b, 0xfd, 0xb4, 0x5e, 0x6a, 0x26, 0xed, 0x32, 0xbd, 0xba, 0xd4, 0xfa, 0x92,
	0xa0, 0x87, 0xf3, 0x04, 0xfa, 0x5a, 0x79, 0xab, 0xb3, 0x0c, 0x6c, 0x41, 0x1e, 0x9f, 0x13, 0xd4,
	0xe0, 0xcb, 0x43, 0x05, 0xf9, 0xbc, 0x39, 0x30, 0x9f, 0x22, 0x76, 0x7a, 0x9b, 0x17, 0xc5, 0xd5,
	0xfa, 0x5e, 0x42, 0x77, 0x5e, 0x49, 0xe7, 0x8b, 0xa5, 0x07, 0x54, 0x0b, 0x8e, 0x9d, 0x0e, 0xb8,
	0x56, 0x42, 0x86, 0xad, 0x7a, 0xd2, 0x4c, 0xda, 0x95, 0xee, 0xf3, 0x1d, 0xc3, 0x9c, 0xdb, 0x1e,
	0xbf, 0x0d, 0x20, 0xfd, 0x05, 0x06, 0xad, 0x4e, 0x56, 0xe6, 0xe9, 0x08, 0xd5, 0x14, 0x1b, 0x83,
	0x33, 0x8c, 0x83, 0x18, 0x84, 0xe1, 0x4c, 0xea, 0x4a, 0xf7, 0xc5, 0xce, 0x6a, 0x44, 0xd7, 0xe1,
	0xd7, 0x4b, 0x9c, 0x30, 0xa2, 0x55, 0xb5, 0x32, 0x4f, 0x4f, 0xd1, 0xd1, 0x47, 0xa9, 0x44, 0xbd,
	0xdc, 0x4c, 0xda, 0xd5, 0x6e, 0x6f, 0x6f, 0xf8, 0x33, 0xa9, 0x04, 0x9d, 0x41, 0xac, 0x7b, 0xe3,
	0x68, 0xd3, 0x1b, 0xdf, 0x12, 0x74, 0xb7, 0x48, 0x5f, 0x67, 0xb4, 0x72, 0x90, 0x3e, 0x42, 0xc7,
	0xf3, 0x7a, 0x88, 0xba, 0xa6, 0x78, 0x5e, 0x29, 0xd8, 0x1a, 0x8e, 0xdf, 0xcd, 0x76, 0x68, 0x3c,
	0x51, 0x54, 0x07, 0xa5, 0x7f, 0x5f, 0x07, 0xad, 0x9f, 0x25, 0x74, 0x3f, 0x64, 0xb2, 0xd5, 0xe3,
	0xd7, 0xce, 0x30, 0x9f, 0xd0, 0x8d, 0xa5, 0x64, 0x20, 0x06, 0xe1, 0x53, 0x8b, 0xe6, 0x39, 0xff,
	0x7b, 0x95, 0x0a, 0xe2, 0xfd, 0xd4, 0x00, 0xad, 0xd9, 0xd5, 0x85, 0xd6, 0x8f, 0x04, 0x3d, 0xd8,
	0x22, 0xfb, 0x01, 0x3e, 0xda, 0xf2, 0x0f, 0x95, 0xfe, 0xf7, 0x3f, 0x74, 0x72, 0xf6, 0xe1, 0x74,
	0x28, 0x7d, 0x7c, 0x7d, 0xf2, 0x9b, 0xf7, 0x31, 0x93, 0x24, 0x34, 0x99, 0x7d, 0x3a, 0xdc, 0xc5,
	0xf1, 0xac, 0xb1, 0x3c, 0xf9, 0x15, 0x00, 0x00, 0xff, 0xff, 0x91, 0xbd, 0xf9, 0x55, 0x03, 0x08,
	0x00, 0x00,
}