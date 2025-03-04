# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: network.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\rnetwork.proto\x12\x06pactus\"\x17\n\x15GetNetworkInfoRequest\"\xbb\x01\n\x16GetNetworkInfoResponse\x12(\n\x10total_sent_bytes\x18\x01 \x01(\x05R\x0etotalSentBytes\x12\x30\n\x14total_received_bytes\x18\x02 \x01(\x05R\x12totalReceivedBytes\x12\x1d\n\nstarted_at\x18\x03 \x01(\x03R\tstartedAt\x12&\n\x05peers\x18\x04 \x03(\x0b\x32\x10.pactus.PeerInfoR\x05peers\"\x14\n\x12GetNodeInfoRequest\"^\n\x13GetNodeInfoResponse\x12\x18\n\x07moniker\x18\x01 \x01(\tR\x07moniker\x12\x14\n\x05\x61gent\x18\x02 \x01(\tR\x05\x61gent\x12\x17\n\x07peer_id\x18\x03 \x01(\x0cR\x06peerId\"\xed\x03\n\x08PeerInfo\x12\x18\n\x07moniker\x18\x01 \x01(\tR\x07moniker\x12\x14\n\x05\x61gent\x18\x02 \x01(\tR\x05\x61gent\x12\x17\n\x07peer_id\x18\x03 \x01(\x0cR\x06peerId\x12%\n\x0e\x63onsensus_keys\x18\x04 \x03(\tR\rconsensusKeys\x12\x14\n\x05\x66lags\x18\x05 \x01(\x05R\x05\x66lags\x12\x16\n\x06height\x18\x06 \x01(\rR\x06height\x12+\n\x11received_messages\x18\x07 \x01(\x05R\x10receivedMessages\x12)\n\x10invalid_messages\x18\x08 \x01(\x05R\x0finvalidMessages\x12%\n\x0ereceived_bytes\x18\t \x01(\x05R\rreceivedBytes\x12\x16\n\x06status\x18\n \x01(\x05R\x06status\x12\x1b\n\tlast_sent\x18\x0b \x01(\x03R\x08lastSent\x12#\n\rlast_received\x18\x0c \x01(\x03R\x0clastReceived\x12!\n\x0csend_success\x18\r \x01(\x05R\x0bsendSuccess\x12\x1f\n\x0bsend_failed\x18\x0e \x01(\x05R\nsendFailed\x12&\n\x0flast_block_hash\x18\x0f \x01(\x0cR\rlastBlockHash2\xa2\x01\n\x07Network\x12O\n\x0eGetNetworkInfo\x12\x1d.pactus.GetNetworkInfoRequest\x1a\x1e.pactus.GetNetworkInfoResponse\x12\x46\n\x0bGetNodeInfo\x12\x1a.pactus.GetNodeInfoRequest\x1a\x1b.pactus.GetNodeInfoResponseBB\n\x0epactus.networkZ0github.com/pactus-project/pactus/www/grpc/pactusb\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'network_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n\016pactus.networkZ0github.com/pactus-project/pactus/www/grpc/pactus'
  _GETNETWORKINFOREQUEST._serialized_start=25
  _GETNETWORKINFOREQUEST._serialized_end=48
  _GETNETWORKINFORESPONSE._serialized_start=51
  _GETNETWORKINFORESPONSE._serialized_end=238
  _GETNODEINFOREQUEST._serialized_start=240
  _GETNODEINFOREQUEST._serialized_end=260
  _GETNODEINFORESPONSE._serialized_start=262
  _GETNODEINFORESPONSE._serialized_end=356
  _PEERINFO._serialized_start=359
  _PEERINFO._serialized_end=852
  _NETWORK._serialized_start=855
  _NETWORK._serialized_end=1017
# @@protoc_insertion_point(module_scope)
