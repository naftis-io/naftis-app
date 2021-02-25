package command

import (
	"context"
	"gitlab.com/naftis/app/naftis/internal/app/naftis/service"
	"gitlab.com/naftis/app/naftis/internal/pkg/validator"
	"gitlab.com/naftis/app/naftis/pkg/protocol/entity"
)

type ScheduleWorkload struct {
	service *service.ScheduledWorkload
}

func NewScheduleWorkload(service *service.ScheduledWorkload) *ScheduleWorkload {
	return &ScheduleWorkload{
		service: service,
	}
}

func (cmd *ScheduleWorkload) Invoke(ctx context.Context, scheduledWorkload entity.ScheduledWorkload) error {
	err := validator.New().Struct(scheduledWorkload)
	if err != nil {
		return err
	}

	err = cmd.service.Schedule(ctx, scheduledWorkload)
	if err != nil {
		return err
	}

	return nil
}
