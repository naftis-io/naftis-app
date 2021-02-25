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

func TestNewObservedWorkload(t *testing.T) {
	s := NewObservedWorkload()

	assert.NotNil(t, s)
}

func TestObservedWorkload_Create(t *testing.T) {
	var err error
	var w *entity.ObservedWorkload

	s := NewObservedWorkload()

	id := uuid.New().String()

	err = s.Create(entity.ObservedWorkload{
		Id: id,
		Spec: &entity.WorkloadSpec{
			Containers: []*entity.WorkloadSpec_Container{},
		},
	})
	assert.NoError(t, err)

	err = s.Create(entity.ObservedWorkload{
		Id: id,
		Spec: &entity.WorkloadSpec{
			Containers: []*entity.WorkloadSpec_Container{},
		},
	})
	assert.Error(t, err, storage.ErrObservedWorkloadIdAlreadyUsed)

	w, err = s.Get(id)
	assert.NotNil(t, w)
	assert.NoError(t, err)
}

func TestObservedWorkload_Get(t *testing.T) {
	var err error
	var w *entity.ObservedWorkload

	s := NewObservedWorkload()

	id := uuid.New().String()

	w, err = s.Get(id)
	assert.Nil(t, w)
	assert.Error(t, err, storage.ErrObservedWorkloadNotFound)

	err = s.Create(entity.ObservedWorkload{
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

func TestObservedWorkload_GetByWorkloadSpecificationMarketId(t *testing.T) {
	var err error
	var w *entity.ObservedWorkload

	s := NewObservedWorkload()

	id := uuid.New().String()
	workloadSpecificationMarketId := randstr.Hex(64)

	w, err = s.Get(id)
	assert.Nil(t, w)
	assert.Error(t, err, storage.ErrObservedWorkloadNotFound)

	err = s.Create(entity.ObservedWorkload{
		Id:                            id,
		WorkloadSpecificationMarketId: workloadSpecificationMarketId,
		Spec: &entity.WorkloadSpec{
			Containers: []*entity.WorkloadSpec_Container{},
		},
	})
	assert.NoError(t, err)

	w, err = s.GetByWorkloadSpecificationMarketId(workloadSpecificationMarketId)
	assert.NotNil(t, w)
	assert.NoError(t, err)
}

func TestObservedWorkload_List(t *testing.T) {
	var err error
	var l []entity.ObservedWorkload

	s := NewObservedWorkload()

	id := uuid.New().String()

	l, err = s.List()
	assert.Len(t, l, 0)
	assert.NoError(t, err)

	err = s.Create(entity.ObservedWorkload{
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

func TestObservedWorkload_SetState(t *testing.T) {
	var err error

	s := NewObservedWorkload()

	id := uuid.New().String()

	err = s.Create(entity.ObservedWorkload{
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

func TestObservedWorkload_SetBackOff(t *testing.T) {
	var err error

	s := NewObservedWorkload()

	id := uuid.NewString()
	err = s.Create(entity.ObservedWorkload{
		Id:    id,
		State: &entity.State{},
	})
	assert.NoError(t, err)

	observedWorkload, err := s.Get(id)
	assert.NoError(t, err)
	assert.NotNil(t, observedWorkload)
	assert.EqualValues(t, observedWorkload.State.BackOffTimestamp, 0)

	err = s.SetBackOff(id, time.Minute)
	assert.NoError(t, err)

	backOff, err := s.GetBackOff(id)
	assert.NoError(t, err)
	assert.NotEqualValues(t, backOff.Unix(), 0)
}

func TestObservedWorkload_ListId(t *testing.T) {
	var err error

	s := NewObservedWorkload()

	id := uuid.NewString()

	list, err := s.ListId()
	assert.NoError(t, err)
	assert.Empty(t, list)

	err = s.Create(entity.ObservedWorkload{Id: id})
	assert.NoError(t, err)

	list, err = s.ListId()
	assert.NoError(t, err)
	assert.NotEmpty(t, list)
	assert.Contains(t, list, id)
}

func TestObservedWorkload_SetPrincipalAcceptance(t *testing.T) {
	var err error

	s := NewObservedWorkload()

	id := uuid.NewString()

	err = s.Create(entity.ObservedWorkload{Id: id})
	assert.NoError(t, err)

	e, err := s.Get(id)
	assert.NoError(t, err)
	assert.Nil(t, e.PrincipalAcceptance)

	err = s.SetPrincipalAcceptance(id, entity.ObservedWorkload_PrincipalAcceptance{
		ContractAcceptMarketId: randstr.Hex(64),
		Accept: &entity.ContractAccept{
			PublicKey: randstr.Hex(128),
		},
	})
	assert.NoError(t, err)

	e, err = s.Get(id)
	assert.NoError(t, err)
	assert.NotNil(t, e.PrincipalAcceptance)
}
