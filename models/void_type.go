package models

type VoidType struct {
	BaseType
}

func NewVoidType() *VoidType {
	return &VoidType{}
}

func (v *VoidType) Size() int64 {
	return 1
}

func (v *VoidType) AllocSize() int64 {
	return v.Size()
}

func (v *VoidType) Alignment() int64 {
	return 1
}

func (v *VoidType) IsSameType(other IType) bool {
	return other.IsVoid()
}

func (v *VoidType) IsCompatible(other IType) bool {
	return other.IsVoid()
}

func (v *VoidType) IsCastableTo(target IType) bool {
	return target.IsVoid()
}

func (v *VoidType) IsVoid() bool {
	return true
}

func (v *VoidType) Equals(other interface{}) bool {
	_, ok := other.(*VoidType)
	return ok
}

func (v *VoidType) String() string {
	return "void"
}
