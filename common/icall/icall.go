package icall

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// InternalCall là client interceptor để tự động thêm metadata vào mỗi lời gọi gRPC
func InternalCall(
	ctx context.Context,
	method string,
	req, reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	// Tạo metadata với "internal-call"
	md := metadata.Pairs("internal-call", "true")

	// Gắn metadata vào context
	ctx = metadata.NewOutgoingContext(ctx, md)

	// Gọi hàm invoker để tiếp tục lời gọi gRPC
	return invoker(ctx, method, req, reply, cc, opts...)
}
