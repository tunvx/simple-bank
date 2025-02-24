package gapi

import (
	"context"

	errdb "github.com/tunvx/simplebank/common/errs/db"

	pb "github.com/tunvx/simplebank/grpc/pb/shardman"
	db "github.com/tunvx/simplebank/shardmansrv/db/sqlc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (service *Service) InsertCustomerShard(ctx context.Context, req *pb.InsertCustomerShardRequest) (*pb.InsertCustomerShardResponse, error) {
	// 1. Validate params

	// 2. Execute the insert of customer shard map into database
	arg := db.InsertCustomerShardMapParams{
		CustomerRid: req.GetCustomerRid(),
		ShardID:     req.GetShardId(),
	}
	_, err := service.store.InsertCustomerShardMap(ctx, arg)
	if err != nil {
		if errdb.ErrorCode(err) == errdb.UniqueViolation {
			return nil, status.Errorf(codes.NotFound, "insert customer shard map for customer rid ( %s ) already exists in database", req.GetCustomerRid())
		}
		return nil, status.Errorf(codes.Internal, "failed to insert customer shard map for customer rid ( %s ) into database: %s", req.GetCustomerRid(), err)
	}

	// 3. Prepare and return the response
	rsp := &pb.InsertCustomerShardResponse{
		IsInserted: true,
	}
	return rsp, nil
}
