package state

import (
	"context"
	"fmt"
	"github.com/looplab/fsm"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"time"
	"unicode"
)

type EntityStateMap map[string]*fsm.FSM

type Manager struct {
	states        EntityStateMap
	specification Specification
	storage       EntityStateStorage
	log           zerolog.Logger
}

func NewManager(specification Specification, storage EntityStateStorage) *Manager {
	return &Manager{
		storage:       storage,
		states:        EntityStateMap{},
		specification: specification,
		log:           log.With().Str("stateManager", specification.Name).Logger(),
	}
}

func (m *Manager) Start(ctx context.Context) error {
	go func(ctx context.Context) {
		for {
			m.loop(ctx)
			time.Sleep(time.Millisecond * 100)
		}
	}(ctx)

	return nil
}

func (m *Manager) Stop() error {
	return nil
}

func (m *Manager) loop(ctx context.Context) {
	storageList, err := m.storage.ListId()
	if err != nil {
		m.log.Error().Err(err).Msg("Unable to list entities from storage.")
		return
	}

	for _, id := range storageList {
		if !m.states.Contains(id) {
			m.addEntity(ctx, id)
		}
	}

	for id, _ := range m.states {
		if !storageList.Contains(id) {
			m.removeEntity(id)
		}
	}

	for id, _ := range m.states {
		m.transitState(ctx, id)
	}
}

func (m *Manager) addEntity(ctx context.Context, id string) {
	log := m.log.With().Str(m.primaryKeyName(), id).Logger()

	currentState, err := m.storage.GetState(id)
	if err != nil {
		log.Error().Str("currentState", currentState).Err(err).Msg("Unable to add entity to state manager.")
		return
	}

	callbacks := fsm.Callbacks{}

	for eventName, callback := range m.specification.EventCallbacks {
		callbackName := fmt.Sprintf("before_%s", eventName)
		callbackFn := callback
		callbacks[callbackName] = func(event *fsm.Event) {
			err = callbackFn(ctx, id)
			if err != nil {
				event.Cancel(err)
			}
		}
	}

	callbacks["enter_state"] = func(event *fsm.Event) {
		err := m.persistState(id, event.Dst)
		if err != nil {
			log.Warn().Err(err).Msg("Unable to persist state.")
		}
		log.Debug().
			Str("previousState", event.Src).
			Str("currentState", event.Dst).
			Msg("State changed.")
	}

	m.states[id] = fsm.NewFSM(currentState, m.specification.Events, callbacks)

	log.Debug().Msg("Entity added to state manager.")
}

func (m *Manager) removeEntity(id string) {
	log := m.log.With().Str(m.primaryKeyName(), id).Logger()

	delete(m.states, id)

	log.Debug().Msg("Entity removed from state manager.")
}

func (m *Manager) transitState(ctx context.Context, id string) {
	var transitionCallback TransitionCallback

	log := m.log.With().Str(m.primaryKeyName(), id).Logger()

	entity, exists := m.states[id]
	if !exists {
		log.Error().Msg("Transition failed. Entity state not found.")
		return
	}

	backOffUntil, err := m.storage.GetBackOff(id)
	if err != nil {
		log.Error().Err(err).Msg("Unable to read back off duration.")
		return
	}

	if time.Now().Before(backOffUntil) {
		return
	}

	previousState := entity.Current()

	log = log.With().Str("previousState", previousState).Logger()

	if transitionCallback, exists = m.specification.TransitionCallbacks[previousState]; !exists {
		log.Warn().Msg("No transition callback from previous state.")
		return
	}

	eventName, err := transitionCallback(ctx, id)

	log = log.With().Str("eventName", eventName).Logger()

	if err != nil {
		log.Error().Err(err).Msg("Transition failed.")
		return
	}

	if len(eventName) == 0 {
		return
	}

	log.Trace().Msg("Transition started.")

	err = entity.Event(eventName)
	if err != nil {
		backOffDuration := time.Second * 2
		m.storage.SetBackOff(id, backOffDuration)
		log.Error().Err(err).Dur("backOffDuration", backOffDuration).Msg("Transition failed. Backing off.")
		return
	}

	log.Trace().Msg("Transition finished.")
}

func (m *Manager) persistState(id string, state string) error {
	err := m.storage.SetState(id, state)

	return err
}

func (m *Manager) primaryKeyName() string {
	primaryKeyName := []rune(fmt.Sprintf("%sId", m.specification.Name))
	primaryKeyName[0] = unicode.ToLower(primaryKeyName[0])
	return string(primaryKeyName)
}

func (l EntityStateMap) Contains(id string) bool {
	_, exists := l[id]

	return exists
}
