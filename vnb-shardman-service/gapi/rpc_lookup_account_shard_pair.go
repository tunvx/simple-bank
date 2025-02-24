package gapi

import (
	"context"

	pb "github.com/tunvx/simplebank/grpc/pb/shardman"
)

func (service *Service) LookupAccountShardPair(ctx context.Context, req *pb.LookupAccountShardPairRequest) (*pb.LookupAccountShardPairResponse, error) {
	return nil, nil
}
