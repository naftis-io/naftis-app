package main

import (
	"context"
	"flag"
	"github.com/google/subcommands"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gitlab.com/naftis/app/naftis/internal/app/naftisctl"
	"gitlab.com/naftis/app/naftis/internal/pkg/env"
	"os"
	"time"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.Stamp}).
		With().
		Str("app", "naftisctl").
		Logger()

	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")

	appCfg := naftisctl.AppConfig{
		GrpcApiServerConnectAddress: "",
		GrpcApiServerConnectPort:    0,
	}

	flag.StringVar(&appCfg.GrpcApiServerConnectAddress,
		"grpc-api-server-connect-address",
		env.String("GRPC_API_SERVER_CONNECT_ADDRESS", "127.0.0.1"),
		"address of grpc api server")

	flag.Uint64Var(&appCfg.GrpcApiServerConnectPort,
		"grpc-api-server-connect-port",
		env.Uint64("GRPC_API_SERVER_CONNECT_PORT", 3009),
		"port of grpc api ")

	flag.Parse()
	ctx := context.Background()

	app, err := naftisctl.NewApp(appCfg)
	if err != nil {
		log.Panic().Err(err).Msg("Application instance create failed.")
		os.Exit(1)
	}

	err = app.Start(ctx)
	if err != nil {
		log.Panic().Err(err).Msg("Application instance start failed.")
		os.Exit(1)
	}

	subcommands.Register(newScheduleWorkload(app), "principal")
	subcommands.Register(newListScheduledWorkloads(app), "principal")

	flag.Parse()

	os.Exit(int(subcommands.Execute(ctx)))
}
