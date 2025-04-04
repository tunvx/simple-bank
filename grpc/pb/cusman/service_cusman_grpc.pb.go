// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.3
// source: cusman/service_cusman.proto

package cusman

import (
	context "context"
	account "github.com/tunvx/simplebank/grpc/pb/cusman/account"
	customer "github.com/tunvx/simplebank/grpc/pb/cusman/customer"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	CustomerManagement_CreateCustomer_FullMethodName   = "/cusman.CustomerManagement/CreateCustomer"
	CustomerManagement_VerifyEmail_FullMethodName      = "/cusman.CustomerManagement/VerifyEmail"
	CustomerManagement_GetCustomerByID_FullMethodName  = "/cusman.CustomerManagement/GetCustomerByID"
	CustomerManagement_GetCustomerByRid_FullMethodName = "/cusman.CustomerManagement/GetCustomerByRid"
	CustomerManagement_CreateAccount_FullMethodName    = "/cusman.CustomerManagement/CreateAccount"
	CustomerManagement_GetAccountByID_FullMethodName   = "/cusman.CustomerManagement/GetAccountByID"
)

// CustomerManagementClient is the client API for CustomerManagement service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CustomerManagementClient interface {
	CreateCustomer(ctx context.Context, in *customer.CreateCustomerRequest, opts ...grpc.CallOption) (*customer.CreateCustomerResponse, error)
	VerifyEmail(ctx context.Context, in *customer.VerifyEmailRequest, opts ...grpc.CallOption) (*customer.VerifyEmailResponse, error)
	GetCustomerByID(ctx context.Context, in *customer.GetCustomerByIDRequest, opts ...grpc.CallOption) (*customer.GetCustomerByIDResponse, error)
	GetCustomerByRid(ctx context.Context, in *customer.GetCustomerByRidRequest, opts ...grpc.CallOption) (*customer.GetCustomerByRidResponse, error)
	CreateAccount(ctx context.Context, in *account.CreateAccountRequest, opts ...grpc.CallOption) (*account.CreateAccountResponse, error)
	GetAccountByID(ctx context.Context, in *account.GetAccountByIDRequest, opts ...grpc.CallOption) (*account.GetAccountByIDResponse, error)
}

type customerManagementClient struct {
	cc grpc.ClientConnInterface
}

func NewCustomerManagementClient(cc grpc.ClientConnInterface) CustomerManagementClient {
	return &customerManagementClient{cc}
}

