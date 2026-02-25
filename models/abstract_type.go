package models

import "fmt"

const SizeUnknown int64 = -1

type IType interface {
	fmt.Stringer
	Size() int64
	AllocSize() int64
	Alignment() int64
	IsSameType(IType) bool
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
	IsCompatible(IType) bool
	IsCastableTo(IType) bool
	ElemType() IType
}

// TODO: these cast method, may need return error, cause java cast can throw exception
func GetIntegerType(t IType) *IntegerType {
	target, ok := t.(*IntegerType)
	if !ok {
		panic("Cast IType to *IntegerType fail")
	}
	return target
}

func GetPointerType(t IType) *PointerType {
	target, ok := t.(*PointerType)
	if !ok {
		panic("Cast IType to *PointerType fail")
	}
	return target
}

func GetCompositeType(t IType) ICompositeType {
	target, ok := t.(ICompositeType)
	if !ok {
		panic("Cast IType to ICompositeType fail")
	}
	return target
}

func GetFunctionType(t IType) *FunctionType {
	target, ok := t.(*FunctionType)
	if !ok {
		panic("Cast IType to *FunctionType fail")
	}
	return target
}

func GetArrayType(t IType) *ArrayType {
	target, ok := t.(*ArrayType)
	if !ok {
		panic("Cast IType to *ArrayType fail")
	}
	return target
}

// TODO
/*
func GetStructType(t IType) *StructType {
	target, ok := t.(*StructType)
	if !ok {
		panic("Cast IType to *StructType fail")
	}
	return target
}
*/

// TODO
/*
func GetUnionType(t IType) *UnionType {
	target, ok := t.(*UnionType)
	if !ok {
		panic("Cast IType to *UnionType fail")
	}
	return target
}
*/

type BaseType struct {
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
	panic("Type#isSigned for non-integer type")
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
	panic("Type#ElemType called for undereferable type")
}
