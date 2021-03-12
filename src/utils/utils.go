package utils

import (
	"reflect"
)

// Sizeof computing byte size of structure
func Sizeof(T interface{}) int {
	return int(reflect.TypeOf(T).Size())
}
