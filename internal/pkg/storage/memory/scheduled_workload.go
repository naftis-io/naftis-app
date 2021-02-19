package memory

import (
	"gitlab.com/naftis/app/naftis/internal/pkg/storage"
	"gitlab.com/naftis/app/naftis/pkg/protocol/entity"
)

type ScheduledWorkload struct {
	data map[string]*entity.ScheduledWorkload
}

func NewScheduledWorkload() *ScheduledWorkload {
	return &ScheduledWorkload{
		data: make(map[string]*entity.ScheduledWorkload, 0),
	}
}

func (w *ScheduledWorkload) Create(entity entity.ScheduledWorkload) error {
	if _, exists := w.data[entity.Id]; exists {
		return storage.ErrScheduledWorkloadIdAlreadyUsed
	}

	w.data[entity.Id] = &entity

	return nil
}

func (w *ScheduledWorkload) Get(id string) (*entity.ScheduledWorkload, error) {
	if entity, exists := w.data[id]; !exists {
		return nil, storage.ErrScheduledWorkloadNotFound
	} else {
		entityCopy := *entity
		return &entityCopy, nil
	}
}

func (w *ScheduledWorkload) List() ([]entity.ScheduledWorkload, error) {
	l := make([]entity.ScheduledWorkload, 0)

	for _, entity := range w.data {
		entityCopy := *entity
		l = append(l, entityCopy)
	}

	return l, nil
}

func (w *ScheduledWorkload) UpdateTxId(id string, txId string) error {
	if entity, exists := w.data[id]; !exists {
		return storage.ErrScheduledWorkloadNotFound
	} else {
		entity.TxId = txId
		return nil
	}
}

func (w *ScheduledWorkload) UpdateState(id string, state string) error {
	if entity, exists := w.data[id]; !exists {
		return storage.ErrScheduledWorkloadNotFound
	} else {
		entity.State = state
		return nil
	}
}
