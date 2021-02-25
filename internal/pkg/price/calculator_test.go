package price

import (
	"github.com/stretchr/testify/assert"
	"gitlab.com/naftis/app/naftis/pkg/protocol/entity"
	"testing"
)

func TestCalculator_GetWorkloadPrice(t *testing.T) {
	var err error
	var ppm uint32

	prices := entity.PriceList{
		CpuPerMinute:    1000,
		MemoryPerMinute: 5,
	}

	zeroContainerWorkload := entity.WorkloadSpec{
		Containers: []*entity.WorkloadSpec_Container{},
	}

	singleContainerWorkload := entity.WorkloadSpec{
		Containers: []*entity.WorkloadSpec_Container{
			{
				Resources: &entity.WorkloadSpec_Container_Resources{
					MemorySize:     1024,
					CpuCount:       1,
					CpuPerformance: 100000,
				},
			},
		},
	}

	multiContainerWorkload := entity.WorkloadSpec{
		Containers: []*entity.WorkloadSpec_Container{
			{
				Resources: &entity.WorkloadSpec_Container_Resources{
					MemorySize:     1024,
					CpuCount:       1,
					CpuPerformance: 100000,
				},
			},
			{
				Resources: &entity.WorkloadSpec_Container_Resources{
					MemorySize:     512,
					CpuCount:       3,
					CpuPerformance: 100000,
				},
			},
		},
	}

	calculator := NewCalculator(prices)

	ppm, err = calculator.GetWorkloadPrice(zeroContainerWorkload)
	assert.Error(t, err)
	assert.Zero(t, ppm)

	ppm, err = calculator.GetWorkloadPrice(singleContainerWorkload)
	assert.NoError(t, err)
	assert.EqualValues(t, 6120, ppm) // (1024 RAM * 5) + (1 CPU * 1000)

	ppm, err = calculator.GetWorkloadPrice(multiContainerWorkload)
	assert.NoError(t, err)
	assert.EqualValues(t, 11680, ppm) // ( (1024 RAM * 5) + (1 CPU * 1000) ) + ( (512 RAM * 5) + (3 CPU * 1000) )

}
