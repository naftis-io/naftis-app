package service

import (
	"context"
	"github.com/rs/zerolog/log"
	"github.com/thanhpk/randstr"
	"gitlab.com/naftis/app/naftis/internal/pkg/contract"
	"gitlab.com/naftis/app/naftis/internal/pkg/market"
	"gitlab.com/naftis/app/naftis/internal/pkg/storage"
	"gitlab.com/naftis/app/naftis/internal/pkg/validator"
	"gitlab.com/naftis/app/naftis/pkg/protocol/entity"
	marketProtocol "gitlab.com/naftis/app/naftis/pkg/protocol/market"
	"time"
)

type ScheduledWorkload struct {
	storage          storage.ScheduledWorkload
	market           market.MessageToken
	contractSelector *contract.Selector
}

func NewScheduledWorkload(storage storage.ScheduledWorkload, market market.MessageToken, contractSelector *contract.Selector) *ScheduledWorkload {
	return &ScheduledWorkload{
		storage:          storage,
		market:           market,
		contractSelector: contractSelector,
	}
}

func (s *ScheduledWorkload) Schedule(_ context.Context, scheduledWorkload entity.ScheduledWorkload) error {
	err := validator.New().Struct(scheduledWorkload)
	if err != nil {
		return err
	}

	scheduledWorkload.State = &entity.State{
		Current:          "new",
		Previous:         "new",
		BackOffTimestamp: 0,
	}

	s.storage.Create(scheduledWorkload)
	if err != nil {
		return err
	}

	log.Info().
		Str("scheduledWorkloadId", scheduledWorkload.Id).
		Msg("Scheduled new workload.")

	return nil
}

// PublishWorkloadSpecification function is used to publish WorkloadSpecification on market.
func (s *ScheduledWorkload) PublishWorkloadSpecification(ctx context.Context, scheduledWorkloadId string) error {
	scheduledWorkload, err := s.storage.Get(scheduledWorkloadId)
	if err != nil {
		return err
	}

	workloadSpecificationMarketId, err := s.market.EmitWorkloadSpecification(ctx, marketProtocol.WorkloadSpecification{
		Spec:              scheduledWorkload.Spec,
		PrincipalProposal: scheduledWorkload.PrincipalProposal,
	})

	err = s.storage.UpdateWorkloadSpecificationMarketId(scheduledWorkloadId, workloadSpecificationMarketId)
	if err != nil {
		return err
	}

	log.Info().
		Str("scheduledWorkloadId", scheduledWorkload.Id).
		Str("workloadSpecificationMarketId", workloadSpecificationMarketId).
		Msg("Scheduled workload published on market.")

	return nil
}

// AddContractProposalFromRunner function adds new ContractProposal from runner, received from market.
func (s *ScheduledWorkload) AddContractProposalFromRunner(_ context.Context, workloadSpecificationMarketId string, contractProposalMarketId string, contractProposal entity.ContractProposal) error {
	err := validator.New().Struct(contractProposal)
	if err != nil {
		return err
	}

	scheduledWorkload, err := s.storage.GetByWorkloadSpecificationMarketId(workloadSpecificationMarketId)
	if err != nil {
		return err
	}

	err = s.storage.AddRunnerContractProposal(scheduledWorkload.Id, entity.ScheduledWorkload_RunnerProposal{
		Proposal:                 &contractProposal,
		ContractProposalMarketId: contractProposalMarketId,
	})
	if err != nil {
		return err
	}

	log.Info().
		Str("scheduledWorkloadId", scheduledWorkload.Id).
		Str("workloadSpecificationMarketId", workloadSpecificationMarketId).
		Str("contractProposalMarketId", contractProposalMarketId).
		Msg("Added new runner ContractProposal to ScheduledWorkload.")

	return nil
}

// WaitForRunnersProposals function wait some time for runners offers.
func (s *ScheduledWorkload) WaitForRunnersProposals(_ context.Context, scheduledWorkloadId string) error {
	err := s.storage.SetBackOff(scheduledWorkloadId, time.Second*10)
	if err != nil {
		return err
	}

	scheduledWorkload, err := s.storage.Get(scheduledWorkloadId)
	if err != nil {
		return err
	}

	log.Info().
		Str("scheduledWorkloadId", scheduledWorkloadId).
		Time("backOffUntil", time.Unix(scheduledWorkload.State.BackOffTimestamp, 0)).
		Msg("Waiting for runners ContractProposal.")

	return nil
}

// AcceptBestRunnerProposal function try to find best contract proposal and accept it, and send confirmation to market.
func (s *ScheduledWorkload) AcceptBestRunnerProposal(ctx context.Context, scheduledWorkloadId string) error {
	scheduledWorkload, err := s.storage.Get(scheduledWorkloadId)
	if err != nil {
		return err
	}

	bestProposal, err := s.contractSelector.Find(ctx, scheduledWorkload.RunnersProposals)
	if err != nil {
		return err
	}

	contractAcceptMarketId, err := s.market.EmitContractAccept(ctx, marketProtocol.ContractAccept{
		WorkloadSpecificationMarketId: scheduledWorkload.WorkloadSpecificationMarketId,
		ContractProposalMarketId:      bestProposal.ContractProposalMarketId,
		Accept: &entity.ContractAccept{
			PublicKey: randstr.Hex(128),
		},
	})
	if err != nil {
		return err
	}

	err = s.storage.SetAcceptedRunnerContractProposal(scheduledWorkloadId, *bestProposal)
	if err != nil {
		return err
	}

	log.Info().
		Str("scheduledWorkloadId", scheduledWorkloadId).
		Str("contractProposalMarketId", bestProposal.ContractProposalMarketId).
		Str("contractAcceptMarketId", contractAcceptMarketId).
		Msg("Contract accepted.")

	return nil
}
