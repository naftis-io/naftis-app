package filter

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/thanhpk/randstr"
	"gitlab.com/naftis/app/naftis/internal/pkg/market"
	memoryStorage "gitlab.com/naftis/app/naftis/internal/pkg/storage/memory"
	"gitlab.com/naftis/app/naftis/pkg/protocol/entity"
	marketProtocol "gitlab.com/naftis/app/naftis/pkg/protocol/market"
	"testing"
)

func TestContractProposalScheduledWorkload_Filter(t *testing.T) {
	storage := memoryStorage.NewScheduledWorkload()

	filter := NewContractProposalScheduledWorkload(storage)

	scheduledWorkloadId := uuid.NewString()
	workloadSpecificationMarketId := randstr.Hex(64)
	scheduledWorkload := entity.ScheduledWorkload{
		Id:                            scheduledWorkloadId,
		WorkloadSpecificationMarketId: workloadSpecificationMarketId,
		State: &entity.State{
			Current:          "new",
			Previous:         "new",
			BackOffTimestamp: 0,
		},
	}

	contractProposal := market.ContractProposal{
		MarketId: randstr.Hex(64),
		Msg: marketProtocol.ContractProposal{
			Proposal:                      nil,
			WorkloadSpecificationMarketId: workloadSpecificationMarketId,
		},
	}

	result := filter.Filter(contractProposal)
	assert.False(t, result)

	err := storage.Create(scheduledWorkload)
	assert.NoError(t, err)

	result = filter.Filter(contractProposal)
	assert.False(t, result)

	err = storage.SetState(scheduledWorkloadId, "waiting_for_runners_proposals")
	assert.NoError(t, err)

	result = filter.Filter(contractProposal)
	assert.True(t, result)
}
