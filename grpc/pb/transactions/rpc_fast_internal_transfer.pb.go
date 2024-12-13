// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: transactions/rpc_fast_internal_transfer.proto

package transactions

import (
	account "github.com/tunvx/simplebank/grpc/pb/manage/account"
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

type FastInternalTransferRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount             int64  `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
	SenderAccNumber    string `protobuf:"bytes,2,opt,name=sender_acc_number,json=senderAccNumber,proto3" json:"sender_acc_number,omitempty"`
	RecipientBankCode  string `protobuf:"bytes,3,opt,name=recipient_bank_code,json=recipientBankCode,proto3" json:"recipient_bank_code,omitempty"`
	RecipientAccNumber string `protobuf:"bytes,4,opt,name=recipient_acc_number,json=recipientAccNumber,proto3" json:"recipient_acc_number,omitempty"`
	RecipientName      string `protobuf:"bytes,5,opt,name=recipient_name,json=recipientName,proto3" json:"recipient_name,omitempty"`
	CurrencyType       string `protobuf:"bytes,6,opt,name=currency_type,json=currencyType,proto3" json:"currency_type,omitempty"`
	Message            string `protobuf:"bytes,7,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *FastInternalTransferRequest) Reset() {
	*x = FastInternalTransferRequest{}
	mi := &file_transactions_rpc_fast_internal_transfer_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FastInternalTransferRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FastInternalTransferRequest) ProtoMessage() {}

func (x *FastInternalTransferRequest) ProtoReflect() protoreflect.Message {
	mi := &file_transactions_rpc_fast_internal_transfer_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FastInternalTransferRequest.ProtoReflect.Descriptor instead.
func (*FastInternalTransferRequest) Descriptor() ([]byte, []int) {
	return file_transactions_rpc_fast_internal_transfer_proto_rawDescGZIP(), []int{0}
}

func (x *FastInternalTransferRequest) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *FastInternalTransferRequest) GetSenderAccNumber() string {
	if x != nil {
		return x.SenderAccNumber
	}
	return ""
}

func (x *FastInternalTransferRequest) GetRecipientBankCode() string {
	if x != nil {
		return x.RecipientBankCode
	}
	return ""
}

func (x *FastInternalTransferRequest) GetRecipientAccNumber() string {
	if x != nil {
		return x.RecipientAccNumber
	}
	return ""
}

func (x *FastInternalTransferRequest) GetRecipientName() string {
	if x != nil {
		return x.RecipientName
	}
	return ""
}

func (x *FastInternalTransferRequest) GetCurrencyType() string {
	if x != nil {
		return x.CurrencyType
	}
	return ""
}

func (x *FastInternalTransferRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type FastInternalTransferResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SenderAccount    *account.Account `protobuf:"bytes,1,opt,name=sender_account,json=senderAccount,proto3" json:"sender_account,omitempty"`
	RecipientAccount *account.Account `protobuf:"bytes,2,opt,name=recipient_account,json=recipientAccount,proto3" json:"recipient_account,omitempty"`
}

func (x *FastInternalTransferResponse) Reset() {
	*x = FastInternalTransferResponse{}
	mi := &file_transactions_rpc_fast_internal_transfer_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FastInternalTransferResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FastInternalTransferResponse) ProtoMessage() {}

func (x *FastInternalTransferResponse) ProtoReflect() protoreflect.Message {
	mi := &file_transactions_rpc_fast_internal_transfer_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FastInternalTransferResponse.ProtoReflect.Descriptor instead.
func (*FastInternalTransferResponse) Descriptor() ([]byte, []int) {
	return file_transactions_rpc_fast_internal_transfer_proto_rawDescGZIP(), []int{1}
}

func (x *FastInternalTransferResponse) GetSenderAccount() *account.Account {
	if x != nil {
		return x.SenderAccount
	}
	return nil
}

func (x *FastInternalTransferResponse) GetRecipientAccount() *account.Account {
	if x != nil {
		return x.RecipientAccount
	}
	return nil
}

var File_transactions_rpc_fast_internal_transfer_proto protoreflect.FileDescriptor

var file_transactions_rpc_fast_internal_transfer_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72,
	0x70, 0x63, 0x5f, 0x66, 0x61, 0x73, 0x74, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1c, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa9, 0x02, 0x0a, 0x1b, 0x46,
	0x61, 0x73, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x11, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x61, 0x63, 0x63,
	0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x41, 0x63, 0x63, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2e,
	0x0a, 0x13, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x62, 0x61, 0x6e, 0x6b,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x72, 0x65, 0x63,
	0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x42, 0x61, 0x6e, 0x6b, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x30,
	0x0a, 0x14, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x63, 0x63, 0x5f,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x72, 0x65,
	0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x41, 0x63, 0x63, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69,
	0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x96, 0x01, 0x0a, 0x1c, 0x46, 0x61, 0x73, 0x74, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x0e, 0x73, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x52, 0x0d, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x3d, 0x0a, 0x11, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x10, 0x72,
	0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42,
	0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x75,
	0x6e, 0x76, 0x78, 0x2f, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x62, 0x61, 0x6e, 0x6b, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x70, 0x62, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_transactions_rpc_fast_internal_transfer_proto_rawDescOnce sync.Once
	file_transactions_rpc_fast_internal_transfer_proto_rawDescData = file_transactions_rpc_fast_internal_transfer_proto_rawDesc
)

func file_transactions_rpc_fast_internal_transfer_proto_rawDescGZIP() []byte {
	file_transactions_rpc_fast_internal_transfer_proto_rawDescOnce.Do(func() {
		file_transactions_rpc_fast_internal_transfer_proto_rawDescData = protoimpl.X.CompressGZIP(file_transactions_rpc_fast_internal_transfer_proto_rawDescData)
	})
	return file_transactions_rpc_fast_internal_transfer_proto_rawDescData
}

var file_transactions_rpc_fast_internal_transfer_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_transactions_rpc_fast_internal_transfer_proto_goTypes = []any{
	(*FastInternalTransferRequest)(nil),  // 0: transaction.FastInternalTransferRequest
	(*FastInternalTransferResponse)(nil), // 1: transaction.FastInternalTransferResponse
	(*account.Account)(nil),              // 2: account.Account
}
var file_transactions_rpc_fast_internal_transfer_proto_depIdxs = []int32{
	2, // 0: transaction.FastInternalTransferResponse.sender_account:type_name -> account.Account
	2, // 1: transaction.FastInternalTransferResponse.recipient_account:type_name -> account.Account
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_transactions_rpc_fast_internal_transfer_proto_init() }
func file_transactions_rpc_fast_internal_transfer_proto_init() {
	if File_transactions_rpc_fast_internal_transfer_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_transactions_rpc_fast_internal_transfer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_transactions_rpc_fast_internal_transfer_proto_goTypes,
		DependencyIndexes: file_transactions_rpc_fast_internal_transfer_proto_depIdxs,
		MessageInfos:      file_transactions_rpc_fast_internal_transfer_proto_msgTypes,
	}.Build()
	File_transactions_rpc_fast_internal_transfer_proto = out.File
	file_transactions_rpc_fast_internal_transfer_proto_rawDesc = nil
	file_transactions_rpc_fast_internal_transfer_proto_goTypes = nil
	file_transactions_rpc_fast_internal_transfer_proto_depIdxs = nil
}
