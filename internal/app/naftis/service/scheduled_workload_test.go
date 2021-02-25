package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/thanhpk/randstr"
	"gitlab.com/naftis/app/naftis/internal/pkg/contract"
	memoryMarket "gitlab.com/naftis/app/naftis/internal/pkg/market/memory"
	memoryStorage "gitlab.com/naftis/app/naftis/internal/pkg/storage/memory"
	"gitlab.com/naftis/app/naftis/pkg/protocol/entity"
	"testing"
	"time"
)

func TestScheduledWorkload_Schedule(t *testing.T) {
	var err error

	ctx := context.TODO()
	market := memoryMarket.NewMessage()
	market.Start(ctx)
	storage := memoryStorage.NewScheduledWorkload()
	contractSelector := contract.NewSelector()
	service := NewScheduledWorkload(storage, market, contractSelector)

	id := uuid.New()

	contractProposal := entity.ContractProposal{
		Contract: &entity.ContractSpecification{
			PricePerMinute:    1000,
			TokenSendInterval: 1,
			Duration:          10,
		},
	}

	container := entity.WorkloadSpec_Container{
		Name:  "random-container",
		Image: "nginx:latest",
		Resources: &entity.WorkloadSpec_Container_Resources{
			MemorySize:     1024,
			CpuCount:       1,
			CpuPerformance: 1000,
		},
		Storage: []*entity.WorkloadSpec_Container_Storage{},
	}

	// Create workload without any container
	err = service.Schedule(ctx, entity.ScheduledWorkload{
		Id: id.String(),
		Spec: &entity.WorkloadSpec{
			Containers:   []*entity.WorkloadSpec_Container{},
			Isolation:    entity.WorkloadSpec_CONTAINER_ISOLATION,
			NodeSelector: []*entity.NodeSelector{},
		},
		PrincipalProposal: &contractProposal,
	})
	assert.Error(t, err)

	// Create workload with two same named containers
	err = service.Schedule(ctx, entity.ScheduledWorkload{
		Id: id.String(),
		Spec: &entity.WorkloadSpec{
			Containers: []*entity.WorkloadSpec_Container{&container, &container},
			Isolation:  entity.WorkloadSpec_CONTAINER_ISOLATION,
		},
		PrincipalProposal: &contractProposal,
	})
	assert.Error(t, err)

	// Create valid workload
	err = service.Schedule(ctx, entity.ScheduledWorkload{
		Id: id.String(),
		Spec: &entity.WorkloadSpec{
			Containers:   []*entity.WorkloadSpec_Container{&container},
			Isolation:    entity.WorkloadSpec_CONTAINER_ISOLATION,
			NodeSelector: []*entity.NodeSelector{},
		},
		PrincipalProposal: &contractProposal,
	})
	assert.NoError(t, err)

	// Retrieve previously created workload
	w, err := storage.Get(id.String())
	assert.NoError(t, err)
	assert.NotNil(t, w)

	// Create again valid workload with same id
	err = service.Schedule(ctx, entity.ScheduledWorkload{
		Id: id.String(),
		Spec: &entity.WorkloadSpec{
			Containers:   []*entity.WorkloadSpec_Container{&container},
			Isolation:    entity.WorkloadSpec_CONTAINER_ISOLATION,
			NodeSelector: []*entity.NodeSelector{},
		},
	})
	assert.Error(t, err)
}

func TestScheduledWorkload_AddContractProposalFromRunner(t *testing.T) {
	var err error

	ctx := context.TODO()
	market := memoryMarket.NewMessage()
	market.Start(ctx)
	storage := memoryStorage.NewScheduledWorkload()
	contractSelector := contract.NewSelector()
	service := NewScheduledWorkload(storage, market, contractSelector)

	scheduledWorkloadId := uuid.NewString()
	scheduledWorkloadTxId := randstr.Hex(64)

	contractProposal := entity.ContractProposal{
		Contract: &entity.ContractSpecification{
			PricePerMinute:    1000,
			TokenSendInterval: 1,
			Duration:          10,
		},
	}
	contractProposalTxId := randstr.Hex(64)

	err = storage.Create(entity.ScheduledWorkload{
		Id:                            scheduledWorkloadId,
		State:                         &entity.State{Current: "new"},
		WorkloadSpecificationMarketId: scheduledWorkloadTxId,
	})
	assert.NoError(t, err)

	err = service.AddContractProposalFromRunner(ctx, scheduledWorkloadTxId, contractProposalTxId, contractProposal)
	assert.NoError(t, err)

	_, err = storage.ListRunnerContractProposals(scheduledWorkloadId)
	assert.NoError(t, err)
}

func TestScheduledWorkload_PublishWorkloadSpecificationOnMarket(t *testing.T) {
	var err error

	ctx := context.TODO()
	market := memoryMarket.NewMessage()
	market.Start(ctx)
	workloadSpecifications := market.ListenWorkloadSpecification(ctx, 1)
	storage := memoryStorage.NewScheduledWorkload()
	contractSelector := contract.NewSelector()
	service := NewScheduledWorkload(storage, market, contractSelector)

	scheduledWorkloadId := uuid.NewString()

	err = storage.Create(entity.ScheduledWorkload{
		Id: scheduledWorkloadId,
	})
	assert.NoError(t, err)

	assert.Len(t, workloadSpecifications, 0)
	err = service.PublishWorkloadSpecification(ctx, scheduledWorkloadId)
	time.Sleep(time.Millisecond) // Dirty, but works, wait for market processing
	assert.NoError(t, err)
	assert.Len(t, workloadSpecifications, 1)
}

func TestScheduledWorkload_AcceptBestRunnerProposal(t *testing.T) {
	var err error

	ctx := context.TODO()
	market := memoryMarket.NewMessage()
	market.Start(ctx)
	storage := memoryStorage.NewScheduledWorkload()
	contractSelector := contract.NewSelector()
	service := NewScheduledWorkload(storage, market, contractSelector)

	scheduledWorkloadId := uuid.NewString()

	runnerProposal := entity.ScheduledWorkload_RunnerProposal{
		Proposal: &entity.ContractProposal{
			Contract: &entity.ContractSpecification{
				PricePerMinute:    1000,
				TokenSendInterval: 1,
				Duration:          10,
			},
		},
		ContractProposalMarketId: randstr.Hex(64),
	}

	scheduledWorkload := entity.ScheduledWorkload{
		Id: scheduledWorkloadId,
		State: &entity.State{
			Current:          "waiting_for_runners_proposals",
			Previous:         "waiting_for_runners_proposals",
			BackOffTimestamp: 0,
		},
		WorkloadSpecificationMarketId: randstr.Hex(64),
		RunnersProposals:              []*entity.ScheduledWorkload_RunnerProposal{},
	}

	err = storage.Create(scheduledWorkload)
	assert.NoError(t, err)

	err = service.AcceptBestRunnerProposal(ctx, scheduledWorkloadId)
	assert.Error(t, err)

	err = storage.AddRunnerContractProposal(scheduledWorkloadId, runnerProposal)
	assert.NoError(t, err)

	err = service.AcceptBestRunnerProposal(ctx, scheduledWorkloadId)
	assert.NoError(t, err)
}
