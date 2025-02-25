package gapi

import (
	"context"

	errdb "github.com/tunvx/simplebank/common/errs/db"
	errga "github.com/tunvx/simplebank/common/errs/gapi"
	"github.com/tunvx/simplebank/common/util"
	pb "github.com/tunvx/simplebank/grpc/pb/cusman/customer"
	"github.com/tunvx/simplebank/grpc/pb/shardman"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (service *Service) GetCustomerByID(ctx context.Context, req *pb.GetCustomerByIDRequest) (*pb.GetCustomerByIDResponse, error) {
	// 1. Authorize the user
	authPayload, err := service.authorizeUser(ctx, []string{util.CustomerRole})
	if err != nil {
		return nil, errga.UnauthenticatedError(err)
	}

	customerId := authPayload.UserID
	shardId := authPayload.ShardID - 1

	// 2. Retrieval customer record
	customer, err := service.stores[shardId-1].GetCustomerByID(ctx, customerId)
	if err != nil {
		if errdb.ErrorCode(err) == errdb.RecordNotFound {
			return nil, status.Errorf(codes.NotFound, "customer ( %d ) not found", customerId)
		}
		return nil, status.Errorf(codes.Internal, "failed to find customer ( %d ): %s", customerId, err)
	}

	// 4. Return response
	rsp := &pb.GetCustomerByIDResponse{
		Customer: convertCustomer(customer),
	}
	return rsp, nil
}

func (service *Service) GetCustomerByRid(ctx context.Context, req *pb.GetCustomerByRidRequest) (*pb.GetCustomerByRidResponse, error) {
	// 1. Authorize the user
	authPayload, err := service.authorizeUser(ctx, []string{util.BankerRole, util.IServiceRole})
	if err != nil {
		return nil, errga.UnauthenticatedError(err)
	}

	// 2. Get customer shard info
	record, err := service.shardmanClient.LookupCustomerShard(ctx, &shardman.LookupCustomerShardRequest{
		CustomerRid: req.GetCustomerRid(),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to find customer ( %s ): %s", req.GetCustomerRid(), err)
	}

	customerId := record.GetCustomerId()
	shardId := record.GetShardId() - 1

	// 3. Retrieval customer record
	customer, err := service.stores[shardId].GetCustomerByID(ctx, customerId)
	if err != nil {
		if errdb.ErrorCode(err) == errdb.RecordNotFound {
			return nil, status.Errorf(codes.NotFound, "customer ( %s ) not found", req.GetCustomerRid())
		}
		return nil, status.Errorf(codes.Internal, "failed to find customer ( %s ): %s", req.GetCustomerRid(), err)
	}

	// 4. Block Customer users from accessing other people's information
	if authPayload.Role == util.CustomerRole && customer.CustomerRid != req.CustomerRid {
		return nil, status.Errorf(codes.PermissionDenied, "no authorized to access this resource")
	}

	// 5. Return response
	rsp := &pb.GetCustomerByRidResponse{
		Customer: convertCustomer(customer),
	}
	return rsp, nil
}
