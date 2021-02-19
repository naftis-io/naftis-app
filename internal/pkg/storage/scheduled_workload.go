package storage

import (
	"errors"
	"gitlab.com/naftis/app/naftis/pkg/protocol/entity"
)

type ScheduledWorkload interface {
	Create(entity entity.ScheduledWorkload) error
	Get(id string) (*entity.ScheduledWorkload, error)
	List() ([]entity.ScheduledWorkload, error)
	UpdateTxId(id string, txId string) error
	UpdateState(id string, state string) error
}

var (
	ErrScheduledWorkloadIdAlreadyUsed = errors.New("workload id already used")
	ErrScheduledWorkloadNotFound      = errors.New("workload not found")
)
