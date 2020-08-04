# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: alameda_api/v1alpha1/datahub/predictions/predictions.proto

from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from alameda_api.v1alpha1.datahub.common import metrics_pb2 as alameda__api_dot_v1alpha1_dot_datahub_dot_common_dot_metrics__pb2
from alameda_api.v1alpha1.datahub.common import rawdata_pb2 as alameda__api_dot_v1alpha1_dot_datahub_dot_common_dot_rawdata__pb2
from alameda_api.v1alpha1.datahub.common import types_pb2 as alameda__api_dot_v1alpha1_dot_datahub_dot_common_dot_types__pb2
from alameda_api.v1alpha1.datahub.predictions import types_pb2 as alameda__api_dot_v1alpha1_dot_datahub_dot_predictions_dot_types__pb2
from alameda_api.v1alpha1.datahub.resources import metadata_pb2 as alameda__api_dot_v1alpha1_dot_datahub_dot_resources_dot_metadata__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='alameda_api/v1alpha1/datahub/predictions/predictions.proto',
  package='containersai.alameda.v1alpha1.datahub.predictions',
  syntax='proto3',
  serialized_options=b'ZEgithub.com/containers-ai/api/alameda_api/v1alpha1/datahub/predictions',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n:alameda_api/v1alpha1/datahub/predictions/predictions.proto\x12\x31\x63ontainersai.alameda.v1alpha1.datahub.predictions\x1a\x31\x61lameda_api/v1alpha1/datahub/common/metrics.proto\x1a\x31\x61lameda_api/v1alpha1/datahub/common/rawdata.proto\x1a/alameda_api/v1alpha1/datahub/common/types.proto\x1a\x34\x61lameda_api/v1alpha1/datahub/predictions/types.proto\x1a\x35\x61lameda_api/v1alpha1/datahub/resources/metadata.proto\"\xc2\x02\n\x13\x43ontainerPrediction\x12\x0c\n\x04name\x18\x01 \x01(\t\x12Y\n\x12predicted_raw_data\x18\x02 \x03(\x0b\x32=.containersai.alameda.v1alpha1.datahub.predictions.MetricData\x12`\n\x19predicted_upperbound_data\x18\x03 \x03(\x0b\x32=.containersai.alameda.v1alpha1.datahub.predictions.MetricData\x12`\n\x19predicted_lowerbound_data\x18\x04 \x03(\x0b\x32=.containersai.alameda.v1alpha1.datahub.predictions.MetricData\"\xc8\x01\n\rPodPrediction\x12P\n\x0bobject_meta\x18\x01 \x01(\x0b\x32;.containersai.alameda.v1alpha1.datahub.resources.ObjectMeta\x12\x65\n\x15\x63ontainer_predictions\x18\x02 \x03(\x0b\x32\x46.containersai.alameda.v1alpha1.datahub.predictions.ContainerPrediction\"\xcc\x03\n\x14\x43ontrollerPrediction\x12P\n\x0bobject_meta\x18\x01 \x01(\x0b\x32;.containersai.alameda.v1alpha1.datahub.resources.ObjectMeta\x12\x43\n\x04kind\x18\x02 \x01(\x0e\x32\x35.containersai.alameda.v1alpha1.datahub.resources.Kind\x12Y\n\x12predicted_raw_data\x18\x03 \x03(\x0b\x32=.containersai.alameda.v1alpha1.datahub.predictions.MetricData\x12`\n\x19predicted_upperbound_data\x18\x04 \x03(\x0b\x32=.containersai.alameda.v1alpha1.datahub.predictions.MetricData\x12`\n\x19predicted_lowerbound_data\x18\x05 \x03(\x0b\x32=.containersai.alameda.v1alpha1.datahub.predictions.MetricData\"\x88\x03\n\x15\x41pplicationPrediction\x12P\n\x0bobject_meta\x18\x01 \x01(\x0b\x32;.containersai.alameda.v1alpha1.datahub.resources.ObjectMeta\x12Y\n\x12predicted_raw_data\x18\x02 \x03(\x0b\x32=.containersai.alameda.v1alpha1.datahub.predictions.MetricData\x12`\n\x19predicted_upperbound_data\x18\x03 \x03(\x0b\x32=.containersai.alameda.v1alpha1.datahub.predictions.MetricData\x12`\n\x19predicted_lowerbound_data\x18\x04 \x03(\x0b\x32=.containersai.alameda.v1alpha1.datahub.predictions.MetricData\"\x86\x03\n\x13NamespacePrediction\x12P\n\x0bobject_meta\x18\x01 \x01(\x0b\x32;.containersai.alameda.v1alpha1.datahub.resources.ObjectMeta\x12Y\n\x12predicted_raw_data\x18\x02 \x03(\x0b\x32=.containersai.alameda.v1alpha1.datahub.predictions.MetricData\x12`\n\x19predicted_upperbound_data\x18\x03 \x03(\x0b\x32=.containersai.alameda.v1alpha1.datahub.predictions.MetricData\x12`\n\x19predicted_lowerbound_data\x18\x04 \x03(\x0b\x32=.containersai.alameda.v1alpha1.datahub.predictions.MetricData\"\x97\x03\n\x0eNodePrediction\x12P\n\x0bobject_meta\x18\x01 \x01(\x0b\x32;.containersai.alameda.v1alpha1.datahub.resources.ObjectMeta\x12\x14\n\x0cis_scheduled\x18\x02 \x01(\x08\x12Y\n\x12predicted_raw_data\x18\x03 \x03(\x0b\x32=.containersai.alameda.v1alpha1.datahub.predictions.MetricData\x12`\n\x19predicted_upperbound_data\x18\x04 \x03(\x0b\x32=.containersai.alameda.v1alpha1.datahub.predictions.MetricData\x12`\n\x19predicted_lowerbound_data\x18\x05 \x03(\x0b\x32=.containersai.alameda.v1alpha1.datahub.predictions.MetricData\"\x84\x03\n\x11\x43lusterPrediction\x12P\n\x0bobject_meta\x18\x01 \x01(\x0b\x32;.containersai.alameda.v1alpha1.datahub.resources.ObjectMeta\x12Y\n\x12predicted_raw_data\x18\x02 \x03(\x0b\x32=.containersai.alameda.v1alpha1.datahub.predictions.MetricData\x12`\n\x19predicted_upperbound_data\x18\x03 \x03(\x0b\x32=.containersai.alameda.v1alpha1.datahub.predictions.MetricData\x12`\n\x19predicted_lowerbound_data\x18\x04 \x03(\x0b\x32=.containersai.alameda.v1alpha1.datahub.predictions.MetricData\"\x88\x02\n\x0fWritePrediction\x12M\n\x0bmetric_type\x18\x01 \x01(\x0e\x32\x38.containersai.alameda.v1alpha1.datahub.common.MetricType\x12Y\n\x11resource_boundary\x18\x02 \x01(\x0e\x32>.containersai.alameda.v1alpha1.datahub.common.ResourceBoundary\x12K\n\nwrite_data\x18\x03 \x01(\x0b\x32\x37.containersai.alameda.v1alpha1.datahub.common.WriteDataBGZEgithub.com/containers-ai/api/alameda_api/v1alpha1/datahub/predictionsb\x06proto3'
  ,
  dependencies=[alameda__api_dot_v1alpha1_dot_datahub_dot_common_dot_metrics__pb2.DESCRIPTOR,alameda__api_dot_v1alpha1_dot_datahub_dot_common_dot_rawdata__pb2.DESCRIPTOR,alameda__api_dot_v1alpha1_dot_datahub_dot_common_dot_types__pb2.DESCRIPTOR,alameda__api_dot_v1alpha1_dot_datahub_dot_predictions_dot_types__pb2.DESCRIPTOR,alameda__api_dot_v1alpha1_dot_datahub_dot_resources_dot_metadata__pb2.DESCRIPTOR,])




