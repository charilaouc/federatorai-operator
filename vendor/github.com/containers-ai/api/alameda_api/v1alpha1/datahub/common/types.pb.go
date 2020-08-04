// This file has messages related general definitions

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v4.0.0
// source: alameda_api/v1alpha1/datahub/common/types.proto

package common

import (
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

type ResourceBoundary int32

const (
	ResourceBoundary_RESOURCE_BOUNDARY_UNDEFINED ResourceBoundary = 0
	ResourceBoundary_RESOURCE_RAW                ResourceBoundary = 1
	ResourceBoundary_RESOURCE_UPPER_BOUND        ResourceBoundary = 2
	ResourceBoundary_RESOURCE_LOWER_BOUND        ResourceBoundary = 3
)

// Enum value maps for ResourceBoundary.
var (
	ResourceBoundary_name = map[int32]string{
		0: "RESOURCE_BOUNDARY_UNDEFINED",
		1: "RESOURCE_RAW",
		2: "RESOURCE_UPPER_BOUND",
		3: "RESOURCE_LOWER_BOUND",
	}
	ResourceBoundary_value = map[string]int32{
		"RESOURCE_BOUNDARY_UNDEFINED": 0,
		"RESOURCE_RAW":                1,
		"RESOURCE_UPPER_BOUND":        2,
		"RESOURCE_LOWER_BOUND":        3,
	}
)

func (x ResourceBoundary) Enum() *ResourceBoundary {
	p := new(ResourceBoundary)
	*p = x
	return p
}

func (x ResourceBoundary) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResourceBoundary) Descriptor() protoreflect.EnumDescriptor {
	return file_alameda_api_v1alpha1_datahub_common_types_proto_enumTypes[0].Descriptor()
}

func (ResourceBoundary) Type() protoreflect.EnumType {
	return &file_alameda_api_v1alpha1_datahub_common_types_proto_enumTypes[0]
}

func (x ResourceBoundary) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResourceBoundary.Descriptor instead.
func (ResourceBoundary) EnumDescriptor() ([]byte, []int) {
	return file_alameda_api_v1alpha1_datahub_common_types_proto_rawDescGZIP(), []int{0}
}

type ResourceQuota int32

const (
	ResourceQuota_RESOURCE_QUOTA_UNDEFINED ResourceQuota = 0
	ResourceQuota_RESOURCE_LIMIT           ResourceQuota = 1
	ResourceQuota_RESOURCE_REQUEST         ResourceQuota = 2
	ResourceQuota_RESOURCE_INITIAL_LIMIT   ResourceQuota = 3
	ResourceQuota_RESOURCE_INITIAL_REQUEST ResourceQuota = 4
)

// Enum value maps for ResourceQuota.
var (
	ResourceQuota_name = map[int32]string{
		0: "RESOURCE_QUOTA_UNDEFINED",
		1: "RESOURCE_LIMIT",
		2: "RESOURCE_REQUEST",
		3: "RESOURCE_INITIAL_LIMIT",
		4: "RESOURCE_INITIAL_REQUEST",
	}
	ResourceQuota_value = map[string]int32{
		"RESOURCE_QUOTA_UNDEFINED": 0,
		"RESOURCE_LIMIT":           1,
		"RESOURCE_REQUEST":         2,
		"RESOURCE_INITIAL_LIMIT":   3,
		"RESOURCE_INITIAL_REQUEST": 4,
	}
)

func (x ResourceQuota) Enum() *ResourceQuota {
	p := new(ResourceQuota)
	*p = x
	return p
}

func (x ResourceQuota) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResourceQuota) Descriptor() protoreflect.EnumDescriptor {
	return file_alameda_api_v1alpha1_datahub_common_types_proto_enumTypes[1].Descriptor()
}

func (ResourceQuota) Type() protoreflect.EnumType {
	return &file_alameda_api_v1alpha1_datahub_common_types_proto_enumTypes[1]
}

func (x ResourceQuota) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResourceQuota.Descriptor instead.
func (ResourceQuota) EnumDescriptor() ([]byte, []int) {
	return file_alameda_api_v1alpha1_datahub_common_types_proto_rawDescGZIP(), []int{1}
}

type DataType int32

