package memory

import (
	"context"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/thanhpk/randstr"
	"gitlab.com/naftis/app/naftis/internal/pkg/market"
	"gitlab.com/naftis/app/naftis/pkg/protocol/blockchain"
)

type Message struct {
	log zerolog.Logger

	contractRequest         chan blockchain.ContractRequest
	contractRequestListener []chan blockchain.ContractRequest

	contractResponse         chan blockchain.ContractResponse
	contractResponseListener []chan blockchain.ContractResponse

	workloadSpecification         chan market.WorkloadSpecification
	workloadSpecificationListener []chan market.WorkloadSpecification
}

func NewMessage() *Message {
	return &Message{
		log: log.With().Str("market", "memory").Logger(),

		contractRequest:         make(chan blockchain.ContractRequest, 16),
		contractRequestListener: make([]chan blockchain.ContractRequest, 0),

		contractResponse:         make(chan blockchain.ContractResponse, 16),
		contractResponseListener: make([]chan blockchain.ContractResponse, 0),

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

	m.log.Info().Msg("Market started.")

	return nil
}

func (m *Message) ListenContractRequest(ctx context.Context, queueSize uint64) <-chan blockchain.ContractRequest {
	panic("implement me")
}

func (m *Message) notifyContractRequestListeners(ctx context.Context, msg blockchain.ContractRequest) {

}

func (m *Message) ListenContractResponse(ctx context.Context, queueSize uint64) <-chan blockchain.ContractResponse {
	panic("implement me")
}

func (m *Message) notifyContractResponseListeners(ctx context.Context, msg blockchain.ContractResponse) {

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

func (m *Message) EmitContractRequest(ctx context.Context, msg blockchain.ContractRequest) (string, error) {
	panic("implement me")
}

func (m *Message) EmitContractResponse(ctx context.Context, msg blockchain.ContractResponse) (string, error) {
	panic("implement me")
}

func (m *Message) EmitWorkloadSpecification(ctx context.Context, msg blockchain.WorkloadSpecification) (string, error) {
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
