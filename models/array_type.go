package models

type ArrayType struct {
	BaseType
	elemType    IType
	length      int64
	pointerSize int64
}
