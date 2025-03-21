package gapi

import (
	"fmt"

	db "github.com/tunvx/simplebank/authsrv/db/sqlc"
	"github.com/tunvx/simplebank/common/icall"
	"github.com/tunvx/simplebank/common/token"
	"github.com/tunvx/simplebank/common/util"
	pb "github.com/tunvx/simplebank/grpc/pb/auth"
	shardmanpb "github.com/tunvx/simplebank/grpc/pb/shardman"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Service struct {
	pb.UnimplementedAuthServiceServer
	config         util.Config
	store          db.Store
	tokenMaker     token.Maker
	shardmanClient shardmanpb.ShardManagementServiceClient
}

// NewService creates new a Grpc service.
func NewService(config util.Config, store db.Store) (*Service, error) {
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
		store:          store,
		tokenMaker:     tokenMaker,
		shardmanClient: shardmanpb.NewShardManagementServiceClient(conn),
	}

	return server, nil
}
