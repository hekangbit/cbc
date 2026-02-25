package models

type FunctionType struct {
	BaseType
	returnType IType
	paramTypes *ParamTypes
}

var _ IType = &FunctionType{}

func NewFunctionType(ret IType, partypes *ParamTypes) *FunctionType {
	return &FunctionType{
		returnType: ret,
		paramTypes: partypes,
	}
}

func (this *FunctionType) IsFunction() bool {
	return true
}

func (this *FunctionType) IsCallable() bool {
	return true
}

func (this *FunctionType) IsSameType(other IType) bool {
	if !other.IsFunction() {
		return false
	}
	t, ok := other.(*FunctionType)
	if !ok {
		return false
	}
	return t.returnType.IsSameType(this.returnType) && t.paramTypes.IsSameType(this.paramTypes)
}

func (this *FunctionType) IsCompatible(target IType) bool {
	if !target.IsFunction() {
		return false
	}
	t, ok := target.(*FunctionType)
	if !ok {
		return false
	}
	return t.returnType.IsCompatible(this.returnType) && t.paramTypes.IsSameType(this.paramTypes) // 注意：原代码这里用 IsSameType
}

func (this *FunctionType) IsCastableTo(target IType) bool {
	return target.IsFunction()
}

func (this *FunctionType) ReturnType() IType {
	return this.returnType
}

func (this *FunctionType) IsVararg() bool {
	return this.paramTypes.IsVararg()
}

func (this *FunctionType) AcceptsArgc(numArgs int) bool {
	if this.paramTypes.IsVararg() {
		return numArgs >= this.paramTypes.MinArgc()
	}
	return numArgs == this.paramTypes.Argc()
}

func (this *FunctionType) ParamTypes() []IType {
	return this.paramTypes.Types()
}

func (this *FunctionType) Alignment() int64 {
	panic("FunctionType#Alignment called")
}

// TODO: java throw error: FunctionType#size called
func (this *FunctionType) Size() int64 {
	panic("FunctionType#Size called")
}

func (this *FunctionType) AllocSize() int64 {
	return this.Size()
}

// TODO: correct string format
func (this *FunctionType) String() string {
	buf := this.returnType.String() + "("
	sep := ""
	for _, t := range this.paramTypes.Types() {
		buf += sep + t.String()
		sep = ", "
	}
	buf += ")"
	return buf
}
