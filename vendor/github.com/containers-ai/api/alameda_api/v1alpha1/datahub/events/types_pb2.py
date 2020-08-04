# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: alameda_api/v1alpha1/datahub/events/types.proto

from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='alameda_api/v1alpha1/datahub/events/types.proto',
  package='containersai.alameda.v1alpha1.datahub.events',
  syntax='proto3',
  serialized_options=b'Z@github.com/containers-ai/api/alameda_api/v1alpha1/datahub/events',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n/alameda_api/v1alpha1/datahub/events/types.proto\x12,containersai.alameda.v1alpha1.datahub.events\".\n\x0b\x45ventSource\x12\x0c\n\x04host\x18\x01 \x01(\t\x12\x11\n\tcomponent\x18\x02 \x01(\t\"X\n\x12K8SObjectReference\x12\x0c\n\x04kind\x18\x01 \x01(\t\x12\x11\n\tnamespace\x18\x02 \x01(\t\x12\x0c\n\x04name\x18\x03 \x01(\t\x12\x13\n\x0b\x61pi_version\x18\x04 \x01(\t*\xdd\x06\n\tEventType\x12\x18\n\x14\x45VENT_TYPE_UNDEFINED\x10\x00\x12$\n EVENT_TYPE_ALAMEDA_SCALER_CREATE\x10\x01\x12$\n EVENT_TYPE_ALAMEDA_SCALER_DELETE\x10\x02\x12\x1c\n\x18\x45VENT_TYPE_NODE_REGISTER\x10\x03\x12\"\n\x1e\x45VENT_TYPE_DEPLOYMENT_REGISTER\x10\x04\x12)\n%EVENT_TYPE_DEPLOYMENT_CONFIG_REGISTER\x10\x05\x12\x1b\n\x17\x45VENT_TYPE_POD_REGISTER\x10\x06\x12\x1e\n\x1a\x45VENT_TYPE_NODE_DEREGISTER\x10\x07\x12$\n EVENT_TYPE_DEPLOYMENT_DEREGISTER\x10\x08\x12+\n\'EVENT_TYPE_DEPLOYMENT_CONFIG_DEREGISTER\x10\t\x12\x1d\n\x19\x45VENT_TYPE_POD_DEREGISTER\x10\n\x12%\n!EVENT_TYPE_NODE_PREDICTION_CREATE\x10\x0b\x12$\n EVENT_TYPE_POD_PREDICTION_CREATE\x10\x0c\x12(\n$EVENT_TYPE_VPA_RECOMMENDATION_CREATE\x10\r\x12(\n$EVENT_TYPE_HPA_RECOMMENDATION_CREATE\x10\x0e\x12)\n%EVENT_TYPE_VPA_RECOMMENDATION_EXECUTE\x10\x0f\x12)\n%EVENT_TYPE_HPA_RECOMMENDATION_EXECUTE\x10\x10\x12$\n EVENT_TYPE_ANOMALY_METRIC_DETECT\x10\x11\x12&\n\"EVENT_TYPE_ANOMALY_ANALYSIS_CREATE\x10\x12\x12\x16\n\x12\x45VENT_TYPE_LICENSE\x10\x13\x12!\n\x1d\x45VENT_TYPE_EMAIL_NOTIFICATION\x10\x14\x12&\n\"EVENT_TYPE_ANOMALY_FORECAST_DETECT\x10\x15\x12&\n\"EVENT_TYPE_ANOMALY_REALTIME_DETECT\x10\x16*A\n\x0c\x45ventVersion\x12\x1b\n\x17\x45VENT_VERSION_UNDEFINED\x10\x00\x12\x14\n\x10\x45VENT_VERSION_V1\x10\x01*\x9b\x01\n\nEventLevel\x12\x19\n\x15\x45VENT_LEVEL_UNDEFINED\x10\x00\x12\x15\n\x11\x45VENT_LEVEL_DEBUG\x10\x01\x12\x14\n\x10\x45VENT_LEVEL_INFO\x10\x02\x12\x17\n\x13\x45VENT_LEVEL_WARNING\x10\x03\x12\x15\n\x11\x45VENT_LEVEL_ERROR\x10\x04\x12\x15\n\x11\x45VENT_LEVEL_FATAL\x10\x05\x42\x42Z@github.com/containers-ai/api/alameda_api/v1alpha1/datahub/eventsb\x06proto3'
)

