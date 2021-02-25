package market

import (
	"context"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"gitlab.com/naftis/app/naftis/internal/app/naftis/service"
	"gitlab.com/naftis/app/naftis/internal/pkg/market"
	"gitlab.com/naftis/app/naftis/internal/pkg/validator"
	"gitlab.com/naftis/app/naftis/pkg/protocol/entity"
)

// Listener listens for market events.
type Listener struct {
	log       zerolog.Logger
	market    market.Message
	services  *service.Container
	queueSize uint64

	workloadSpecification        <-chan market.WorkloadSpecification
	workloadSpecificationFilters []WorkloadSpecificationFilter
	contractProposal             <-chan market.ContractProposal
	contractProposalFilters      []ContractProposalFilter
	contractAccept               <-chan market.ContractAccept
	contractAcceptFilters        []ContractAcceptFilter
}

func NewMarket(market market.Message, services *service.Container, queueSize uint64) *Listener {
	return &Listener{
		log:                          log.With().Str("listener", "market").Logger(),
		market:                       market,
		services:                     services,
		queueSize:                    queueSize,
		workloadSpecificationFilters: []WorkloadSpecificationFilter{},
		contractProposalFilters:      []ContractProposalFilter{},
		contractAcceptFilters:        []ContractAcceptFilter{},
	}
}

func (m *Listener) Start(ctx context.Context) error {
	m.workloadSpecification = m.market.ListenWorkloadSpecification(ctx, m.queueSize)
	m.contractProposal = m.market.ListenContractProposal(ctx, m.queueSize)
	m.contractAccept = m.market.ListenContractAccept(ctx, m.queueSize)

	go func(ctx context.Context) {
		m.loop(ctx)
	}(ctx)

	m.log.Info().Msg("Listener started.")
	return nil
}

func (m *Listener) AttachWorkloadSpecificationFilter(filter WorkloadSpecificationFilter) {
	m.workloadSpecificationFilters = append(m.workloadSpecificationFilters, filter)
}

func (m *Listener) AttachContractProposalFilter(filter ContractProposalFilter) {
	m.contractProposalFilters = append(m.contractProposalFilters, filter)
}

func (m *Listener) AttachContractAcceptFilter(filter ContractAcceptFilter) {
	m.contractAcceptFilters = append(m.contractAcceptFilters, filter)
}

func (m *Listener) loop(ctx context.Context) {
	for {
		select {
		case msg := <-m.workloadSpecification:
			m.processWorkloadSpecification(ctx, msg)
		case msg := <-m.contractProposal:
			m.processContractProposal(ctx, msg)
		case msg := <-m.contractAccept:
			m.processContractAccept(ctx, msg)
		case <-ctx.Done():
			m.log.Info().Msg("Listener loop finished.")
			return
		}
	}
}

func (m *Listener) processContractAccept(ctx context.Context, msg market.ContractAccept) {
	log := m.log.With().Str("contractAcceptMarketId", msg.MarketId).Logger()

	err := validator.New().Struct(msg)
	if err != nil {
		log.Error().Err(err).Msg("ContractAccept rejected during validation.")
		return
	}

	for _, filter := range m.contractAcceptFilters {
		if accept := filter.Filter(msg); !accept {
			log.Trace().Msg("ContractAccept rejected by filter.")
			return
		}
	}

	err = m.services.ObservedWorkload().ConfirmPrincipalAcceptance(ctx, msg.Msg.WorkloadSpecificationMarketId, msg.MarketId, *msg.Msg.Accept)
	if err != nil {
		log.Error().Err(err).Msg("Unable to select ObservedWorkload for run.")
	}

	log.Trace().Msg("Processed ContractAccept from market.")
}

func (m *Listener) processContractProposal(ctx context.Context, msg market.ContractProposal) {
	log := m.log.With().Str("contractProposalMarketId", msg.MarketId).Logger()

	err := validator.New().Struct(msg)
	if err != nil {
		log.Error().Err(err).Msg("ContractProposal rejected during validation.")
		return
	}

	for _, filter := range m.contractProposalFilters {
		if accept := filter.Filter(msg); !accept {
			log.Debug().Msg("ContractProposal rejected by filter.")
			return
		}
	}

	err = m.services.ScheduledWorkload().AddContractProposalFromRunner(ctx, msg.Msg.WorkloadSpecificationMarketId, msg.MarketId, *msg.Msg.Proposal)
	if err != nil {
		log.Error().Err(err).Msg("Unable to add runner contract proposal to scheduled workload.")
	}

	log.Trace().Msg("Processed ContractProposal from market.")
}

func (m *Listener) processWorkloadSpecification(ctx context.Context, msg market.WorkloadSpecification) {
	log := m.log.With().Str("workloadSpecificationMarketId", msg.MarketId).Logger()

	err := validator.New().Struct(msg)
	if err != nil {
		log.Error().Err(err).Msg("WorkloadSpecification rejected during validation.")
		return
	}

	for _, filter := range m.workloadSpecificationFilters {
		if accept := filter.Filter(msg); !accept {
			log.Debug().Msg("WorkloadSpecification rejected by filter.")
			return
		}
	}

	observedWorkload := entity.ObservedWorkload{
		Id:                            uuid.NewString(),
		WorkloadSpecificationMarketId: msg.MarketId,
		Spec:                          msg.Msg.Spec,
		PrincipalProposal:             msg.Msg.PrincipalProposal,
	}

	err = m.services.ObservedWorkload().Observe(ctx, observedWorkload)
	if err != nil {
		log.Error().Err(err).Msg("Unable to observe workload.")
	}

	log.Trace().Msg("Processed WorkloadSpecification from market.")
}
