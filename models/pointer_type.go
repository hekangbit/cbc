package models

import "fmt"

type PointerType struct {
	BaseType
	size     int64
	elemType IType
}

func NewPointerType(size int64, elemType IType) *PointerType {
	return &PointerType{
		size:     size,
		elemType: elemType,
	}
}

func (this *PointerType) Size() int64 {
	return this.size
}

func (this *PointerType) AllocSize() int64 {
	return this.size
}

func (this *PointerType) Alignment() int64 {
	return this.size
}

func (this *PointerType) IsPointer() bool {
	return true
}

func (this *PointerType) IsScalar() bool {
	return true
}

func (this *PointerType) IsSigned() bool {
	return false
}

func (this *PointerType) IsCallable() bool {
	return false
}

func (this *PointerType) ElemType() IType {
	return this.elemType
}

func (this *PointerType) Equals(other interface{}) bool {
	otherPtr, ok := other.(*PointerType)
	if !ok {
		return false
	}
	// compare elem type
	return this.elemType.IsSameType(otherPtr.elemType)
}

func (this *PointerType) IsSameType(other IType) bool {
	if !other.IsPointer() {
		return false
	}
	otherElemType := other.ElemType()
	return this.elemType.IsSameType(otherElemType)
}

func (this *PointerType) IsCompatible(other IType) bool {
	if !other.IsPointer() {
		return false
	}

	otherElemType := other.ElemType()

	if this.elemType.IsVoid() || otherElemType.IsVoid() {
		return true
	}

	return this.elemType.IsCompatible(otherElemType)
}

func (this *PointerType) GetPointerType() *PointerType {
	return this
}

func (this *PointerType) IsCastableTo(target IType) bool {
	return target.IsPointer() || target.IsInteger()
}

func (this *PointerType) String() string {
	return fmt.Sprintf("%s*", this.elemType)
}
