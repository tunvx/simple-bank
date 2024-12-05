package gapi

import (
	"fmt"

	db "github.com/tunvx/simplebank/auth/db/sqlc"
	pb "github.com/tunvx/simplebank/grpc/pb/auth"
	manpb "github.com/tunvx/simplebank/grpc/pb/manage"
	"github.com/tunvx/simplebank/pkg/icall"
	"github.com/tunvx/simplebank/pkg/token"
	"github.com/tunvx/simplebank/pkg/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Service struct {
	pb.UnimplementedAuthServiceServer
	config       util.Config
	store        db.Store
	tokenMaker   token.Maker
	manageClient manpb.ManageServiceClient
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
		config.InternalManageServiceAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		icallInterceptor,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Management Service: %w", err)
	}

	server := &Service{
		config:       config,
		store:        store,
		tokenMaker:   tokenMaker,
		manageClient: manpb.NewManageServiceClient(conn),
	}

	return server, nil
}
