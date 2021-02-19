package market

import (
	"context"
	"errors"
	"gitlab.com/naftis/app/naftis/pkg/protocol/blockchain"
)

type WorkloadSpecification struct {
	TxId string
	Msg  blockchain.WorkloadSpecification
}

// Message interface describe minimum function set that allow sending messages to market.
type Message interface {
	Start(ctx context.Context) error

	EmitContractRequest(ctx context.Context, msg blockchain.ContractRequest) (string, error)
	EmitContractResponse(ctx context.Context, msg blockchain.ContractResponse) (string, error)
	EmitWorkloadSpecification(ctx context.Context, msg blockchain.WorkloadSpecification) (string, error)

	ListenContractRequest(ctx context.Context, queueSize uint64) <-chan blockchain.ContractRequest
	ListenContractResponse(ctx context.Context, queueSize uint64) <-chan blockchain.ContractResponse
	ListenWorkloadSpecification(ctx context.Context, queueSize uint64) <-chan WorkloadSpecification
}

var (
	ErrInterfaceBusy     = errors.New("market interface is busy")
	ErrOperationCanceled = errors.New("market operation canceled")
)
