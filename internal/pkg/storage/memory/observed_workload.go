package memory

import (
	"gitlab.com/naftis/app/naftis/internal/pkg/state"
	"gitlab.com/naftis/app/naftis/internal/pkg/storage"
	"gitlab.com/naftis/app/naftis/pkg/protocol/entity"
	"sync"
	"time"
)

type ObservedWorkload struct {
	data  map[string]*entity.ObservedWorkload
	mutex *sync.Mutex
}

func NewObservedWorkload() *ObservedWorkload {
	return &ObservedWorkload{
		data:  make(map[string]*entity.ObservedWorkload, 0),
		mutex: &sync.Mutex{},
	}
}

func (o *ObservedWorkload) Create(entity entity.ObservedWorkload) error {
	defer o.mutex.Unlock()
	o.mutex.Lock()

	if _, exists := o.data[entity.Id]; exists {
		return storage.ErrObservedWorkloadIdAlreadyUsed
	}

	o.data[entity.Id] = &entity

	return nil
}

func (o *ObservedWorkload) Get(id string) (*entity.ObservedWorkload, error) {
	defer o.mutex.Unlock()
	o.mutex.Lock()

	if entity, exists := o.data[id]; !exists {
		return nil, storage.ErrObservedWorkloadNotFound
	} else {
		entityCopy := *entity
		return &entityCopy, nil
	}
}

func (o *ObservedWorkload) GetByWorkloadSpecificationMarketId(marketId string) (*entity.ObservedWorkload, error) {
	defer o.mutex.Unlock()
	o.mutex.Lock()

	for _, entity := range o.data {
		if entity.WorkloadSpecificationMarketId != marketId {
			continue
		}
		entityCopy := *entity
		return &entityCopy, nil
	}

	return nil, storage.ErrObservedWorkloadNotFound
}

func (o *ObservedWorkload) List() ([]entity.ObservedWorkload, error) {
	defer o.mutex.Unlock()
	o.mutex.Lock()

	l := make([]entity.ObservedWorkload, 0)

	for _, entity := range o.data {
		entityCopy := *entity
		l = append(l, entityCopy)
	}

	return l, nil
}

func (o *ObservedWorkload) SetState(id string, state string) error {
	defer o.mutex.Unlock()
	o.mutex.Lock()

	if entity, exists := o.data[id]; !exists {
		return storage.ErrObservedWorkloadNotFound
	} else {
		entity.State.Previous = entity.State.Current
		entity.State.Current = state
		return nil
	}
}

func (o *ObservedWorkload) ListId() (state.EntityList, error) {
	defer o.mutex.Unlock()
	o.mutex.Lock()

	list := []string{}

	for id, _ := range o.data {
		list = append(list, id)
	}

	return list, nil
}

func (o *ObservedWorkload) GetState(id string) (string, error) {
	defer o.mutex.Unlock()
	o.mutex.Lock()

	if entity, exists := o.data[id]; !exists {
		return "", storage.ErrObservedWorkloadNotFound
	} else {
		return entity.State.Current, nil
	}
}

func (o *ObservedWorkload) SetBackOff(id string, duration time.Duration) error {
	defer o.mutex.Unlock()
	o.mutex.Lock()

	if entity, exists := o.data[id]; !exists {
		return storage.ErrObservedWorkloadNotFound
	} else {
		entity.State.BackOffTimestamp = time.Now().Add(duration).Unix()
	}

	return nil
}

func (o *ObservedWorkload) GetBackOff(id string) (time.Time, error) {
	defer o.mutex.Unlock()
	o.mutex.Lock()

	if entity, exists := o.data[id]; !exists {
		return time.Unix(0, 0), storage.ErrObservedWorkloadNotFound
	} else {
		return time.Unix(entity.State.BackOffTimestamp, 0), nil
	}
}

func (o *ObservedWorkload) SetPrincipalAcceptance(id string, acceptance entity.ObservedWorkload_PrincipalAcceptance) error {
	defer o.mutex.Unlock()
	o.mutex.Lock()

	_, exists := o.data[id]
	if !exists {
		return storage.ErrObservedWorkloadNotFound
	}

	o.data[id].PrincipalAcceptance = &acceptance

	return nil
}
