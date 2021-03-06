# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: alameda_api/v1alpha1/datahub/applications/types.proto

from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from alameda_api.v1alpha1.datahub.common import rawdata_pb2 as alameda__api_dot_v1alpha1_dot_datahub_dot_common_dot_rawdata__pb2
from alameda_api.v1alpha1.datahub.schemas import types_pb2 as alameda__api_dot_v1alpha1_dot_datahub_dot_schemas_dot_types__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='alameda_api/v1alpha1/datahub/applications/types.proto',
  package='containersai.alameda.v1alpha1.datahub.applications',
  syntax='proto3',
  serialized_options=b'ZFgithub.com/containers-ai/api/alameda_api/v1alpha1/datahub/applications',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n5alameda_api/v1alpha1/datahub/applications/types.proto\x12\x32\x63ontainersai.alameda.v1alpha1.datahub.applications\x1a\x31\x61lameda_api/v1alpha1/datahub/common/rawdata.proto\x1a\x30\x61lameda_api/v1alpha1/datahub/schemas/types.proto\"\xbc\x01\n\x0b\x41pplication\x12N\n\x0bschema_meta\x18\x01 \x01(\x0b\x32\x39.containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta\x12]\n\x10\x61pplication_data\x18\x02 \x03(\x0b\x32\x43.containersai.alameda.v1alpha1.datahub.applications.ApplicationData\"q\n\x0f\x41pplicationData\x12\x13\n\x0bmeasurement\x18\x01 \x01(\t\x12I\n\tread_data\x18\x02 \x01(\x0b\x32\x36.containersai.alameda.v1alpha1.datahub.common.ReadDataBHZFgithub.com/containers-ai/api/alameda_api/v1alpha1/datahub/applicationsb\x06proto3'
  ,
  dependencies=[alameda__api_dot_v1alpha1_dot_datahub_dot_common_dot_rawdata__pb2.DESCRIPTOR,alameda__api_dot_v1alpha1_dot_datahub_dot_schemas_dot_types__pb2.DESCRIPTOR,])




_APPLICATION = _descriptor.Descriptor(
  name='Application',
  full_name='containersai.alameda.v1alpha1.datahub.applications.Application',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='schema_meta', full_name='containersai.alameda.v1alpha1.datahub.applications.Application.schema_meta', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='application_data', full_name='containersai.alameda.v1alpha1.datahub.applications.Application.application_data', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=211,
  serialized_end=399,
)


_APPLICATIONDATA = _descriptor.Descriptor(
  name='ApplicationData',
  full_name='containersai.alameda.v1alpha1.datahub.applications.ApplicationData',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='measurement', full_name='containersai.alameda.v1alpha1.datahub.applications.ApplicationData.measurement', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='read_data', full_name='containersai.alameda.v1alpha1.datahub.applications.ApplicationData.read_data', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=401,
  serialized_end=514,
)

_APPLICATION.fields_by_name['schema_meta'].message_type = alameda__api_dot_v1alpha1_dot_datahub_dot_schemas_dot_types__pb2._SCHEMAMETA
_APPLICATION.fields_by_name['application_data'].message_type = _APPLICATIONDATA
_APPLICATIONDATA.fields_by_name['read_data'].message_type = alameda__api_dot_v1alpha1_dot_datahub_dot_common_dot_rawdata__pb2._READDATA
DESCRIPTOR.message_types_by_name['Application'] = _APPLICATION
DESCRIPTOR.message_types_by_name['ApplicationData'] = _APPLICATIONDATA
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Application = _reflection.GeneratedProtocolMessageType('Application', (_message.Message,), {
  'DESCRIPTOR' : _APPLICATION,
  '__module__' : 'alameda_api.v1alpha1.datahub.applications.types_pb2'
  # @@protoc_insertion_point(class_scope:containersai.alameda.v1alpha1.datahub.applications.Application)
  })
_sym_db.RegisterMessage(Application)

ApplicationData = _reflection.GeneratedProtocolMessageType('ApplicationData', (_message.Message,), {
  'DESCRIPTOR' : _APPLICATIONDATA,
  '__module__' : 'alameda_api.v1alpha1.datahub.applications.types_pb2'
  # @@protoc_insertion_point(class_scope:containersai.alameda.v1alpha1.datahub.applications.ApplicationData)
  })
_sym_db.RegisterMessage(ApplicationData)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)
