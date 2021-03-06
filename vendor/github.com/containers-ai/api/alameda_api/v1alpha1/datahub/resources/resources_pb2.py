# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: alameda_api/v1alpha1/datahub/resources/resources.proto

from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from alameda_api.v1alpha1.datahub.resources import metadata_pb2 as alameda__api_dot_v1alpha1_dot_datahub_dot_resources_dot_metadata__pb2
from alameda_api.v1alpha1.datahub.resources import status_pb2 as alameda__api_dot_v1alpha1_dot_datahub_dot_resources_dot_status__pb2
from alameda_api.v1alpha1.datahub.resources import types_pb2 as alameda__api_dot_v1alpha1_dot_datahub_dot_resources_dot_types__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='alameda_api/v1alpha1/datahub/resources/resources.proto',
  package='containersai.alameda.v1alpha1.datahub.resources',
  syntax='proto3',
  serialized_options=b'ZCgithub.com/containers-ai/api/alameda_api/v1alpha1/datahub/resources',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n6alameda_api/v1alpha1/datahub/resources/resources.proto\x12/containersai.alameda.v1alpha1.datahub.resources\x1a\x35\x61lameda_api/v1alpha1/datahub/resources/metadata.proto\x1a\x33\x61lameda_api/v1alpha1/datahub/resources/status.proto\x1a\x32\x61lameda_api/v1alpha1/datahub/resources/types.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"\xc5\x01\n\tContainer\x12\x0c\n\x04name\x18\x01 \x01(\t\x12X\n\tresources\x18\x02 \x01(\x0b\x32\x45.containersai.alameda.v1alpha1.datahub.resources.ResourceRequirements\x12P\n\x06status\x18\x03 \x01(\x0b\x32@.containersai.alameda.v1alpha1.datahub.resources.ContainerStatus\"\x91\x04\n\x03Pod\x12P\n\x0bobject_meta\x18\x01 \x01(\x0b\x32;.containersai.alameda.v1alpha1.datahub.resources.ObjectMeta\x12.\n\nstart_time\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\x15\n\rresource_link\x18\x03 \x01(\t\x12\x10\n\x08\x61pp_name\x18\x04 \x01(\t\x12\x13\n\x0b\x61pp_part_of\x18\x05 \x01(\t\x12Y\n\x10\x61lameda_pod_spec\x18\x06 \x01(\x0b\x32?.containersai.alameda.v1alpha1.datahub.resources.AlamedaPodSpec\x12S\n\x0etop_controller\x18\x07 \x01(\x0b\x32;.containersai.alameda.v1alpha1.datahub.resources.Controller\x12J\n\x06status\x18\x08 \x01(\x0b\x32:.containersai.alameda.v1alpha1.datahub.resources.PodStatus\x12N\n\ncontainers\x18\t \x03(\x0b\x32:.containersai.alameda.v1alpha1.datahub.resources.Container\"\xb5\x02\n\nController\x12P\n\x0bobject_meta\x18\x01 \x01(\x0b\x32;.containersai.alameda.v1alpha1.datahub.resources.ObjectMeta\x12\x43\n\x04kind\x18\x02 \x01(\x0e\x32\x35.containersai.alameda.v1alpha1.datahub.resources.Kind\x12\x10\n\x08replicas\x18\x03 \x01(\x05\x12\x15\n\rspec_replicas\x18\x04 \x01(\x05\x12g\n\x17\x61lameda_controller_spec\x18\x05 \x01(\x0b\x32\x46.containersai.alameda.v1alpha1.datahub.resources.AlamedaControllerSpec\"\x9c\x02\n\x0b\x41pplication\x12P\n\x0bobject_meta\x18\x01 \x01(\x0b\x32;.containersai.alameda.v1alpha1.datahub.resources.ObjectMeta\x12i\n\x18\x61lameda_application_spec\x18\x02 \x01(\x0b\x32G.containersai.alameda.v1alpha1.datahub.resources.AlamedaApplicationSpec\x12P\n\x0b\x63ontrollers\x18\x03 \x03(\x0b\x32;.containersai.alameda.v1alpha1.datahub.resources.Controller\"]\n\tNamespace\x12P\n\x0bobject_meta\x18\x01 \x01(\x0b\x32;.containersai.alameda.v1alpha1.datahub.resources.ObjectMeta\"\xea\x02\n\x04Node\x12P\n\x0bobject_meta\x18\x01 \x01(\x0b\x32;.containersai.alameda.v1alpha1.datahub.resources.ObjectMeta\x12.\n\nstart_time\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\x36\n\x12machine_start_time\x18\x03 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12K\n\x08\x63\x61pacity\x18\x04 \x01(\x0b\x32\x39.containersai.alameda.v1alpha1.datahub.resources.Capacity\x12[\n\x11\x61lameda_node_spec\x18\x05 \x01(\x0b\x32@.containersai.alameda.v1alpha1.datahub.resources.AlamedaNodeSpec\"[\n\x07\x43luster\x12P\n\x0bobject_meta\x18\x01 \x01(\x0b\x32;.containersai.alameda.v1alpha1.datahub.resources.ObjectMetaBEZCgithub.com/containers-ai/api/alameda_api/v1alpha1/datahub/resourcesb\x06proto3'
  ,
  dependencies=[alameda__api_dot_v1alpha1_dot_datahub_dot_resources_dot_metadata__pb2.DESCRIPTOR,alameda__api_dot_v1alpha1_dot_datahub_dot_resources_dot_status__pb2.DESCRIPTOR,alameda__api_dot_v1alpha1_dot_datahub_dot_resources_dot_types__pb2.DESCRIPTOR,google_dot_protobuf_dot_timestamp__pb2.DESCRIPTOR,])




