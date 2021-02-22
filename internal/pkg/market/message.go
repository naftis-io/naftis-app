package market

import (
	"context"
	"errors"
	"gitlab.com/naftis/app/naftis/pkg/protocol/market"
)

type WorkloadSpecification struct {
	TxId string
	Msg  market.WorkloadSpecification
}

// Message interface describe minimum function set that allow sending messages to market.
type Message interface {
	Start(ctx context.Context) error

	EmitContractRequest(ctx context.Context, msg market.ContractAccept) (string, error)
	EmitContractResponse(ctx context.Context, msg market.ContractProposal) (string, error)
	EmitWorkloadSpecification(ctx context.Context, msg market.WorkloadSpecification) (string, error)

	ListenContractRequest(ctx context.Context, queueSize uint64) <-chan market.ContractAccept
	ListenContractResponse(ctx context.Context, queueSize uint64) <-chan market.ContractProposal
	ListenWorkloadSpecification(ctx context.Context, queueSize uint64) <-chan WorkloadSpecification
}

var (
	ErrInterfaceBusy     = errors.New("market interface is busy")
	ErrOperationCanceled = errors.New("market operation canceled")
)
