package naftisctl

import (
	"context"
	"github.com/olekukonko/tablewriter"
	"gitlab.com/naftis/app/naftis/internal/pkg/validator"
	"gitlab.com/naftis/app/naftis/pkg/protocol/api"
	"os"
)

type ListScheduledWorkloadsAppConfig struct {
}

type ListScheduledWorkloadsApp struct {
	config ListScheduledWorkloadsAppConfig
	app    *App
}

func NewListScheduledWorkloadsApp(config ListScheduledWorkloadsAppConfig, app *App) (*ListScheduledWorkloadsApp, error) {
	err := validator.New().Struct(config)
	if err != nil {
		return nil, err
	}

	return &ListScheduledWorkloadsApp{
		config: config,
		app:    app,
	}, nil
}

func (a *ListScheduledWorkloadsApp) Start(ctx context.Context) error {
	l, err := a.app.api.ListScheduledWorkloads(ctx, &api.ListScheduledWorkloadsRequest{})
	if err != nil {
		return err
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Status"})

	for _, item := range l.List {
		table.Append([]string{
			item.Id,
			item.State.Current,
		})
	}

	table.Render()

	return nil
}
