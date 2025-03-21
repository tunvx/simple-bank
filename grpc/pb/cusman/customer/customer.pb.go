// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.3
// source: cusman/customer/customer.proto

package customer

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

type Customer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId       int64  `protobuf:"varint,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	CustomerRid      string `protobuf:"bytes,2,opt,name=customer_rid,json=customerRid,proto3" json:"customer_rid,omitempty"`
	FullName         string `protobuf:"bytes,3,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	DateOfBirth      string `protobuf:"bytes,4,opt,name=date_of_birth,json=dateOfBirth,proto3" json:"date_of_birth,omitempty"`
	PermanentAddress string `protobuf:"bytes,5,opt,name=permanent_address,json=permanentAddress,proto3" json:"permanent_address,omitempty"`
	PhoneNumber      string `protobuf:"bytes,6,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	EmailAddress     string `protobuf:"bytes,7,opt,name=email_address,json=emailAddress,proto3" json:"email_address,omitempty"`
	CustomerTier     string `protobuf:"bytes,8,opt,name=customer_tier,json=customerTier,proto3" json:"customer_tier,omitempty"`
	CustomerSegment  string `protobuf:"bytes,9,opt,name=customer_segment,json=customerSegment,proto3" json:"customer_segment,omitempty"`
	FinancialStatus  string `protobuf:"bytes,10,opt,name=financial_status,json=financialStatus,proto3" json:"financial_status,omitempty"`
}

func (x *Customer) Reset() {
	*x = Customer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cusman_customer_customer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Customer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Customer) ProtoMessage() {}

func (x *Customer) ProtoReflect() protoreflect.Message {
	mi := &file_cusman_customer_customer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Customer.ProtoReflect.Descriptor instead.
func (*Customer) Descriptor() ([]byte, []int) {
	return file_cusman_customer_customer_proto_rawDescGZIP(), []int{0}
}

func (x *Customer) GetCustomerId() int64 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *Customer) GetCustomerRid() string {
	if x != nil {
		return x.CustomerRid
	}
	return ""
}

func (x *Customer) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *Customer) GetDateOfBirth() string {
	if x != nil {
		return x.DateOfBirth
	}
	return ""
}

func (x *Customer) GetPermanentAddress() string {
	if x != nil {
		return x.PermanentAddress
	}
	return ""
}

func (x *Customer) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *Customer) GetEmailAddress() string {
	if x != nil {
		return x.EmailAddress
	}
	return ""
}

func (x *Customer) GetCustomerTier() string {
	if x != nil {
		return x.CustomerTier
	}
	return ""
}

func (x *Customer) GetCustomerSegment() string {
	if x != nil {
		return x.CustomerSegment
	}
	return ""
}

func (x *Customer) GetFinancialStatus() string {
	if x != nil {
		return x.FinancialStatus
	}
	return ""
}

var File_cusman_customer_customer_proto protoreflect.FileDescriptor

var file_cusman_customer_customer_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x63, 0x75, 0x73, 0x6d, 0x61, 0x6e, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x22, 0xff, 0x02, 0x0a, 0x08, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x72, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66,
	0x75, 0x6c, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x64, 0x61, 0x74, 0x65,
	0x5f, 0x6f, 0x66, 0x5f, 0x62, 0x69, 0x72, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x66, 0x42, 0x69, 0x72, 0x74, 0x68, 0x12, 0x2b, 0x0a, 0x11,
	0x70, 0x65, 0x72, 0x6d, 0x61, 0x6e, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x70, 0x65, 0x72, 0x6d, 0x61, 0x6e, 0x65,
	0x6e, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x0d,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x74, 0x69,
	0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x54, 0x69, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x5f, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x29, 0x0a, 0x10, 0x66, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x69, 0x61, 0x6c, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x66, 0x69, 0x6e,
	0x61, 0x6e, 0x63, 0x69, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x35, 0x5a, 0x33,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x75, 0x6e, 0x76, 0x78,
	0x2f, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x62, 0x61, 0x6e, 0x6b, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x2f, 0x70, 0x62, 0x2f, 0x63, 0x75, 0x73, 0x6d, 0x61, 0x6e, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cusman_customer_customer_proto_rawDescOnce sync.Once
	file_cusman_customer_customer_proto_rawDescData = file_cusman_customer_customer_proto_rawDesc
)

func file_cusman_customer_customer_proto_rawDescGZIP() []byte {
	file_cusman_customer_customer_proto_rawDescOnce.Do(func() {
		file_cusman_customer_customer_proto_rawDescData = protoimpl.X.CompressGZIP(file_cusman_customer_customer_proto_rawDescData)
	})
	return file_cusman_customer_customer_proto_rawDescData
}

var file_cusman_customer_customer_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_cusman_customer_customer_proto_goTypes = []any{
	(*Customer)(nil), // 0: customer.Customer
}
var file_cusman_customer_customer_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cusman_customer_customer_proto_init() }
func file_cusman_customer_customer_proto_init() {
	if File_cusman_customer_customer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cusman_customer_customer_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Customer); i {
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
			RawDescriptor: file_cusman_customer_customer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cusman_customer_customer_proto_goTypes,
		DependencyIndexes: file_cusman_customer_customer_proto_depIdxs,
		MessageInfos:      file_cusman_customer_customer_proto_msgTypes,
	}.Build()
	File_cusman_customer_customer_proto = out.File
	file_cusman_customer_customer_proto_rawDesc = nil
	file_cusman_customer_customer_proto_goTypes = nil
	file_cusman_customer_customer_proto_depIdxs = nil
}
