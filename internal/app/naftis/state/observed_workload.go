package state

import (
	"context"
	"github.com/looplab/fsm"
	"gitlab.com/naftis/app/naftis/internal/app/naftis/service"
	"gitlab.com/naftis/app/naftis/internal/pkg/state"
	"gitlab.com/naftis/app/naftis/internal/pkg/storage"
)

type ObservedWorkload struct {
	state.Manager
	storage storage.ObservedWorkload
	service *service.ObservedWorkload
}

func NewObservedWorkload(storage storage.ObservedWorkload, service *service.ObservedWorkload) *ObservedWorkload {
	var observedWorkload *ObservedWorkload
	observedWorkload = &ObservedWorkload{
		Manager: *state.NewManager(state.Specification{
			Name: "ObservedWorkload",
			Events: fsm.Events{
				{Name: "send_proposal", Src: []string{"new"}, Dst: "proposal_sent"},
				{Name: "wait_for_acceptance", Src: []string{"proposal_sent"}, Dst: "waiting_for_acceptance"},
				{Name: "confirm_principal_acceptance", Src: []string{"waiting_for_acceptance"}, Dst: "principal_acceptance_confirmed"},
				{Name: "prepare_to_run", Src: []string{"principal_acceptance_confirmed"}, Dst: "preparing_to_run"},
			},
			EventCallbacks: state.EventCallbacks{
				"send_proposal": func(ctx context.Context, id string) error {
					return observedWorkload.onSendProposal(ctx, id)
				},
				"prepare_to_run": func(ctx context.Context, id string) error {
					return observedWorkload.onPrepareToRun(ctx, id)
				},
			},
			TransitionCallbacks: state.TransitionCallbacks{
				"new": func(ctx context.Context, id string) (string, error) {
					return observedWorkload.transitFromNew(ctx, id)
				},
				"proposal_sent": func(ctx context.Context, id string) (string, error) {
					return observedWorkload.transitFromProposalSent(ctx, id)
				},
				"waiting_for_acceptance": func(ctx context.Context, id string) (string, error) {
					return observedWorkload.transitFromWaitingForAcceptance(ctx, id)
				},
				"principal_acceptance_confirmed": func(ctx context.Context, id string) (string, error) {
					return observedWorkload.transitFromPrincipalAcceptanceConfirmed(ctx, id)
				},
				"preparing_to_run": func(ctx context.Context, id string) (string, error) {
					return observedWorkload.transitFromPreparingToRun(ctx, id)
				},
			},
		}, storage),
		storage: storage,
		service: service,
	}

	return observedWorkload
}

func (s *ObservedWorkload) transitFromNew(ctx context.Context, id string) (string, error) {
	return "send_proposal", nil
}

func (s *ObservedWorkload) transitFromProposalSent(ctx context.Context, id string) (string, error) {
	return "wait_for_acceptance", nil
}

func (s *ObservedWorkload) transitFromWaitingForAcceptance(ctx context.Context, id string) (string, error) {
	observedWorkload, err := s.storage.Get(id)
	if err != nil {
		return state.NoTransition, err
	}

	if observedWorkload.PrincipalAcceptance == nil {
		return state.NoTransition, nil
	}

	return "confirm_principal_acceptance", nil
}

func (s *ObservedWorkload) transitFromPrincipalAcceptanceConfirmed(ctx context.Context, id string) (string, error) {
	return "prepare_to_run", nil
}

func (s *ObservedWorkload) transitFromPreparingToRun(ctx context.Context, id string) (string, error) {
	return state.NoTransition, nil
}

func (s *ObservedWorkload) onSendProposal(ctx context.Context, id string) error {
	err := s.service.PublishContractProposal(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *ObservedWorkload) onPrepareToRun(ctx context.Context, id string) error {
	return nil
}
