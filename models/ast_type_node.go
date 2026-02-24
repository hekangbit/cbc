package models

type TypeNode struct {
	typeRef ITypeRef
	typ     IType
}

var _ INode = &TypeNode{}

func NewTypeNodeFromRef(ref ITypeRef) *TypeNode {
	return &TypeNode{
		typeRef: ref,
		typ:     nil,
	}
}

func NewTypeNodeFromType(typ IType) *TypeNode {
	return &TypeNode{
		typeRef: nil,
		typ:     typ,
	}
}

func (this *TypeNode) TypeRef() ITypeRef {
	return this.typeRef
}

func (this *TypeNode) Type() IType {
	if this.typ == nil {
		panic("TypeNode not resolved and no typeRef available")
	}
	return this.typ
}

func (this *TypeNode) IsResolved() bool {
	return this.typ != nil
}

func (this *TypeNode) SetType(typ IType) {
	if this.typ != nil {
		panic("TypeNode#SetType called twice")
	}
	this.typ = typ
}

func (this *TypeNode) Location() *Location {
	if this.typeRef != nil {
		return this.typeRef.Location()
	}
	return nil
}

func (this *TypeNode) Dump(d *Dumper) {
	d.PrintMemberTypeRef("typeRef", this.typeRef)
	d.PrintMemberType("type", this.typ)
}
