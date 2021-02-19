package command

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/thanhpk/randstr"
	"gitlab.com/naftis/app/naftis/internal/pkg/storage/memory"
	"gitlab.com/naftis/app/naftis/pkg/protocol/entity"
	"testing"
)

func TestObserveWorkload_Invoke(t *testing.T) {
	var err error

	storage := memory.NewObservedWorkload()
	cmd := NewObserveWorkload(storage)

	id := uuid.New()
	txId := randstr.Hex(64)

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
	err = cmd.Invoke(entity.ObservedWorkload{
		Id:   id.String(),
		TxId: txId,
		Spec: &entity.WorkloadSpec{
			Containers: []*entity.WorkloadSpec_Container{},
			Isolation:  entity.WorkloadSpec_CONTAINER_ISOLATION,
		},
	})
	assert.Error(t, err)

	// Create workload with two same named containers
	err = cmd.Invoke(entity.ObservedWorkload{
		Id:   id.String(),
		TxId: txId,
		Spec: &entity.WorkloadSpec{
			Containers: []*entity.WorkloadSpec_Container{&container, &container},
			Isolation:  entity.WorkloadSpec_CONTAINER_ISOLATION,
		},
	})
	assert.Error(t, err)

	// Create workload without txid
	err = cmd.Invoke(entity.ObservedWorkload{
		Id: id.String(),
		Spec: &entity.WorkloadSpec{
			Containers: []*entity.WorkloadSpec_Container{&container},
			Isolation:  entity.WorkloadSpec_CONTAINER_ISOLATION,
		},
	})
	assert.Error(t, err)

	// Create valid workload
	err = cmd.Invoke(entity.ObservedWorkload{
		Id:   id.String(),
		TxId: txId,
		Spec: &entity.WorkloadSpec{
			Containers: []*entity.WorkloadSpec_Container{&container},
			Isolation:  entity.WorkloadSpec_CONTAINER_ISOLATION,
		},
	})
	assert.NoError(t, err)

	// Retrieve previously created workload
	w, err := storage.Get(id.String())
	assert.NoError(t, err)
	assert.NotNil(t, w)

	// Create again valid workload with same id
	err = cmd.Invoke(entity.ObservedWorkload{
		Id:   id.String(),
		TxId: txId,
		Spec: &entity.WorkloadSpec{
			Containers: []*entity.WorkloadSpec_Container{&container},
			Isolation:  entity.WorkloadSpec_CONTAINER_ISOLATION,
		},
	})
	assert.Error(t, err)

}
