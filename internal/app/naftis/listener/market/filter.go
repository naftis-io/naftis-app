package market

import "gitlab.com/naftis/app/naftis/internal/pkg/market"

type WorkloadSpecificationFilter interface {
	Filter(msg market.WorkloadSpecification) bool
}

type ContractProposalFilter interface {
	Filter(msg market.ContractProposal) bool
}

type ContractAcceptFilter interface {
	Filter(msg market.ContractAccept) bool
}
