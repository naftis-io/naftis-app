package command

import (
	"gitlab.com/naftis/app/naftis/internal/app/naftis/service"
)

type Factory struct {
	services *service.Container
}

func NewFactory(services *service.Container) *Factory {
	return &Factory{
		services: services,
	}
}

func (f *Factory) ScheduleWorkload() *ScheduleWorkload {
	return NewScheduleWorkload(f.services.ScheduledWorkload())
}
