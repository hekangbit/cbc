package models

import (
	"reflect"
	"strings"
)

// FunctionTypeRef 函数类型引用
type FunctionTypeRef struct {
	*BaseTypeRef
	retTypeRef ITypeRef
	params     *ParamTypeRefs
}

// NewFunctionTypeRef 创建函数类型引用
func NewFunctionTypeRef(retTypeRef ITypeRef, params *ParamTypeRefs) *FunctionTypeRef {
	return &FunctionTypeRef{
		BaseTypeRef: NewBaseTypeRef(retTypeRef.Location()),
		retTypeRef:  retTypeRef,
		params:      params,
	}
}

// IsFunction 检查是否为函数类型引用
func (f *FunctionTypeRef) IsFunction() bool {
	return true
}

// Equals 检查两个函数类型引用是否相等（重载版本1）
func (f *FunctionTypeRef) Equals(other interface{}) bool {
	otherRef, ok := other.(*FunctionTypeRef)
	if !ok {
		return false
	}
	return f.Equals_(otherRef)
}

// equals 检查两个函数类型引用是否相等（重载版本2）
func (f *FunctionTypeRef) Equals_(other *FunctionTypeRef) bool {
	return reflect.DeepEqual(f.retTypeRef, other.retTypeRef) && f.params.Equals(other.params)
}

// ReturnType 返回返回类型
func (f *FunctionTypeRef) ReturnType() ITypeRef {
	return f.retTypeRef
}

// Params 返回参数列表
func (f *FunctionTypeRef) Params() *ParamTypeRefs {
	return f.params
}

// String 返回字符串表示
func (f *FunctionTypeRef) String() string {
	var buf strings.Builder
	buf.WriteString(f.retTypeRef.String())
	buf.WriteString(" (")

	sep := ""
	for _, ref := range f.params.TypeRefs() {
		buf.WriteString(sep)
		buf.WriteString(ref.String())
		sep = ", "
	}

	buf.WriteString(")")
	return buf.String()
}
