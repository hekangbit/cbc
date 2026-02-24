package models

import "strconv"

type ArrayType struct {
	BaseType
	elemType    IType
	length      int64
	pointerSize int64
}

var _ IType = &ArrayType{}

func NewArrayType(t IType, pointerSize int64) *ArrayType {
	return &ArrayType{
		elemType:    t,
		length:      undefined,
		pointerSize: pointerSize,
	}
}

func NewArrayTypeWithLen(t IType, len int64, pointerSize int64) *ArrayType {
	return &ArrayType{
		elemType:    t,
		length:      len,
		pointerSize: pointerSize,
	}
}

func (this *ArrayType) IsArray() bool {
	return true
}

func (this *ArrayType) IsAllocatedArray() bool {
	return this.length != undefined && (!this.elemType.IsArray() || this.elemType.IsAllocatedArray())
}

func (this *ArrayType) IsIncompleteArray() bool {
	if !this.elemType.IsArray() {
		return false
	}
	return !(this.elemType.IsAllocatedArray())
}

func (this *ArrayType) ElemType() IType {
	return this.elemType
}

func (this *ArrayType) Length() int64 {
	return this.length
}

// Value size as pointer
func (this *ArrayType) Size() int64 {
	return this.pointerSize
}

// Value size as allocated array
func (this *ArrayType) AllocSize() int64 {
	if this.length == undefined {
		return this.Size()
	} else {
		return this.elemType.AllocSize() * this.length
	}
}

func (this *ArrayType) Alignment() int64 {
	return this.elemType.Alignment()
}

func (this *ArrayType) Equals(other interface{}) bool {
	t, ok := other.(*ArrayType)
	if !ok {
		return false
	}
	// TODO: java equals is pointer address compare, how about in golang
	return (this.elemType == t.elemType) && (this.length == t.length)
}

func (this *ArrayType) IsSameType(other IType) bool {
	// length is not important
	if !other.IsPointer() && !other.IsArray() {
		return false
	}
	return this.elemType.IsSameType(other.ElemType())
}

func (this *ArrayType) IsCompatible(target IType) bool {
	if !target.IsPointer() && !target.IsArray() {
		return false
	}
	if target.ElemType().IsVoid() {
		return true
	}
	return this.elemType.IsCompatible(target.ElemType()) && this.elemType.Size() == target.ElemType().Size()
}

func (this *ArrayType) IsCasstableTo(target IType) bool {
	return target.IsPointer() || target.IsArray()
}

func (this *ArrayType) String() string {
	if this.length < 0 {
		return this.elemType.String() + "[]"
	}
	return this.elemType.String() + "[" + strconv.FormatInt(this.length, 10) + "]"
}
