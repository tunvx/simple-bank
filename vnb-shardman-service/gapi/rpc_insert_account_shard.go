package gapi

import (
	"context"

	errdb "github.com/tunvx/simplebank/common/errs/db"

	pb "github.com/tunvx/simplebank/grpc/pb/shardman"
	db "github.com/tunvx/simplebank/shardmansrv/db/sqlc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (service *Service) InsertAccountShard(ctx context.Context, req *pb.InsertAccountShardRequest) (*pb.InsertAccountShardResponse, error) {
	// 1. Validate params

	// 2. Execute the insert of account shard map into database
	arg := db.InsertAccountShardMapParams{
		AccountID:  req.GetAccountId(),
		CustomerID: req.GetCustomerId(),
		ShardID:    req.GetShardId(),
	}
	_, err := service.store.InsertAccountShardMap(ctx, arg)
	if err != nil {
		if errdb.ErrorCode(err) == errdb.UniqueViolation {
			return nil, status.Errorf(codes.NotFound, "insert account shard map for account id ( %s ) already exists in database", req.GetAccountId())
		}
		return nil, status.Errorf(codes.Internal, "failed to insert account shard map for account id ( %s ) into database: %s", req.GetAccountId(), err)
	}

	// 3. Prepare and return the response
	rsp := &pb.InsertAccountShardResponse{
		IsInserted: true,
	}
	return rsp, nil
}
