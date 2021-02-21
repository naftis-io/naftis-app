package filter

import (
	"gitlab.com/naftis/app/naftis/internal/app/naftis/label"
	"gitlab.com/naftis/app/naftis/internal/pkg/market"
)

// MarketWorkloadSpecificationNodeSelector is filter used to filter out workloads that not meet Runner labels.
type MarketWorkloadSpecificationNodeSelector struct {
	labels *label.Container
}

func NewMarketWorkloadSpecificationNodeSelector(labels *label.Container) *MarketWorkloadSpecificationNodeSelector {
	return &MarketWorkloadSpecificationNodeSelector{
		labels: labels,
	}
}

func (f *MarketWorkloadSpecificationNodeSelector) Filter(msg market.WorkloadSpecification) bool {
	labels := f.labels.Map()

	for _, selector := range msg.Msg.Spec.NodeSelector {
		if !selector.IsConstraint {
			continue
		}

		if label, exists := labels[selector.Label.Key]; !exists {
			return false
		} else {
			if label.Value != selector.Label.Value {
				return false
			}
		}
	}

	return true
}
