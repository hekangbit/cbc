package models

import (
	"math"
)

type IntegerType struct {
	*BaseType
	size     int64
	isSigned bool
	name     string
}

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

func (i *IntegerType) IsSigned() (bool, error) {
	return i.isSigned, nil
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

func (i *IntegerType) IsSameType(other IType) bool {
	if !other.IsInteger() {
		return false
	}
	otherInt := other.GetIntegerType()
	return i.Equals(otherInt)
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

func (i *IntegerType) GetIntegerType() *IntegerType {
	return i
}

func (i *IntegerType) String() string {
	return i.name
}
