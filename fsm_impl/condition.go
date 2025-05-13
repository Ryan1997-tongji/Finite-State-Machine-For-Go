package fsm_impl

import (
	"fmt"
	"github.com/fsm/runtime"
	"github.com/fsm/utils"
)

type SourceStateValidator struct {
	isForce               bool
	allowedSourceStateIDs []int64
}

func NewSourceStateValidator(allowedSourceStateIDs []int64) *SourceStateValidator {
	return &SourceStateValidator{
		allowedSourceStateIDs: allowedSourceStateIDs,
	}
}

func (t *SourceStateValidator) IsSatisfied(tCtx *runtime.TransitionCtx) (bool, error) {
	if tCtx.SourceStateID == nil {
		return false, fmt.Errorf("source state is empty")
	}
	if utils.SliceContains(t.allowedSourceStateIDs, *tCtx.SourceStateID) {
		return true, nil
	}
	return false, fmt.Errorf("source state: %v,not allowed", *tCtx.SourceStateID)
}
