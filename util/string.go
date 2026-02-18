package util

import (
	"fmt"
	"reflect"
)

// TODO: need check again
func ToString(obj interface{}) string {
	if obj == nil {
		return "nil"
	}

	// get actual type
	t := reflect.TypeOf(obj)
	typeName := t.String()

	// Get hash based on address
	val := reflect.ValueOf(obj)
	var hashCode int32
	switch val.Kind() {
	case reflect.Ptr, reflect.Slice, reflect.Map, reflect.Chan, reflect.Func, reflect.UnsafePointer:
		hashCode = int32(val.Pointer() & 0x7fffffff)
	default:
		hashCode = HashCode(fmt.Sprintf("%v", obj))
	}

	return fmt.Sprintf("%s@%x", typeName, uint32(hashCode))
}

func HashCode(s string) int32 {
	var h int32 = 0
	for i := 0; i < len(s); i++ {
		h = 31*h + int32(s[i])
	}
	return h
}
