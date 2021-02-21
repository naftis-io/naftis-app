package naftis

import (
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	"gitlab.com/naftis/app/naftis/internal/app/naftis/api"
	"gitlab.com/naftis/app/naftis/internal/app/naftis/command"
	"gitlab.com/naftis/app/naftis/internal/app/naftis/label"
	"gitlab.com/naftis/app/naftis/internal/app/naftis/label/source"
	marketListener "gitlab.com/naftis/app/naftis/internal/app/naftis/listener/market"
	"gitlab.com/naftis/app/naftis/internal/app/naftis/listener/market/filter"
	"gitlab.com/naftis/app/naftis/internal/app/naftis/query"
	"gitlab.com/naftis/app/naftis/internal/app/naftis/reconciler"
	"gitlab.com/naftis/app/naftis/internal/pkg/market"
	memoryMarket "gitlab.com/naftis/app/naftis/internal/pkg/market/memory"
	"gitlab.com/naftis/app/naftis/internal/pkg/storage"
	"gitlab.com/naftis/app/naftis/internal/pkg/storage/memory"
	"gitlab.com/naftis/app/naftis/internal/pkg/validator"
)

type AppConfig struct {
	GrpcApiServerListenAddress string `validate:"required,ip"`
	GrpcApiServerListenPort    uint64 `validate:"required,max=65535"`
}

type App struct {
	config            AppConfig
	cmd               *command.Factory
	query             *query.Factory
	labels            *label.Container
	market            market.MessageToken
	marketListener    *marketListener.Listener
	storage           storage.Container
	apiServer         *api.Server
	apiService        *api.ApiService
	scheduledWorkload *reconciler.ScheduledWorkload
}

func NewApp(config AppConfig) (*App, error) {
	var err error

	err = validator.New().Struct(config)
	if err != nil {
		return nil, err
	}

	storage := memory.NewContainer()
	cmd := command.NewFactory(storage)
	query := query.NewFactory(storage)

	labels := label.NewContainer()
	labels.AttachSource(source.NewOs())

	market := memoryMarket.NewMessage()

	marketListener := marketListener.NewMarket(cmd, market, 64)
	marketListener.AttachWorkloadSpecificationFilter(filter.NewMarketWorkloadSpecificationNodeSelector(labels))

	scheduledWorkload := reconciler.NewScheduledWorkload(storage, market)

	apiService := api.NewApiService(cmd, query)

	apiServer, err := api.NewServer(api.ServerParams{
		GrpcListenAddress: config.GrpcApiServerListenAddress,
		GrpcListenPort:    config.GrpcApiServerListenPort,
	}, apiService)
	if err != nil {
		return nil, err
	}

	a := &App{
		config:            config,
		cmd:               cmd,
		query:             query,
		labels:            labels,
		storage:           storage,
		apiServer:         apiServer,
		scheduledWorkload: scheduledWorkload,
		market:            market,
		marketListener:    marketListener,
	}

	return a, nil
}

func (a *App) Start(ctx context.Context) error {
	var err error

	err = a.labels.Refresh()
	if err != nil {
		return err
	}

	a.logNodeLabels()

	err = a.market.Start(ctx)
	if err != nil {
		return err
	}

	err = a.marketListener.Start(ctx)
	if err != nil {
		return err
	}

	err = a.apiServer.Start(ctx)
	if err != nil {
		return err
	}

	err = a.scheduledWorkload.Start(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) logNodeLabels() {
	for _, label := range a.labels.List() {
		log.Info().
			Str("label", fmt.Sprintf("%s: %s", label.Key, label.Value)).
			Msg("Discovered node label.")
	}
}
