package query

import (
	"gitlab.com/naftis/app/naftis/internal/pkg/storage"
	"gitlab.com/naftis/app/naftis/pkg/protocol/entity"
)

type ListObservedWorkloads struct {
	storage storage.ObservedWorkload
}

func NewListObservedWorkloads(storage storage.ObservedWorkload) *ListObservedWorkloads {
	return &ListObservedWorkloads{
		storage: storage,
	}
}

func (cmd *ListObservedWorkloads) Query() ([]entity.ObservedWorkload, error) {
	list, err := cmd.storage.List()
	if err != nil {
		return []entity.ObservedWorkload{}, err
	}

	return list, nil
}
