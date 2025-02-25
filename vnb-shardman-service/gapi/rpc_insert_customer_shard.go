package gapi

import (
	"context"

	errdb "github.com/tunvx/simplebank/common/errs/db"
	errga "github.com/tunvx/simplebank/common/errs/gapi"
	"github.com/tunvx/simplebank/common/util"
	pb "github.com/tunvx/simplebank/grpc/pb/shardman"
	db "github.com/tunvx/simplebank/shardmansrv/db/sqlc"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (service *Service) InsertCustomerShard(ctx context.Context, req *pb.InsertCustomerShardRequest) (*pb.InsertCustomerShardResponse, error) {
	// 1. Validate params
	violations := validateInsertCustomerShardRequest(req)
	if violations != nil {
		return nil, errga.InvalidArgumentError(violations)
	}

	// 2. Check if the record already exists
	record, err := service.store.GetShardByCustomeRid(ctx, req.GetCustomerRid())
	if err == nil {
		// Record already exists, return it immediately
		return &pb.InsertCustomerShardResponse{
			CustomerId: record.CustomerID,
			ShardId:    record.ShardID,
		}, nil
	} else if errdb.ErrorCode(err) != errdb.RecordNotFound {
		// If the error is something other than "not found," return an error
		return nil, status.Errorf(codes.Internal, "failed to check customer shard for customer rid ( %s ): %s", req.GetCustomerRid(), err)
	}

	// 3. Execute the insert since the record doesn't exist
	arg := db.InsertCustomerShardMapParams{
		CustomerRid: req.GetCustomerRid(),
		ShardID:     1,
	}
	newRecord, err := service.store.InsertCustomerShardMap(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to insert customer shard map for customer rid ( %s ): %s", req.GetCustomerRid(), err)
	}

	// 4. Update shard id if needed
	updatedShardID := util.GenShardID(newRecord.CustomerID, service.config.ShardVolume)
	newRecord, err = service.store.UpdateCustomerShardMap(ctx, db.UpdateCustomerShardMapParams{
		CustomerID: newRecord.CustomerID,
		ShardID:    updatedShardID,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update shard ID for customer rid ( %s ): %s", req.GetCustomerRid(), err)
	}

	// 5. Prepare and return the response
	rsp := &pb.InsertCustomerShardResponse{
		CustomerId: newRecord.CustomerID,
		ShardId:    newRecord.ShardID,
	}
	return rsp, nil
}

func validateInsertCustomerShardRequest(req *pb.InsertCustomerShardRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	// Validate Shard ID
	// if err := val.ValidateShardID(req.GetShardId()); err != nil {
	// 	violations = append(violations, errga.FieldViolation("shard_id", err))
	// }
	// return violations
	return nil
}
