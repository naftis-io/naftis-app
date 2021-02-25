package memory

import (
	"context"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/thanhpk/randstr"
	"gitlab.com/naftis/app/naftis/internal/pkg/market"
	marketProtocol "gitlab.com/naftis/app/naftis/pkg/protocol/market"
)

type Message struct {
	log zerolog.Logger

	contractAccept         chan market.ContractAccept
	contractAcceptListener []chan market.ContractAccept

	contractProposal         chan market.ContractProposal
	contractProposalListener []chan market.ContractProposal

	workloadSpecification         chan market.WorkloadSpecification
	workloadSpecificationListener []chan market.WorkloadSpecification
}

func NewMessage() *Message {
	return &Message{
		log: log.With().Str("market", "memory").Logger(),

		contractAccept:         make(chan market.ContractAccept, 16),
		contractAcceptListener: make([]chan market.ContractAccept, 0),

		contractProposal:         make(chan market.ContractProposal, 16),
		contractProposalListener: make([]chan market.ContractProposal, 0),

		workloadSpecification:         make(chan market.WorkloadSpecification, 16),
		workloadSpecificationListener: make([]chan market.WorkloadSpecification, 0),
	}
}

func (m *Message) Start(ctx context.Context) error {
	go func() {
		for {
			select {
			case msg := <-m.workloadSpecification:
				m.notifyWorkloadSpecificationListeners(ctx, msg)
			case msg := <-m.contractAccept:
				m.notifyContractAcceptListeners(ctx, msg)
			case msg := <-m.contractProposal:
				m.notifyContractProposalListeners(ctx, msg)
			}
		}
	}()

	m.log.Info().Msg("Listener started.")

	return nil
}

func (m *Message) ListenContractAccept(ctx context.Context, queueSize uint64) <-chan market.ContractAccept {
	ch := make(chan market.ContractAccept, queueSize)

	m.contractAcceptListener = append(m.contractAcceptListener, ch)

	return ch
}

func (m *Message) notifyContractAcceptListeners(ctx context.Context, msg market.ContractAccept) {
	for _, l := range m.contractAcceptListener {
		select {
		case l <- msg:
		case <-ctx.Done():
		default:
			m.log.Warn().Msg("Listener queue is full. ContractAccept will be NOT delivered to this listener.")
		}
	}
}

func (m *Message) ListenContractProposal(ctx context.Context, queueSize uint64) <-chan market.ContractProposal {
	ch := make(chan market.ContractProposal, queueSize)

	m.contractProposalListener = append(m.contractProposalListener, ch)

	return ch
}

func (m *Message) notifyContractProposalListeners(ctx context.Context, msg market.ContractProposal) {
	for _, l := range m.contractProposalListener {
		select {
		case l <- msg:
		case <-ctx.Done():
		default:
			m.log.Warn().Msg("Listener queue is full. ContractProposal will be NOT delivered to this listener.")
		}
	}
}

func (m *Message) ListenWorkloadSpecification(ctx context.Context, queueSize uint64) <-chan market.WorkloadSpecification {
	ch := make(chan market.WorkloadSpecification, queueSize)

	m.workloadSpecificationListener = append(m.workloadSpecificationListener, ch)

	return ch
}

func (m *Message) notifyWorkloadSpecificationListeners(ctx context.Context, msg market.WorkloadSpecification) {
	for _, l := range m.workloadSpecificationListener {
		select {
		case l <- msg:
		case <-ctx.Done():
		default:
			m.log.Warn().Msg("Listener queue is full. WorkloadSpecification will be NOT delivered to this listener.")
		}
	}
}

func (m *Message) EmitContractAccept(ctx context.Context, msg marketProtocol.ContractAccept) (string, error) {
	marketId := randstr.Hex(64)

	select {
	case m.contractAccept <- market.ContractAccept{MarketId: marketId, Msg: msg}:
		return marketId, nil
	case <-ctx.Done():
		return "", market.ErrOperationCanceled
	default:
		return "", market.ErrInterfaceBusy
	}
}

func (m *Message) EmitContractProposal(ctx context.Context, msg marketProtocol.ContractProposal) (string, error) {
	marketId := randstr.Hex(64)

	select {
	case m.contractProposal <- market.ContractProposal{MarketId: marketId, Msg: msg}:
		return marketId, nil
	case <-ctx.Done():
		return "", market.ErrOperationCanceled
	default:
		return "", market.ErrInterfaceBusy
	}
}

func (m *Message) EmitWorkloadSpecification(ctx context.Context, msg marketProtocol.WorkloadSpecification) (string, error) {
	marketId := randstr.Hex(64)

	select {
	case m.workloadSpecification <- market.WorkloadSpecification{MarketId: marketId, Msg: msg}:
		return marketId, nil
	case <-ctx.Done():
		return "", market.ErrOperationCanceled
	default:
		return "", market.ErrInterfaceBusy
	}
}
