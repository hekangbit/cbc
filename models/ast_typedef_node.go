package models

type ASTTypedefNode struct {
	ASTAbstractTypeDefinitionNode
	real *ASTTypeNode
}

var _ IASTAbstractTypeDefinitionNode = &ASTTypedefNode{}

func NewASTTypedefNode(loc *Location, realTyRef ITypeRef, name string) *ASTTypedefNode {
	p := &ASTTypedefNode{
		ASTAbstractTypeDefinitionNode: ASTAbstractTypeDefinitionNode{
			location: loc,
			typeNode: NewASTTypeNodeFromRef(NewUserTypeRef(name)),
			name:     name,
		},
		real: NewASTTypeNodeFromRef(realTyRef),
	}
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
