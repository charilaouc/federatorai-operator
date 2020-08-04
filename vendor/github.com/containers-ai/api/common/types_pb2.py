# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: common/types.proto

from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='common/types.proto',
  package='containersai.common',
  syntax='proto3',
  serialized_options=b'Z#github.com/containers-ai/api/common',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x12\x63ommon/types.proto\x12\x13\x63ontainersai.common*\xbd\x02\n\x08\x44\x61taType\x12\x16\n\x12\x44\x41TATYPE_UNDEFINED\x10\x00\x12\x11\n\rDATATYPE_BOOL\x10\x01\x12\x10\n\x0c\x44\x41TATYPE_INT\x10\x02\x12\x11\n\rDATATYPE_INT8\x10\x03\x12\x12\n\x0e\x44\x41TATYPE_INT16\x10\x04\x12\x12\n\x0e\x44\x41TATYPE_INT32\x10\x05\x12\x12\n\x0e\x44\x41TATYPE_INT64\x10\x06\x12\x11\n\rDATATYPE_UINT\x10\x07\x12\x12\n\x0e\x44\x41TATYPE_UINT8\x10\x08\x12\x13\n\x0f\x44\x41TATYPE_UINT16\x10\t\x12\x13\n\x0f\x44\x41TATYPE_UINT32\x10\n\x12\x13\n\x0f\x44\x41TATYPE_UTIN64\x10\x0b\x12\x14\n\x10\x44\x41TATYPE_FLOAT32\x10\x0c\x12\x14\n\x10\x44\x41TATYPE_FLOAT64\x10\r\x12\x13\n\x0f\x44\x41TATYPE_STRING\x10\x0e*O\n\nColumnType\x12\x17\n\x13\x43OLUMNTYPE_UDEFINED\x10\x00\x12\x12\n\x0e\x43OLUMNTYPE_TAG\x10\x01\x12\x14\n\x10\x43OLUMNTYPE_FIELD\x10\x02\x42%Z#github.com/containers-ai/api/commonb\x06proto3'
)

_DATATYPE = _descriptor.EnumDescriptor(
  name='DataType',
  full_name='containersai.common.DataType',
  filename=None,
  file=DESCRIPTOR,
  create_key=_descriptor._internal_create_key,
  values=[
    _descriptor.EnumValueDescriptor(
      name='DATATYPE_UNDEFINED', index=0, number=0,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='DATATYPE_BOOL', index=1, number=1,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='DATATYPE_INT', index=2, number=2,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='DATATYPE_INT8', index=3, number=3,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='DATATYPE_INT16', index=4, number=4,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='DATATYPE_INT32', index=5, number=5,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='DATATYPE_INT64', index=6, number=6,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='DATATYPE_UINT', index=7, number=7,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='DATATYPE_UINT8', index=8, number=8,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='DATATYPE_UINT16', index=9, number=9,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='DATATYPE_UINT32', index=10, number=10,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='DATATYPE_UTIN64', index=11, number=11,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='DATATYPE_FLOAT32', index=12, number=12,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='DATATYPE_FLOAT64', index=13, number=13,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='DATATYPE_STRING', index=14, number=14,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=44,
  serialized_end=361,
)
_sym_db.RegisterEnumDescriptor(_DATATYPE)

DataType = enum_type_wrapper.EnumTypeWrapper(_DATATYPE)
_COLUMNTYPE = _descriptor.EnumDescriptor(
  name='ColumnType',
  full_name='containersai.common.ColumnType',
  filename=None,
  file=DESCRIPTOR,
  create_key=_descriptor._internal_create_key,
  values=[
    _descriptor.EnumValueDescriptor(
      name='COLUMNTYPE_UDEFINED', index=0, number=0,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='COLUMNTYPE_TAG', index=1, number=1,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='COLUMNTYPE_FIELD', index=2, number=2,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=363,
  serialized_end=442,
)
_sym_db.RegisterEnumDescriptor(_COLUMNTYPE)

ColumnType = enum_type_wrapper.EnumTypeWrapper(_COLUMNTYPE)
DATATYPE_UNDEFINED = 0
DATATYPE_BOOL = 1
DATATYPE_INT = 2
DATATYPE_INT8 = 3
DATATYPE_INT16 = 4
DATATYPE_INT32 = 5
DATATYPE_INT64 = 6
DATATYPE_UINT = 7
DATATYPE_UINT8 = 8
DATATYPE_UINT16 = 9
DATATYPE_UINT32 = 10
DATATYPE_UTIN64 = 11
DATATYPE_FLOAT32 = 12
DATATYPE_FLOAT64 = 13
DATATYPE_STRING = 14
COLUMNTYPE_UDEFINED = 0
COLUMNTYPE_TAG = 1
COLUMNTYPE_FIELD = 2


DESCRIPTOR.enum_types_by_name['DataType'] = _DATATYPE
DESCRIPTOR.enum_types_by_name['ColumnType'] = _COLUMNTYPE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)
