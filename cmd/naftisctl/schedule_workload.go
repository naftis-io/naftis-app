package main

import (
	"context"
	"flag"
	"github.com/google/subcommands"
	"github.com/rs/zerolog/log"
	"gitlab.com/naftis/app/naftis/internal/app/naftisctl"
)

type scheduleWorkload struct {
	config naftisctl.ScheduleWorkloadAppConfig
	app    *naftisctl.App
}

func newScheduleWorkload(app *naftisctl.App) *scheduleWorkload {
	return &scheduleWorkload{
		app: app,
	}
}

func (s *scheduleWorkload) Name() string {
	return "schedule-workload"
}

func (s *scheduleWorkload) Synopsis() string {
	return "schedule new workload"
}

func (s *scheduleWorkload) Usage() string {
	return `naftisctl schedule-workload [-workload-spec-file]
Schedule new workload based on JSON specification.

`
}

func (s *scheduleWorkload) SetFlags(f *flag.FlagSet) {
	f.StringVar(&s.config.WorkloadSpecFile, "workload-spec-file", "xx", "file containing json workload specification")
}

func (s *scheduleWorkload) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	subcommandApp, err := naftisctl.NewScheduleWorkloadApp(s.config, s.app)
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
