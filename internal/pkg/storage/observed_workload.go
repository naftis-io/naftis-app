package storage

import (
	"errors"
	"gitlab.com/naftis/app/naftis/pkg/protocol/entity"
)

type ObservedWorkload interface {
	Create(entity entity.ObservedWorkload) error
	Get(id string) (*entity.ObservedWorkload, error)
	GetByTxId(txId string) (*entity.ObservedWorkload, error)
	List() ([]entity.ObservedWorkload, error)
	UpdateState(id string, state string) error
}

var (
	ErrObservedWorkloadIdAlreadyUsed = errors.New("workload id already used")
	ErrObservedWorkloadNotFound      = errors.New("workload not found")
)
