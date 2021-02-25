package contract

import (
	"context"
	"errors"
	"gitlab.com/naftis/app/naftis/pkg/protocol/entity"
)

type Selector struct {
}

func NewSelector() *Selector {
	return &Selector{}
}

func (s *Selector) Find(ctx context.Context, proposals []*entity.ScheduledWorkload_RunnerProposal) (*entity.ScheduledWorkload_RunnerProposal, error) {
	if len(proposals) == 0 {
		return nil, ErrNoRunnersProposals
	}

	return proposals[0], nil
}

var ErrNoRunnersProposals = errors.New("no proposals from runners")
