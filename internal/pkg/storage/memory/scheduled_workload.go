package memory

import (
	"gitlab.com/naftis/app/naftis/internal/pkg/state"
	"gitlab.com/naftis/app/naftis/internal/pkg/storage"
	"gitlab.com/naftis/app/naftis/pkg/protocol/entity"
	"sync"
	"time"
)

type ScheduledWorkload struct {
	data  map[string]*entity.ScheduledWorkload
	mutex *sync.Mutex
}

func NewScheduledWorkload() *ScheduledWorkload {
	return &ScheduledWorkload{
		data:  make(map[string]*entity.ScheduledWorkload, 0),
		mutex: &sync.Mutex{},
	}
}

func (w *ScheduledWorkload) Create(newEntity entity.ScheduledWorkload) error {
	defer w.mutex.Unlock()
	w.mutex.Lock()

	if _, exists := w.data[newEntity.Id]; exists {
		return storage.ErrScheduledWorkloadIdAlreadyUsed
	}

	entityCopy := newEntity

	w.data[entityCopy.Id] = &entityCopy

	return nil
}

func (w *ScheduledWorkload) Get(id string) (*entity.ScheduledWorkload, error) {
	defer w.mutex.Unlock()
	w.mutex.Lock()

	if entity, exists := w.data[id]; !exists {
		return nil, storage.ErrScheduledWorkloadNotFound
	} else {
		entityCopy := *entity
		return &entityCopy, nil
	}
}

func (w *ScheduledWorkload) List() ([]entity.ScheduledWorkload, error) {
	defer w.mutex.Unlock()
	w.mutex.Lock()

	l := make([]entity.ScheduledWorkload, 0)

	for _, entity := range w.data {
		entityCopy := *entity
		l = append(l, entityCopy)
	}

	return l, nil
}

func (w *ScheduledWorkload) UpdateWorkloadSpecificationMarketId(id string, marketId string) error {
	defer w.mutex.Unlock()
	w.mutex.Lock()

	if entity, exists := w.data[id]; !exists {
		return storage.ErrScheduledWorkloadNotFound
	} else {
		entity.WorkloadSpecificationMarketId = marketId
		return nil
	}
}

func (w *ScheduledWorkload) SetState(id string, state string) error {
	defer w.mutex.Unlock()
	w.mutex.Lock()

	if entity, exists := w.data[id]; !exists {
		return storage.ErrScheduledWorkloadNotFound
	} else {
		entity.State.Previous = entity.State.Current
		entity.State.Current = state
		return nil
	}
}

func (w *ScheduledWorkload) AddRunnerContractProposal(id string, proposal entity.ScheduledWorkload_RunnerProposal) error {
	defer w.mutex.Unlock()
	w.mutex.Lock()

	_, exists := w.data[id]
	if !exists {
		return storage.ErrScheduledWorkloadNotFound
	}

	w.data[id].RunnersProposals = append(w.data[id].RunnersProposals, &proposal)

	return nil
}

func (w *ScheduledWorkload) ListRunnerContractProposals(id string) ([]*entity.ScheduledWorkload_RunnerProposal, error) {
	defer w.mutex.Unlock()
	w.mutex.Lock()

	if item, exists := w.data[id]; !exists {
		return []*entity.ScheduledWorkload_RunnerProposal{}, storage.ErrScheduledWorkloadNotFound
	} else {
		return item.RunnersProposals, nil
	}
}

func (w *ScheduledWorkload) GetByWorkloadSpecificationMarketId(marketId string) (*entity.ScheduledWorkload, error) {
	defer w.mutex.Unlock()
	w.mutex.Lock()

	for _, entity := range w.data {
		if entity.WorkloadSpecificationMarketId != marketId {
			continue
		}

		entityCopy := *entity
		return &entityCopy, nil
	}

	return nil, storage.ErrScheduledWorkloadNotFound
}

func (w *ScheduledWorkload) SetBackOff(id string, duration time.Duration) error {
	defer w.mutex.Unlock()
	w.mutex.Lock()

	if entity, exists := w.data[id]; !exists {
		return storage.ErrScheduledWorkloadNotFound
	} else {
		entity.State.BackOffTimestamp = time.Now().Add(duration).Unix()
	}

	return nil
}

func (w *ScheduledWorkload) ListId() (state.EntityList, error) {
	defer w.mutex.Unlock()
	w.mutex.Lock()

	list := []string{}

	for id, _ := range w.data {
		list = append(list, id)
	}

	return list, nil
}

func (w *ScheduledWorkload) GetState(id string) (string, error) {
	defer w.mutex.Unlock()
	w.mutex.Lock()

	if entity, exists := w.data[id]; !exists {
		return "", storage.ErrScheduledWorkloadNotFound
	} else {
		return entity.State.Current, nil
	}
}

func (w *ScheduledWorkload) GetBackOff(id string) (time.Time, error) {
	defer w.mutex.Unlock()
	w.mutex.Lock()

	if entity, exists := w.data[id]; !exists {
		return time.Unix(0, 0), storage.ErrScheduledWorkloadNotFound
	} else {
		return time.Unix(entity.State.BackOffTimestamp, 0), nil
	}
}

func (w *ScheduledWorkload) SetAcceptedRunnerContractProposal(id string, proposal entity.ScheduledWorkload_RunnerProposal) error {
	defer w.mutex.Unlock()
	w.mutex.Lock()

	if _, exists := w.data[id]; !exists {
		return storage.ErrScheduledWorkloadNotFound
	}

	w.data[id].AcceptedRunnerProposal = &proposal

	return nil
}
