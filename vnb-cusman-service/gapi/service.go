package gapi

import (
	"fmt"

	"github.com/tunvx/simplebank/common/icall"
	"github.com/tunvx/simplebank/common/token"
	"github.com/tunvx/simplebank/common/util"
	db "github.com/tunvx/simplebank/cusmansrv/db/sqlc"
	"github.com/tunvx/simplebank/cusmansrv/worker"
	pb "github.com/tunvx/simplebank/grpc/pb/cusman"
	"github.com/tunvx/simplebank/grpc/pb/shardman"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Service struct {
	pb.UnimplementedCustomerManagementServiceServer
	config          util.Config
	stores          []db.Store
	tokenMaker      token.Maker
	shardmanClient  shardman.ShardManagementServiceClient
	taskDistributor worker.TaskDistributor
}

// NewService creates new a Grpc service
func NewService(config util.Config, stores []db.Store, taskDistributor worker.TaskDistributor) (*Service, error) {
	tokenMaker, err := token.NewPasetoMaker(config.PublicKeyBase64, config.PrivateKeyBase64)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	// Create interceptor to pass metadata "internal-call"
	icallInterceptor := grpc.WithUnaryInterceptor(icall.InternalCall)

	// Dial the ShardMan Service (Insecure for local dev environments)
	// Using insecure credentials for local development
	conn, err := grpc.NewClient(
		config.GRPCShardManServiceAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		icallInterceptor,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Management Service: %w", err)
	}

	server := &Service{
		config:          config,
		stores:          stores,
		tokenMaker:      tokenMaker,
		shardmanClient:  shardman.NewShardManagementServiceClient(conn),
		taskDistributor: taskDistributor,
	}

	return server, nil
}
