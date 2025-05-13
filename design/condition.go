package design

import "github.com/fsm/runtime"

type Condition interface {
	IsSatisfied(tCtx *runtime.TransitionCtx) (bool,error)
}
