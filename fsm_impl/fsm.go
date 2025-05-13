package fsm_impl

import (
	"fmt"
	"github.com/fsm/design"
	"github.com/fsm/runtime"
	"sync"
)

type DefaultFSM struct {
	mutex       sync.RWMutex
	transitions map[string]*design.Transition // map[event_name]*Transition

}

func NewDefaultFSM() *DefaultFSM {
	return &DefaultFSM{
		mutex:       sync.RWMutex{},
		transitions: make(map[string]*design.Transition),
	}

}

func (f *DefaultFSM) Init(constructors []*design.TransitionConstructor) {
	for _, _constructor := range constructors {
		constructor := _constructor
		if constructor == nil {
			continue
		}
		f.mutex.RLock()
		if _, ok := f.transitions[constructor.EventType]; ok {
			continue
		}

		var conditions []design.Condition
		conditions = append(conditions, constructor.GlobalConditions...)
		if len(constructor.AllowedSourceStateIDs) > 0 {
			conditions = append(conditions, NewSourceStateValidator(constructor.AllowedSourceStateIDs))
		}
		conditions = append(conditions, constructor.TransitionConditions...)
		var callbacks []design.Action
		if len(constructor.AllowedDstStateIDs) > 0 {
			callbacks = append(callbacks, NewTargetStateValidator(constructor.IsForce, constructor.AllowedDstStateIDs))
		}
		callbacks = append(callbacks, constructor.TransitionCallbacks...)
		callbacks = append(callbacks, constructor.GlobalCallbacks...)

		f.transitions[constructor.EventType] = &design.Transition{
			EventType:  constructor.EventType,
			DstStateID: nil,
			IsForce:    constructor.IsForce,
			DstState:   constructor.DstState,
			Callbacks:  callbacks,
			Conditions: conditions,
		}
		f.mutex.RUnlock()
	}
}

// Transition State transition
func (f *DefaultFSM) Transition(tCtx *runtime.TransitionCtx) (err error) {
	if tCtx == nil {
		return fmt.Errorf("tCtx is empty")
	}
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("transition panic: %v", r)
		}
	}()
	f.mutex.RLock()
	eventName := tCtx.EventName
	transition, ok := f.transitions[eventName]
	if !ok {
		f.mutex.RUnlock()
		return fmt.Errorf("event %s not found", eventName)
	}
	f.mutex.RUnlock()

	for _, condition := range transition.Conditions {
		var isSatisfied bool
		isSatisfied, err = condition.IsSatisfied(tCtx)
		if err != nil {
			return err
		}
		if !isSatisfied {
			return fmt.Errorf("condition not satisfied,err: %v", err)
		}
	}

	transition.DstStateID, err = transition.DstState.GetDstState(tCtx)
	if transition.DstStateID == nil || *transition.DstStateID <= 0 {
		return fmt.Errorf("getTargetState,DesStateID<=0")
	}

	if err != nil {
		return fmt.Errorf("getTargetState,err: %v", err)
	}

	for _, callback := range transition.Callbacks {
		err = callback.Execute(tCtx)
		if err != nil {
			return err
		}
	}
	return nil
}
