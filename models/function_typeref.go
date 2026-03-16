package models

import (
	"strings"
)

type FunctionTypeRef struct {
	BaseTypeRef
	retTypeRef ITypeRef
	params     *ParamTypeRefs
}

var _ ITypeRef = &FunctionTypeRef{}

func NewFunctionTypeRef(retTypeRef ITypeRef, params *ParamTypeRefs) *FunctionTypeRef {
	return &FunctionTypeRef{
		BaseTypeRef: BaseTypeRef{location: retTypeRef.Location()},
		retTypeRef:  retTypeRef,
		params:      params,
	}
}

func (this *FunctionTypeRef) IsFunction() bool {
	return true
}

func (this *FunctionTypeRef) ReturnType() ITypeRef {
	return this.retTypeRef
}

func (this *FunctionTypeRef) Params() *ParamTypeRefs {
	return this.params
}

func (this *FunctionTypeRef) String() string {
	var buf strings.Builder
	buf.WriteString(this.retTypeRef.String())
	buf.WriteString(" (")

	sep := ""
	for _, ref := range this.params.TypeRefs() {
		buf.WriteString(sep)
		buf.WriteString(ref.String())
		sep = ", "
	}

	buf.WriteString(")")
	return buf.String()
}
