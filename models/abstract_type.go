package models

import "fmt"

const SizeUnknown int64 = -1

// Type 接口 - 对应Java中的抽象类
type IType interface {
	fmt.Stringer
	Size() int64
	AllocSize() int64
	Alignment() int64
	IsSameType(other IType) bool
	IsVoid() bool
	IsInt() bool
	IsInteger() bool
	IsSigned() (bool, error)
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
	// 类型转换方法
	GetIntegerType() *IntegerType
	GetPointerType() *PointerType
	GetCompositeType() ICompositeType

	// GetIntegerType() (*IntegerType, error)
	// GetPointerType() (*PointerType, error)
	// GetFunctionType() (*FunctionType, error)
	// GetStructType() (*StructType, error)
	// GetUnionType() (*UnionType, error)
	// GetCompositeType() (*CompositeType, error)
	// GetArrayType() (*ArrayType, error)
}

// BaseType 基础结构体 - 提供默认实现
type BaseType struct{}

func (b *BaseType) Size() int64 {
	panic("Size method must be implemented by concrete type")
}

// 默认方法实现
func (b *BaseType) AllocSize() int64 {
	panic("AllocSize method must be implemented by concrete type")
}

func (b *BaseType) Alignment() int64 {
	panic("Alignment method must be implemented by concrete type")
}

func (b *BaseType) IsSameType(other IType) bool {
	panic("IsSameType method must be implemented by concrete type")
}

func (b *BaseType) IsVoid() bool {
	return false
}

func (b *BaseType) IsInt() bool {
	return false
}

func (b *BaseType) IsInteger() bool {
	return false
}

func (b *BaseType) IsSigned() bool {
	panic("#isSigned for non-integer type")
}

func (b *BaseType) IsPointer() bool {
	return false
}

func (b *BaseType) IsArray() bool {
	return false
}

func (b *BaseType) IsCompositeType() bool {
	return false
}

func (b *BaseType) IsStruct() bool {
	return false
}

func (b *BaseType) IsUnion() bool {
	return false
}

func (b *BaseType) IsUserType() bool {
	return false
}

func (b *BaseType) IsFunction() bool {
	return false
}

func (b *BaseType) IsAllocatedArray() bool {
	return false
}

func (b *BaseType) IsIncompleteArray() bool {
	return false
}

func (b *BaseType) IsScalar() bool {
	return false
}

func (b *BaseType) IsCallable() bool {
	return false
}

func (b *BaseType) ElemType() IType {
	panic("#baseType called for undereferable type")
}

// 类型转换方法的默认实现（返回错误）
func (b *BaseType) GetIntegerType() *IntegerType {
	panic("#not an integer type")
}

func (b *BaseType) GetPointerType() *PointerType {
	panic("not a pointer type")
}

func (b *BaseType) GetCompositeType() ICompositeType {
	panic("not a pointer type")
}

// // 类型转换方法的默认实现（返回错误）
// func (b *BaseType) GetIntegerType() (*IntegerType, error) {
// 	return nil, errors.New("not an integer type")
// }

// func (b *BaseType) GetPointerType() (*PointerType, error) {
// 	return nil, errors.New("not a pointer type")
// }

// func (b *BaseType) GetFunctionType() (*FunctionType, error) {
// 	return nil, errors.New("not a function type")
// }

// func (b *BaseType) GetStructType() (*StructType, error) {
// 	return nil, errors.New("not a struct type")
// }

// func (b *BaseType) GetUnionType() (*UnionType, error) {
// 	return nil, errors.New("not a union type")
// }

// func (b *BaseType) GetCompositeType() (*CompositeType, error) {
// 	return nil, errors.New("not a composite type")
// }

// func (b *BaseType) GetArrayType() (*ArrayType, error) {
// 	return nil, errors.New("not an array type")
// }

// 以下需要具体类型实现的方法，在BaseType中不提供实现
// 具体类型必须实现这些方法
// func (b *BaseType) Size() int64
// func (b *BaseType) IsSameType(other Type) bool
// func (b *BaseType) IsCompatible(other Type) bool
// func (b *BaseType) IsCastableTo(target Type) bool
