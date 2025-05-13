package main

import (
	"fmt"
	"github.com/fsm/fsm_impl"
	"github.com/fsm/runtime"
	"github.com/fsm/state_service"
)

func main() {
	fsm := fsm_impl.NewDefaultFSM()
	fsm.Init(nil)
	tCtx := &runtime.TransitionCtx{}
	err := state_service.Transition(fsm, tCtx)
	if err != nil {
		fmt.Println(err.Error())
	}
}
