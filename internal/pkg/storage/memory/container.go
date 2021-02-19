package memory

import "gitlab.com/naftis/app/naftis/internal/pkg/storage"

type Container struct {
	scheduledWorkload *ScheduledWorkload
	observedWorkload  *ObservedWorkload
}

func NewContainer() *Container {
	return &Container{
		scheduledWorkload: NewScheduledWorkload(),
		observedWorkload:  NewObservedWorkload(),
	}
}

func (c *Container) ScheduledWorkload() storage.ScheduledWorkload {
	return c.scheduledWorkload
}

func (c *Container) ObservedWorkload() storage.ObservedWorkload {
	return c.observedWorkload
}
