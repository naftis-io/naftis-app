package command

import (
	"github.com/rs/zerolog/log"
	"gitlab.com/naftis/app/naftis/internal/pkg/storage"
	"gitlab.com/naftis/app/naftis/internal/pkg/validator"
	"gitlab.com/naftis/app/naftis/pkg/protocol/entity"
)

type ScheduleWorkload struct {
	storage storage.ScheduledWorkload
}

func NewScheduleWorkload(storage storage.ScheduledWorkload) *ScheduleWorkload {
	return &ScheduleWorkload{
		storage: storage,
	}
}

func (cmd *ScheduleWorkload) Invoke(workload entity.ScheduledWorkload) error {
	err := validator.New().Struct(workload)

	if err != nil {
		return err
	}

	workload.State = "new"

	err = cmd.storage.Create(workload)
	if err != nil {
		return err
	}

	log.Info().
		Str("scheduledWorkloadId", workload.Id).
		Msg("Scheduled new workload.")

	return nil
}
