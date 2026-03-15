package models

type ASTTypedefNode struct {
	ASTAbstractTypeDefinitionNode
	real *ASTTypeNode
}

var _ IASTAbstractTypeDefinitionNode = &ASTTypedefNode{}

func NewASTTypedefNode(loc *Location, realTyRef ITypeRef, name string) *ASTTypedefNode {
	p := new(ASTTypedefNode)
	p.location = loc
	p.typeNode = NewASTTypeNodeFromRef(NewUserTypeRef(name))
	p.name = name
	p._impl = p
	return p
}

func (this *ASTTypedefNode) IsUserType() bool {
	return true
}

func (this *ASTTypedefNode) RealTypeNode() *ASTTypeNode {
	return this.real
}

func (this *ASTTypedefNode) RealType() IType {
	return this.real.Type()
}

func (this *ASTTypedefNode) RealTypeRef() ITypeRef {
	return this.real.TypeRef()
}

func (this *ASTTypedefNode) DefiningType() IType {
	return NewUserType(this.Name(), this.RealTypeNode(), this.Location())
}

func (this *ASTTypedefNode) _Dump(d *Dumper) {
	d.PrintMemberStringNotResolved("name", this.name)
	d.PrintMemberDumpable("typeNode", this.typeNode)
}

func (this *ASTTypedefNode) Accept(visitor IDeclarationVisitor) any {
	return visitor.VisitTypedefNode(this)
}
