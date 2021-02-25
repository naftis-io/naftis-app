package naftisctl

import (
	"context"
	"github.com/olekukonko/tablewriter"
	"gitlab.com/naftis/app/naftis/internal/pkg/validator"
	"gitlab.com/naftis/app/naftis/pkg/protocol/api"
	"os"
)

type ListObservedWorkloadsAppConfig struct {
}

type ListObservedWorkloadsApp struct {
	config ListObservedWorkloadsAppConfig
	app    *App
}

func NewListObservedWorkloadsApp(config ListObservedWorkloadsAppConfig, app *App) (*ListObservedWorkloadsApp, error) {
	err := validator.New().Struct(config)
	if err != nil {
		return nil, err
	}

	return &ListObservedWorkloadsApp{
		config: config,
		app:    app,
	}, nil
}

func (a *ListObservedWorkloadsApp) Start(ctx context.Context) error {
	l, err := a.app.api.ListObservedWorkloads(ctx, &api.ListObservedWorkloadsRequest{})
	if err != nil {
		return err
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Status", "WorkloadSpecificationMarketId"})

	for _, item := range l.List {
		table.Append([]string{
			item.Id,
			item.State.Current,
			item.WorkloadSpecificationMarketId,
		})
	}

	table.Render()

	return nil
}
