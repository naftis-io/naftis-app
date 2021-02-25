package market

import (
	"context"
	"errors"
	marketProtocol "gitlab.com/naftis/app/naftis/pkg/protocol/market"
)

type WorkloadSpecification struct {
	MarketId string
	Msg      marketProtocol.WorkloadSpecification
}

type ContractAccept struct {
	MarketId string
	Msg      marketProtocol.ContractAccept
}

type ContractProposal struct {
	MarketId string
	Msg      marketProtocol.ContractProposal
}

// Message interface describe minimum function set that allow sending messages to marketProtocol.
type Message interface {
	Start(ctx context.Context) error

	EmitContractAccept(ctx context.Context, msg marketProtocol.ContractAccept) (string, error)
	EmitContractProposal(ctx context.Context, msg marketProtocol.ContractProposal) (string, error)
	EmitWorkloadSpecification(ctx context.Context, msg marketProtocol.WorkloadSpecification) (string, error)

	ListenContractAccept(ctx context.Context, queueSize uint64) <-chan ContractAccept
	ListenContractProposal(ctx context.Context, queueSize uint64) <-chan ContractProposal
	ListenWorkloadSpecification(ctx context.Context, queueSize uint64) <-chan WorkloadSpecification
}

var (
	ErrInterfaceBusy     = errors.New("marketProtocol interface is busy")
	ErrOperationCanceled = errors.New("marketProtocol operation canceled")
)
