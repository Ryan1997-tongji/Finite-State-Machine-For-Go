package fsm_impl

import (
	"fmt"
	"github.com/fsm/utils"
)

func GetFact(input interface{}) (*Object, error) {
	if !utils.IsPointer(input) {
		return nil, fmt.Errorf("input is not a pointer")
	}
	if utils.IsNilPointer(input) {
		return nil, fmt.Errorf("input is nil")
	}
	obj, ok := input.(*Object)
	if !ok {
		return nil, fmt.Errorf("input is not a *Object")
	}
	if utils.IsNilPointer(obj) {
		return nil, fmt.Errorf("obj is nil")
	}
	return obj, nil
}
