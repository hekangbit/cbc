package models

import "fmt"

type ArrayTypeRef struct {
	*BaseTypeRef
	elemTypeRef ITypeRef
	length      int64
}

const undefined = -1

var _ ITypeRef = (*ArrayTypeRef)(nil)

func NewArrayTypeRef(elemTypeRef ITypeRef) *ArrayTypeRef {
	return &ArrayTypeRef{
		BaseTypeRef: NewBaseTypeRef(elemTypeRef.Location()),
		elemTypeRef: elemTypeRef,
		length:      undefined,
	}
}

func NewArrayTypeRefWithLen(elemTypeRef ITypeRef, length int64) *ArrayTypeRef {
	if length < 0 {
		panic("negative array length")
	}
	return &ArrayTypeRef{
		BaseTypeRef: NewBaseTypeRef(elemTypeRef.Location()),
		elemTypeRef: elemTypeRef,
		length:      length,
	}
}

func NewArrayTypeRefEmptyElem() *ArrayTypeRef {
	return &ArrayTypeRef{
		BaseTypeRef: NewBaseTypeRef(nil),
		elemTypeRef: nil,
		length:      undefined,
	}
}

func (this *ArrayTypeRef) isArray() bool {
	return true
}

func (this *ArrayTypeRef) Equals(other interface{}) bool {
	otherRef, ok := other.(*ArrayTypeRef)
	if !ok {
		return false
	}
	return this.length == otherRef.length
}

func (this *ArrayTypeRef) ElemType() ITypeRef {
	return this.elemTypeRef
}

func (this *ArrayTypeRef) Length() int64 {
	return this.length
}

func (this *ArrayTypeRef) IsLengthUndefined() bool {
	return this.length == undefined
}

func (this *ArrayTypeRef) String() string {
	if this.length == undefined {
		return this.elemTypeRef.String() + "[]"
	}
	return fmt.Sprintf("%s[%d]", this.elemTypeRef.String(), this.length)
}

func (this *ArrayTypeRef) SetElemType(ref ITypeRef) {
	this.elemTypeRef = ref
	this.location = ref.Location()
}