const (
	DataType_DATATYPE_UNDEFINED DataType = 0
	DataType_DATATYPE_BOOL      DataType = 1
	DataType_DATATYPE_INT       DataType = 2
	DataType_DATATYPE_INT8      DataType = 3
	DataType_DATATYPE_INT16     DataType = 4
	DataType_DATATYPE_INT32     DataType = 5
	DataType_DATATYPE_INT64     DataType = 6
	DataType_DATATYPE_UINT      DataType = 7
	DataType_DATATYPE_UINT8     DataType = 8
	DataType_DATATYPE_UINT16    DataType = 9
	DataType_DATATYPE_UINT32    DataType = 10
	DataType_DATATYPE_UTIN64    DataType = 11
	DataType_DATATYPE_FLOAT32   DataType = 12
	DataType_DATATYPE_FLOAT64   DataType = 13
	DataType_DATATYPE_STRING    DataType = 14
)

// Enum value maps for DataType.
var (
	DataType_name = map[int32]string{
		0:  "DATATYPE_UNDEFINED",
		1:  "DATATYPE_BOOL",
		2:  "DATATYPE_INT",
		3:  "DATATYPE_INT8",
		4:  "DATATYPE_INT16",
		5:  "DATATYPE_INT32",
		6:  "DATATYPE_INT64",
		7:  "DATATYPE_UINT",
		8:  "DATATYPE_UINT8",
		9:  "DATATYPE_UINT16",
		10: "DATATYPE_UINT32",
		11: "DATATYPE_UTIN64",
		12: "DATATYPE_FLOAT32",
		13: "DATATYPE_FLOAT64",
		14: "DATATYPE_STRING",
	}
	DataType_value = map[string]int32{
		"DATATYPE_UNDEFINED": 0,
		"DATATYPE_BOOL":      1,
		"DATATYPE_INT":       2,
		"DATATYPE_INT8":      3,
		"DATATYPE_INT16":     4,
		"DATATYPE_INT32":     5,
		"DATATYPE_INT64":     6,
		"DATATYPE_UINT":      7,
		"DATATYPE_UINT8":     8,
		"DATATYPE_UINT16":    9,
		"DATATYPE_UINT32":    10,
		"DATATYPE_UTIN64":    11,
		"DATATYPE_FLOAT32":   12,
		"DATATYPE_FLOAT64":   13,
		"DATATYPE_STRING":    14,
	}
)

func (x DataType) Enum() *DataType {
	p := new(DataType)
	*p = x
	return p
}

func (x DataType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DataType) Descriptor() protoreflect.EnumDescriptor {
	return file_alameda_api_v1alpha1_datahub_common_types_proto_enumTypes[2].Descriptor()
}

func (DataType) Type() protoreflect.EnumType {
	return &file_alameda_api_v1alpha1_datahub_common_types_proto_enumTypes[2]
}