_CONTAINER = _descriptor.Descriptor(
  name='Container',
  full_name='containersai.alameda.v1alpha1.datahub.resources.Container',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='containersai.alameda.v1alpha1.datahub.resources.Container.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='resources', full_name='containersai.alameda.v1alpha1.datahub.resources.Container.resources', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='status', full_name='containersai.alameda.v1alpha1.datahub.resources.Container.status', index=2,
      number=3, type=11, cpp_type=10, label=1,
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
  serialized_start=301,
  serialized_end=498,
)


_POD = _descriptor.Descriptor(
  name='Pod',
  full_name='containersai.alameda.v1alpha1.datahub.resources.Pod',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='object_meta', full_name='containersai.alameda.v1alpha1.datahub.resources.Pod.object_meta', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='start_time', full_name='containersai.alameda.v1alpha1.datahub.resources.Pod.start_time', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='resource_link', full_name='containersai.alameda.v1alpha1.datahub.resources.Pod.resource_link', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='app_name', full_name='containersai.alameda.v1alpha1.datahub.resources.Pod.app_name', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='app_part_of', full_name='containersai.alameda.v1alpha1.datahub.resources.Pod.app_part_of', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='alameda_pod_spec', full_name='containersai.alameda.v1alpha1.datahub.resources.Pod.alameda_pod_spec', index=5,
      number=6, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='top_controller', full_name='containersai.alameda.v1alpha1.datahub.resources.Pod.top_controller', index=6,
      number=7, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='status', full_name='containersai.alameda.v1alpha1.datahub.resources.Pod.status', index=7,
      number=8, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='containers', full_name='containersai.alameda.v1alpha1.datahub.resources.Pod.containers', index=8,
      number=9, type=11, cpp_type=10, label=3,
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
  serialized_start=501,
  serialized_end=1030,
)


_CONTROLLER = _descriptor.Descriptor(
  name='Controller',
  full_name='containersai.alameda.v1alpha1.datahub.resources.Controller',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='object_meta', full_name='containersai.alameda.v1alpha1.datahub.resources.Controller.object_meta', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='kind', full_name='containersai.alameda.v1alpha1.datahub.resources.Controller.kind', index=1,
      number=2, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='replicas', full_name='containersai.alameda.v1alpha1.datahub.resources.Controller.replicas', index=2,
      number=3, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='spec_replicas', full_name='containersai.alameda.v1alpha1.datahub.resources.Controller.spec_replicas', index=3,
      number=4, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='alameda_controller_spec', full_name='containersai.alameda.v1alpha1.datahub.resources.Controller.alameda_controller_spec', index=4,
      number=5, type=11, cpp_type=10, label=1,
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
  serialized_start=1033,
  serialized_end=1342,
)


