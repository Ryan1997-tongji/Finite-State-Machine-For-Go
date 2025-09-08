package runtime

import (
	"context"
)

// TransitionCtx Runtime Ctx for State Transition
type TransitionCtx struct {
	Ctx              context.Context
	EventName        string
	SourceStateIDPtr *int64
	TargetStateIDPtr *int64 // Externally specified target state
	Fact             interface{}
	InputExtra       map[string]string
	OutputExtra      map[string]string
}
