package models

import (
	"math"
)

// IntegerType 结构体
type IntegerType struct {
	*BaseType
	size     int64
	isSigned bool
	name     string
}

// 构造函数
func NewIntegerType(size int64, isSigned bool, name string) *IntegerType {
	return &IntegerType{
		size:     size,
		isSigned: isSigned,
		name:     name,
	}
}

// 重写接口方法
func (i *IntegerType) IsInteger() bool {
	return true
}

func (i *IntegerType) IsSigned() (bool, error) {
	return i.isSigned, nil
}

func (i *IntegerType) IsScalar() bool {
	return true
}

// MinValue 返回最小值
func (i *IntegerType) MinValue() int64 {
	if i.isSigned {
		return -int64(math.Pow(2, float64(i.size*8-1)))
	}
	return 0
}

// MaxValue 返回最大值
func (i *IntegerType) MaxValue() int64 {
	if i.isSigned {
		return int64(math.Pow(2, float64(i.size*8-1))) - 1
	}
	return int64(math.Pow(2, float64(i.size*8))) - 1
}

// IsInDomain 检查值是否在范围内
func (i *IntegerType) IsInDomain(value int64) bool {
	return i.MinValue() <= value && value <= i.MaxValue()
}

// IsSameType 检查类型是否相同
func (i *IntegerType) IsSameType(other IType) bool {
	if !other.IsInteger() {
		return false
	}

	// 获取整数类型进行比较
	otherInt := other.GetIntegerType()
	return i.Equals(otherInt)
}

// Equals 检查两个IntegerType是否相等
func (i *IntegerType) Equals(other interface{}) bool {
	otherType, ok := other.(*IntegerType)
	if !ok {
		return false
	}

	return i.size == otherType.size && i.isSigned == otherType.isSigned && i.name == otherType.name
}

// IsCompatible 检查类型是否兼容
func (i *IntegerType) IsCompatible(other IType) bool {
	if !other.IsInteger() {
		return false
	}

	otherSize := other.Size()
	return i.size <= otherSize
}

// IsCastableTo 检查是否可以强制转换到目标类型
func (i *IntegerType) IsCastableTo(target IType) bool {
	return target.IsInteger() || target.IsPointer()
}

// Size 返回类型大小
func (i *IntegerType) Size() int64 {
	return i.size
}

// AllocSize 分配大小
func (i *IntegerType) AllocSize() int64 {
	return i.size
}

// Alignment 对齐大小
func (i *IntegerType) Alignment() int64 {
	return i.size
}

// GetIntegerType 类型转换方法
func (i *IntegerType) GetIntegerType() *IntegerType {
	return i
}

// func (i *IntegerType) GetIntegerType() (*IntegerType, error) {
// 	return i, nil
// }

// String 返回字符串表示
func (i *IntegerType) String() string {
	return i.name
}

// 辅助函数，用于创建常见的整数类型
func NewInt8Type() *IntegerType {
	return NewIntegerType(1, true, "int8")
}

func NewUInt8Type() *IntegerType {
	return NewIntegerType(1, false, "uint8")
}

func NewInt16Type() *IntegerType {
	return NewIntegerType(2, true, "int16")
}

func NewUInt16Type() *IntegerType {
	return NewIntegerType(2, false, "uint16")
}

func NewInt32Type() *IntegerType {
	return NewIntegerType(4, true, "int32")
}

func NewUInt32Type() *IntegerType {
	return NewIntegerType(4, false, "uint32")
}

func NewInt64Type() *IntegerType {
	return NewIntegerType(8, true, "int64")
}

func NewUInt64Type() *IntegerType {
	return NewIntegerType(8, false, "uint64")
}
