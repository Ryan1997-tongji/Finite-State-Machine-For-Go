package design

import "github.com/fsm/runtime"

type Action interface {
	Execute(tCtx *runtime.TransitionCtx) error
}


