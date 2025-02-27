package gapi

import (
	"context"

	"github.com/tunvx/simplebank/common/util"
	pb "github.com/tunvx/simplebank/grpc/pb/moneytransfer"
)

const okayResponse = "Ok"

func (service *Service) TestEmptyGET(ctx context.Context, req *pb.Empty) (*pb.OkayResponse, error) {
	return &pb.OkayResponse{Ok: okayResponse}, nil
}

func (service *Service) TestEmptyPOST(ctx context.Context, req *pb.Empty) (*pb.OkayResponse, error) {
	return &pb.OkayResponse{Ok: okayResponse}, nil
}

func (service *Service) TestCheckAccountNoProcessing(ctx context.Context, req *pb.CheckAccountRequest) (*pb.OkayResponse, error) {
	// In this test, we're simply validating the request parsing and network handling, no processing done
	// Return the same request back as a response
	return &pb.OkayResponse{Ok: okayResponse}, nil
}

func (service *Service) TestCheckAccountJustAuth(ctx context.Context, req *pb.CheckAccountRequest) (*pb.OkayResponse, error) {
	_, err := service.authorizeUser(ctx, []string{util.BankerRole, util.CustomerRole})
	if err != nil {
		return nil, err
	}

	// In this test, we're simply validating the request parsing, network handling, just processing authentication step
	// Return the same request back as a response
	return &pb.OkayResponse{Ok: okayResponse}, nil
}

func (service *Service) TestInternalTransferMoneyNoProcessing(ctx context.Context, req *pb.InternalTransferMoneyRequest) (*pb.OkayResponse, error) {
	// Similar to the above, this function tests request parsing and network handling only
	// Return the request object back as a response (no transfer processing)
	return &pb.OkayResponse{Ok: okayResponse}, nil
}

func (service *Service) TestInternalTransferMoneyJustAuth(ctx context.Context, req *pb.InternalTransferMoneyRequest) (*pb.OkayResponse, error) {
	_, err := service.authorizeUser(ctx, []string{util.BankerRole, util.CustomerRole})
	if err != nil {
		return nil, err
	}

	// Similar to the above, this function tests request parsing, network handling only, and just processing authentication ste
	// Return the request object back as a response (no transfer processing)
	return &pb.OkayResponse{Ok: okayResponse}, nil
}
