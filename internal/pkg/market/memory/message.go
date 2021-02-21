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

	contractRequest         chan marketProtocol.ContractRequest
	contractRequestListener []chan marketProtocol.ContractRequest

	contractResponse         chan marketProtocol.ContractResponse
	contractResponseListener []chan marketProtocol.ContractResponse

	workloadSpecification         chan market.WorkloadSpecification
	workloadSpecificationListener []chan market.WorkloadSpecification
}

func NewMessage() *Message {
	return &Message{
		log: log.With().Str("market", "memory").Logger(),

		contractRequest:         make(chan marketProtocol.ContractRequest, 16),
		contractRequestListener: make([]chan marketProtocol.ContractRequest, 0),

		contractResponse:         make(chan marketProtocol.ContractResponse, 16),
		contractResponseListener: make([]chan marketProtocol.ContractResponse, 0),

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
			case msg := <-m.contractRequest:
				m.notifyContractRequestListeners(ctx, msg)
			case msg := <-m.contractResponse:
				m.notifyContractResponseListeners(ctx, msg)
			}
		}
	}()

	m.log.Info().Msg("Listener started.")

	return nil
}

func (m *Message) ListenContractRequest(ctx context.Context, queueSize uint64) <-chan marketProtocol.ContractRequest {
	panic("implement me")
}

func (m *Message) notifyContractRequestListeners(ctx context.Context, msg marketProtocol.ContractRequest) {

}

func (m *Message) ListenContractResponse(ctx context.Context, queueSize uint64) <-chan marketProtocol.ContractResponse {
	panic("implement me")
}

func (m *Message) notifyContractResponseListeners(ctx context.Context, msg marketProtocol.ContractResponse) {

}

func (m *Message) ListenWorkloadSpecification(ctx context.Context, queueSize uint64) <-chan market.WorkloadSpecification {
	ch := make(chan market.WorkloadSpecification, queueSize)

	m.workloadSpecificationListener = append(m.workloadSpecificationListener, ch)

	return ch
}

func (m *Message) notifyWorkloadSpecificationListeners(ctx context.Context, msg market.WorkloadSpecification) {
	var deliveryCount uint8

	for _, l := range m.workloadSpecificationListener {
		select {
		case l <- msg:
			deliveryCount++
		case <-ctx.Done():
		default:
			m.log.Warn().Msg("Listener queue is full. MessageToken will be NOT delivered to this listener.")
		}
	}
}

func (m *Message) EmitContractRequest(ctx context.Context, msg marketProtocol.ContractRequest) (string, error) {
	panic("implement me")
}

func (m *Message) EmitContractResponse(ctx context.Context, msg marketProtocol.ContractResponse) (string, error) {
	panic("implement me")
}

func (m *Message) EmitWorkloadSpecification(ctx context.Context, msg marketProtocol.WorkloadSpecification) (string, error) {
	txId := randstr.Hex(64)

	select {
	case m.workloadSpecification <- market.WorkloadSpecification{TxId: txId, Msg: msg}:
		return txId, nil
	case <-ctx.Done():
		return "", market.ErrOperationCanceled
	default:
		return "", market.ErrInterfaceBusy
	}
}
