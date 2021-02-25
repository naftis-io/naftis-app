package storage

import (
	"errors"
	"gitlab.com/naftis/app/naftis/internal/pkg/state"
	"gitlab.com/naftis/app/naftis/pkg/protocol/entity"
)

type ScheduledWorkload interface {
	state.EntityStateStorage
	Create(entity entity.ScheduledWorkload) error
	Get(id string) (*entity.ScheduledWorkload, error)
	GetByWorkloadSpecificationMarketId(marketId string) (*entity.ScheduledWorkload, error)
	List() ([]entity.ScheduledWorkload, error)
	UpdateWorkloadSpecificationMarketId(id string, marketId string) error
	AddRunnerContractProposal(id string, proposal entity.ScheduledWorkload_RunnerProposal) error
	ListRunnerContractProposals(id string) ([]*entity.ScheduledWorkload_RunnerProposal, error)
	SetAcceptedRunnerContractProposal(id string, proposal entity.ScheduledWorkload_RunnerProposal) error
}

var (
	ErrScheduledWorkloadIdAlreadyUsed = errors.New("scheduled workload id already used")
	ErrScheduledWorkloadNotFound      = errors.New("scheduled workload not found")
)
