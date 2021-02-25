package state

import "time"

type EntityList []string

type EntityStateStorage interface {
	ListId() (EntityList, error)
	SetState(id string, state string) error
	GetState(id string) (string, error)
	SetBackOff(id string, duration time.Duration) error
	GetBackOff(id string) (time.Time, error)
}

func (l EntityList) Contains(id string) bool {
	for _, listId := range l {
		if listId == id {
			return true
		}
	}

	return false
}