_APPLICATION = _descriptor.Descriptor(
  name='Application',
  full_name='containersai.alameda.v1alpha1.datahub.resources.Application',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='object_meta', full_name='containersai.alameda.v1alpha1.datahub.resources.Application.object_meta', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='alameda_application_spec', full_name='containersai.alameda.v1alpha1.datahub.resources.Application.alameda_application_spec', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='controllers', full_name='containersai.alameda.v1alpha1.datahub.resources.Application.controllers', index=2,
      number=3, type=11, cpp_type=10, label=3,
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
  serialized_start=1345,
  serialized_end=1629,
)


_NAMESPACE = _descriptor.Descriptor(
  name='Namespace',
  full_name='containersai.alameda.v1alpha1.datahub.resources.Namespace',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='object_meta', full_name='containersai.alameda.v1alpha1.datahub.resources.Namespace.object_meta', index=0,
      number=1, type=11, cpp_type=10, label=1,
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
  serialized_start=1631,
  serialized_end=1724,
)


_NODE = _descriptor.Descriptor(
  name='Node',
  full_name='containersai.alameda.v1alpha1.datahub.resources.Node',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='object_meta', full_name='containersai.alameda.v1alpha1.datahub.resources.Node.object_meta', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='start_time', full_name='containersai.alameda.v1alpha1.datahub.resources.Node.start_time', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='machine_start_time', full_name='containersai.alameda.v1alpha1.datahub.resources.Node.machine_start_time', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='capacity', full_name='containersai.alameda.v1alpha1.datahub.resources.Node.capacity', index=3,
      number=4, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='alameda_node_spec', full_name='containersai.alameda.v1alpha1.datahub.resources.Node.alameda_node_spec', index=4,
      number=5, type=11, cpp_type=10, label=1,
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
  serialized_start=1727,
  serialized_end=2089,
)


_CLUSTER = _descriptor.Descriptor(
  name='Cluster',
  full_name='containersai.alameda.v1alpha1.datahub.resources.Cluster',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='object_meta', full_name='containersai.alameda.v1alpha1.datahub.resources.Cluster.object_meta', index=0,
      number=1, type=11, cpp_type=10, label=1,
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
  serialized_start=2091,
  serialized_end=2182,
)

