// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.3
// source: auth/rpc_login_customer.proto

package auth

import (
	_ "github.com/tunvx/simplebank/grpc/pb/cusman/customer"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type LoginCustomerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LoginCustomerRequest) Reset() {
	*x = LoginCustomerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_rpc_login_customer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginCustomerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginCustomerRequest) ProtoMessage() {}

func (x *LoginCustomerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_rpc_login_customer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginCustomerRequest.ProtoReflect.Descriptor instead.
func (*LoginCustomerRequest) Descriptor() ([]byte, []int) {
	return file_auth_rpc_login_customer_proto_rawDescGZIP(), []int{0}
}

func (x *LoginCustomerRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *LoginCustomerRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginCustomerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId            int64                  `protobuf:"varint,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	ShardId               int32                  `protobuf:"varint,2,opt,name=shard_id,json=shardId,proto3" json:"shard_id,omitempty"`
	SessionId             string                 `protobuf:"bytes,3,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	AccessToken           string                 `protobuf:"bytes,4,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	RefreshToken          string                 `protobuf:"bytes,5,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	AccessTokenExpiresAt  *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=access_token_expires_at,json=accessTokenExpiresAt,proto3" json:"access_token_expires_at,omitempty"`
	RefreshTokenExpiresAt *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=refresh_token_expires_at,json=refreshTokenExpiresAt,proto3" json:"refresh_token_expires_at,omitempty"`
}

func (x *LoginCustomerResponse) Reset() {
	*x = LoginCustomerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_rpc_login_customer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginCustomerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginCustomerResponse) ProtoMessage() {}

func (x *LoginCustomerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_rpc_login_customer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginCustomerResponse.ProtoReflect.Descriptor instead.
func (*LoginCustomerResponse) Descriptor() ([]byte, []int) {
	return file_auth_rpc_login_customer_proto_rawDescGZIP(), []int{1}
}

func (x *LoginCustomerResponse) GetCustomerId() int64 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *LoginCustomerResponse) GetShardId() int32 {
	if x != nil {
		return x.ShardId
	}
	return 0
}

func (x *LoginCustomerResponse) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *LoginCustomerResponse) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *LoginCustomerResponse) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

func (x *LoginCustomerResponse) GetAccessTokenExpiresAt() *timestamppb.Timestamp {
	if x != nil {
		return x.AccessTokenExpiresAt
	}
	return nil
}

func (x *LoginCustomerResponse) GetRefreshTokenExpiresAt() *timestamppb.Timestamp {
	if x != nil {
		return x.RefreshTokenExpiresAt
	}
	return nil
}

var File_auth_rpc_login_customer_proto protoreflect.FileDescriptor

var file_auth_rpc_login_customer_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x6e,
	0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x04, 0x61, 0x75, 0x74, 0x68, 0x1a, 0x1e, 0x63, 0x75, 0x73, 0x6d, 0x61, 0x6e, 0x2f, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4e, 0x0a, 0x14, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0xe2, 0x02, 0x0a, 0x15, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x68, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x73, 0x68, 0x61, 0x72, 0x64, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x23,
	0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x51, 0x0a, 0x17, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x5f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x61, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x14, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x45, 0x78, 0x70,
	0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x12, 0x53, 0x0a, 0x18, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f,
	0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x15, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x42, 0x2a, 0x5a, 0x28, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x75, 0x6e, 0x76, 0x78, 0x2f,
	0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x62, 0x61, 0x6e, 0x6b, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f,
	0x70, 0x62, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_auth_rpc_login_customer_proto_rawDescOnce sync.Once
	file_auth_rpc_login_customer_proto_rawDescData = file_auth_rpc_login_customer_proto_rawDesc
)

func file_auth_rpc_login_customer_proto_rawDescGZIP() []byte {
	file_auth_rpc_login_customer_proto_rawDescOnce.Do(func() {
		file_auth_rpc_login_customer_proto_rawDescData = protoimpl.X.CompressGZIP(file_auth_rpc_login_customer_proto_rawDescData)
	})
	return file_auth_rpc_login_customer_proto_rawDescData
}

var file_auth_rpc_login_customer_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_auth_rpc_login_customer_proto_goTypes = []any{
	(*LoginCustomerRequest)(nil),  // 0: auth.LoginCustomerRequest
	(*LoginCustomerResponse)(nil), // 1: auth.LoginCustomerResponse
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_auth_rpc_login_customer_proto_depIdxs = []int32{
	2, // 0: auth.LoginCustomerResponse.access_token_expires_at:type_name -> google.protobuf.Timestamp
	2, // 1: auth.LoginCustomerResponse.refresh_token_expires_at:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_auth_rpc_login_customer_proto_init() }
func file_auth_rpc_login_customer_proto_init() {
	if File_auth_rpc_login_customer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_auth_rpc_login_customer_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*LoginCustomerRequest); i {
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
		file_auth_rpc_login_customer_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*LoginCustomerResponse); i {
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
			RawDescriptor: file_auth_rpc_login_customer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_auth_rpc_login_customer_proto_goTypes,
		DependencyIndexes: file_auth_rpc_login_customer_proto_depIdxs,
		MessageInfos:      file_auth_rpc_login_customer_proto_msgTypes,
	}.Build()
	File_auth_rpc_login_customer_proto = out.File
	file_auth_rpc_login_customer_proto_rawDesc = nil
	file_auth_rpc_login_customer_proto_goTypes = nil
	file_auth_rpc_login_customer_proto_depIdxs = nil
}
