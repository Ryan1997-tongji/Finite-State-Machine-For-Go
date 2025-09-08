package fsm_impl

import (
	"fmt"
	"github.com/fsm/design"
	"github.com/fsm/runtime"
	"sync"
)

type DefaultFSM struct {
	mutex       sync.RWMutex
	transitions map[int64]map[string]*design.Transition // map[int64][event_name]*Transition

}

func NewDefaultFSM() *DefaultFSM {
	return &DefaultFSM{
		mutex:       sync.RWMutex{},
		transitions: make(map[int64]map[string]*design.Transition),
	}

}

func (f *DefaultFSM) Init(constructors []*design.TransitionConstructor) {
	for _, _constructor := range constructors {
		constructor := _constructor
		if constructor == nil {
			continue
		}
		f.mutex.Lock()
		for _, _sourceID := range constructor.AllowedSourceStateIDs {
			sourceID := _sourceID
			if _, ok := f.transitions[sourceID]; !ok {
				f.transitions[sourceID] = make(map[string]*design.Transition)
			}
			eventToTransition := f.transitions[sourceID]
			if _, ok := eventToTransition[constructor.EventType]; ok {
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
			eventToTransition[constructor.EventType] = &design.Transition{
				EventType:  constructor.EventType,
				DstStateID: nil,
				IsForce:    constructor.IsForce,
				DstState:   constructor.DstState,
				Callbacks:  callbacks,
				Conditions: conditions,
			}
		}

		f.mutex.Unlock()
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
	eventName := tCtx.EventName
	sourceStateID := f.DefaultInitialState()
	if tCtx.SourceStateIDPtr != nil {
		sourceStateID = *tCtx.SourceStateIDPtr
	}

	f.mutex.RLock()
	transition, ok := f.transitions[sourceStateID]
	if !ok {
		f.mutex.RUnlock()
		return fmt.Errorf("sourceStateID %v not found", sourceStateID)
	}
	eventToTransition, ok := transition[eventName]
	if !ok {
		f.mutex.RUnlock()
		return fmt.Errorf("event %v not found", eventName)
	}
	f.mutex.RUnlock()

	for _, condition := range eventToTransition.Conditions {
		var isSatisfied bool
		isSatisfied, err = condition.IsSatisfied(tCtx)
		if err != nil {
			return err
		}
		if !isSatisfied {
			return fmt.Errorf("condition not satisfied,err: %v", err)
		}
	}

	eventToTransition.DstStateID, err = eventToTransition.DstState.GetDstState(tCtx)
	if eventToTransition.DstStateID == nil || *eventToTransition.DstStateID <= 0 {
		return fmt.Errorf("getTargetState,DesStateID<=0")
	}

	if err != nil {
		return fmt.Errorf("getTargetState,err: %v", err)
	}

	for _, callback := range eventToTransition.Callbacks {
		err = callback.Execute(tCtx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (f *DefaultFSM) DefaultInitialState() int64 {
	return 0
}
