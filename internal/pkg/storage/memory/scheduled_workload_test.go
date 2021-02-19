package memory

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/thanhpk/randstr"
	"gitlab.com/naftis/app/naftis/internal/pkg/storage"
	"gitlab.com/naftis/app/naftis/pkg/protocol/entity"
	"testing"
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

func TestScheduledWorkload_UpdateTxId(t *testing.T) {
	var err error

	s := NewScheduledWorkload()

	id := uuid.New().String()

	err = s.Create(entity.ScheduledWorkload{
		Id: id,
		Spec: &entity.WorkloadSpec{
			Containers: []*entity.WorkloadSpec_Container{},
		},
	})
	assert.NoError(t, err)

	e, err := s.Get(id)
	assert.NoError(t, err)
	assert.NotNil(t, e)
	assert.Empty(t, e.TxId)

	txId := randstr.Hex(64)

	err = s.UpdateTxId(id, txId)
	assert.NoError(t, err)

	e, err = s.Get(id)
	assert.NoError(t, err)
	assert.NotNil(t, e)
	assert.Equal(t, txId, e.TxId)
}

func TestScheduledWorkload_UpdateState(t *testing.T) {
	var err error

	s := NewScheduledWorkload()

	id := uuid.New().String()

	err = s.Create(entity.ScheduledWorkload{
		Id: id,
		Spec: &entity.WorkloadSpec{
			Containers: []*entity.WorkloadSpec_Container{},
		},
	})
	assert.NoError(t, err)

	e, err := s.Get(id)
	assert.NoError(t, err)
	assert.NotNil(t, e)
	assert.Empty(t, e.State)

	newState := "test"

	err = s.UpdateState(id, newState)
	assert.NoError(t, err)

	e, err = s.Get(id)
	assert.NoError(t, err)
	assert.NotNil(t, e)
	assert.Equal(t, newState, e.State)
}