_EVENTTYPE = _descriptor.EnumDescriptor(
  name='EventType',
  full_name='containersai.alameda.v1alpha1.datahub.events.EventType',
  filename=None,
  file=DESCRIPTOR,
  create_key=_descriptor._internal_create_key,
  values=[
    _descriptor.EnumValueDescriptor(
      name='EVENT_TYPE_UNDEFINED', index=0, number=0,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='EVENT_TYPE_ALAMEDA_SCALER_CREATE', index=1, number=1,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='EVENT_TYPE_ALAMEDA_SCALER_DELETE', index=2, number=2,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='EVENT_TYPE_NODE_REGISTER', index=3, number=3,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='EVENT_TYPE_DEPLOYMENT_REGISTER', index=4, number=4,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='EVENT_TYPE_DEPLOYMENT_CONFIG_REGISTER', index=5, number=5,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='EVENT_TYPE_POD_REGISTER', index=6, number=6,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='EVENT_TYPE_NODE_DEREGISTER', index=7, number=7,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='EVENT_TYPE_DEPLOYMENT_DEREGISTER', index=8, number=8,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='EVENT_TYPE_DEPLOYMENT_CONFIG_DEREGISTER', index=9, number=9,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='EVENT_TYPE_POD_DEREGISTER', index=10, number=10,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='EVENT_TYPE_NODE_PREDICTION_CREATE', index=11, number=11,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='EVENT_TYPE_POD_PREDICTION_CREATE', index=12, number=12,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='EVENT_TYPE_VPA_RECOMMENDATION_CREATE', index=13, number=13,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='EVENT_TYPE_HPA_RECOMMENDATION_CREATE', index=14, number=14,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='EVENT_TYPE_VPA_RECOMMENDATION_EXECUTE', index=15, number=15,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='EVENT_TYPE_HPA_RECOMMENDATION_EXECUTE', index=16, number=16,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='EVENT_TYPE_ANOMALY_METRIC_DETECT', index=17, number=17,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='EVENT_TYPE_ANOMALY_ANALYSIS_CREATE', index=18, number=18,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='EVENT_TYPE_LICENSE', index=19, number=19,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='EVENT_TYPE_EMAIL_NOTIFICATION', index=20, number=20,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='EVENT_TYPE_ANOMALY_FORECAST_DETECT', index=21, number=21,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='EVENT_TYPE_ANOMALY_REALTIME_DETECT', index=22, number=22,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=236,
  serialized_end=1097,
)
_sym_db.RegisterEnumDescriptor(_EVENTTYPE)

EventType = enum_type_wrapper.EnumTypeWrapper(_EVENTTYPE)
_EVENTVERSION = _descriptor.EnumDescriptor(
  name='EventVersion',
  full_name='containersai.alameda.v1alpha1.datahub.events.EventVersion',
  filename=None,
  file=DESCRIPTOR,
  create_key=_descriptor._internal_create_key,
  values=[
    _descriptor.EnumValueDescriptor(
      name='EVENT_VERSION_UNDEFINED', index=0, number=0,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='EVENT_VERSION_V1', index=1, number=1,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=1099,
  serialized_end=1164,
)
_sym_db.RegisterEnumDescriptor(_EVENTVERSION)

EventVersion = enum_type_wrapper.EnumTypeWrapper(_EVENTVERSION)
_EVENTLEVEL = _descriptor.EnumDescriptor(
  name='EventLevel',
  full_name='containersai.alameda.v1alpha1.datahub.events.EventLevel',
  filename=None,
  file=DESCRIPTOR,
  create_key=_descriptor._internal_create_key,
  values=[
    _descriptor.EnumValueDescriptor(
      name='EVENT_LEVEL_UNDEFINED', index=0, number=0,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='EVENT_LEVEL_DEBUG', index=1, number=1,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='EVENT_LEVEL_INFO', index=2, number=2,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='EVENT_LEVEL_WARNING', index=3, number=3,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='EVENT_LEVEL_ERROR', index=4, number=4,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='EVENT_LEVEL_FATAL', index=5, number=5,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=1167,
  serialized_end=1322,
)
_sym_db.RegisterEnumDescriptor(_EVENTLEVEL)

EventLevel = enum_type_wrapper.EnumTypeWrapper(_EVENTLEVEL)
EVENT_TYPE_UNDEFINED = 0
EVENT_TYPE_ALAMEDA_SCALER_CREATE = 1
EVENT_TYPE_ALAMEDA_SCALER_DELETE = 2
EVENT_TYPE_NODE_REGISTER = 3
EVENT_TYPE_DEPLOYMENT_REGISTER = 4
EVENT_TYPE_DEPLOYMENT_CONFIG_REGISTER = 5
EVENT_TYPE_POD_REGISTER = 6
EVENT_TYPE_NODE_DEREGISTER = 7
EVENT_TYPE_DEPLOYMENT_DEREGISTER = 8
EVENT_TYPE_DEPLOYMENT_CONFIG_DEREGISTER = 9
EVENT_TYPE_POD_DEREGISTER = 10
EVENT_TYPE_NODE_PREDICTION_CREATE = 11
EVENT_TYPE_POD_PREDICTION_CREATE = 12
EVENT_TYPE_VPA_RECOMMENDATION_CREATE = 13
EVENT_TYPE_HPA_RECOMMENDATION_CREATE = 14
EVENT_TYPE_VPA_RECOMMENDATION_EXECUTE = 15
EVENT_TYPE_HPA_RECOMMENDATION_EXECUTE = 16
EVENT_TYPE_ANOMALY_METRIC_DETECT = 17
EVENT_TYPE_ANOMALY_ANALYSIS_CREATE = 18
EVENT_TYPE_LICENSE = 19
EVENT_TYPE_EMAIL_NOTIFICATION = 20
EVENT_TYPE_ANOMALY_FORECAST_DETECT = 21
EVENT_TYPE_ANOMALY_REALTIME_DETECT = 22
EVENT_VERSION_UNDEFINED = 0
EVENT_VERSION_V1 = 1
EVENT_LEVEL_UNDEFINED = 0
EVENT_LEVEL_DEBUG = 1
EVENT_LEVEL_INFO = 2
EVENT_LEVEL_WARNING = 3
EVENT_LEVEL_ERROR = 4
EVENT_LEVEL_FATAL = 5



_EVENTSOURCE = _descriptor.Descriptor(
  name='EventSource',
  full_name='containersai.alameda.v1alpha1.datahub.events.EventSource',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='host', full_name='containersai.alameda.v1alpha1.datahub.events.EventSource.host', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='component', full_name='containersai.alameda.v1alpha1.datahub.events.EventSource.component', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
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
  serialized_start=97,
  serialized_end=143,
)


_K8SOBJECTREFERENCE = _descriptor.Descriptor(
  name='K8SObjectReference',
  full_name='containersai.alameda.v1alpha1.datahub.events.K8SObjectReference',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='kind', full_name='containersai.alameda.v1alpha1.datahub.events.K8SObjectReference.kind', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='namespace', full_name='containersai.alameda.v1alpha1.datahub.events.K8SObjectReference.namespace', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='name', full_name='containersai.alameda.v1alpha1.datahub.events.K8SObjectReference.name', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='api_version', full_name='containersai.alameda.v1alpha1.datahub.events.K8SObjectReference.api_version', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
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
  serialized_start=145,
  serialized_end=233,
)

DESCRIPTOR.message_types_by_name['EventSource'] = _EVENTSOURCE
DESCRIPTOR.message_types_by_name['K8SObjectReference'] = _K8SOBJECTREFERENCE
DESCRIPTOR.enum_types_by_name['EventType'] = _EVENTTYPE
DESCRIPTOR.enum_types_by_name['EventVersion'] = _EVENTVERSION
DESCRIPTOR.enum_types_by_name['EventLevel'] = _EVENTLEVEL
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

EventSource = _reflection.GeneratedProtocolMessageType('EventSource', (_message.Message,), {
  'DESCRIPTOR' : _EVENTSOURCE,
  '__module__' : 'alameda_api.v1alpha1.datahub.events.types_pb2'
  # @@protoc_insertion_point(class_scope:containersai.alameda.v1alpha1.datahub.events.EventSource)
  })
_sym_db.RegisterMessage(EventSource)

K8SObjectReference = _reflection.GeneratedProtocolMessageType('K8SObjectReference', (_message.Message,), {
  'DESCRIPTOR' : _K8SOBJECTREFERENCE,
  '__module__' : 'alameda_api.v1alpha1.datahub.events.types_pb2'
  # @@protoc_insertion_point(class_scope:containersai.alameda.v1alpha1.datahub.events.K8SObjectReference)
  })
_sym_db.RegisterMessage(K8SObjectReference)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)
