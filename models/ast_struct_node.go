package models

type ASTStructNode struct {
	ASTAbstractCompositeTypeDefinitionNode
}

var _ IASTAbstractCompositeTypeDefinitionNode = &ASTStructNode{}

func NewASTStructNode(loc *Location, ref ITypeRef, name string, members []*Slot) *ASTStructNode {
	p := new(ASTStructNode)
	p.name = name
	p.location = loc
	p.typeNode = NewASTTypeNodeFromRef(ref)
	p.members = members
	p._impl = p
	return p
}

func (this *ASTStructNode) Kind() string {
	return "struct"
}

// Used by type resolver
func (this *ASTStructNode) DefiningType() IType {
	return NewStructType(this.Name(), this.Members(), this.Location())
}

func (this *ASTStructNode) Accept(visitor IDeclarationVisitor) any {
	return visitor.VisitStructNode(this)
}
