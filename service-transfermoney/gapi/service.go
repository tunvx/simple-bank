package gapi

import (
	"fmt"

	"github.com/tunvx/simplebank/common/token"
	"github.com/tunvx/simplebank/common/util"
	pb "github.com/tunvx/simplebank/grpc/pb/transactions"
	db "github.com/tunvx/simplebank/management/db/sqlc"
	worker "github.com/tunvx/simplebank/notification/redis"
	"github.com/tunvx/simplebank/transfermoney/cache"
)

type Service struct {
	pb.UnimplementedTransactionServiceServer
	config          util.Config
	stores          []db.Store
	cache           cache.Cache
	tokenMaker      token.Maker
	taskDistributor worker.TaskDistributor
}

// NewService creates new a Grpc service
func NewService(config util.Config, stores []db.Store, cache cache.Cache, taskDistributor worker.TaskDistributor) (*Service, error) {
	tokenMaker, err := token.NewPasetoMaker(config.PublicKeyBase64, config.PrivateKeyBase64)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Service{
		config:          config,
		stores:          stores,
		cache:           cache,
		tokenMaker:      tokenMaker,
		taskDistributor: taskDistributor,
	}

	return server, nil
}
