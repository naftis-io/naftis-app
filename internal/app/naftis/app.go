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
	"gitlab.com/naftis/app/naftis/internal/app/naftis/service"
	"gitlab.com/naftis/app/naftis/internal/app/naftis/state"
	"gitlab.com/naftis/app/naftis/internal/pkg/contract"
	"gitlab.com/naftis/app/naftis/internal/pkg/market"
	memoryMarket "gitlab.com/naftis/app/naftis/internal/pkg/market/memory"
	"gitlab.com/naftis/app/naftis/internal/pkg/price"
	"gitlab.com/naftis/app/naftis/internal/pkg/storage"
	"gitlab.com/naftis/app/naftis/internal/pkg/storage/memory"
	"gitlab.com/naftis/app/naftis/internal/pkg/validator"
	"gitlab.com/naftis/app/naftis/pkg/protocol/entity"
)

type AppConfig struct {
	GrpcApiServerListenAddress string           `validate:"required,ip"`
	GrpcApiServerListenPort    uint64           `validate:"required,max=65535"`
	PriceList                  entity.PriceList `validate:"required"`
}

type App struct {
	config            AppConfig
	cmd               *command.Factory
	query             *query.Factory
	labels            *label.Container
	market            market.MessageToken
	marketListener    *marketListener.Listener
	storage           storage.Container
	priceCalculator   *price.Calculator
	contractSelector  *contract.Selector
	apiServer         *api.Server
	apiService        *api.ApiService
	scheduledWorkload *state.ScheduledWorkload
	observedWorkload  *state.ObservedWorkload
}

func NewApp(config AppConfig) (*App, error) {
	var err error

	err = validator.New().Struct(config)
	if err != nil {
		return nil, err
	}

	market := memoryMarket.NewMessage()
	storage := memory.NewContainer()

	contractSelector := contract.NewSelector()
	priceCalculator := price.NewCalculator(config.PriceList)

	services := service.NewContainer(storage, market, priceCalculator, contractSelector)

	cmd := command.NewFactory(services)
	query := query.NewFactory(storage)

	labels := label.NewContainer()
	labels.AttachSource(source.NewOs())

	marketListener := marketListener.NewMarket(market, services, 64)
	marketListener.AttachWorkloadSpecificationFilter(filter.NewMarketWorkloadSpecificationNodeSelector(labels))
	marketListener.AttachContractProposalFilter(filter.NewContractProposalScheduledWorkload(storage.ScheduledWorkload()))

	scheduledWorkload := state.NewScheduledWorkload(storage.ScheduledWorkload(), services.ScheduledWorkload())
	observedWorkload := state.NewObservedWorkload(storage.ObservedWorkload(), services.ObservedWorkload())

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
		priceCalculator:   priceCalculator,
		storage:           storage,
		apiServer:         apiServer,
		scheduledWorkload: scheduledWorkload,
		observedWorkload:  observedWorkload,
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

	err = a.observedWorkload.Start(ctx)
	if err != nil {
		return err
	}

	a.logNodeLabels()
	a.logNodePrices()

	return nil
}

func (a *App) logNodeLabels() {
	for _, label := range a.labels.List() {
		log.Info().
			Str("label", fmt.Sprintf("%s: %s", label.Key, label.Value)).
			Msg("Node label.")
	}
}

func (a *App) logNodePrices() {
	log.Info().
		Str("kind", "MemoryPerMinute").
		Uint32("pricePerMinute", a.config.PriceList.MemoryPerMinute).
		Msg("Price list.")

	log.Info().
		Str("kind", "CpuPerMinute").
		Uint32("pricePerMinute", a.config.PriceList.CpuPerMinute).
		Msg("Price list.")

}
