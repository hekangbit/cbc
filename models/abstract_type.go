package models

import (
	"fmt"
)

const SizeUnknown int64 = -1

type IType interface {
	fmt.Stringer
	Size() int64
	AllocSize() int64
	Alignment() int64
	IsSameType(other IType) bool
	IsVoid() bool
	IsInt() bool
	IsInteger() bool
	IsSigned() bool
	IsPointer() bool
	IsArray() bool
	IsCompositeType() bool
	IsStruct() bool
	IsUnion() bool
	IsUserType() bool
	IsFunction() bool
	IsAllocatedArray() bool
	IsIncompleteArray() bool
	IsScalar() bool
	IsCallable() bool
	IsCompatible(other IType) bool
	IsCastableTo(target IType) bool
	ElemType() IType
	GetIntegerType() *IntegerType
	GetPointerType() *PointerType
	GetFunctionType() *FunctionType
	GetCompositeType() ICompositeType
	// GetIntegerType() (*IntegerType, error)
	// GetPointerType() (*PointerType, error)
	// GetFunctionType() (*FunctionType, error)
	// GetCompositeType() (*CompositeType, error)
	// GetStructType() (*StructType, error)
	// GetUnionType() (*UnionType, error)
	// GetArrayType() (*ArrayType, error)
}

type BaseType struct{
}

func (this *BaseType) Size() int64 {
	panic("Size method must be implemented by concrete type")
}

func (this *BaseType) AllocSize() int64 {
	panic("AllocSize method must be implemented by concrete type")
}

func (this *BaseType) Alignment() int64 {
	panic("Alignment method must be implemented by concrete type")
}

func (this *BaseType) IsSameType(other IType) bool {
	panic("IsSameType method must be implemented by concrete type")
}

func (this *BaseType) IsVoid() bool {
	return false
}

func (this *BaseType) IsInt() bool {
	return false
}

func (this *BaseType) IsInteger() bool {
	return false
}

func (this *BaseType) IsSigned() bool {
	panic("#isSigned for non-integer type")
}

func (this *BaseType) IsPointer() bool {
	return false
}

func (this *BaseType) IsArray() bool {
	return false
}

func (this *BaseType) IsCompositeType() bool {
	return false
}

func (this *BaseType) IsStruct() bool {
	return false
}

func (this *BaseType) IsUnion() bool {
	return false
}

func (this *BaseType) IsUserType() bool {
	return false
}

func (this *BaseType) IsFunction() bool {
	return false
}

func (this *BaseType) IsAllocatedArray() bool {
	return false
}

func (this *BaseType) IsIncompleteArray() bool {
	return false
}

func (this *BaseType) IsScalar() bool {
	return false
}

func (this *BaseType) IsCallable() bool {
	return false
}

func (this *BaseType) ElemType() IType {
	panic("#baseType called for undereferable type")
}

func (this *BaseType) GetIntegerType() *IntegerType {
	panic("#not an integer type")
}

func (this *BaseType) GetPointerType() *PointerType {
	panic("not a pointer type")
}

func (this *BaseType) GetCompositeType() ICompositeType {
	panic("not a pointer type")
}

func (this *BaseType) GetFunctionType() *FunctionType {
	panic("not a function type")
}
