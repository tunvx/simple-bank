package gapi

import (
	"context"

	errdb "github.com/tunvx/simplebank/common/errs/db"
	errga "github.com/tunvx/simplebank/common/errs/gapi"
	"github.com/tunvx/simplebank/common/util"
	"github.com/tunvx/simplebank/cusmansrv/val"
	pb "github.com/tunvx/simplebank/grpc/pb/cusman/account"
	"github.com/tunvx/simplebank/grpc/pb/shardman"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (service *Service) GetAccountByAccNumber(ctx context.Context, req *pb.GetAccountByIDRequest) (*pb.GetAccountByIDResponse, error) {
	violations := validateGetAccountByAccNumber(req)
	if violations != nil {
		return nil, errga.InvalidArgumentError(violations)
	}

	accountID := util.ConvertAccNumberToInt64(req.GetAccountNumber())
	accShard, internal_err := service.shardmanClient.LookupAccountShard(ctx, 
		&shardman.LookupAccountShardRequest{
			AccountId: accountID,
		})
	if internal_err != nil {
		return nil, internal_err
	}

	shardID := accShard.ShardId - 1
	account, err := service.stores[shardID].GetAccountByID(ctx, accountID)
	if err != nil {
		if errdb.ErrorCode(err) == errdb.RecordNotFound {
			return nil, status.Errorf(codes.NotFound, "account ( %s ) not found in database", req.GetAccountNumber())
		}
		return nil, status.Errorf(codes.Internal, "failed to retrieve account ( %s ) in database: %s", req.GetAccountNumber(), err)
	}

	response := &pb.GetAccountByIDResponse{
		Account: convertAccount(account),
	}
	return response, nil
}

func validateGetAccountByAccNumber(req *pb.GetAccountByIDRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	// Validate Customer Real ID
	if err := val.ValidateAccountNumber(req.GetAccountNumber()); err != nil {
		violations = append(violations, errga.FieldViolation("account_number", err))
	}

	return violations
}
