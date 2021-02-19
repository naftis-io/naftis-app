package command

import "gitlab.com/naftis/app/naftis/internal/pkg/storage"

type Factory struct {
	storage storage.Container
}

func NewFactory(storage storage.Container) *Factory {
	return &Factory{
		storage: storage,
	}
}

func (f *Factory) ScheduleWorkload() *ScheduleWorkload {
	return NewScheduleWorkload(f.storage.ScheduledWorkload())
}

func (f *Factory) SetReplicaCount() *SetReplicaCount {
	return NewSetReplicaCount()
}

func (f *Factory) ForgetScheduledWorkload() *ForgetScheduledWorkload {
	return NewForgetScheduleWorkload()
}

func (f *Factory) RunWorkload() *RunWorkload {
	return NewRunWorkload()
}

func (f *Factory) ObserveWorkload() *ObserveWorkload {
	return NewObserveWorkload(f.storage.ObservedWorkload())
}