func (x DataType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DataType.Descriptor instead.
func (DataType) EnumDescriptor() ([]byte, []int) {
	return file_alameda_api_v1alpha1_datahub_common_types_proto_rawDescGZIP(), []int{2}
}

type ColumnType int32

const (
	ColumnType_COLUMNTYPE_UDEFINED ColumnType = 0
	ColumnType_COLUMNTYPE_TAG      ColumnType = 1
	ColumnType_COLUMNTYPE_FIELD    ColumnType = 2
)

// Enum value maps for ColumnType.
var (
	ColumnType_name = map[int32]string{
		0: "COLUMNTYPE_UDEFINED",
		1: "COLUMNTYPE_TAG",
		2: "COLUMNTYPE_FIELD",
	}
	ColumnType_value = map[string]int32{
		"COLUMNTYPE_UDEFINED": 0,
		"COLUMNTYPE_TAG":      1,
		"COLUMNTYPE_FIELD":    2,
	}
)

func (x ColumnType) Enum() *ColumnType {
	p := new(ColumnType)
	*p = x
	return p
}

func (x ColumnType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ColumnType) Descriptor() protoreflect.EnumDescriptor {
	return file_alameda_api_v1alpha1_datahub_common_types_proto_enumTypes[3].Descriptor()
}

func (ColumnType) Type() protoreflect.EnumType {
	return &file_alameda_api_v1alpha1_datahub_common_types_proto_enumTypes[3]
}

func (x ColumnType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ColumnType.Descriptor instead.
func (ColumnType) EnumDescriptor() ([]byte, []int) {
	return file_alameda_api_v1alpha1_datahub_common_types_proto_rawDescGZIP(), []int{3}
}

type FunctionType int32

const (
	FunctionType_FUNCTIONTYPE_UNDEFINED FunctionType = 0
	// Aggregation function
	FunctionType_FUNCTIONTYPE_COUNT    FunctionType = 1 // Returns the number of non-null field values
	FunctionType_FUNCTIONTYPE_DISTINCT FunctionType = 2 // Returns the list of unique field values
	FunctionType_FUNCTIONTYPE_INTEGRAL FunctionType = 3 // Returns the area under the curve for subsequent field values
	FunctionType_FUNCTIONTYPE_MEAN     FunctionType = 4 // Returns the arithmetic mean (average) of field values.
	FunctionType_FUNCTIONTYPE_MEDIAN   FunctionType = 5 // Returns the middle value from a sorted list of field values.
	FunctionType_FUNCTIONTYPE_MODE     FunctionType = 6 // Returns the most frequent value in a list of field values
	FunctionType_FUNCTIONTYPE_SPREAD   FunctionType = 7 // Returns the difference between the minimum and maximum field values
	FunctionType_FUNCTIONTYPE_STDDEV   FunctionType = 8 // Returns the standard deviation of field values.
	FunctionType_FUNCTIONTYPE_SUM      FunctionType = 9 // Returns the sum of field values.
	// Selector function
	FunctionType_FUNCTIONTYPE_BOTTOM     FunctionType = 10 // Returns the smallest N field values.
	FunctionType_FUNCTIONTYPE_FIRST      FunctionType = 11 // Returns the field value with the oldest timestamp.
	FunctionType_FUNCTIONTYPE_LAST       FunctionType = 12 // Returns the field value with the most recent timestamp.
	FunctionType_FUNCTIONTYPE_MAX        FunctionType = 13 // Returns the greatest field value.
	FunctionType_FUNCTIONTYPE_MIN        FunctionType = 14 // Returns the lowest field value.
	FunctionType_FUNCTIONTYPE_PERCENTILE FunctionType = 15 // Returns the Nth percentile field value
	FunctionType_FUNCTIONTYPE_SAMPLE     FunctionType = 16 // Returns a random sample of N field values. SAMPLE() uses reservoir sampling to generate the random points
	FunctionType_FUNCTIONTYPE_TOP        FunctionType = 17 // Returns the greatest N field values.
	// Transformation function
	FunctionType_FUNCTIONTYPE_DERIVATIVE FunctionType = 18 // Returns the rate of change between subsequent field values
)

// Enum value maps for FunctionType.
var (
	FunctionType_name = map[int32]string{
		0:  "FUNCTIONTYPE_UNDEFINED",
		1:  "FUNCTIONTYPE_COUNT",
		2:  "FUNCTIONTYPE_DISTINCT",
		3:  "FUNCTIONTYPE_INTEGRAL",
		4:  "FUNCTIONTYPE_MEAN",
		5:  "FUNCTIONTYPE_MEDIAN",
		6:  "FUNCTIONTYPE_MODE",
		7:  "FUNCTIONTYPE_SPREAD",
		8:  "FUNCTIONTYPE_STDDEV",
		9:  "FUNCTIONTYPE_SUM",
		10: "FUNCTIONTYPE_BOTTOM",
		11: "FUNCTIONTYPE_FIRST",
		12: "FUNCTIONTYPE_LAST",
		13: "FUNCTIONTYPE_MAX",
		14: "FUNCTIONTYPE_MIN",
		15: "FUNCTIONTYPE_PERCENTILE",
		16: "FUNCTIONTYPE_SAMPLE",
		17: "FUNCTIONTYPE_TOP",
		18: "FUNCTIONTYPE_DERIVATIVE",
	}
	FunctionType_value = map[string]int32{
		"FUNCTIONTYPE_UNDEFINED":  0,
		"FUNCTIONTYPE_COUNT":      1,
		"FUNCTIONTYPE_DISTINCT":   2,
		"FUNCTIONTYPE_INTEGRAL":   3,
		"FUNCTIONTYPE_MEAN":       4,
		"FUNCTIONTYPE_MEDIAN":     5,
		"FUNCTIONTYPE_MODE":       6,
		"FUNCTIONTYPE_SPREAD":     7,
		"FUNCTIONTYPE_STDDEV":     8,
		"FUNCTIONTYPE_SUM":        9,
		"FUNCTIONTYPE_BOTTOM":     10,
		"FUNCTIONTYPE_FIRST":      11,
		"FUNCTIONTYPE_LAST":       12,
		"FUNCTIONTYPE_MAX":        13,
		"FUNCTIONTYPE_MIN":        14,
		"FUNCTIONTYPE_PERCENTILE": 15,
		"FUNCTIONTYPE_SAMPLE":     16,
		"FUNCTIONTYPE_TOP":        17,
		"FUNCTIONTYPE_DERIVATIVE": 18,
	}
)

func (x FunctionType) Enum() *FunctionType {
	p := new(FunctionType)
	*p = x
	return p
}

func (x FunctionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FunctionType) Descriptor() protoreflect.EnumDescriptor {
	return file_alameda_api_v1alpha1_datahub_common_types_proto_enumTypes[4].Descriptor()
}

func (FunctionType) Type() protoreflect.EnumType {
	return &file_alameda_api_v1alpha1_datahub_common_types_proto_enumTypes[4]
}

func (x FunctionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FunctionType.Descriptor instead.
func (FunctionType) EnumDescriptor() ([]byte, []int) {
	return file_alameda_api_v1alpha1_datahub_common_types_proto_rawDescGZIP(), []int{4}
}

var File_alameda_api_v1alpha1_datahub_common_types_proto protoreflect.FileDescriptor

var file_alameda_api_v1alpha1_datahub_common_types_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x2c, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x61, 0x69, 0x2e,
	0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2a,
	0x79, 0x0a, 0x10, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x6f, 0x75, 0x6e, 0x64,
	0x61, 0x72, 0x79, 0x12, 0x1f, 0x0a, 0x1b, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f,
	0x42, 0x4f, 0x55, 0x4e, 0x44, 0x41, 0x52, 0x59, 0x5f, 0x55, 0x4e, 0x44, 0x45, 0x46, 0x49, 0x4e,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45,
	0x5f, 0x52, 0x41, 0x57, 0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52,
	0x43, 0x45, 0x5f, 0x55, 0x50, 0x50, 0x45, 0x52, 0x5f, 0x42, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x02,
	0x12, 0x18, 0x0a, 0x14, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x4c, 0x4f, 0x57,
	0x45, 0x52, 0x5f, 0x42, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x03, 0x2a, 0x91, 0x01, 0x0a, 0x0d, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x18,
	0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x51, 0x55, 0x4f, 0x54, 0x41, 0x5f, 0x55,
	0x4e, 0x44, 0x45, 0x46, 0x49, 0x4e, 0x45, 0x44, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x52, 0x45,
	0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x4c, 0x49, 0x4d, 0x49, 0x54, 0x10, 0x01, 0x12, 0x14,
	0x0a, 0x10, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45,
	0x53, 0x54, 0x10, 0x02, 0x12, 0x1a, 0x0a, 0x16, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45,
	0x5f, 0x49, 0x4e, 0x49, 0x54, 0x49, 0x41, 0x4c, 0x5f, 0x4c, 0x49, 0x4d, 0x49, 0x54, 0x10, 0x03,
	0x12, 0x1c, 0x0a, 0x18, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x49, 0x4e, 0x49,
	0x54, 0x49, 0x41, 0x4c, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x04, 0x2a, 0xbd,
	0x02, 0x0a, 0x08, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x12, 0x44,
	0x41, 0x54, 0x41, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x44, 0x45, 0x46, 0x49, 0x4e, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x44, 0x41, 0x54, 0x41, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x42, 0x4f, 0x4f, 0x4c, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x44, 0x41, 0x54, 0x41, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x49, 0x4e, 0x54, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x44, 0x41, 0x54, 0x41,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4e, 0x54, 0x38, 0x10, 0x03, 0x12, 0x12, 0x0a, 0x0e, 0x44,
	0x41, 0x54, 0x41, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4e, 0x54, 0x31, 0x36, 0x10, 0x04, 0x12,
	0x12, 0x0a, 0x0e, 0x44, 0x41, 0x54, 0x41, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4e, 0x54, 0x33,
	0x32, 0x10, 0x05, 0x12, 0x12, 0x0a, 0x0e, 0x44, 0x41, 0x54, 0x41, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x49, 0x4e, 0x54, 0x36, 0x34, 0x10, 0x06, 0x12, 0x11, 0x0a, 0x0d, 0x44, 0x41, 0x54, 0x41, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x55, 0x49, 0x4e, 0x54, 0x10, 0x07, 0x12, 0x12, 0x0a, 0x0e, 0x44, 0x41,
	0x54, 0x41, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x49, 0x4e, 0x54, 0x38, 0x10, 0x08, 0x12, 0x13,
	0x0a, 0x0f, 0x44, 0x41, 0x54, 0x41, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x49, 0x4e, 0x54, 0x31,
	0x36, 0x10, 0x09, 0x12, 0x13, 0x0a, 0x0f, 0x44, 0x41, 0x54, 0x41, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x55, 0x49, 0x4e, 0x54, 0x33, 0x32, 0x10, 0x0a, 0x12, 0x13, 0x0a, 0x0f, 0x44, 0x41, 0x54, 0x41,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x54, 0x49, 0x4e, 0x36, 0x34, 0x10, 0x0b, 0x12, 0x14, 0x0a,
	0x10, 0x44, 0x41, 0x54, 0x41, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x46, 0x4c, 0x4f, 0x41, 0x54, 0x33,
	0x32, 0x10, 0x0c, 0x12, 0x14, 0x0a, 0x10, 0x44, 0x41, 0x54, 0x41, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x46, 0x4c, 0x4f, 0x41, 0x54, 0x36, 0x34, 0x10, 0x0d, 0x12, 0x13, 0x0a, 0x0f, 0x44, 0x41, 0x54,
	0x41, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x10, 0x0e, 0x2a, 0x4f,
	0x0a, 0x0a, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x17, 0x0a, 0x13,
	0x43, 0x4f, 0x4c, 0x55, 0x4d, 0x4e, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x44, 0x45, 0x46, 0x49,
	0x4e, 0x45, 0x44, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x4f, 0x4c, 0x55, 0x4d, 0x4e, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x54, 0x41, 0x47, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x43, 0x4f, 0x4c,
	0x55, 0x4d, 0x4e, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x10, 0x02, 0x2a,
	0xe4, 0x03, 0x0a, 0x0c, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1a, 0x0a, 0x16, 0x46, 0x55, 0x4e, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x55, 0x4e, 0x44, 0x45, 0x46, 0x49, 0x4e, 0x45, 0x44, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12,
	0x46, 0x55, 0x4e, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x4f, 0x55,
	0x4e, 0x54, 0x10, 0x01, 0x12, 0x19, 0x0a, 0x15, 0x46, 0x55, 0x4e, 0x43, 0x54, 0x49, 0x4f, 0x4e,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x49, 0x53, 0x54, 0x49, 0x4e, 0x43, 0x54, 0x10, 0x02, 0x12,
	0x19, 0x0a, 0x15, 0x46, 0x55, 0x4e, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x49, 0x4e, 0x54, 0x45, 0x47, 0x52, 0x41, 0x4c, 0x10, 0x03, 0x12, 0x15, 0x0a, 0x11, 0x46, 0x55,
	0x4e, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d, 0x45, 0x41, 0x4e, 0x10,
	0x04, 0x12, 0x17, 0x0a, 0x13, 0x46, 0x55, 0x4e, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x4d, 0x45, 0x44, 0x49, 0x41, 0x4e, 0x10, 0x05, 0x12, 0x15, 0x0a, 0x11, 0x46, 0x55,
	0x4e, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x10,
	0x06, 0x12, 0x17, 0x0a, 0x13, 0x46, 0x55, 0x4e, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x53, 0x50, 0x52, 0x45, 0x41, 0x44, 0x10, 0x07, 0x12, 0x17, 0x0a, 0x13, 0x46, 0x55,
	0x4e, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x54, 0x44, 0x44, 0x45,
	0x56, 0x10, 0x08, 0x12, 0x14, 0x0a, 0x10, 0x46, 0x55, 0x4e, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x53, 0x55, 0x4d, 0x10, 0x09, 0x12, 0x17, 0x0a, 0x13, 0x46, 0x55, 0x4e,
	0x43, 0x54, 0x49, 0x4f, 0x4e, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x42, 0x4f, 0x54, 0x54, 0x4f, 0x4d,
	0x10, 0x0a, 0x12, 0x16, 0x0a, 0x12, 0x46, 0x55, 0x4e, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x46, 0x49, 0x52, 0x53, 0x54, 0x10, 0x0b, 0x12, 0x15, 0x0a, 0x11, 0x46, 0x55,
	0x4e, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4c, 0x41, 0x53, 0x54, 0x10,
	0x0c, 0x12, 0x14, 0x0a, 0x10, 0x46, 0x55, 0x4e, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x4d, 0x41, 0x58, 0x10, 0x0d, 0x12, 0x14, 0x0a, 0x10, 0x46, 0x55, 0x4e, 0x43, 0x54,
	0x49, 0x4f, 0x4e, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d, 0x49, 0x4e, 0x10, 0x0e, 0x12, 0x1b, 0x0a,
	0x17, 0x46, 0x55, 0x4e, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x45,
	0x52, 0x43, 0x45, 0x4e, 0x54, 0x49, 0x4c, 0x45, 0x10, 0x0f, 0x12, 0x17, 0x0a, 0x13, 0x46, 0x55,
	0x4e, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x41, 0x4d, 0x50, 0x4c,
	0x45, 0x10, 0x10, 0x12, 0x14, 0x0a, 0x10, 0x46, 0x55, 0x4e, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x54, 0x4f, 0x50, 0x10, 0x11, 0x12, 0x1b, 0x0a, 0x17, 0x46, 0x55, 0x4e,
	0x43, 0x54, 0x49, 0x4f, 0x4e, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x45, 0x52, 0x49, 0x56, 0x41,
	0x54, 0x49, 0x56, 0x45, 0x10, 0x12, 0x42, 0x42, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x2d,
	0x61, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6c, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x5f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61,
	0x68, 0x75, 0x62, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_alameda_api_v1alpha1_datahub_common_types_proto_rawDescOnce sync.Once
	file_alameda_api_v1alpha1_datahub_common_types_proto_rawDescData = file_alameda_api_v1alpha1_datahub_common_types_proto_rawDesc
)

func file_alameda_api_v1alpha1_datahub_common_types_proto_rawDescGZIP() []byte {
	file_alameda_api_v1alpha1_datahub_common_types_proto_rawDescOnce.Do(func() {
		file_alameda_api_v1alpha1_datahub_common_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_alameda_api_v1alpha1_datahub_common_types_proto_rawDescData)
	})
	return file_alameda_api_v1alpha1_datahub_common_types_proto_rawDescData
}

var file_alameda_api_v1alpha1_datahub_common_types_proto_enumTypes = make([]protoimpl.EnumInfo, 5)
var file_alameda_api_v1alpha1_datahub_common_types_proto_goTypes = []interface{}{
	(ResourceBoundary)(0), // 0: containersai.alameda.v1alpha1.datahub.common.ResourceBoundary
	(ResourceQuota)(0),    // 1: containersai.alameda.v1alpha1.datahub.common.ResourceQuota
	(DataType)(0),         // 2: containersai.alameda.v1alpha1.datahub.common.DataType
	(ColumnType)(0),       // 3: containersai.alameda.v1alpha1.datahub.common.ColumnType
	(FunctionType)(0),     // 4: containersai.alameda.v1alpha1.datahub.common.FunctionType
}
var file_alameda_api_v1alpha1_datahub_common_types_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_alameda_api_v1alpha1_datahub_common_types_proto_init() }
func file_alameda_api_v1alpha1_datahub_common_types_proto_init() {
	if File_alameda_api_v1alpha1_datahub_common_types_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_alameda_api_v1alpha1_datahub_common_types_proto_rawDesc,
			NumEnums:      5,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_alameda_api_v1alpha1_datahub_common_types_proto_goTypes,
		DependencyIndexes: file_alameda_api_v1alpha1_datahub_common_types_proto_depIdxs,
		EnumInfos:         file_alameda_api_v1alpha1_datahub_common_types_proto_enumTypes,
	}.Build()
	File_alameda_api_v1alpha1_datahub_common_types_proto = out.File
	file_alameda_api_v1alpha1_datahub_common_types_proto_rawDesc = nil
	file_alameda_api_v1alpha1_datahub_common_types_proto_goTypes = nil
	file_alameda_api_v1alpha1_datahub_common_types_proto_depIdxs = nil
}
