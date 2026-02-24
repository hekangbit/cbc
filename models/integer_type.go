package models

import (
	"math"
)

type IntegerType struct {
	BaseType
	size     int64
	isSigned bool
	name     string
}

var _ IType = &IntegerType{}

func NewIntegerType(size int64, isSigned bool, name string) *IntegerType {
	return &IntegerType{
		size:     size,
		isSigned: isSigned,
		name:     name,
	}
}

func (i *IntegerType) IsInteger() bool {
	return true
}

func (i *IntegerType) IsSigned() bool {
	return i.isSigned
}

func (i *IntegerType) IsScalar() bool {
	return true
}

func (i *IntegerType) MinValue() int64 {
	if i.isSigned {
		return -int64(math.Pow(2, float64(i.size*8-1)))
	}
	return 0
}

func (i *IntegerType) MaxValue() int64 {
	if i.isSigned {
		return int64(math.Pow(2, float64(i.size*8-1))) - 1
	}
	return int64(math.Pow(2, float64(i.size*8))) - 1
}

func (i *IntegerType) IsInDomain(value int64) bool {
	return i.MinValue() <= value && value <= i.MaxValue()
}

// TODO: Cast may return error? IsSame means Ptr same
// use typeTable, same type always points to same object, singleton mode
func (i *IntegerType) IsSameType(other IType) bool {
	if !other.IsInteger() {
		return false
	}
	t := GetIntegerType(other)
	return i == t
}

func (i *IntegerType) Equals(other interface{}) bool {
	otherType, ok := other.(*IntegerType)
	if !ok {
		return false
	}

	return i.size == otherType.size && i.isSigned == otherType.isSigned && i.name == otherType.name
}

func (i *IntegerType) IsCompatible(other IType) bool {
	if !other.IsInteger() {
		return false
	}

	otherSize := other.Size()
	return i.size <= otherSize
}

func (i *IntegerType) IsCastableTo(target IType) bool {
	return target.IsInteger() || target.IsPointer()
}

func (i *IntegerType) Size() int64 {
	return i.size
}

func (i *IntegerType) AllocSize() int64 {
	return i.size
}

func (i *IntegerType) Alignment() int64 {
	return i.size
}

func (i *IntegerType) String() string {
	return i.name
}
