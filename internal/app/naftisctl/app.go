package naftisctl

import (
	"context"
	"fmt"
	"gitlab.com/naftis/app/naftis/internal/pkg/validator"
	"gitlab.com/naftis/app/naftis/pkg/protocol/api"
	"google.golang.org/grpc"
)

type AppConfig struct {
	GrpcApiServerConnectAddress string `validate:"required,ip"`
	GrpcApiServerConnectPort    uint64 `validate:"required,min=1,max=65535"`
}

type App struct {
	config AppConfig
	api    api.ApiClient
}

func NewApp(config AppConfig) (*App, error) {
	err := validator.New().Struct(config)
	if err != nil {
		return nil, err
	}

	return &App{
		config: config,
	}, nil
}

func (a *App) Start(ctx context.Context) error {
	serverConnection, err := grpc.DialContext(ctx,
		fmt.Sprintf("%s:%d", a.config.GrpcApiServerConnectAddress, a.config.GrpcApiServerConnectPort),
		grpc.WithInsecure())
	if err != nil {
		return err
	}

	a.api = api.NewApiClient(serverConnection)

	return nil
}
