package gapi

import (
	"fmt"

	"github.com/tunvx/simplebank/common/token"
	"github.com/tunvx/simplebank/common/util"
	pb "github.com/tunvx/simplebank/grpc/pb/shardman"
	db "github.com/tunvx/simplebank/shardmansrv/db/sqlc"
)

type Service struct {
	pb.UnimplementedShardManagementServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

// NewService creates new a Grpc service.
func NewService(config util.Config, store db.Store) (*Service, error) {
	tokenMaker, err := token.NewPasetoMaker(config.PublicKeyBase64, config.PrivateKeyBase64)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Service{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}
