# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: accessgraph/v1alpha/query.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1f\x61\x63\x63\x65ssgraph/v1alpha/query.proto\x12\x13\x61\x63\x63\x65ssgraph.v1alpha\"\xef\x01\n\x04Node\x12\x0e\n\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n\x04kind\x18\x02 \x01(\tR\x04kind\x12\x19\n\x08sub_kind\x18\x03 \x01(\tR\x07subKind\x12\x12\n\x04name\x18\x04 \x01(\tR\x04name\x12=\n\x06labels\x18\x05 \x03(\x0b\x32%.accessgraph.v1alpha.Node.LabelsEntryR\x06labels\x12\x1a\n\x08hostname\x18\x06 \x01(\tR\x08hostname\x1a\x39\n\x0bLabelsEntry\x12\x10\n\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n\x05value\x18\x02 \x01(\tR\x05value:\x02\x38\x01\">\n\x04\x45\x64ge\x12\x12\n\x04\x66rom\x18\x01 \x01(\tR\x04\x66rom\x12\x0e\n\x02to\x18\x02 \x01(\tR\x02to\x12\x12\n\x04type\x18\x03 \x01(\tR\x04type\"$\n\x0cQueryRequest\x12\x14\n\x05query\x18\x01 \x01(\tR\x05query\"q\n\rQueryResponse\x12/\n\x05nodes\x18\x01 \x03(\x0b\x32\x19.accessgraph.v1alpha.NodeR\x05nodes\x12/\n\x05\x65\x64ges\x18\x02 \x03(\x0b\x32\x19.accessgraph.v1alpha.EdgeR\x05\x65\x64ges\",\n\x0eGetFileRequest\x12\x1a\n\x08\x66ilepath\x18\x01 \x01(\tR\x08\x66ilepath\"%\n\x0fGetFileResponse\x12\x12\n\x04\x64\x61ta\x18\x01 \x01(\x0cR\x04\x64\x61ta2\xba\x01\n\x12\x41\x63\x63\x65ssGraphService\x12N\n\x05Query\x12!.accessgraph.v1alpha.QueryRequest\x1a\".accessgraph.v1alpha.QueryResponse\x12T\n\x07GetFile\x12#.accessgraph.v1alpha.GetFileRequest\x1a$.accessgraph.v1alpha.GetFileResponseBWZUgithub.com/gravitational/teleport/gen/proto/go/accessgraph/v1alpha;accessgraphv1alphab\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'accessgraph.v1alpha.query_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'ZUgithub.com/gravitational/teleport/gen/proto/go/accessgraph/v1alpha;accessgraphv1alpha'
  _NODE_LABELSENTRY._options = None
  _NODE_LABELSENTRY._serialized_options = b'8\001'
  _NODE._serialized_start=57
  _NODE._serialized_end=296
  _NODE_LABELSENTRY._serialized_start=239
  _NODE_LABELSENTRY._serialized_end=296
  _EDGE._serialized_start=298
  _EDGE._serialized_end=360
  _QUERYREQUEST._serialized_start=362
  _QUERYREQUEST._serialized_end=398
  _QUERYRESPONSE._serialized_start=400
  _QUERYRESPONSE._serialized_end=513
  _GETFILEREQUEST._serialized_start=515
  _GETFILEREQUEST._serialized_end=559
  _GETFILERESPONSE._serialized_start=561
  _GETFILERESPONSE._serialized_end=598
  _ACCESSGRAPHSERVICE._serialized_start=601
  _ACCESSGRAPHSERVICE._serialized_end=787
# @@protoc_insertion_point(module_scope)
