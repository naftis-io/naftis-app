package main

import (
	"context"
	"flag"
	"github.com/google/subcommands"
	"github.com/rs/zerolog/log"
	"gitlab.com/naftis/app/naftis/internal/app/naftisctl"
)

type listObservedWorkloads struct {
	config naftisctl.ListObservedWorkloadsAppConfig
	app    *naftisctl.App
}

func newListObservedWorkloads(app *naftisctl.App) *listObservedWorkloads {
	return &listObservedWorkloads{
		app: app,
	}
}

func (s *listObservedWorkloads) Name() string {
	return "list-observed-workloads"
}

func (s *listObservedWorkloads) Synopsis() string {
	return "list observed workloads"
}

func (s *listObservedWorkloads) Usage() string {
	return `naftisctl list-observed-workloads
Lists observed workloads.

`
}

func (s *listObservedWorkloads) SetFlags(_ *flag.FlagSet) {

}

func (s *listObservedWorkloads) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	subcommandApp, err := naftisctl.NewListObservedWorkloadsApp(s.config, s.app)
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
