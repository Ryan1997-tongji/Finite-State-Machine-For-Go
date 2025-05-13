package state_service

import (
	"github.com/fsm/design"
	"github.com/fsm/runtime"
)

func Transition(fsm design.FSM, tCtx *runtime.TransitionCtx) error {
	return fsm.Transition(tCtx)
}
