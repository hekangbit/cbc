package models

type ASTUnionNode struct {
	ASTAbstractCompositeTypeDefinitionNode
}

var _ IASTAbstractCompositeTypeDefinitionNode = &ASTUnionNode{}

func NewASTUnionNode(loc *Location, ref ITypeRef, name string, members []*Slot) *ASTUnionNode {
	p := &ASTUnionNode{
		ASTAbstractCompositeTypeDefinitionNode: ASTAbstractCompositeTypeDefinitionNode{
			ASTAbstractTypeDefinitionNode: ASTAbstractTypeDefinitionNode{
				name:     name,
				location: loc,
				typeNode: NewASTTypeNodeFromRef(ref),
			},
			members: members,
		},
	}

	p._impl = p
	return p
}

func (this *ASTUnionNode) Kind() string {
	return "union"
}

// Used by type resolver
func (this *ASTUnionNode) DefiningType() IType {
	return NewUnionType(this.Name(), this.Members(), this.Location())
}

func (this *ASTUnionNode) Accept(visitor IDeclarationVisitor) any {
	return visitor.VisitUnionNode(this)
}
