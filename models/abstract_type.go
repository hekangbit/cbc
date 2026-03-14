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
	GetIntegerType() *IntegerType
	GetPointerType() *PointerType
	GetCompositeType() ICompositeType
	GetFunctionType() *FunctionType
	GetArrayType() *ArrayType
	GetStructType() *StructType
	GetUnionType() *UnionType
}

type BaseType struct {
	_impl IType
}

func (this *BaseType) AllocSize() int64 {
	return this._impl.Size()
}

func (this *BaseType) Alignment() int64 {
	return this._impl.AllocSize()
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
	panic("#ElemType called for undereferable type")
}

func (this *BaseType) GetIntegerType() *IntegerType {
	target, ok := this._impl.(*IntegerType)
	if !ok {
		return nil
	}
	return target
}

func (this *BaseType) GetPointerType() *PointerType {
	target, ok := this._impl.(*PointerType)
	if !ok {
		return nil
	}
	return target
}

func (this *BaseType) GetCompositeType() ICompositeType {
	target, ok := this._impl.(ICompositeType)
	if !ok {
		return nil
	}
	return target
}

func (this *BaseType) GetFunctionType() *FunctionType {
	target, ok := this._impl.(*FunctionType)
	if !ok {
		return nil
	}
	return target
}

func (this *BaseType) GetArrayType() *ArrayType {
	target, ok := this._impl.(*ArrayType)
	if !ok {
		return nil
	}
	return target
}

func (this *BaseType) GetStructType() *StructType {
	target, ok := this._impl.(*StructType)
	if !ok {
		return nil
	}
	return target
}

func (this *BaseType) GetUnionType() *UnionType {
	target, ok := this._impl.(*UnionType)
	if !ok {
		return nil
	}
	return target
}
