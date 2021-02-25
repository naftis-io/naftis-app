package service

import (
	"gitlab.com/naftis/app/naftis/internal/pkg/contract"
	"gitlab.com/naftis/app/naftis/internal/pkg/market"
	"gitlab.com/naftis/app/naftis/internal/pkg/price"
	"gitlab.com/naftis/app/naftis/internal/pkg/storage"
)

type Container struct {
	priceCalculator   *price.Calculator
	scheduledWorkload *ScheduledWorkload
	observedWorkload  *ObservedWorkload
	contractSelector  *contract.Selector
}

func NewContainer(storage storage.Container, market market.MessageToken, priceCalculator *price.Calculator, contractSelector *contract.Selector) *Container {
	return &Container{
		priceCalculator:   priceCalculator,
		contractSelector:  contractSelector,
		scheduledWorkload: NewScheduledWorkload(storage.ScheduledWorkload(), market, contractSelector),
		observedWorkload:  NewObservedWorkload(storage.ObservedWorkload(), market, priceCalculator),
	}
}

func (c *Container) ObservedWorkload() *ObservedWorkload {
	return c.observedWorkload
}

func (c *Container) ScheduledWorkload() *ScheduledWorkload {
	return c.scheduledWorkload
}

func (c *Container) PriceCalculator() *price.Calculator {
	return c.priceCalculator
}
