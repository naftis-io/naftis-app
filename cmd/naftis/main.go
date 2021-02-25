package main

import (
	"context"
	"flag"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gitlab.com/naftis/app/naftis/internal/app/naftis"
	"gitlab.com/naftis/app/naftis/internal/pkg/env"
	"gitlab.com/naftis/app/naftis/pkg/protocol/entity"
	"os"
	"time"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.Stamp}).
		With().
		Str("app", "naftis").
		Logger()

	ctx := context.Background()

	cfg := naftis.AppConfig{
		GrpcApiServerListenAddress: "",
		GrpcApiServerListenPort:    0,
		PriceList: entity.PriceList{
			CpuPerMinute:    1000,
			MemoryPerMinute: 50,
		},
	}

	flag.StringVar(&cfg.GrpcApiServerListenAddress,
		"grpc-api-server-listen-address",
		env.String("GRPC_API_SERVER_LISTEN_ADDRESS", "127.0.0.1"),
		"listen address of GRPC API server")

	flag.Uint64Var(&cfg.GrpcApiServerListenPort,
		"grpc-api-server-listen-port",
		env.Uint64("GRPC_API_SERVER_LISTEN_PORT", 3009),
		"listen port of GRPC API server")

	log.Info().Msg("Starting application.")

	app, err := naftis.NewApp(cfg)
	if err != nil {
		log.Panic().Err(err).Msg("Application start failed.")
		os.Exit(1)
	}

	err = app.Start(ctx)
	if err != nil {
		log.Panic().Err(err)
		os.Exit(1)
	}

	log.Info().Msg("Application started.")

	select {}

	os.Exit(1)
}
