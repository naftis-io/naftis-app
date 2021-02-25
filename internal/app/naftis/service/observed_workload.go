package service

import (
	"context"
	"errors"
	"github.com/rs/zerolog/log"
	"gitlab.com/naftis/app/naftis/internal/pkg/market"
	"gitlab.com/naftis/app/naftis/internal/pkg/price"
	"gitlab.com/naftis/app/naftis/internal/pkg/storage"
	"gitlab.com/naftis/app/naftis/internal/pkg/validator"
	"gitlab.com/naftis/app/naftis/pkg/protocol/entity"
	marketProtocol "gitlab.com/naftis/app/naftis/pkg/protocol/market"
)

type ObservedWorkload struct {
	storage         storage.ObservedWorkload
	market          market.MessageToken
	priceCalculator *price.Calculator
}

func NewObservedWorkload(storage storage.ObservedWorkload, market market.MessageToken, priceCalculator *price.Calculator) *ObservedWorkload {
	return &ObservedWorkload{
		storage:         storage,
		market:          market,
		priceCalculator: priceCalculator,
	}
}

// Observe add new workload for watching.
func (s *ObservedWorkload) Observe(_ context.Context, observedWorkload entity.ObservedWorkload) error {
	err := validator.New().Struct(observedWorkload)
	if err != nil {
		return err
	}

	observedWorkload.State = &entity.State{
		Current:          "new",
		Previous:         "new",
		BackOffTimestamp: 0,
	}

	err = s.storage.Create(observedWorkload)
	if err != nil {
		return err
	}

	log.Info().
		Str("observedWorkloadId", observedWorkload.Id).
		Str("workloadSpecificationMarketId", observedWorkload.WorkloadSpecificationMarketId).
		Msg("Observing new workload.")

	return nil
}

// PublishContractProposal function publishes new contract proposal on market.
func (s *ObservedWorkload) PublishContractProposal(ctx context.Context, observedWorkloadId string) error {
	observedWorkload, err := s.storage.Get(observedWorkloadId)
	if err != nil {
		return err
	}

	pricePerMinute, err := s.priceCalculator.GetWorkloadPrice(*observedWorkload.Spec)
	if err != nil {
		return err
	}

	contractSpecification := entity.ContractSpecification{
		PricePerMinute:    pricePerMinute,
		TokenSendInterval: observedWorkload.PrincipalProposal.Contract.TokenSendInterval,
		Duration:          observedWorkload.PrincipalProposal.Contract.Duration,
	}

	contractProposalMarketId, err := s.market.EmitContractProposal(ctx, marketProtocol.ContractProposal{
		Proposal: &entity.ContractProposal{
			Contract: &contractSpecification,
		},
		WorkloadSpecificationMarketId: observedWorkload.WorkloadSpecificationMarketId,
	})
	if err != nil {
		return err
	}

	log.Info().
		Str("observedWorkloadId", observedWorkloadId).
		Str("workloadSpecificationMarketId", observedWorkload.WorkloadSpecificationMarketId).
		Str("contractProposalMarketId", contractProposalMarketId).
		Msg("Published ContractProposal on market.")

	return nil
}

// ConfirmPrincipalAcceptance function confirms that ObservedWorkload is accepted
func (s *ObservedWorkload) ConfirmPrincipalAcceptance(ctx context.Context, workloadSpecificationMarketId string, contractAcceptMarketId string, contractAccept entity.ContractAccept) error {
	log := log.With().
		Str("workloadSpecificationMarketId", workloadSpecificationMarketId).
		Str("contractAcceptMarketId", contractAcceptMarketId).
		Logger()

	observedWorkload, err := s.storage.GetByWorkloadSpecificationMarketId(workloadSpecificationMarketId)
	if err != nil {
		return err
	}

	log = log.With().Str("observedWorkloadId", observedWorkload.Id).Logger()

	if observedWorkload.State.Current != "waiting_for_acceptance" {
		return ErrInvalidObservedWorkloadState
	}

	if observedWorkload.PrincipalAcceptance != nil {
		return ErrPrincipalAcceptanceExists
	}

	err = s.storage.SetPrincipalAcceptance(observedWorkload.Id, entity.ObservedWorkload_PrincipalAcceptance{
		ContractAcceptMarketId: contractAcceptMarketId,
		Accept:                 &contractAccept,
	})
	if err != nil {
		return err
	}

	log.Info().Msg("Principal acceptance confirmed.")

	return nil
}

var ErrInvalidObservedWorkloadState = errors.New("invalid observed workload state")
var ErrPrincipalAcceptanceExists = errors.New("principal acceptance exists")
