package storage

type Container interface {
	ScheduledWorkload() ScheduledWorkload
	ObservedWorkload() ObservedWorkload
}