_CONTAINERPREDICTION = _descriptor.Descriptor(
  name='ContainerPrediction',
  full_name='containersai.alameda.v1alpha1.datahub.predictions.ContainerPrediction',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='containersai.alameda.v1alpha1.datahub.predictions.ContainerPrediction.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='predicted_raw_data', full_name='containersai.alameda.v1alpha1.datahub.predictions.ContainerPrediction.predicted_raw_data', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='predicted_upperbound_data', full_name='containersai.alameda.v1alpha1.datahub.predictions.ContainerPrediction.predicted_upperbound_data', index=2,
      number=3, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='predicted_lowerbound_data', full_name='containersai.alameda.v1alpha1.datahub.predictions.ContainerPrediction.predicted_lowerbound_data', index=3,
      number=4, type=11, cpp_type=10, label=3,
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
  serialized_start=374,
  serialized_end=696,
)


_PODPREDICTION = _descriptor.Descriptor(
  name='PodPrediction',
  full_name='containersai.alameda.v1alpha1.datahub.predictions.PodPrediction',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='object_meta', full_name='containersai.alameda.v1alpha1.datahub.predictions.PodPrediction.object_meta', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='container_predictions', full_name='containersai.alameda.v1alpha1.datahub.predictions.PodPrediction.container_predictions', index=1,
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
  serialized_start=699,
  serialized_end=899,
)