func (c *customerManagementClient) CreateCustomer(ctx context.Context, in *customer.CreateCustomerRequest, opts ...grpc.CallOption) (*customer.CreateCustomerResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(customer.CreateCustomerResponse)
	err := c.cc.Invoke(ctx, CustomerManagement_CreateCustomer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerManagementClient) VerifyEmail(ctx context.Context, in *customer.VerifyEmailRequest, opts ...grpc.CallOption) (*customer.VerifyEmailResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(customer.VerifyEmailResponse)
	err := c.cc.Invoke(ctx, CustomerManagement_VerifyEmail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerManagementClient) GetCustomerByID(ctx context.Context, in *customer.GetCustomerByIDRequest, opts ...grpc.CallOption) (*customer.GetCustomerByIDResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(customer.GetCustomerByIDResponse)
	err := c.cc.Invoke(ctx, CustomerManagement_GetCustomerByID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerManagementClient) GetCustomerByRid(ctx context.Context, in *customer.GetCustomerByRidRequest, opts ...grpc.CallOption) (*customer.GetCustomerByRidResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(customer.GetCustomerByRidResponse)
	err := c.cc.Invoke(ctx, CustomerManagement_GetCustomerByRid_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerManagementClient) CreateAccount(ctx context.Context, in *account.CreateAccountRequest, opts ...grpc.CallOption) (*account.CreateAccountResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(account.CreateAccountResponse)
	err := c.cc.Invoke(ctx, CustomerManagement_CreateAccount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerManagementClient) GetAccountByID(ctx context.Context, in *account.GetAccountByIDRequest, opts ...grpc.CallOption) (*account.GetAccountByIDResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(account.GetAccountByIDResponse)
	err := c.cc.Invoke(ctx, CustomerManagement_GetAccountByID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerManagementServer is the server API for CustomerManagement service.
// All implementations must embed UnimplementedCustomerManagementServer
// for forward compatibility.
type CustomerManagementServer interface {
	CreateCustomer(context.Context, *customer.CreateCustomerRequest) (*customer.CreateCustomerResponse, error)
	VerifyEmail(context.Context, *customer.VerifyEmailRequest) (*customer.VerifyEmailResponse, error)
	GetCustomerByID(context.Context, *customer.GetCustomerByIDRequest) (*customer.GetCustomerByIDResponse, error)
	GetCustomerByRid(context.Context, *customer.GetCustomerByRidRequest) (*customer.GetCustomerByRidResponse, error)
	CreateAccount(context.Context, *account.CreateAccountRequest) (*account.CreateAccountResponse, error)
	GetAccountByID(context.Context, *account.GetAccountByIDRequest) (*account.GetAccountByIDResponse, error)
	mustEmbedUnimplementedCustomerManagementServer()
}

// UnimplementedCustomerManagementServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCustomerManagementServer struct{}

func (UnimplementedCustomerManagementServer) CreateCustomer(context.Context, *customer.CreateCustomerRequest) (*customer.CreateCustomerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCustomer not implemented")
}
func (UnimplementedCustomerManagementServer) VerifyEmail(context.Context, *customer.VerifyEmailRequest) (*customer.VerifyEmailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyEmail not implemented")
}
func (UnimplementedCustomerManagementServer) GetCustomerByID(context.Context, *customer.GetCustomerByIDRequest) (*customer.GetCustomerByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCustomerByID not implemented")
}
func (UnimplementedCustomerManagementServer) GetCustomerByRid(context.Context, *customer.GetCustomerByRidRequest) (*customer.GetCustomerByRidResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCustomerByRid not implemented")
}
func (UnimplementedCustomerManagementServer) CreateAccount(context.Context, *account.CreateAccountRequest) (*account.CreateAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccount not implemented")
}
func (UnimplementedCustomerManagementServer) GetAccountByID(context.Context, *account.GetAccountByIDRequest) (*account.GetAccountByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountByID not implemented")
}
func (UnimplementedCustomerManagementServer) mustEmbedUnimplementedCustomerManagementServer() {}
func (UnimplementedCustomerManagementServer) testEmbeddedByValue()                            {}

// UnsafeCustomerManagementServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CustomerManagementServer will
// result in compilation errors.
type UnsafeCustomerManagementServer interface {
	mustEmbedUnimplementedCustomerManagementServer()
}

func RegisterCustomerManagementServer(s grpc.ServiceRegistrar, srv CustomerManagementServer) {
	// If the following call pancis, it indicates UnimplementedCustomerManagementServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CustomerManagement_ServiceDesc, srv)
}

func _CustomerManagement_CreateCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(customer.CreateCustomerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerManagementServer).CreateCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CustomerManagement_CreateCustomer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerManagementServer).CreateCustomer(ctx, req.(*customer.CreateCustomerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerManagement_VerifyEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(customer.VerifyEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerManagementServer).VerifyEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CustomerManagement_VerifyEmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerManagementServer).VerifyEmail(ctx, req.(*customer.VerifyEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerManagement_GetCustomerByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(customer.GetCustomerByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerManagementServer).GetCustomerByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CustomerManagement_GetCustomerByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerManagementServer).GetCustomerByID(ctx, req.(*customer.GetCustomerByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerManagement_GetCustomerByRid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(customer.GetCustomerByRidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerManagementServer).GetCustomerByRid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CustomerManagement_GetCustomerByRid_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerManagementServer).GetCustomerByRid(ctx, req.(*customer.GetCustomerByRidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerManagement_CreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(account.CreateAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerManagementServer).CreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CustomerManagement_CreateAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerManagementServer).CreateAccount(ctx, req.(*account.CreateAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerManagement_GetAccountByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(account.GetAccountByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerManagementServer).GetAccountByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CustomerManagement_GetAccountByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerManagementServer).GetAccountByID(ctx, req.(*account.GetAccountByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CustomerManagement_ServiceDesc is the grpc.ServiceDesc for CustomerManagement service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CustomerManagement_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cusman.CustomerManagement",
	HandlerType: (*CustomerManagementServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCustomer",
			Handler:    _CustomerManagement_CreateCustomer_Handler,
		},
		{
			MethodName: "VerifyEmail",
			Handler:    _CustomerManagement_VerifyEmail_Handler,
		},
		{
			MethodName: "GetCustomerByID",
			Handler:    _CustomerManagement_GetCustomerByID_Handler,
		},
		{
			MethodName: "GetCustomerByRid",
			Handler:    _CustomerManagement_GetCustomerByRid_Handler,
		},
		{
			MethodName: "CreateAccount",
			Handler:    _CustomerManagement_CreateAccount_Handler,
		},
		{
			MethodName: "GetAccountByID",
			Handler:    _CustomerManagement_GetAccountByID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cusman/service_cusman.proto",
}
