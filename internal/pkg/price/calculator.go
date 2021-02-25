package price

import (
	"errors"
	"gitlab.com/naftis/app/naftis/pkg/protocol/entity"
)

// Calculator is struct used to calculate prices
type Calculator struct {
	priceList entity.PriceList
}

func NewCalculator(priceList entity.PriceList) *Calculator {
	return &Calculator{
		priceList: priceList,
	}
}

// GetWorkloadPrice returns price-per-minute of given workload
func (c *Calculator) GetWorkloadPrice(specification entity.WorkloadSpec) (uint32, error) {
	if len(specification.Containers) == 0 {
		return 0, ErrNoContainers
	}

	var totalPrice uint32
	totalPrice = 0

	for _, container := range specification.Containers {
		totalPrice += c.priceList.CpuPerMinute * container.Resources.CpuCount
		totalPrice += c.priceList.MemoryPerMinute * container.Resources.MemorySize
	}

	return totalPrice, nil
}

var (
	ErrNoContainers = errors.New("missing containers from workload")
)