_CONTROLLERPREDICTION = _descriptor.Descriptor(
  name='ControllerPrediction',
  full_name='containersai.alameda.v1alpha1.datahub.predictions.ControllerPrediction',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='object_meta', full_name='containersai.alameda.v1alpha1.datahub.predictions.ControllerPrediction.object_meta', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='kind', full_name='containersai.alameda.v1alpha1.datahub.predictions.ControllerPrediction.kind', index=1,
      number=2, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='predicted_raw_data', full_name='containersai.alameda.v1alpha1.datahub.predictions.ControllerPrediction.predicted_raw_data', index=2,
      number=3, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='predicted_upperbound_data', full_name='containersai.alameda.v1alpha1.datahub.predictions.ControllerPrediction.predicted_upperbound_data', index=3,
      number=4, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='predicted_lowerbound_data', full_name='containersai.alameda.v1alpha1.datahub.predictions.ControllerPrediction.predicted_lowerbound_data', index=4,
      number=5, type=11, cpp_type=10, label=3,
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
  serialized_start=902,
  serialized_end=1362,
)


_APPLICATIONPREDICTION = _descriptor.Descriptor(
  name='ApplicationPrediction',
  full_name='containersai.alameda.v1alpha1.datahub.predictions.ApplicationPrediction',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='object_meta', full_name='containersai.alameda.v1alpha1.datahub.predictions.ApplicationPrediction.object_meta', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='predicted_raw_data', full_name='containersai.alameda.v1alpha1.datahub.predictions.ApplicationPrediction.predicted_raw_data', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='predicted_upperbound_data', full_name='containersai.alameda.v1alpha1.datahub.predictions.ApplicationPrediction.predicted_upperbound_data', index=2,
      number=3, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='predicted_lowerbound_data', full_name='containersai.alameda.v1alpha1.datahub.predictions.ApplicationPrediction.predicted_lowerbound_data', index=3,
      number=4, type=11, cpp_type=10, label=3,
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
  serialized_start=1365,
  serialized_end=1757,
)


_NAMESPACEPREDICTION = _descriptor.Descriptor(
  name='NamespacePrediction',
  full_name='containersai.alameda.v1alpha1.datahub.predictions.NamespacePrediction',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='object_meta', full_name='containersai.alameda.v1alpha1.datahub.predictions.NamespacePrediction.object_meta', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='predicted_raw_data', full_name='containersai.alameda.v1alpha1.datahub.predictions.NamespacePrediction.predicted_raw_data', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='predicted_upperbound_data', full_name='containersai.alameda.v1alpha1.datahub.predictions.NamespacePrediction.predicted_upperbound_data', index=2,
      number=3, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='predicted_lowerbound_data', full_name='containersai.alameda.v1alpha1.datahub.predictions.NamespacePrediction.predicted_lowerbound_data', index=3,
      number=4, type=11, cpp_type=10, label=3,
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
  serialized_start=1760,
  serialized_end=2150,
)


