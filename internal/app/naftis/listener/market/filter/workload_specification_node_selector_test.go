package filter

import (
	"github.com/stretchr/testify/assert"
	"gitlab.com/naftis/app/naftis/internal/app/naftis/label"
	"gitlab.com/naftis/app/naftis/internal/app/naftis/label/source"
	"gitlab.com/naftis/app/naftis/internal/pkg/market"
	"gitlab.com/naftis/app/naftis/pkg/protocol/blockchain"
	"gitlab.com/naftis/app/naftis/pkg/protocol/entity"
	"testing"
)

func TestMarketWorkloadSpecificationNodeSelector_Filter(t *testing.T) {
	labels := label.NewContainer()

	invalidWithoutConstraintMsg := market.WorkloadSpecification{
		Msg: blockchain.WorkloadSpecification{
			Spec: &entity.WorkloadSpec{
				NodeSelector: []*entity.NodeSelector{
					&entity.NodeSelector{
						Label: &entity.NodeLabel{
							Key:   "foo.bar",
							Value: "invalid",
						},
						IsConstraint: false,
					},
				},
			},
		},
	}

	validMsg := market.WorkloadSpecification{
		Msg: blockchain.WorkloadSpecification{
			Spec: &entity.WorkloadSpec{
				NodeSelector: []*entity.NodeSelector{
					&entity.NodeSelector{
						Label: &entity.NodeLabel{
							Key:   "foo.bar",
							Value: "test",
						},
						IsConstraint: true,
					},
				},
			},
		},
	}

	invalidMsg := market.WorkloadSpecification{
		Msg: blockchain.WorkloadSpecification{
			Spec: &entity.WorkloadSpec{
				NodeSelector: []*entity.NodeSelector{
					&entity.NodeSelector{
						Label: &entity.NodeLabel{
							Key:   "foo.bar",
							Value: "invalid",
						},
						IsConstraint: true,
					},
				},
			},
		},
	}

	labels.AttachSource(source.NewStatic(label.SourceLabels{
		"foo.bar": "test",
	}))

	labels.Refresh()

	filter := NewMarketWorkloadSpecificationNodeSelector(labels)

	ok := filter.Filter(validMsg)
	assert.True(t, ok)

	ok = filter.Filter(invalidMsg)
	assert.False(t, ok)

	ok = filter.Filter(invalidWithoutConstraintMsg)
	assert.True(t, ok)
}
