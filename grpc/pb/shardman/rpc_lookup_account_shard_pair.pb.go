// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.3
// source: shardman/rpc_lookup_account_shard_pair.proto

package shardman

import (
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

type LookupAccountShardPairRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account1Id int64 `protobuf:"varint,1,opt,name=account1_id,json=account1Id,proto3" json:"account1_id,omitempty"`
	Account2Id int64 `protobuf:"varint,2,opt,name=account2_id,json=account2Id,proto3" json:"account2_id,omitempty"`
}

func (x *LookupAccountShardPairRequest) Reset() {
	*x = LookupAccountShardPairRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shardman_rpc_lookup_account_shard_pair_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LookupAccountShardPairRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LookupAccountShardPairRequest) ProtoMessage() {}

func (x *LookupAccountShardPairRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shardman_rpc_lookup_account_shard_pair_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LookupAccountShardPairRequest.ProtoReflect.Descriptor instead.
func (*LookupAccountShardPairRequest) Descriptor() ([]byte, []int) {
	return file_shardman_rpc_lookup_account_shard_pair_proto_rawDescGZIP(), []int{0}
}

func (x *LookupAccountShardPairRequest) GetAccount1Id() int64 {
	if x != nil {
		return x.Account1Id
	}
	return 0
}

func (x *LookupAccountShardPairRequest) GetAccount2Id() int64 {
	if x != nil {
		return x.Account2Id
	}
	return 0
}

type LookupAccountShardPairResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account1Id    int64 `protobuf:"varint,1,opt,name=account1_id,json=account1Id,proto3" json:"account1_id,omitempty"`
	Account1Shard int32 `protobuf:"varint,2,opt,name=account1_shard,json=account1Shard,proto3" json:"account1_shard,omitempty"`
	Account2Id    int64 `protobuf:"varint,3,opt,name=account2_id,json=account2Id,proto3" json:"account2_id,omitempty"`
	Account2Shard int32 `protobuf:"varint,4,opt,name=account2_shard,json=account2Shard,proto3" json:"account2_shard,omitempty"`
}

func (x *LookupAccountShardPairResponse) Reset() {
	*x = LookupAccountShardPairResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shardman_rpc_lookup_account_shard_pair_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LookupAccountShardPairResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LookupAccountShardPairResponse) ProtoMessage() {}

func (x *LookupAccountShardPairResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shardman_rpc_lookup_account_shard_pair_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LookupAccountShardPairResponse.ProtoReflect.Descriptor instead.
func (*LookupAccountShardPairResponse) Descriptor() ([]byte, []int) {
	return file_shardman_rpc_lookup_account_shard_pair_proto_rawDescGZIP(), []int{1}
}

func (x *LookupAccountShardPairResponse) GetAccount1Id() int64 {
	if x != nil {
		return x.Account1Id
	}
	return 0
}

func (x *LookupAccountShardPairResponse) GetAccount1Shard() int32 {
	if x != nil {
		return x.Account1Shard
	}
	return 0
}

func (x *LookupAccountShardPairResponse) GetAccount2Id() int64 {
	if x != nil {
		return x.Account2Id
	}
	return 0
}

func (x *LookupAccountShardPairResponse) GetAccount2Shard() int32 {
	if x != nil {
		return x.Account2Shard
	}
	return 0
}

var File_shardman_rpc_lookup_account_shard_pair_proto protoreflect.FileDescriptor

var file_shardman_rpc_lookup_account_shard_pair_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x73, 0x68, 0x61, 0x72, 0x64, 0x6d, 0x61, 0x6e, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x6c,
	0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x73, 0x68,
	0x61, 0x72, 0x64, 0x5f, 0x70, 0x61, 0x69, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08,
	0x73, 0x68, 0x61, 0x72, 0x64, 0x6d, 0x61, 0x6e, 0x22, 0x61, 0x0a, 0x1d, 0x4c, 0x6f, 0x6f, 0x6b,
	0x75, 0x70, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x68, 0x61, 0x72, 0x64, 0x50, 0x61,
	0x69, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x31, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x31, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0x49, 0x64, 0x22, 0xb0, 0x01, 0x0a, 0x1e,
	0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x68, 0x61,
	0x72, 0x64, 0x50, 0x61, 0x69, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f,
	0x0a, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x31, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x31, 0x49, 0x64, 0x12,
	0x25, 0x0a, 0x0e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x31, 0x5f, 0x73, 0x68, 0x61, 0x72,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x31, 0x53, 0x68, 0x61, 0x72, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x32, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x32, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x32, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0x53, 0x68, 0x61, 0x72, 0x64, 0x42, 0x2e,
	0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x75, 0x6e,
	0x76, 0x78, 0x2f, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x62, 0x61, 0x6e, 0x6b, 0x2f, 0x67, 0x72,
	0x70, 0x63, 0x2f, 0x70, 0x62, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x64, 0x6d, 0x61, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_shardman_rpc_lookup_account_shard_pair_proto_rawDescOnce sync.Once
	file_shardman_rpc_lookup_account_shard_pair_proto_rawDescData = file_shardman_rpc_lookup_account_shard_pair_proto_rawDesc
)

func file_shardman_rpc_lookup_account_shard_pair_proto_rawDescGZIP() []byte {
	file_shardman_rpc_lookup_account_shard_pair_proto_rawDescOnce.Do(func() {
		file_shardman_rpc_lookup_account_shard_pair_proto_rawDescData = protoimpl.X.CompressGZIP(file_shardman_rpc_lookup_account_shard_pair_proto_rawDescData)
	})
	return file_shardman_rpc_lookup_account_shard_pair_proto_rawDescData
}

var file_shardman_rpc_lookup_account_shard_pair_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_shardman_rpc_lookup_account_shard_pair_proto_goTypes = []any{
	(*LookupAccountShardPairRequest)(nil),  // 0: shardman.LookupAccountShardPairRequest
	(*LookupAccountShardPairResponse)(nil), // 1: shardman.LookupAccountShardPairResponse
}
var file_shardman_rpc_lookup_account_shard_pair_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_shardman_rpc_lookup_account_shard_pair_proto_init() }
func file_shardman_rpc_lookup_account_shard_pair_proto_init() {
	if File_shardman_rpc_lookup_account_shard_pair_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_shardman_rpc_lookup_account_shard_pair_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*LookupAccountShardPairRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_shardman_rpc_lookup_account_shard_pair_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*LookupAccountShardPairResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_shardman_rpc_lookup_account_shard_pair_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_shardman_rpc_lookup_account_shard_pair_proto_goTypes,
		DependencyIndexes: file_shardman_rpc_lookup_account_shard_pair_proto_depIdxs,
		MessageInfos:      file_shardman_rpc_lookup_account_shard_pair_proto_msgTypes,
	}.Build()
	File_shardman_rpc_lookup_account_shard_pair_proto = out.File
	file_shardman_rpc_lookup_account_shard_pair_proto_rawDesc = nil
	file_shardman_rpc_lookup_account_shard_pair_proto_goTypes = nil
	file_shardman_rpc_lookup_account_shard_pair_proto_depIdxs = nil
}
