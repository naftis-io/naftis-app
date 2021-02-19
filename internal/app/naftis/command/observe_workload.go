package command

import (
	"github.com/rs/zerolog/log"
	"gitlab.com/naftis/app/naftis/internal/pkg/storage"
	"gitlab.com/naftis/app/naftis/internal/pkg/validator"
	"gitlab.com/naftis/app/naftis/pkg/protocol/entity"
)

type ObserveWorkload struct {
	storage storage.ObservedWorkload
}

func NewObserveWorkload(storage storage.ObservedWorkload) *ObserveWorkload {
	return &ObserveWorkload{
		storage: storage,
	}
}

func (cmd *ObserveWorkload) Invoke(workload entity.ObservedWorkload) error {
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
		Str("observedWorkloadId", workload.Id).
		Str("txId", workload.TxId).
		Msg("Observing new workload.")

	return nil
}
