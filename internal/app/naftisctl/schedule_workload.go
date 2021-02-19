package naftisctl

import (
	"context"
	"github.com/golang/protobuf/jsonpb"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"gitlab.com/naftis/app/naftis/internal/pkg/validator"
	"gitlab.com/naftis/app/naftis/pkg/protocol/api"
	"gitlab.com/naftis/app/naftis/pkg/protocol/entity"
	"os"
)

type ScheduleWorkloadAppConfig struct {
	WorkloadSpecFile string `validate:"required"`
}

type ScheduleWorkloadApp struct {
	config ScheduleWorkloadAppConfig
	app    *App
}

func NewScheduleWorkloadApp(config ScheduleWorkloadAppConfig, app *App) (*ScheduleWorkloadApp, error) {
	err := validator.New().Struct(config)
	if err != nil {
		return nil, err
	}

	return &ScheduleWorkloadApp{
		config: config,
		app:    app,
	}, nil
}

func (a *ScheduleWorkloadApp) Start(ctx context.Context) error {
	workloadSpecFile, err := os.Open(a.config.WorkloadSpecFile)
	if err != nil {
		return err
	}
	defer workloadSpecFile.Close()

	workloadSpec := entity.WorkloadSpec{}

	if err := jsonpb.Unmarshal(workloadSpecFile, &workloadSpec); err != nil {
		return err
	}

	if err := validator.New().Struct(workloadSpec); err != nil {
		return err
	}

	id := uuid.New()

	_, err = a.app.api.ScheduleWorkload(ctx, &api.ScheduleWorkloadRequest{
		Spec: &entity.ScheduledWorkload{
			Id:   id.String(),
			Spec: &workloadSpec,
		},
	})
	if err != nil {
		return err
	}

	log.Info().Str("scheduledWorkloadId", id.String()).
		Msg("Workload scheduled.")

	return nil
}
