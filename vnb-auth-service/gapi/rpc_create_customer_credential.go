package gapi

import (
	"context"
	"fmt"

	db "github.com/tunvx/simplebank/authsrv/db/sqlc"
	"github.com/tunvx/simplebank/authsrv/val"
	errdb "github.com/tunvx/simplebank/common/errs/db"
	errga "github.com/tunvx/simplebank/common/errs/gapi"
	"github.com/tunvx/simplebank/common/util"
	pb "github.com/tunvx/simplebank/grpc/pb/auth"
	"github.com/tunvx/simplebank/grpc/pb/shardman"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (service *Service) CreateCustomerCredential(ctx context.Context, req *pb.CreateCustomerCredentialRequest) (*pb.CreateCustomerCredentialResponse, error) {
	// 1. Validate params
	violations := validateCreateCustomerCredentialRequest(req)
	if violations != nil {
		return nil, errga.InvalidArgumentError(violations)
	}

	// 2. Get relative info of customer
	cusShard, err := service.shardmanClient.LookupCustomerShard(ctx, &shardman.LookupCustomerShardRequest{
		CustomerRid: req.GetCustomerRid(),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to lool up customer shard for customer ( %s ) into database: %s", req.GetCustomerRid(), err)
	}

	fmt.Println("Customer Shard: ", cusShard)

	// 3. Hash the password
	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to hash password: %s", err)
	}

	// 4. Execute the creation of customer credential in the database
	arg := db.CreateCustomerCredentialParams{
		CustomerID:     cusShard.GetCustomerId(),
		ShardID:        cusShard.GetShardId(),
		Username:       req.GetUsername(),
		HashedPassword: hashedPassword,
	}
	_, err = service.store.CreateCustomerCredential(ctx, arg)
	if err != nil {
		if errdb.ErrorCode(err) == errdb.UniqueViolation {
			return nil, status.Errorf(codes.NotFound, "customer credential for user ( %s ) already exists in database", req.Username)
		}
		return nil, status.Errorf(codes.Internal, "failed to create customer credential for user ( %s ) into database: %s", req.Username, err)
	}

	// 5. Prepare and return the response
	rsp := &pb.CreateCustomerCredentialResponse{
		IsCreated: true,
	}
	return rsp, nil
}

func validateCreateCustomerCredentialRequest(req *pb.CreateCustomerCredentialRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	// >>> Customer Rid

	// Validate User Name
	if err := val.ValidateUsername(req.GetUsername()); err != nil {
		violations = append(violations, errga.FieldViolation("user_name", err))
	}

	// Validate Password
	if err := val.ValidatePassword(req.GetPassword()); err != nil {
		violations = append(violations, errga.FieldViolation("password", err))
	}

	return violations
}
