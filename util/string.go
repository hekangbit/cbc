package util

import (
	"fmt"
	"reflect"
)

// JavaStyleToString 简洁的工具函数
func ToString(obj interface{}) string {
	if obj == nil {
		return "nil"
	}

	// 获取类型
	t := reflect.TypeOf(obj)
	typeName := t.String()

	// 获取哈希（基于地址）
	val := reflect.ValueOf(obj)
	var hashCode int32

	switch val.Kind() {
	case reflect.Ptr, reflect.Slice, reflect.Map, reflect.Chan, reflect.Func, reflect.UnsafePointer:
		// 引用类型：使用指针地址
		hashCode = int32(val.Pointer() & 0x7fffffff)
	default:
		// 值类型：使用内容的哈希
		// 这里简单使用值的字符串表示
		hashCode = HashCode(fmt.Sprintf("%v", obj))
	}

	return fmt.Sprintf("%s@%x", typeName, uint32(hashCode))
}

// 字符串哈希函数
func HashCode(s string) int32 {
	var h int32 = 0
	for i := 0; i < len(s); i++ {
		h = 31*h + int32(s[i])
	}
	return h
}
