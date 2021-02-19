package query

import "gitlab.com/naftis/app/naftis/internal/pkg/storage"

type Factory struct {
	storage storage.Container
}

func NewFactory(storage storage.Container) *Factory {
	return &Factory{
		storage: storage,
	}
}

func (f *Factory) ListScheduledWorkloads() *ListScheduledWorkloads {
	return NewListScheduledWorkloads(f.storage.ScheduledWorkload())
}
