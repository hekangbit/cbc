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

func (t *PointerType) Size() int64 {
	return t.size
}

func (t *PointerType) AllocSize() int64 {
	return t.size
}

func (t *PointerType) Alignment() int64 {
	return t.size
}

func (t *PointerType) IsPointer() bool {
	return true
}

func (t *PointerType) IsScalar() bool {
	return true
}

func (t *PointerType) IsSigned() bool {
	return false
}

func (t *PointerType) IsCallable() bool {
	return false
}

func (t *PointerType) ElemType() IType {
	return t.elemType
}

func (t *PointerType) Equals(other interface{}) bool {
	otherPtr, ok := other.(*PointerType)
	if !ok {
		return false
	}

	// 比较基类型是否相等
	return t.elemType.IsSameType(otherPtr.elemType)
}

// IsSameType 检查类型是否相同
func (t *PointerType) IsSameType(other IType) bool {
	if !other.IsPointer() {
		return false
	}
	otherElemType := other.ElemType()
	return t.elemType.IsSameType(otherElemType)
}

// IsCompatible 检查类型是否兼容
func (t *PointerType) IsCompatible(other IType) bool {
	if !other.IsPointer() {
		return false
	}

	otherElemType := other.ElemType()

	// 如果当前指针指向void，或者目标指针指向void，都兼容
	if t.elemType.IsVoid() || otherElemType.IsVoid() {
		return true
	}

	return t.elemType.IsCompatible(otherElemType)
}

func (t *PointerType) GetPointerType() *PointerType {
	return t
}

// IsCastableTo 检查是否可以强制转换到目标类型
func (t *PointerType) IsCastableTo(target IType) bool {
	return target.IsPointer() || target.IsInteger()
}

// String 返回字符串表示
func (t *PointerType) String() string {
	return fmt.Sprintf("%s*", t.elemType)
}
