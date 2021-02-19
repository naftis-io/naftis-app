package amqp

import (
	"context"
	"gitlab.com/naftis/app/naftis/pkg/protocol/blockchain"
)

type Message struct {
}

func NewMessage() *Message {
	return &Message{}
}

func (m *Message) Start(ctx context.Context) error {
	panic("implement me")
}

func (m *Message) ListenContractRequest(ctx context.Context, queueSize uint64) <-chan blockchain.ContractRequest {
	panic("implement me")
}

func (m *Message) ListenContractResponse(ctx context.Context, queueSize uint64) <-chan blockchain.ContractResponse {
	panic("implement me")
}

func (m *Message) ListenWorkloadSpecification(ctx context.Context, queueSize uint64) <-chan blockchain.WorkloadSpecification {
	panic("implement me")
}

func (m *Message) EmitContractRequest(ctx context.Context, msg blockchain.ContractRequest) (string, error) {
	panic("implement me")
}

func (m *Message) EmitContractResponse(ctx context.Context, msg blockchain.ContractResponse) (string, error) {
	panic("implement me")
}

func (m *Message) EmitWorkloadSpecification(ctx context.Context, msg blockchain.WorkloadSpecification) (string, error) {
	panic("implement me")
}
