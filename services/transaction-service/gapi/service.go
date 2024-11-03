package gapi

import (
	"fmt"

	pb "github.com/tunvx/simplebank/grpc/pb/transactions"
	db "github.com/tunvx/simplebank/manage/db/sqlc"
	worker "github.com/tunvx/simplebank/notification/redis"
	"github.com/tunvx/simplebank/pkg/token"
	"github.com/tunvx/simplebank/pkg/util"
	"github.com/tunvx/simplebank/transactions/cache"
)

type Service struct {
	pb.UnimplementedTransactionServiceServer
	config          util.Config
	store           db.Store
	cache           cache.Cache
	tokenMaker      token.Maker
	taskDistributor worker.TaskDistributor
}

// NewService creates new a Grpc service
func NewService(config util.Config, store db.Store, cache cache.Cache, taskDistributor worker.TaskDistributor) (*Service, error) {
	tokenMaker, err := token.NewPasetoMaker(config.PublicKeyBase64, config.PrivateKeyBase64)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Service{
		config:          config,
		store:           store,
		cache:           cache,
		tokenMaker:      tokenMaker,
		taskDistributor: taskDistributor,
	}

	return server, nil
}
