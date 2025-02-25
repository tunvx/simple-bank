package gapi

import (
	"context"
	"fmt"

	errdb "github.com/tunvx/simplebank/common/errs/db"

	pb "github.com/tunvx/simplebank/grpc/pb/shardman"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (service *Service) LookupCustomerShard(ctx context.Context, req *pb.LookupCustomerShardRequest) (*pb.LookupCustomerShardResponse, error) {
	// 1. Validate params

	// 2. Retrieve customer shard map
	row, err := service.store.GetShardByCustomeRid(ctx, req.GetCustomerRid())
	if err != nil {
		if errdb.ErrorCode(err) == errdb.RecordNotFound {
			return nil, status.Errorf(codes.NotFound, "customer shard map for customer rid ( %s ) not found in database: ", req.GetCustomerRid())
		}
		return nil, status.Errorf(codes.Internal, "failed to retrieve customer shard map of customer rid ( %s ) in database: %s", req.GetCustomerRid(), err)
	}

	fmt.Println(row)

	// 3. Prepare and return the response
	rsp := &pb.LookupCustomerShardResponse{
		CustomerId: row.CustomerID,
		ShardId:    row.ShardID,
	}
	return rsp, nil
}
