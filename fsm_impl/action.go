package fsm_impl

import (
	"fmt"
	"github.com/fsm/runtime"
	"github.com/fsm/utils"
)

type TargetStateValidator struct {
	isForce            bool
	allowedDstStateIDs []int64
}

func NewTargetStateValidator(isForce bool, allowedDstStateIDs []int64) *TargetStateValidator {
	return &TargetStateValidator{
		isForce:            isForce,
		allowedDstStateIDs: allowedDstStateIDs,
	}
}

func (t *TargetStateValidator) Execute(tCtx *runtime.TransitionCtx) error {
	if t.isForce {
		return nil
	}
	if tCtx.TargetStateID == nil {
		return nil
	}
	if utils.SliceContains(t.allowedDstStateIDs, *tCtx.TargetStateID) {
		return nil
	}
	return fmt.Errorf("target state: %v,not allowed", *tCtx.TargetStateID)
}
