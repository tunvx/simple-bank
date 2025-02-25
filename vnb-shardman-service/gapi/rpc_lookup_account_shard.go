package gapi

import (
	"context"
	"fmt"

	errdb "github.com/tunvx/simplebank/common/errs/db"

	pb "github.com/tunvx/simplebank/grpc/pb/shardman"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (service *Service) LookupAccountShard(ctx context.Context, req *pb.LookupAccountShardRequest) (*pb.LookupAccountShardResponse, error) {
	// 1. Validate params

	// 2. Retrieve account shard map
	shardID, err := service.store.GetShardByAccountID(ctx, req.GetAccountId())
	if err != nil {
		if errdb.ErrorCode(err) == errdb.RecordNotFound {
			return nil, status.Errorf(codes.NotFound, "account shard map for account id ( %s ) not found in database: ", req.GetAccountId())
		}
		return nil, status.Errorf(codes.Internal, "failed to retrieve account shard map of account id ( %s ) in database: %s", req.GetAccountId(), err)
	}

	fmt.Println("Shard ID: ", shardID)

	// 3. Prepare and return the response
	rsp := &pb.LookupAccountShardResponse{
		AccountId: req.GetAccountId(),
		ShardId:   shardID,
	}
	fmt.Println("Returned Response: ", rsp)
	return rsp, nil
}
