package api

import (
	"context"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gitlab.com/naftis/app/naftis/internal/app/naftis/command"
	"gitlab.com/naftis/app/naftis/internal/app/naftis/query"
	"gitlab.com/naftis/app/naftis/pkg/protocol/api"
	"gitlab.com/naftis/app/naftis/pkg/protocol/entity"
)

type ApiService struct {
	cmd   *command.Factory
	query *query.Factory
	log   zerolog.Logger
}

func NewApiService(cmd *command.Factory, query *query.Factory) *ApiService {
	return &ApiService{
		log:   log.With().Str("service", "api").Logger(),
		cmd:   cmd,
		query: query,
	}
}

func (a *ApiService) ScheduleWorkload(_ context.Context, req *api.ScheduleWorkloadRequest) (*api.ScheduleWorkloadResponse, error) {
	log := a.log.With().Str("method", "ScheduleWorkload").Logger()

	err := a.cmd.ScheduleWorkload().Invoke(*req.Spec)
	if err != nil {
		log.Warn().Err(err).Msg("Method call failed.")
		return nil, err
	}

	log.Trace().Msg("Method call success.")

	return &api.ScheduleWorkloadResponse{}, nil
}

func (a *ApiService) ListScheduledWorkloads(context.Context, *api.ListScheduledWorkloadsRequest) (*api.ListScheduledWorkloadsResponse, error) {
	log := a.log.With().Str("method", "ListScheduledWorkloads").Logger()

	list, err := a.query.ListScheduledWorkloads().Query()
	if err != nil {
		log.Warn().Err(err).Msg("Method call failed.")
		return nil, err
	}

	resList := make([]*entity.ScheduledWorkload, 0)

	for _, item := range list {
		itemCopy := item
		resList = append(resList, &itemCopy)
	}

	log.Trace().Msg("Method call success.")

	return &api.ListScheduledWorkloadsResponse{
		List: resList,
	}, nil
}
