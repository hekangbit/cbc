package models

type UnionTypeRef struct {
	BaseTypeRef
	name string
}

var _ ITypeRef = &UnionTypeRef{}

func NewUnionTypeRef(name string) *UnionTypeRef {
	p := new(UnionTypeRef)
	p.location = nil
	p.name = name
	return p
}

func NewUnionTypeRefWithLoc(loc *Location, name string) *UnionTypeRef {
	p := new(UnionTypeRef)
	p.location = loc
	p.name = name
	return p
}

func (this *UnionTypeRef) IsUnion() bool {
	return true
}

func (this *UnionTypeRef) Name() string {
	return this.name
}

func (this *UnionTypeRef) String() string {
	return "union " + this.name
}
