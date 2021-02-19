package listener

import (
	"context"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gitlab.com/naftis/app/naftis/internal/app/naftis/command"
	"gitlab.com/naftis/app/naftis/internal/pkg/market"
	"gitlab.com/naftis/app/naftis/pkg/protocol/entity"
)

type WorkloadSpecificationFilterCallback func(msg market.WorkloadSpecification) bool

type Market struct {
	log       zerolog.Logger
	cmd       *command.Factory
	market    market.MessageToken
	queueSize uint64

	workloadSpecification        <-chan market.WorkloadSpecification
	workloadSpecificationFilters []WorkloadSpecificationFilterCallback
}

func NewMarket(cmd *command.Factory, market market.MessageToken, queueSize uint64) *Market {
	return &Market{
		log:                          log.With().Str("listener", "market").Logger(),
		cmd:                          cmd,
		market:                       market,
		queueSize:                    queueSize,
		workloadSpecificationFilters: make([]WorkloadSpecificationFilterCallback, 0),
	}
}

func (m *Market) Start(ctx context.Context) error {
	m.workloadSpecification = m.market.ListenWorkloadSpecification(ctx, m.queueSize)

	go func(ctx context.Context) {
		m.loop(ctx)
	}(ctx)

	m.log.Info().Msg("Listener started.")
	return nil
}

func (m *Market) AttachWorkloadSpecificationFilter(callback WorkloadSpecificationFilterCallback) {
	m.workloadSpecificationFilters = append(m.workloadSpecificationFilters, callback)
}

func (m *Market) loop(ctx context.Context) {
	for {
		select {
		case msg := <-m.workloadSpecification:
			m.processWorkloadSpecification(msg)
		case <-ctx.Done():
			m.log.Info().Msg("Listener loop finished.")
			return
		}
	}
}

func (m *Market) processWorkloadSpecification(msg market.WorkloadSpecification) {
	log := m.log.With().Str("txId", msg.TxId).Logger()

	for _, filter := range m.workloadSpecificationFilters {
		if accept := filter(msg); !accept {
			log.Debug().Msg("Workload specification rejected by filter.")
			return
		}
	}

	observedWorkload := entity.ObservedWorkload{
		Id:   uuid.New().String(),
		TxId: msg.TxId,
		Spec: msg.Msg.Spec,
	}

	err := m.cmd.ObserveWorkload().Invoke(observedWorkload)
	if err != nil {
		log.Error().Err(err).Msg("Unable to observe workload.")
	}
}
