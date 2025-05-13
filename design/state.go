package design

import "github.com/fsm/runtime"

type State interface {
	GetDstState(tCtx *runtime.TransitionCtx) (*int64, error)
}