_NODEPREDICTION = _descriptor.Descriptor(
  name='NodePrediction',
  full_name='containersai.alameda.v1alpha1.datahub.predictions.NodePrediction',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='object_meta', full_name='containersai.alameda.v1alpha1.datahub.predictions.NodePrediction.object_meta', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='is_scheduled', full_name='containersai.alameda.v1alpha1.datahub.predictions.NodePrediction.is_scheduled', index=1,
      number=2, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='predicted_raw_data', full_name='containersai.alameda.v1alpha1.datahub.predictions.NodePrediction.predicted_raw_data', index=2,
      number=3, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='predicted_upperbound_data', full_name='containersai.alameda.v1alpha1.datahub.predictions.NodePrediction.predicted_upperbound_data', index=3,
      number=4, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='predicted_lowerbound_data', full_name='containersai.alameda.v1alpha1.datahub.predictions.NodePrediction.predicted_lowerbound_data', index=4,
      number=5, type=11, cpp_type=10, label=3,
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
  serialized_start=2153,
  serialized_end=2560,
)


_CLUSTERPREDICTION = _descriptor.Descriptor(
  name='ClusterPrediction',
  full_name='containersai.alameda.v1alpha1.datahub.predictions.ClusterPrediction',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='object_meta', full_name='containersai.alameda.v1alpha1.datahub.predictions.ClusterPrediction.object_meta', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='predicted_raw_data', full_name='containersai.alameda.v1alpha1.datahub.predictions.ClusterPrediction.predicted_raw_data', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='predicted_upperbound_data', full_name='containersai.alameda.v1alpha1.datahub.predictions.ClusterPrediction.predicted_upperbound_data', index=2,
      number=3, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='predicted_lowerbound_data', full_name='containersai.alameda.v1alpha1.datahub.predictions.ClusterPrediction.predicted_lowerbound_data', index=3,
      number=4, type=11, cpp_type=10, label=3,
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
  serialized_start=2563,
  serialized_end=2951,
)


_WRITEPREDICTION = _descriptor.Descriptor(
  name='WritePrediction',
  full_name='containersai.alameda.v1alpha1.datahub.predictions.WritePrediction',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='metric_type', full_name='containersai.alameda.v1alpha1.datahub.predictions.WritePrediction.metric_type', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='resource_boundary', full_name='containersai.alameda.v1alpha1.datahub.predictions.WritePrediction.resource_boundary', index=1,
      number=2, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='write_data', full_name='containersai.alameda.v1alpha1.datahub.predictions.WritePrediction.write_data', index=2,
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
  serialized_start=2954,
  serialized_end=3218,
)

