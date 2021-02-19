package memory

import (
	"gitlab.com/naftis/app/naftis/internal/pkg/storage"
	"gitlab.com/naftis/app/naftis/pkg/protocol/entity"
)

type ObservedWorkload struct {
	data map[string]*entity.ObservedWorkload
}

func NewObservedWorkload() *ObservedWorkload {
	return &ObservedWorkload{
		data: make(map[string]*entity.ObservedWorkload, 0),
	}
}

func (o *ObservedWorkload) Create(entity entity.ObservedWorkload) error {
	if _, exists := o.data[entity.Id]; exists {
		return storage.ErrObservedWorkloadIdAlreadyUsed
	}

	o.data[entity.Id] = &entity

	return nil
}

func (o *ObservedWorkload) Get(id string) (*entity.ObservedWorkload, error) {
	if entity, exists := o.data[id]; !exists {
		return nil, storage.ErrObservedWorkloadNotFound
	} else {
		entityCopy := *entity
		return &entityCopy, nil
	}
}

func (o *ObservedWorkload) GetByTxId(txId string) (*entity.ObservedWorkload, error) {
	for _, entity := range o.data {
		if entity.TxId != txId {
			continue
		}
		entityCopy := *entity
		return &entityCopy, nil
	}

	return nil, storage.ErrObservedWorkloadNotFound
}

func (o *ObservedWorkload) List() ([]entity.ObservedWorkload, error) {
	l := make([]entity.ObservedWorkload, 0)

	for _, entity := range o.data {
		entityCopy := *entity
		l = append(l, entityCopy)
	}

	return l, nil
}

func (o *ObservedWorkload) UpdateState(id string, state string) error {
	if entity, exists := o.data[id]; !exists {
		return storage.ErrObservedWorkloadNotFound
	} else {
		entity.State = state
		return nil
	}
}
