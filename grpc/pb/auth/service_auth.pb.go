// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.3
// source: auth/service_auth.proto

package auth

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_auth_service_auth_proto protoreflect.FileDescriptor

var file_auth_service_auth_proto_rawDesc = []byte{
	0x0a, 0x17, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x61, 0x75, 0x74, 0x68, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70,
	0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x61,
	0x75, 0x74, 0x68, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x61, 0x75,
	0x74, 0x68, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x72, 0x70,
	0x63, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x5f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x29, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x84, 0x07,
	0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xdc, 0x01,
	0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x12, 0x25, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x26, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x71, 0x92, 0x41, 0x4a, 0x12, 0x1a,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x20,
	0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x1a, 0x2c, 0x55, 0x73, 0x65, 0x20,
	0x74, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x6f, 0x20, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x20, 0x61, 0x20, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x20, 0x63, 0x72,
	0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x3a, 0x01,
	0x2a, 0x22, 0x19, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73,
	0x2f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x12, 0xde, 0x01, 0x0a,
	0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x43,
	0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x12, 0x25, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x43,
	0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x26, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x73, 0x92, 0x41, 0x4c, 0x12, 0x1a, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x20, 0x63,
	0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x1a, 0x2e, 0x55, 0x73, 0x65, 0x20, 0x74,
	0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x6f, 0x20, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x20, 0x74, 0x68, 0x65, 0x20, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x20, 0x63,
	0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x3a,
	0x01, 0x2a, 0x32, 0x19, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x73, 0x2f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x12, 0xac, 0x01,
	0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12,
	0x1a, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x62, 0x92, 0x41, 0x41, 0x12, 0x0e, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x20, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x1a, 0x2f, 0x55,
	0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x6f, 0x20, 0x6c,
	0x6f, 0x67, 0x69, 0x6e, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x67, 0x65,
	0x74, 0x20, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x20, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x18, 0x3a, 0x01, 0x2a, 0x22, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x85, 0x02, 0x0a,
	0x17, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x54, 0x65, 0x73, 0x74, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x24, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x54, 0x65, 0x73, 0x74, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x54, 0x65,
	0x73, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x9c, 0x01, 0x92, 0x41, 0x73, 0x12, 0x21, 0x47, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x20, 0x62, 0x61, 0x6e, 0x6b, 0x65, 0x72, 0x20, 0x72, 0x6f, 0x6c,
	0x65, 0x20, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x20, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x4e,
	0x55, 0x73, 0x65, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x6f, 0x20,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x20, 0x61, 0x6e, 0x20, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x20, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x20, 0x77, 0x69, 0x74, 0x68, 0x20, 0x62, 0x61,
	0x6e, 0x6b, 0x65, 0x72, 0x20, 0x72, 0x6f, 0x6c, 0x65, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x74, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x67, 0x20, 0x70, 0x75, 0x72, 0x70, 0x6f, 0x73, 0x65, 0x73, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x20, 0x12, 0x1e, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x82, 0x01, 0x92, 0x41, 0x55, 0x12, 0x53, 0x0a, 0x0c, 0x43, 0x6f,
	0x72, 0x65, 0x20, 0x42, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x22, 0x3e, 0x0a, 0x0d, 0x4e, 0x67,
	0x75, 0x79, 0x65, 0x6e, 0x20, 0x56, 0x61, 0x6e, 0x20, 0x54, 0x75, 0x12, 0x18, 0x68, 0x74, 0x74,
	0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x74, 0x75, 0x6e, 0x76, 0x78, 0x1a, 0x13, 0x74, 0x75, 0x6e, 0x76, 0x2e, 0x74, 0x6f, 0x64, 0x6f,
	0x40, 0x67, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x63, 0x6f, 0x6d, 0x32, 0x03, 0x31, 0x2e, 0x31, 0x5a,
	0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x75, 0x6e, 0x76,
	0x78, 0x2f, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x62, 0x61, 0x6e, 0x6b, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x2f, 0x70, 0x62, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var file_auth_service_auth_proto_goTypes = []any{
	(*CreateCustomerCredentialRequest)(nil),  // 0: auth.CreateCustomerCredentialRequest
	(*UpdateCustomerCredentialRequest)(nil),  // 1: auth.UpdateCustomerCredentialRequest
	(*LoginCustomerRequest)(nil),             // 2: auth.LoginCustomerRequest
	(*GenerateTestAccessTokenRequest)(nil),   // 3: auth.GenerateTestAccessTokenRequest
	(*CreateCustomerCredentialResponse)(nil), // 4: auth.CreateCustomerCredentialResponse
	(*UpdateCustomerCredentialResponse)(nil), // 5: auth.UpdateCustomerCredentialResponse
	(*LoginCustomerResponse)(nil),            // 6: auth.LoginCustomerResponse
	(*GenerateTestAccessTokenResponse)(nil),  // 7: auth.GenerateTestAccessTokenResponse
}
var file_auth_service_auth_proto_depIdxs = []int32{
	0, // 0: auth.AuthService.CreateCustomerCredential:input_type -> auth.CreateCustomerCredentialRequest
	1, // 1: auth.AuthService.UpdateCustomerCredential:input_type -> auth.UpdateCustomerCredentialRequest
	2, // 2: auth.AuthService.LoginCustomer:input_type -> auth.LoginCustomerRequest
	3, // 3: auth.AuthService.GenerateTestAccessToken:input_type -> auth.GenerateTestAccessTokenRequest
	4, // 4: auth.AuthService.CreateCustomerCredential:output_type -> auth.CreateCustomerCredentialResponse
	5, // 5: auth.AuthService.UpdateCustomerCredential:output_type -> auth.UpdateCustomerCredentialResponse
	6, // 6: auth.AuthService.LoginCustomer:output_type -> auth.LoginCustomerResponse
	7, // 7: auth.AuthService.GenerateTestAccessToken:output_type -> auth.GenerateTestAccessTokenResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_auth_service_auth_proto_init() }
func file_auth_service_auth_proto_init() {
	if File_auth_service_auth_proto != nil {
		return
	}
	file_auth_rpc_login_customer_proto_init()
	file_auth_rpc_create_customer_credential_proto_init()
	file_auth_rpc_update_customer_credential_proto_init()
	file_auth_rpc_generate_test_access_token_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_auth_service_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_auth_service_auth_proto_goTypes,
		DependencyIndexes: file_auth_service_auth_proto_depIdxs,
	}.Build()
	File_auth_service_auth_proto = out.File
	file_auth_service_auth_proto_rawDesc = nil
	file_auth_service_auth_proto_goTypes = nil
	file_auth_service_auth_proto_depIdxs = nil
}
