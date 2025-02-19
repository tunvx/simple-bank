package gapi

import (
	"context"

	errdb "github.com/tunvx/simplebank/common/errs/db"
	errga "github.com/tunvx/simplebank/common/errs/gapi"
	"github.com/tunvx/simplebank/common/util"
	cuspb "github.com/tunvx/simplebank/grpc/pb/management/customer"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (service *Service) IGetCustomerByID(ctx context.Context, req *cuspb.IGetCustomerByIDRequest) (*cuspb.IGetCustomerByIDResponse, error) {
	// 1. Authorize the user
	authPayload, err := service.authorizeUser(ctx, []string{util.BankerRole, util.CustomerRole, util.IServiceRole})
	if err != nil {
		return nil, errga.UnauthenticatedError(err)
	}
	if authPayload.Role == util.CustomerRole && authPayload.UserID != req.CustomerId {
		return nil, status.Errorf(codes.PermissionDenied, "no authorized to access this resource")
	}

	userID := authPayload.UserID
	shardID := util.ExtractShardID(userID)

	// 2. Retrieval customer record
	customer, err := service.stores[shardID].GetCustomerByID(ctx, req.CustomerId)
	if err != nil {
		if errdb.ErrorCode(err) == errdb.RecordNotFound {
			return nil, status.Errorf(codes.NotFound, "iget customer by id not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to ifind customer by id")
	}

	// 3. Return response
	rsp := &cuspb.IGetCustomerByIDResponse{
		Customer: convertCustomer(customer),
	}
	return rsp, nil
}

func (service *Service) IGetCustomerByRid(ctx context.Context, req *cuspb.IGetCustomerByRidRequest) (*cuspb.IGetCustomerByRidResponse, error) {
	// 1. Authorize the user
	authPayload, err := service.authorizeUser(ctx, []string{util.BankerRole, util.IServiceRole})
	if err != nil {
		return nil, errga.UnauthenticatedError(err)
	}

	userID := authPayload.UserID
	shardID := util.ExtractShardID(userID)

	// 2. Retrieval customer record
	customer, err := service.stores[shardID].GetCustomerByRid(ctx, req.CustomerRid)
	if err != nil {
		if errdb.ErrorCode(err) == errdb.RecordNotFound {
			return nil, status.Errorf(codes.NotFound, "iget customer by rid not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to ifind customer by rid")
	}

	// 3. Block Customer users from accessing other people's information
	if authPayload.Role == util.CustomerRole && customer.CustomerRid != req.CustomerRid {
		return nil, status.Errorf(codes.PermissionDenied, "no authorized to access this resource")
	}

	// 4. Return response
	rsp := &cuspb.IGetCustomerByRidResponse{
		Customer: convertICustomer(customer),
	}
	return rsp, nil
}

func (service *Service) GetCustomerByRid(ctx context.Context, req *cuspb.GetCustomerByRidRequest) (*cuspb.GetCustomerByRidResponse, error) {
	// 1. Authorize the user
	authPayload, err := service.authorizeUser(ctx, []string{util.BankerRole, util.IServiceRole})
	if err != nil {
		return nil, errga.UnauthenticatedError(err)
	}

	userID := authPayload.UserID
	shardID := util.ExtractShardID(userID)

	// 2. Retrieval customer record
	customer, err := service.stores[shardID].GetCustomerByRid(ctx, req.CustomerRid)
	if err != nil {
		if errdb.ErrorCode(err) == errdb.RecordNotFound {
			return nil, status.Errorf(codes.NotFound, "customer ( %s ) not found", req.GetCustomerRid())
		}
		return nil, status.Errorf(codes.Internal, "failed to find customer ( %s ): %s", req.GetCustomerRid(), err)
	}

	// 3. Block Customer users from accessing other people's information
	if authPayload.Role == util.CustomerRole && customer.CustomerRid != req.CustomerRid {
		return nil, status.Errorf(codes.PermissionDenied, "no authorized to access this resource")
	}

	// 4. Return response
	rsp := &cuspb.GetCustomerByRidResponse{
		Customer: convertCustomer(customer),
	}
	return rsp, nil
}
