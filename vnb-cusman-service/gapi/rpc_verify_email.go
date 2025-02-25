package gapi

import (
	"context"

	"github.com/tunvx/simplebank/common/util"
	db "github.com/tunvx/simplebank/cusmansrv/db/sqlc"
	pb "github.com/tunvx/simplebank/grpc/pb/cusman/customer"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (service *Service) VerifyEmail(ctx context.Context, req *pb.VerifyEmailRequest) (*pb.VerifyEmailResponse, error) {
	shardID := req.GetShardId() - 1
	emailIDStr := req.GetEmailId()
	secretCode := req.GetSecretCode()

	emailUUID, err := util.ConvertStringToUUID(emailIDStr)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid email UUID: %v", err)
	}

	// Update verify status WHEN: email_id = @email_id AND secret_code = @secret_code
	txResult, err := service.stores[shardID].VerifyEmailTx(ctx, db.VerifyEmailTxParams{
		EmailId:    emailUUID,
		SecretCode: secretCode,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to verify email id ( %s ): %v", req.GetEmailId(), err)
	}

	rsp := &pb.VerifyEmailResponse{
		IsVerified: txResult.Customer.IsEmailVerified,
	}
	return rsp, nil
}
