# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: witness.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


import method_pb2 as method__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='witness.proto',
  package='api_spec',
  syntax='proto3',
  serialized_options=b'Z*akitasoftware.com/superstar/pb/go/api_spec',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\rwitness.proto\x12\x08\x61pi_spec\x1a\x0cmethod.proto\"+\n\x07Witness\x12 \n\x06method\x18\x01 \x01(\x0b\x32\x10.api_spec.MethodB,Z*akitasoftware.com/superstar/pb/go/api_specb\x06proto3'
  ,
  dependencies=[method__pb2.DESCRIPTOR,])




_WITNESS = _descriptor.Descriptor(
  name='Witness',
  full_name='api_spec.Witness',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='method', full_name='api_spec.Witness.method', index=0,
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
  serialized_start=41,
  serialized_end=84,
)

_WITNESS.fields_by_name['method'].message_type = method__pb2._METHOD
DESCRIPTOR.message_types_by_name['Witness'] = _WITNESS
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Witness = _reflection.GeneratedProtocolMessageType('Witness', (_message.Message,), {
  'DESCRIPTOR' : _WITNESS,
  '__module__' : 'witness_pb2'
  # @@protoc_insertion_point(class_scope:api_spec.Witness)
  })
_sym_db.RegisterMessage(Witness)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)
