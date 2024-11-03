package gapi

import (
	"fmt"

	pb "github.com/tunvx/simplebank/grpc/pb/manage"
	db "github.com/tunvx/simplebank/manage/db/sqlc"
	"github.com/tunvx/simplebank/notification/redis"
	"github.com/tunvx/simplebank/pkg/token"
	"github.com/tunvx/simplebank/pkg/util"
)

type Service struct {
	pb.UnimplementedManageServiceServer
	config          util.Config
	store           db.Store
	tokenMaker      token.Maker
	taskDistributor redis.TaskDistributor
}

// NewService creates new a Grpc service
func NewService(config util.Config, store db.Store, taskDistributor redis.TaskDistributor) (*Service, error) {
	tokenMaker, err := token.NewPasetoMaker(config.PublicKeyBase64, config.PrivateKeyBase64)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Service{
		config:          config,
		store:           store,
		tokenMaker:      tokenMaker,
		taskDistributor: taskDistributor,
	}

	return server, nil
}
