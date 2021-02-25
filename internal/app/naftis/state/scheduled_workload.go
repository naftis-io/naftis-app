package state

import (
	"context"
	"github.com/looplab/fsm"
	"gitlab.com/naftis/app/naftis/internal/app/naftis/service"
	"gitlab.com/naftis/app/naftis/internal/pkg/state"
	"gitlab.com/naftis/app/naftis/internal/pkg/storage"
)

type ScheduledWorkload struct {
	state.Manager
	storage storage.ScheduledWorkload
	service *service.ScheduledWorkload
}

func NewScheduledWorkload(storage storage.ScheduledWorkload, service *service.ScheduledWorkload) *ScheduledWorkload {
	var scheduledWorkload *ScheduledWorkload
	scheduledWorkload = &ScheduledWorkload{
		Manager: *state.NewManager(state.Specification{
			Name: "ScheduledWorkload",
			Events: fsm.Events{
				{Name: "publish_on_market", Src: []string{"new"}, Dst: "published_on_market"},
				{Name: "wait_for_runners_proposals", Src: []string{"published_on_market"}, Dst: "waiting_for_runners_proposals"},
				{Name: "accept_runner_proposal", Src: []string{"waiting_for_runners_proposals"}, Dst: "runner_proposal_accepted"},
				{Name: "control", Src: []string{"runner_proposal_accepted"}, Dst: "controlling"},
			},
			EventCallbacks: state.EventCallbacks{
				"publish_on_market": func(ctx context.Context, id string) error {
					return scheduledWorkload.onPublishOnMarket(ctx, id)
				},
				"wait_for_runners_proposals": func(ctx context.Context, id string) error {
					return scheduledWorkload.onWaitForRunnersProposals(ctx, id)
				},
				"accept_runner_proposal": func(ctx context.Context, id string) error {
					return scheduledWorkload.onAcceptRunnerProposal(ctx, id)
				},
				"control": func(ctx context.Context, id string) error {
					return scheduledWorkload.onControl(ctx, id)
				},
			},
			TransitionCallbacks: state.TransitionCallbacks{
				"new": func(ctx context.Context, id string) (string, error) {
					return scheduledWorkload.transitFromNew(ctx, id)
				},
				"published_on_market": func(ctx context.Context, id string) (string, error) {
					return scheduledWorkload.transitFromPublishedOnMarket(ctx, id)
				},
				"waiting_for_runners_proposals": func(ctx context.Context, id string) (string, error) {
					return scheduledWorkload.transitFromWaitingForRunnersProposals(ctx, id)
				},
				"runner_proposal_accepted": func(ctx context.Context, id string) (string, error) {
					return scheduledWorkload.transitFromRunnerProposalAccepted(ctx, id)
				},
				"controlling": func(ctx context.Context, id string) (string, error) {
					return scheduledWorkload.transitFromControlling(ctx, id)
				},
			},
		}, storage),
		storage: storage,
		service: service,
	}

	return scheduledWorkload
}

func (s *ScheduledWorkload) transitFromNew(ctx context.Context, id string) (string, error) {
	return "publish_on_market", nil
}

func (s *ScheduledWorkload) transitFromPublishedOnMarket(ctx context.Context, id string) (string, error) {
	return "wait_for_runners_proposals", nil
}

func (s *ScheduledWorkload) transitFromWaitingForRunnersProposals(ctx context.Context, id string) (string, error) {
	runnersProposals, err := s.storage.ListRunnerContractProposals(id)
	if err != nil {
		return "nil", err
	}

	if len(runnersProposals) == 0 {
		return state.NoTransition, nil
	}

	return "accept_runner_proposal", nil
}

func (s *ScheduledWorkload) transitFromRunnerProposalAccepted(ctx context.Context, id string) (string, error) {
	return "control", nil
}

func (s *ScheduledWorkload) transitFromControlling(ctx context.Context, id string) (string, error) {
	return state.NoTransition, nil
}

func (s *ScheduledWorkload) onPublishOnMarket(ctx context.Context, id string) error {
	err := s.service.PublishWorkloadSpecification(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *ScheduledWorkload) onWaitForRunnersProposals(ctx context.Context, id string) error {
	err := s.service.WaitForRunnersProposals(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *ScheduledWorkload) onAcceptRunnerProposal(ctx context.Context, id string) error {
	err := s.service.AcceptBestRunnerProposal(ctx, id)
	if err != nil {

		return err
	}

	return nil
}

func (s *ScheduledWorkload) onControl(ctx context.Context, id string) error {
	return nil
}
