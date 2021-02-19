package query

import (
	"gitlab.com/naftis/app/naftis/internal/pkg/storage"
	"gitlab.com/naftis/app/naftis/pkg/protocol/entity"
)

type ListScheduledWorkloads struct {
	storage storage.ScheduledWorkload
}

func NewListScheduledWorkloads(storage storage.ScheduledWorkload) *ListScheduledWorkloads {
	return &ListScheduledWorkloads{
		storage: storage,
	}
}

func (cmd *ListScheduledWorkloads) Query() ([]entity.ScheduledWorkload, error) {
	list, err := cmd.storage.List()
	if err != nil {
		return []entity.ScheduledWorkload{}, err
	}

	return list, nil
}
