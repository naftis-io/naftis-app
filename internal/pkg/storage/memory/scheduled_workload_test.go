package memory

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/thanhpk/randstr"
	"gitlab.com/naftis/app/naftis/internal/pkg/storage"
	"gitlab.com/naftis/app/naftis/pkg/protocol/entity"
	"testing"
	"time"
)

func TestNewScheduledWorkload(t *testing.T) {
	s := NewScheduledWorkload()

	assert.NotNil(t, s)
}

func TestScheduledWorkload_Create(t *testing.T) {
	var err error
	var w *entity.ScheduledWorkload

	s := NewScheduledWorkload()

	id := uuid.New().String()

	err = s.Create(entity.ScheduledWorkload{
		Id: id,
		Spec: &entity.WorkloadSpec{
			Containers: []*entity.WorkloadSpec_Container{},
		},
	})
	assert.NoError(t, err)

	err = s.Create(entity.ScheduledWorkload{
		Id: id,
		Spec: &entity.WorkloadSpec{
			Containers: []*entity.WorkloadSpec_Container{},
		},
	})
	assert.Error(t, err, storage.ErrScheduledWorkloadIdAlreadyUsed)

	w, err = s.Get(id)
	assert.NotNil(t, w)
	assert.NoError(t, err)
}

func TestScheduledWorkload_Get(t *testing.T) {
	var err error
	var w *entity.ScheduledWorkload

	s := NewScheduledWorkload()

	id := uuid.New().String()

	w, err = s.Get(id)
	assert.Nil(t, w)
	assert.Error(t, err, storage.ErrScheduledWorkloadNotFound)

	err = s.Create(entity.ScheduledWorkload{
		Id: id,
		Spec: &entity.WorkloadSpec{
			Containers: []*entity.WorkloadSpec_Container{},
		},
	})
	assert.NoError(t, err)

	w, err = s.Get(id)
	assert.NotNil(t, w)
	assert.NoError(t, err)
}

func TestScheduledWorkload_List(t *testing.T) {
	var err error
	var l []entity.ScheduledWorkload

	s := NewScheduledWorkload()

	id := uuid.New().String()

	l, err = s.List()
	assert.Len(t, l, 0)
	assert.NoError(t, err)

	err = s.Create(entity.ScheduledWorkload{
		Id: id,
		Spec: &entity.WorkloadSpec{
			Containers: []*entity.WorkloadSpec_Container{},
		},
	})
	assert.NoError(t, err)

	l, err = s.List()
	assert.Len(t, l, 1)
	assert.NoError(t, err)
}

func TestScheduledWorkload_UpdateWorkloadSpecificationMarketIdId(t *testing.T) {
	var err error

	s := NewScheduledWorkload()

	scheduledWorkloadId := uuid.New().String()

	err = s.Create(entity.ScheduledWorkload{
		Id: scheduledWorkloadId,
		Spec: &entity.WorkloadSpec{
			Containers: []*entity.WorkloadSpec_Container{},
		},
	})
	assert.NoError(t, err)

	e, err := s.Get(scheduledWorkloadId)
	assert.NoError(t, err)
	assert.NotNil(t, e)
	assert.Empty(t, e.WorkloadSpecificationMarketId)

	workloadSpecificationMarketId := randstr.Hex(64)

	err = s.UpdateWorkloadSpecificationMarketId(scheduledWorkloadId, workloadSpecificationMarketId)
	assert.NoError(t, err)

	e, err = s.Get(scheduledWorkloadId)
	assert.NoError(t, err)
	assert.NotNil(t, e)
	assert.Equal(t, workloadSpecificationMarketId, e.WorkloadSpecificationMarketId)
}

func TestScheduledWorkload_SetState(t *testing.T) {
	var err error

	s := NewScheduledWorkload()

	id := uuid.New().String()

	err = s.Create(entity.ScheduledWorkload{
		Id: id,
		State: &entity.State{
			Current:          "new",
			Previous:         "new",
			BackOffTimestamp: 0,
		},
		Spec: &entity.WorkloadSpec{
			Containers: []*entity.WorkloadSpec_Container{},
		},
	})
	assert.NoError(t, err)

	e, err := s.Get(id)
	assert.NoError(t, err)
	assert.NotNil(t, e)
	assert.EqualValues(t, e.State.Current, "new")

	newState := "test"

	err = s.SetState(id, newState)
	assert.NoError(t, err)

	e, err = s.Get(id)
	assert.NoError(t, err)
	assert.NotNil(t, e)
	assert.Equal(t, newState, e.State.Current)

	state, err := s.GetState(id)
	assert.NoError(t, err)
	assert.EqualValues(t, newState, state)
}

