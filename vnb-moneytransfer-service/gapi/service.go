package gapi

import (
	"fmt"

	"github.com/tunvx/simplebank/common/token"
	"github.com/tunvx/simplebank/common/util"
	db "github.com/tunvx/simplebank/cusmansrv/db/sqlc"
	pb "github.com/tunvx/simplebank/grpc/pb/moneytransfer"
	"github.com/tunvx/simplebank/moneytransfersrv/cache"
	worker "github.com/tunvx/simplebank/notificationsrv/redis"
)

type Service struct {
	pb.UnimplementedMoneyTransferServiceServer
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
