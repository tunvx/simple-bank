package gapi

import (
	"fmt"

	"github.com/tunvx/simplebank/common/token"
	"github.com/tunvx/simplebank/common/util"
	pb "github.com/tunvx/simplebank/grpc/pb/manage"
	db "github.com/tunvx/simplebank/management/db/sqlc"
	"github.com/tunvx/simplebank/notification/redis"
)

type Service struct {
	pb.UnimplementedManageServiceServer
	config          util.Config
	stores          []db.Store
	tokenMaker      token.Maker
	taskDistributor redis.TaskDistributor
}

// NewService creates new a Grpc service
func NewService(config util.Config, stores []db.Store, taskDistributor redis.TaskDistributor) (*Service, error) {
	tokenMaker, err := token.NewPasetoMaker(config.PublicKeyBase64, config.PrivateKeyBase64)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Service{
		config:          config,
		stores:          stores,
		tokenMaker:      tokenMaker,
		taskDistributor: taskDistributor,
	}

	return server, nil
}