_CONTAINERPREDICTION.fields_by_name['predicted_raw_data'].message_type = alameda__api_dot_v1alpha1_dot_datahub_dot_predictions_dot_types__pb2._METRICDATA
_CONTAINERPREDICTION.fields_by_name['predicted_upperbound_data'].message_type = alameda__api_dot_v1alpha1_dot_datahub_dot_predictions_dot_types__pb2._METRICDATA
_CONTAINERPREDICTION.fields_by_name['predicted_lowerbound_data'].message_type = alameda__api_dot_v1alpha1_dot_datahub_dot_predictions_dot_types__pb2._METRICDATA
_PODPREDICTION.fields_by_name['object_meta'].message_type = alameda__api_dot_v1alpha1_dot_datahub_dot_resources_dot_metadata__pb2._OBJECTMETA
_PODPREDICTION.fields_by_name['container_predictions'].message_type = _CONTAINERPREDICTION
_CONTROLLERPREDICTION.fields_by_name['object_meta'].message_type = alameda__api_dot_v1alpha1_dot_datahub_dot_resources_dot_metadata__pb2._OBJECTMETA
_CONTROLLERPREDICTION.fields_by_name['kind'].enum_type = alameda__api_dot_v1alpha1_dot_datahub_dot_resources_dot_metadata__pb2._KIND
_CONTROLLERPREDICTION.fields_by_name['predicted_raw_data'].message_type = alameda__api_dot_v1alpha1_dot_datahub_dot_predictions_dot_types__pb2._METRICDATA
_CONTROLLERPREDICTION.fields_by_name['predicted_upperbound_data'].message_type = alameda__api_dot_v1alpha1_dot_datahub_dot_predictions_dot_types__pb2._METRICDATA
_CONTROLLERPREDICTION.fields_by_name['predicted_lowerbound_data'].message_type = alameda__api_dot_v1alpha1_dot_datahub_dot_predictions_dot_types__pb2._METRICDATA
_APPLICATIONPREDICTION.fields_by_name['object_meta'].message_type = alameda__api_dot_v1alpha1_dot_datahub_dot_resources_dot_metadata__pb2._OBJECTMETA
_APPLICATIONPREDICTION.fields_by_name['predicted_raw_data'].message_type = alameda__api_dot_v1alpha1_dot_datahub_dot_predictions_dot_types__pb2._METRICDATA
_APPLICATIONPREDICTION.fields_by_name['predicted_upperbound_data'].message_type = alameda__api_dot_v1alpha1_dot_datahub_dot_predictions_dot_types__pb2._METRICDATA
_APPLICATIONPREDICTION.fields_by_name['predicted_lowerbound_data'].message_type = alameda__api_dot_v1alpha1_dot_datahub_dot_predictions_dot_types__pb2._METRICDATA
_NAMESPACEPREDICTION.fields_by_name['object_meta'].message_type = alameda__api_dot_v1alpha1_dot_datahub_dot_resources_dot_metadata__pb2._OBJECTMETA
_NAMESPACEPREDICTION.fields_by_name['predicted_raw_data'].message_type = alameda__api_dot_v1alpha1_dot_datahub_dot_predictions_dot_types__pb2._METRICDATA
_NAMESPACEPREDICTION.fields_by_name['predicted_upperbound_data'].message_type = alameda__api_dot_v1alpha1_dot_datahub_dot_predictions_dot_types__pb2._METRICDATA
_NAMESPACEPREDICTION.fields_by_name['predicted_lowerbound_data'].message_type = alameda__api_dot_v1alpha1_dot_datahub_dot_predictions_dot_types__pb2._METRICDATA
_NODEPREDICTION.fields_by_name['object_meta'].message_type = alameda__api_dot_v1alpha1_dot_datahub_dot_resources_dot_metadata__pb2._OBJECTMETA
_NODEPREDICTION.fields_by_name['predicted_raw_data'].message_type = alameda__api_dot_v1alpha1_dot_datahub_dot_predictions_dot_types__pb2._METRICDATA
_NODEPREDICTION.fields_by_name['predicted_upperbound_data'].message_type = alameda__api_dot_v1alpha1_dot_datahub_dot_predictions_dot_types__pb2._METRICDATA
_NODEPREDICTION.fields_by_name['predicted_lowerbound_data'].message_type = alameda__api_dot_v1alpha1_dot_datahub_dot_predictions_dot_types__pb2._METRICDATA
_CLUSTERPREDICTION.fields_by_name['object_meta'].message_type = alameda__api_dot_v1alpha1_dot_datahub_dot_resources_dot_metadata__pb2._OBJECTMETA
_CLUSTERPREDICTION.fields_by_name['predicted_raw_data'].message_type = alameda__api_dot_v1alpha1_dot_datahub_dot_predictions_dot_types__pb2._METRICDATA
_CLUSTERPREDICTION.fields_by_name['predicted_upperbound_data'].message_type = alameda__api_dot_v1alpha1_dot_datahub_dot_predictions_dot_types__pb2._METRICDATA
_CLUSTERPREDICTION.fields_by_name['predicted_lowerbound_data'].message_type = alameda__api_dot_v1alpha1_dot_datahub_dot_predictions_dot_types__pb2._METRICDATA
_WRITEPREDICTION.fields_by_name['metric_type'].enum_type = alameda__api_dot_v1alpha1_dot_datahub_dot_common_dot_metrics__pb2._METRICTYPE
_WRITEPREDICTION.fields_by_name['resource_boundary'].enum_type = alameda__api_dot_v1alpha1_dot_datahub_dot_common_dot_types__pb2._RESOURCEBOUNDARY
_WRITEPREDICTION.fields_by_name['write_data'].message_type = alameda__api_dot_v1alpha1_dot_datahub_dot_common_dot_rawdata__pb2._WRITEDATA
DESCRIPTOR.message_types_by_name['ContainerPrediction'] = _CONTAINERPREDICTION
DESCRIPTOR.message_types_by_name['PodPrediction'] = _PODPREDICTION
DESCRIPTOR.message_types_by_name['ControllerPrediction'] = _CONTROLLERPREDICTION
DESCRIPTOR.message_types_by_name['ApplicationPrediction'] = _APPLICATIONPREDICTION
DESCRIPTOR.message_types_by_name['NamespacePrediction'] = _NAMESPACEPREDICTION
DESCRIPTOR.message_types_by_name['NodePrediction'] = _NODEPREDICTION
DESCRIPTOR.message_types_by_name['ClusterPrediction'] = _CLUSTERPREDICTION
DESCRIPTOR.message_types_by_name['WritePrediction'] = _WRITEPREDICTION
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

