package amqp

import (
	"context"
	"gitlab.com/naftis/app/naftis/pkg/protocol/market"
)

type Message struct {
}

func NewMessage() *Message {
	return &Message{}
}

func (m *Message) Start(ctx context.Context) error {
	panic("implement me")
}

func (m *Message) ListenContractRequest(ctx context.Context, queueSize uint64) <-chan market.ContractAccept {
	panic("implement me")
}

func (m *Message) ListenContractResponse(ctx context.Context, queueSize uint64) <-chan market.ContractProposal {
	panic("implement me")
}

func (m *Message) ListenWorkloadSpecification(ctx context.Context, queueSize uint64) <-chan market.WorkloadSpecification {
	panic("implement me")
}

func (m *Message) EmitContractRequest(ctx context.Context, msg market.ContractAccept) (string, error) {
	panic("implement me")
}

func (m *Message) EmitContractResponse(ctx context.Context, msg market.ContractProposal) (string, error) {
	panic("implement me")
}

func (m *Message) EmitWorkloadSpecification(ctx context.Context, msg market.WorkloadSpecification) (string, error) {
	panic("implement me")
}
