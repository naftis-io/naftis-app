package reconciler

import (
	"context"
	"github.com/looplab/fsm"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gitlab.com/naftis/app/naftis/internal/pkg/market"
	"gitlab.com/naftis/app/naftis/internal/pkg/storage"
	"gitlab.com/naftis/app/naftis/pkg/protocol/entity"
	marketProtocol "gitlab.com/naftis/app/naftis/pkg/protocol/market"
	"time"
)

type scheduledWorkloadState struct {
	fsm.FSM
	reconciler *ScheduledWorkload
	entity     entity.ScheduledWorkload
	log        zerolog.Logger
}

type ScheduledWorkload struct {
	ctx     context.Context
	log     zerolog.Logger
	storage storage.Container
	market  market.MessageToken
	state   map[string]*scheduledWorkloadState
}

func NewScheduledWorkload(storage storage.Container, market market.MessageToken) *ScheduledWorkload {
	return &ScheduledWorkload{
		log:     log.With().Str("reconciler", "ScheduledWorkload").Logger(),
		storage: storage,
		market:  market,
		ctx:     context.TODO(),
		state:   make(map[string]*scheduledWorkloadState, 0),
	}
}

func (p *ScheduledWorkload) Start(ctx context.Context) error {
	p.ctx = ctx
	go func() {
		for {
			err := p.loop()
			if err != nil {
				p.log.Warn().Err(err).Msg("Error while reconcilation loop pass.")
			}

			time.Sleep(time.Millisecond * 100)
		}
	}()

	p.log.Info().Msg("Reconciler started.")

	return nil
}

func (p *ScheduledWorkload) Stop() error {
	return nil
}

func (p *ScheduledWorkload) loop() error {
	l, err := p.storage.ScheduledWorkload().List()
	if err != nil {
		return err
	}

	for _, entity := range l {
		err = p.loopEntity(entity)
		if err != nil {
			p.log.Warn().
				Str("scheduledWorkloadId", entity.Id).
				Err(err).
				Msg("Failed to reconcile entity.")
		}
	}

	for _, state := range p.state {
		err = p.loopState(state)
		if err != nil {
			p.log.Warn().
				Str("scheduledWorkloadId", state.entity.Id).
				Err(err).
				Msg("Failed to reconcile entity state.")
		}
	}

	return nil
}

func (p *ScheduledWorkload) loopEntity(entity entity.ScheduledWorkload) error {
	log := p.log.With().
		Str("scheduledEntityId", entity.Id).
		Logger()

	if _, exists := p.state[entity.Id]; !exists {
		p.state[entity.Id] = createScheduledWorkloadState(entity, p)
		log.Debug().Msg("Entity added to reconciler.")
		return nil
	}

	return nil
}

func (p *ScheduledWorkload) loopState(state *scheduledWorkloadState) error {
	if state.Current() == "new" {
		return state.Event("publish_on_market")
	}

	return nil
}

func createScheduledWorkloadState(entity entity.ScheduledWorkload, reconciler *ScheduledWorkload) *scheduledWorkloadState {
	state := &scheduledWorkloadState{}
	state.entity = entity
	state.reconciler = reconciler
	state.log = reconciler.log.With().Str("scheduledWorkloadId", entity.Id).Logger()
	state.FSM = *fsm.NewFSM(entity.State,
		fsm.Events{
			{Name: "publish_on_market", Src: []string{"new"}, Dst: "published_on_market"},
		}, fsm.Callbacks{
			"enter_state":              func(event *fsm.Event) { state.onEnterState(event) },
			"before_publish_on_market": func(event *fsm.Event) { state.onPublishOnMarket(event) },
		})

	return state
}

func (s *scheduledWorkloadState) onEnterState(event *fsm.Event) {
	log := s.log.With().
		Str("previousState", event.Src).
		Str("currentState", event.Dst).
		Logger()

	err := s.reconciler.storage.ScheduledWorkload().UpdateState(s.entity.Id, event.Dst)
	if err != nil {
		log.Err(err).Msg("Failed to persist changed state.")
		return
	}

	log.Debug().Msg("State changed.")
}

func (s *scheduledWorkloadState) onPublishOnMarket(event *fsm.Event) {
	txId, err := s.reconciler.market.EmitWorkloadSpecification(s.reconciler.ctx, marketProtocol.WorkloadSpecification{
		Spec:               s.entity.Spec,
		PrincipalPublicKey: "",
	})
	if err != nil {
		event.Cancel(err)
		return
	}

	err = s.reconciler.storage.ScheduledWorkload().UpdateTxId(s.entity.Id, txId)
	if err != nil {
		event.Cancel(err)
		return
	}

	s.log.Info().Str("txId", txId).Msg("Published on market.")
}
