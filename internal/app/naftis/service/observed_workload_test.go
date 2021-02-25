package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/thanhpk/randstr"
	memoryMarket "gitlab.com/naftis/app/naftis/internal/pkg/market/memory"
	"gitlab.com/naftis/app/naftis/internal/pkg/price"
	memoryStorage "gitlab.com/naftis/app/naftis/internal/pkg/storage/memory"
	"gitlab.com/naftis/app/naftis/pkg/protocol/entity"
	"testing"
	"time"
)

func TestObservedWorkload_Observe(t *testing.T) {
	var err error

	ctx := context.TODO()
	market := memoryMarket.NewMessage()
	market.Start(ctx)
	storage := memoryStorage.NewObservedWorkload()
	priceCalculator := price.NewCalculator(entity.PriceList{
		CpuPerMinute:    1000,
		MemoryPerMinute: 10,
	})
	service := NewObservedWorkload(storage, market, priceCalculator)

	id := uuid.New()
	workloadSpecificationMarketId := randstr.Hex(64)

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
	err = service.Observe(ctx, entity.ObservedWorkload{
		Id:                            id.String(),
		WorkloadSpecificationMarketId: workloadSpecificationMarketId,
		Spec: &entity.WorkloadSpec{
			Containers: []*entity.WorkloadSpec_Container{},
			Isolation:  entity.WorkloadSpec_CONTAINER_ISOLATION,
		},
		PrincipalProposal: &contractProposal,
	})
	assert.Error(t, err)

	// Create workload with two same named containers
	err = service.Observe(ctx, entity.ObservedWorkload{
		Id:                            id.String(),
		WorkloadSpecificationMarketId: workloadSpecificationMarketId,
		Spec: &entity.WorkloadSpec{
			Containers: []*entity.WorkloadSpec_Container{&container, &container},
			Isolation:  entity.WorkloadSpec_CONTAINER_ISOLATION,
		},
		PrincipalProposal: &contractProposal,
	})
	assert.Error(t, err)

	// Create workload without txid
	err = service.Observe(ctx, entity.ObservedWorkload{
		Id: id.String(),
		Spec: &entity.WorkloadSpec{
			Containers: []*entity.WorkloadSpec_Container{&container},
			Isolation:  entity.WorkloadSpec_CONTAINER_ISOLATION,
		},
	})
	assert.Error(t, err)

	// Create valid workload
	err = service.Observe(ctx, entity.ObservedWorkload{
		Id:                            id.String(),
		WorkloadSpecificationMarketId: workloadSpecificationMarketId,
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
	err = service.Observe(ctx, entity.ObservedWorkload{
		Id:                            id.String(),
		WorkloadSpecificationMarketId: workloadSpecificationMarketId,
		Spec: &entity.WorkloadSpec{
			Containers: []*entity.WorkloadSpec_Container{&container},
			Isolation:  entity.WorkloadSpec_CONTAINER_ISOLATION,
		},
		PrincipalProposal: &contractProposal,
	})
	assert.Error(t, err)
}

func TestObservedWorkload_PublishContractProposal(t *testing.T) {
	var err error

	ctx := context.TODO()
	market := memoryMarket.NewMessage()
	market.Start(ctx)
	contactProposals := market.ListenContractProposal(ctx, 1)
	storage := memoryStorage.NewObservedWorkload()
	priceCalculator := price.NewCalculator(entity.PriceList{
		CpuPerMinute:    1000,
		MemoryPerMinute: 10,
	})
	service := NewObservedWorkload(storage, market, priceCalculator)

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

	observedWorkloadId := uuid.NewString()
	observedWorkloadTxId := randstr.Hex(64)

	observedWorkload := entity.ObservedWorkload{
		Id:                            observedWorkloadId,
		WorkloadSpecificationMarketId: observedWorkloadTxId,
		Spec: &entity.WorkloadSpec{
			Containers:   []*entity.WorkloadSpec_Container{&container},
			Isolation:    entity.WorkloadSpec_CONTAINER_ISOLATION,
			NodeSelector: []*entity.NodeSelector{},
		},
		PrincipalProposal: &contractProposal,
	}

	err = storage.Create(observedWorkload)
	assert.NoError(t, err)

	assert.Len(t, contactProposals, 0)
	err = service.PublishContractProposal(ctx, observedWorkloadId)
	time.Sleep(time.Millisecond) // Dirty, but works, wait for market processing
	assert.NoError(t, err)
	assert.Len(t, contactProposals, 1)
}
