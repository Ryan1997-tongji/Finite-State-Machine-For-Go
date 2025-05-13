package utils

import "reflect"

func SliceContains[T comparable](slice []T, item T) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func IsPointer(i interface{}) bool {
	t := reflect.TypeOf(i)
	if t == nil {
		return false
	}
	return t.Kind() == reflect.Ptr
}

func IsNilPointer(i interface{}) bool {
	if !IsPointer(i) {
		return false
	}
	v := reflect.ValueOf(i)
	return v.IsNil()
}
