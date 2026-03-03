package models

// cb typedef
type UserType struct {
	NamedType
	real *ASTTypeNode
}

// NewUserType 创建 UserType 实例。
func NewUserType(name string, real *ASTTypeNode, loc *Location) *UserType {
	p := &UserType{
		NamedType: NamedType{name: name, location: loc},
		real:      real,
	}
	p._impl = p
	return p
}

func (this *UserType) RealType() IType {
	return this.real.Type()
}

func (this *UserType) String() string {
	return this.name
}

func (this *UserType) Size() int64 {
	return this.RealType().Size()
}

func (this *UserType) AllocSize() int64 {
	return this.RealType().AllocSize()
}

func (this *UserType) Alignment() int64 {
	return this.RealType().Alignment()
}

func (this *UserType) IsVoid() bool {
	return this.RealType().IsVoid()
}

func (this *UserType) IsInt() bool {
	return this.RealType().IsInt()
}

func (this *UserType) IsInteger() bool {
	return this.RealType().IsInteger()
}

func (this *UserType) IsSigned() bool {
	return this.RealType().IsSigned()
}

func (this *UserType) IsPointer() bool {
	return this.RealType().IsPointer()
}

func (this *UserType) IsArray() bool {
	return this.RealType().IsArray()
}

func (this *UserType) IsAllocatedArray() bool {
	return this.RealType().IsAllocatedArray()
}

func (this *UserType) IsCompositeType() bool {
	return this.RealType().IsCompositeType()
}

func (this *UserType) IsStruct() bool {
	return this.RealType().IsStruct()
}

func (this *UserType) IsUnion() bool {
	return this.RealType().IsUnion()
}

func (this *UserType) IsUserType() bool {
	return true
}

func (this *UserType) IsFunction() bool {
	return this.RealType().IsFunction()
}

func (this *UserType) IsCallable() bool {
	return this.RealType().IsCallable()
}

func (this *UserType) IsScalar() bool {
	return this.RealType().IsScalar()
}

func (this *UserType) ElemType() IType {
	return this.RealType().ElemType()
}

func (this *UserType) IsSameType(other IType) bool {
	return this.RealType().IsSameType(other)
}

func (this *UserType) IsCompatible(other IType) bool {
	return this.RealType().IsCompatible(other)
}

func (this *UserType) IsCastableTo(other IType) bool {
	return this.RealType().IsCastableTo(other)
}

func (this *UserType) GetIntegerType() *IntegerType {
	return this.RealType().GetIntegerType()
}

func (this *UserType) GetCompositeType() ICompositeType {
	return this.RealType().GetCompositeType()
}

func (this *UserType) GetStructType() *StructType {
	return this.RealType().GetStructType()
}

func (this *UserType) GetUnionType() *UnionType {
	return this.RealType().GetUnionType()
}

func (this *UserType) GetArrayType() *ArrayType {
	return this.RealType().GetArrayType()
}

func (this *UserType) GetPointerType() *PointerType {
	return this.RealType().GetPointerType()
}

func (this *UserType) GetFunctionType() *FunctionType {
	return this.RealType().GetFunctionType()
}
