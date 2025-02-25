package gapi

import (
	"context"

	"github.com/tunvx/simplebank/common/util"
	pb "github.com/tunvx/simplebank/grpc/pb/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (service *Service) GenerateTestAccessToken(ctx context.Context, req *pb.GenerateTestAccessTokenRequest) (*pb.GenerateTestAccessTokenResponse, error) {
	// 4. Create tokens
	accessToken, accessPayload, err := service.tokenMaker.CreateToken(
		1,1,
		util.BankerRole,
		service.config.AccessTokenDuration,
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create access token: %s", err)
	}

	rsp := &pb.GenerateTestAccessTokenResponse{
		AccessToken:          accessToken,
		AccessTokenExpiresAt: timestamppb.New(accessPayload.ExpiredAt),
	}
	return rsp, nil
}
