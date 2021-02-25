package filter

import (
	"gitlab.com/naftis/app/naftis/internal/pkg/market"
	"gitlab.com/naftis/app/naftis/internal/pkg/storage"
)

type ContractProposalScheduledWorkload struct {
	storage storage.ScheduledWorkload
}

func NewContractProposalScheduledWorkload(storage storage.ScheduledWorkload) *ContractProposalScheduledWorkload {
	return &ContractProposalScheduledWorkload{
		storage: storage,
	}
}

func (c *ContractProposalScheduledWorkload) Filter(msg market.ContractProposal) bool {
	workloadSpecificationMarketId := msg.Msg.WorkloadSpecificationMarketId

	scheduledWorkload, err := c.storage.GetByWorkloadSpecificationMarketId(workloadSpecificationMarketId)
	if err != nil {
		return false
	}

	if scheduledWorkload.State.Current != "waiting_for_runners_proposals" {
		return false
	}

	return true
}