_CONTAINER.fields_by_name['resources'].message_type = alameda__api_dot_v1alpha1_dot_datahub_dot_resources_dot_types__pb2._RESOURCEREQUIREMENTS
_CONTAINER.fields_by_name['status'].message_type = alameda__api_dot_v1alpha1_dot_datahub_dot_resources_dot_status__pb2._CONTAINERSTATUS
_POD.fields_by_name['object_meta'].message_type = alameda__api_dot_v1alpha1_dot_datahub_dot_resources_dot_metadata__pb2._OBJECTMETA
_POD.fields_by_name['start_time'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_POD.fields_by_name['alameda_pod_spec'].message_type = alameda__api_dot_v1alpha1_dot_datahub_dot_resources_dot_types__pb2._ALAMEDAPODSPEC
_POD.fields_by_name['top_controller'].message_type = _CONTROLLER
_POD.fields_by_name['status'].message_type = alameda__api_dot_v1alpha1_dot_datahub_dot_resources_dot_status__pb2._PODSTATUS
_POD.fields_by_name['containers'].message_type = _CONTAINER
_CONTROLLER.fields_by_name['object_meta'].message_type = alameda__api_dot_v1alpha1_dot_datahub_dot_resources_dot_metadata__pb2._OBJECTMETA
_CONTROLLER.fields_by_name['kind'].enum_type = alameda__api_dot_v1alpha1_dot_datahub_dot_resources_dot_metadata__pb2._KIND
_CONTROLLER.fields_by_name['alameda_controller_spec'].message_type = alameda__api_dot_v1alpha1_dot_datahub_dot_resources_dot_types__pb2._ALAMEDACONTROLLERSPEC
_APPLICATION.fields_by_name['object_meta'].message_type = alameda__api_dot_v1alpha1_dot_datahub_dot_resources_dot_metadata__pb2._OBJECTMETA
_APPLICATION.fields_by_name['alameda_application_spec'].message_type = alameda__api_dot_v1alpha1_dot_datahub_dot_resources_dot_types__pb2._ALAMEDAAPPLICATIONSPEC
_APPLICATION.fields_by_name['controllers'].message_type = _CONTROLLER
_NAMESPACE.fields_by_name['object_meta'].message_type = alameda__api_dot_v1alpha1_dot_datahub_dot_resources_dot_metadata__pb2._OBJECTMETA
_NODE.fields_by_name['object_meta'].message_type = alameda__api_dot_v1alpha1_dot_datahub_dot_resources_dot_metadata__pb2._OBJECTMETA
_NODE.fields_by_name['start_time'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_NODE.fields_by_name['machine_start_time'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_NODE.fields_by_name['capacity'].message_type = alameda__api_dot_v1alpha1_dot_datahub_dot_resources_dot_types__pb2._CAPACITY
_NODE.fields_by_name['alameda_node_spec'].message_type = alameda__api_dot_v1alpha1_dot_datahub_dot_resources_dot_types__pb2._ALAMEDANODESPEC
_CLUSTER.fields_by_name['object_meta'].message_type = alameda__api_dot_v1alpha1_dot_datahub_dot_resources_dot_metadata__pb2._OBJECTMETA
DESCRIPTOR.message_types_by_name['Container'] = _CONTAINER
DESCRIPTOR.message_types_by_name['Pod'] = _POD
DESCRIPTOR.message_types_by_name['Controller'] = _CONTROLLER
DESCRIPTOR.message_types_by_name['Application'] = _APPLICATION
DESCRIPTOR.message_types_by_name['Namespace'] = _NAMESPACE
DESCRIPTOR.message_types_by_name['Node'] = _NODE
DESCRIPTOR.message_types_by_name['Cluster'] = _CLUSTER
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Container = _reflection.GeneratedProtocolMessageType('Container', (_message.Message,), {
  'DESCRIPTOR' : _CONTAINER,
  '__module__' : 'alameda_api.v1alpha1.datahub.resources.resources_pb2'
  # @@protoc_insertion_point(class_scope:containersai.alameda.v1alpha1.datahub.resources.Container)
  })
_sym_db.RegisterMessage(Container)

Pod = _reflection.GeneratedProtocolMessageType('Pod', (_message.Message,), {
  'DESCRIPTOR' : _POD,
  '__module__' : 'alameda_api.v1alpha1.datahub.resources.resources_pb2'
  # @@protoc_insertion_point(class_scope:containersai.alameda.v1alpha1.datahub.resources.Pod)
  })
_sym_db.RegisterMessage(Pod)

Controller = _reflection.GeneratedProtocolMessageType('Controller', (_message.Message,), {
  'DESCRIPTOR' : _CONTROLLER,
  '__module__' : 'alameda_api.v1alpha1.datahub.resources.resources_pb2'
  # @@protoc_insertion_point(class_scope:containersai.alameda.v1alpha1.datahub.resources.Controller)
  })
_sym_db.RegisterMessage(Controller)

Application = _reflection.GeneratedProtocolMessageType('Application', (_message.Message,), {
  'DESCRIPTOR' : _APPLICATION,
  '__module__' : 'alameda_api.v1alpha1.datahub.resources.resources_pb2'
  # @@protoc_insertion_point(class_scope:containersai.alameda.v1alpha1.datahub.resources.Application)
  })
_sym_db.RegisterMessage(Application)

Namespace = _reflection.GeneratedProtocolMessageType('Namespace', (_message.Message,), {
  'DESCRIPTOR' : _NAMESPACE,
  '__module__' : 'alameda_api.v1alpha1.datahub.resources.resources_pb2'
  # @@protoc_insertion_point(class_scope:containersai.alameda.v1alpha1.datahub.resources.Namespace)
  })
_sym_db.RegisterMessage(Namespace)

Node = _reflection.GeneratedProtocolMessageType('Node', (_message.Message,), {
  'DESCRIPTOR' : _NODE,
  '__module__' : 'alameda_api.v1alpha1.datahub.resources.resources_pb2'
  # @@protoc_insertion_point(class_scope:containersai.alameda.v1alpha1.datahub.resources.Node)
  })
_sym_db.RegisterMessage(Node)

Cluster = _reflection.GeneratedProtocolMessageType('Cluster', (_message.Message,), {
  'DESCRIPTOR' : _CLUSTER,
  '__module__' : 'alameda_api.v1alpha1.datahub.resources.resources_pb2'
  # @@protoc_insertion_point(class_scope:containersai.alameda.v1alpha1.datahub.resources.Cluster)
  })
_sym_db.RegisterMessage(Cluster)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)
