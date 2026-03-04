package models

type ASTUnionNode struct {
	ASTCompositeTypeDefinition
}

var _ IASTCompositeTypeDefinition = &ASTUnionNode{}

func NewASTUnionNode(loc *Location, ref ITypeRef, name string, members []*Slot) *ASTUnionNode {
	p := &ASTUnionNode{
		ASTCompositeTypeDefinition: ASTCompositeTypeDefinition{
			ASTTypeDefinition: ASTTypeDefinition{
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

func (this *ASTUnionNode) IsUnion() bool {
	return true
}

// Used by type resolver
func (this *ASTUnionNode) DefiningType() IType {
	return NewUnionType(this.Name(), this.Members(), this.Location())
}

func (this *ASTUnionNode) Accept(visitor IDeclarationVisitor) any {
	return visitor.VisitUnionNode(this)
}
