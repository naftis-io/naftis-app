package state

import (
	"context"
	"github.com/looplab/fsm"
)

const NoTransition string = ""

type TransitionCallback func(ctx context.Context, id string) (string, error)
type TransitionCallbacks map[string]TransitionCallback
type EventCallback func(ctx context.Context, id string) error
type EventCallbacks map[string]EventCallback

type Specification struct {
	Name                string
	Events              fsm.Events
	EventCallbacks      EventCallbacks
	TransitionCallbacks TransitionCallbacks
}