func TestScheduledWorkload_AddRunnerContractProposal(t *testing.T) {
	var err error

	s := NewScheduledWorkload()

	id := uuid.New().String()

	contractProposal := entity.ContractProposal{
		Contract: &entity.ContractSpecification{
			PricePerMinute:    1000,
			TokenSendInterval: 1,
			Duration:          10,
		},
	}
	contractProposalTxId := randstr.Hex(64)

	err = s.Create(entity.ScheduledWorkload{
		Id: id,
		Spec: &entity.WorkloadSpec{
			Containers: []*entity.WorkloadSpec_Container{},
		},
	})
	assert.NoError(t, err)

	err = s.AddRunnerContractProposal(id, entity.ScheduledWorkload_RunnerProposal{
		Proposal:                 &contractProposal,
		ContractProposalMarketId: contractProposalTxId,
	})
	assert.NoError(t, err)

	list, err := s.ListRunnerContractProposals(id)
	assert.NoError(t, err)
	assert.Len(t, list, 1)
}

func TestScheduledWorkload_GetByTxId(t *testing.T) {
	var err error

	s := NewScheduledWorkload()

	id := uuid.New().String()
	txId := randstr.Hex(64)

	e, err := s.GetByWorkloadSpecificationMarketId(txId)
	assert.Nil(t, e)
	assert.ErrorIs(t, err, storage.ErrScheduledWorkloadNotFound)

	err = s.Create(entity.ScheduledWorkload{
		Id:                            id,
		WorkloadSpecificationMarketId: txId,
	})
	assert.NoError(t, err)

	e, err = s.GetByWorkloadSpecificationMarketId(txId)
	assert.NotNil(t, e)
	assert.NoError(t, err)

}

func TestScheduledWorkload_SetBackOff(t *testing.T) {
	var err error

	s := NewScheduledWorkload()

	id := uuid.NewString()
	err = s.Create(entity.ScheduledWorkload{
		Id:    id,
		State: &entity.State{},
	})
	assert.NoError(t, err)

	scheduledWorkload, err := s.Get(id)
	assert.NoError(t, err)
	assert.NotNil(t, scheduledWorkload)
	assert.EqualValues(t, scheduledWorkload.State.BackOffTimestamp, 0)

	err = s.SetBackOff(id, time.Minute)
	assert.NoError(t, err)

	backOff, err := s.GetBackOff(id)
	assert.NoError(t, err)
	assert.NotEqualValues(t, backOff.Unix(), 0)
}

func TestScheduledWorkload_ListId(t *testing.T) {
	var err error

	s := NewScheduledWorkload()

	id := uuid.NewString()

	list, err := s.ListId()
	assert.NoError(t, err)
	assert.Empty(t, list)

	err = s.Create(entity.ScheduledWorkload{Id: id})
	assert.NoError(t, err)

	list, err = s.ListId()
	assert.NoError(t, err)
	assert.NotEmpty(t, list)
	assert.Contains(t, list, id)
}

func TestScheduledWorkload_SetAcceptedRunnerContractProposal(t *testing.T) {
	var err error

	s := NewScheduledWorkload()

	id := uuid.NewString()

	runnerProposal := entity.ScheduledWorkload_RunnerProposal{
		Proposal:                 &entity.ContractProposal{},
		ContractProposalMarketId: randstr.Hex(64),
	}

	err = s.Create(entity.ScheduledWorkload{Id: id})
	assert.NoError(t, err)

	scheduledWorkload, err := s.Get(id)
	assert.NoError(t, err)
	assert.Nil(t, scheduledWorkload.AcceptedRunnerProposal)

	err = s.SetAcceptedRunnerContractProposal(id, runnerProposal)
	assert.NoError(t, err)

	scheduledWorkload, err = s.Get(id)
	assert.NoError(t, err)
	assert.NotNil(t, scheduledWorkload.AcceptedRunnerProposal)
}
