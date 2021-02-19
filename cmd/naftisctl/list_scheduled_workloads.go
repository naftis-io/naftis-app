package main

import (
	"context"
	"flag"
	"github.com/google/subcommands"
	"github.com/rs/zerolog/log"
	"gitlab.com/naftis/app/naftis/internal/app/naftisctl"
)

type listScheduledWorkloads struct {
	config naftisctl.ListScheduledWorkloadsAppConfig
	app    *naftisctl.App
}

func newListScheduledWorkloads(app *naftisctl.App) *listScheduledWorkloads {
	return &listScheduledWorkloads{
		app: app,
	}
}

func (s *listScheduledWorkloads) Name() string {
	return "list-scheduled-workloads"
}

func (s *listScheduledWorkloads) Synopsis() string {
	return "list scheduled workloads"
}

func (s *listScheduledWorkloads) Usage() string {
	return `naftisctl list-scheduled-workloads
Lists scheduled workloads.

`
}

func (s *listScheduledWorkloads) SetFlags(_ *flag.FlagSet) {

}

func (s *listScheduledWorkloads) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	subcommandApp, err := naftisctl.NewListScheduledWorkloadsApp(s.config, s.app)
	if err != nil {
		log.Panic().
			Err(err).
			Msg("Application instance create failed.")
		return subcommands.ExitFailure
	}

	err = subcommandApp.Start(ctx)
	if err != nil {
		log.Err(err).
			Msg("Execution failed.")
		return subcommands.ExitFailure
	}

	return subcommands.ExitSuccess
}
