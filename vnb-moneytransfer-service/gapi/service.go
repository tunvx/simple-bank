package gapi

import (
	"fmt"

	"github.com/tunvx/simplebank/common/icall"
	"github.com/tunvx/simplebank/common/token"
	"github.com/tunvx/simplebank/common/util"
	db "github.com/tunvx/simplebank/cusmansrv/db/sqlc"
	pb "github.com/tunvx/simplebank/grpc/pb/moneytransfer"
	shardmanpb "github.com/tunvx/simplebank/grpc/pb/shardman"
	"github.com/tunvx/simplebank/moneytransfersrv/cache"
	worker "github.com/tunvx/simplebank/moneytransfersrv/worker/kafka"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Service struct {
	pb.UnimplementedMoneyTransferServiceServer
	config         util.Config
	stores         []db.Store
	cache          cache.Cache
	tokenMaker     token.Maker
	taskProducer   worker.TaskProducer
	shardmanClient shardmanpb.ShardManagementServiceClient
}

// NewService creates new a Grpc service
func NewService(config util.Config, stores []db.Store, cache cache.Cache, taskProducer worker.TaskProducer) (*Service, error) {
	tokenMaker, err := token.NewPasetoMaker(config.PublicKeyBase64, config.PrivateKeyBase64)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	// Create interceptor to pass metadata "internal-call"
	icallInterceptor := grpc.WithUnaryInterceptor(icall.InternalCall)

	// Dial the Management Service (Insecure for local dev environments)
	// Using insecure credentials for local development
	conn, err := grpc.NewClient(
		config.InternalShardManServiceAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		icallInterceptor,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Management Service: %w", err)
	}

	server := &Service{
		config:         config,
		stores:         stores,
		cache:          cache,
		tokenMaker:     tokenMaker,
		taskProducer:   taskProducer,
		shardmanClient: shardmanpb.NewShardManagementServiceClient(conn),
	}

	return server, nil
}
