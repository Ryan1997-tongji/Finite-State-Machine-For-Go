package design

import "github.com/fsm/runtime"

type FSM interface {
	// NewFsm
	Init(constructors []*TransitionConstructor)

	// Transition Exec state transitions based on event context
	Transition(tCtx *runtime.TransitionCtx) (err error)
}
