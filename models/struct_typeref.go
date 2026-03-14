package models

type StructTypeRef struct {
	BaseTypeRef
	name string
}

var _ ITypeRef = &StructTypeRef{}

func NewStructTypeRef(name string) *StructTypeRef {
	p := new(StructTypeRef)
	p.location = nil
	p.name = name
	return p
}

func NewStructTypeRefWithLoc(loc *Location, name string) *StructTypeRef {
	p := new(StructTypeRef)
	p.location = loc
	p.name = name
	return p
}

func (this *StructTypeRef) IsStruct() bool {
	return true
}

func (this *StructTypeRef) Name() string {
	return this.name
}

func (this *StructTypeRef) String() string {
	return "struct " + this.name
}
