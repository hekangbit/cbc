package models

type ASTTypeNode struct {
	Node
	typeRef ITypeRef
	typ     IType
}

var _ INode = &ASTTypeNode{}

func NewTypeNodeFromRef(ref ITypeRef) *ASTTypeNode {
	p := &ASTTypeNode{typeRef: ref, typ: nil}
	p._impl = p
	return p
}

func NewTypeNodeFromType(typ IType) *ASTTypeNode {
	p := &ASTTypeNode{typeRef: nil, typ: typ}
	p._impl = p
	return p
}

func (this *ASTTypeNode) TypeRef() ITypeRef {
	return this.typeRef
}

func (this *ASTTypeNode) Type() IType {
	if this.typ == nil {
		panic("ASTTypeNode not resolved and no typeRef available")
	}
	return this.typ
}

func (this *ASTTypeNode) IsResolved() bool {
	return this.typ != nil
}

func (this *ASTTypeNode) SetType(typ IType) {
	if this.typ != nil {
		panic("ASTTypeNode#SetType called twice")
	}
	this.typ = typ
}

func (this *ASTTypeNode) Location() *Location {
	if this.typeRef != nil {
		return this.typeRef.Location()
	}
	return nil
}

func (this *ASTTypeNode) _Dump(d *Dumper) {
	d.PrintMemberTypeRef("typeRef", this.typeRef)
	d.PrintMemberType("type", this.typ)
}
