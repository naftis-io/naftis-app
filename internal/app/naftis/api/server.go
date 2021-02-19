package api

import (
	"context"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gitlab.com/naftis/app/naftis/internal/pkg/validator"
	"gitlab.com/naftis/app/naftis/pkg/protocol/api"
	"google.golang.org/grpc"
	"net"
	"time"
)

type ServerParams struct {
	GrpcListenAddress string `validate:"required,ip"`
	GrpcListenPort    uint64 `validate:"required,max=65535"`
}

type Server struct {
	params     ServerParams
	log        zerolog.Logger
	grpcSocket net.Listener
	grpc       *grpc.Server
	apiService *ApiService
}

func NewServer(params ServerParams, apiService *ApiService) (*Server, error) {
	err := validator.New().Struct(params)
	if err != nil {
		return nil, err
	}

	s := &Server{
		log:        log.With().Str("server", "api").Logger(),
		params:     params,
		apiService: apiService,
	}

	return s, nil
}

func (s *Server) Start(ctx context.Context) error {
	var err error

	s.log.Info().
		Uint64("listenPort", s.params.GrpcListenPort).
		Str("listenAddress", s.params.GrpcListenAddress).
		Msg("Starting server.")

	s.grpcSocket, err = net.Listen("tcp", fmt.Sprintf("%s:%d", s.params.GrpcListenAddress, s.params.GrpcListenPort))
	if err != nil {
		s.log.Fatal().
			Err(err).
			Msg("Failed to start TCP server for gRPC.")

		return err
	}

	s.grpc = grpc.NewServer()

	api.RegisterApiServer(s.grpc, s.apiService)

	go func() {
		for {
			err = s.grpc.Serve(s.grpcSocket)
			if err != nil {
				s.log.Fatal().
					Err(err).
					Msg("Failed to serve gRPC.")
			}

			time.Sleep(time.Second * 5)
		}
	}()

	return nil
}

func (s *Server) Stop() error {

	return nil
}