ContainerPrediction = _reflection.GeneratedProtocolMessageType('ContainerPrediction', (_message.Message,), {
  'DESCRIPTOR' : _CONTAINERPREDICTION,
  '__module__' : 'alameda_api.v1alpha1.datahub.predictions.predictions_pb2'
  # @@protoc_insertion_point(class_scope:containersai.alameda.v1alpha1.datahub.predictions.ContainerPrediction)
  })
_sym_db.RegisterMessage(ContainerPrediction)

PodPrediction = _reflection.GeneratedProtocolMessageType('PodPrediction', (_message.Message,), {
  'DESCRIPTOR' : _PODPREDICTION,
  '__module__' : 'alameda_api.v1alpha1.datahub.predictions.predictions_pb2'
  # @@protoc_insertion_point(class_scope:containersai.alameda.v1alpha1.datahub.predictions.PodPrediction)
  })
_sym_db.RegisterMessage(PodPrediction)

ControllerPrediction = _reflection.GeneratedProtocolMessageType('ControllerPrediction', (_message.Message,), {
  'DESCRIPTOR' : _CONTROLLERPREDICTION,
  '__module__' : 'alameda_api.v1alpha1.datahub.predictions.predictions_pb2'
  # @@protoc_insertion_point(class_scope:containersai.alameda.v1alpha1.datahub.predictions.ControllerPrediction)
  })
_sym_db.RegisterMessage(ControllerPrediction)

ApplicationPrediction = _reflection.GeneratedProtocolMessageType('ApplicationPrediction', (_message.Message,), {
  'DESCRIPTOR' : _APPLICATIONPREDICTION,
  '__module__' : 'alameda_api.v1alpha1.datahub.predictions.predictions_pb2'
  # @@protoc_insertion_point(class_scope:containersai.alameda.v1alpha1.datahub.predictions.ApplicationPrediction)
  })
_sym_db.RegisterMessage(ApplicationPrediction)

NamespacePrediction = _reflection.GeneratedProtocolMessageType('NamespacePrediction', (_message.Message,), {
  'DESCRIPTOR' : _NAMESPACEPREDICTION,
  '__module__' : 'alameda_api.v1alpha1.datahub.predictions.predictions_pb2'
  # @@protoc_insertion_point(class_scope:containersai.alameda.v1alpha1.datahub.predictions.NamespacePrediction)
  })
_sym_db.RegisterMessage(NamespacePrediction)

NodePrediction = _reflection.GeneratedProtocolMessageType('NodePrediction', (_message.Message,), {
  'DESCRIPTOR' : _NODEPREDICTION,
  '__module__' : 'alameda_api.v1alpha1.datahub.predictions.predictions_pb2'
  # @@protoc_insertion_point(class_scope:containersai.alameda.v1alpha1.datahub.predictions.NodePrediction)
  })
_sym_db.RegisterMessage(NodePrediction)

ClusterPrediction = _reflection.GeneratedProtocolMessageType('ClusterPrediction', (_message.Message,), {
  'DESCRIPTOR' : _CLUSTERPREDICTION,
  '__module__' : 'alameda_api.v1alpha1.datahub.predictions.predictions_pb2'
  # @@protoc_insertion_point(class_scope:containersai.alameda.v1alpha1.datahub.predictions.ClusterPrediction)
  })
_sym_db.RegisterMessage(ClusterPrediction)

WritePrediction = _reflection.GeneratedProtocolMessageType('WritePrediction', (_message.Message,), {
  'DESCRIPTOR' : _WRITEPREDICTION,
  '__module__' : 'alameda_api.v1alpha1.datahub.predictions.predictions_pb2'
  # @@protoc_insertion_point(class_scope:containersai.alameda.v1alpha1.datahub.predictions.WritePrediction)
  })
_sym_db.RegisterMessage(WritePrediction)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)
