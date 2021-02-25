package storage

import (
	"errors"
	"gitlab.com/naftis/app/naftis/internal/pkg/state"
	"gitlab.com/naftis/app/naftis/pkg/protocol/entity"
)

type ObservedWorkload interface {
	state.EntityStateStorage
	Get(id string) (*entity.ObservedWorkload, error)
	GetByWorkloadSpecificationMarketId(marketId string) (*entity.ObservedWorkload, error)
	Create(entity entity.ObservedWorkload) error
	List() ([]entity.ObservedWorkload, error)
	SetPrincipalAcceptance(id string, acceptance entity.ObservedWorkload_PrincipalAcceptance) error
}

var (
	ErrObservedWorkloadIdAlreadyUsed = errors.New("workload id already used")
	ErrObservedWorkloadNotFound      = errors.New("workload not found")
)
