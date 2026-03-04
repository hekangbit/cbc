package models

type ASTStructNode struct {
	ASTCompositeTypeDefinition
}

var _ IASTCompositeTypeDefinition = &ASTStructNode{}

func NewASTStructNode(loc *Location, ref ITypeRef, name string, members []*Slot) *ASTStructNode {
	p := &ASTStructNode{
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

func (this *ASTStructNode) Kind() string {
	return "struct"
}

func (this *ASTStructNode) IsStruct() bool {
	return true
}

// Used by type resolver
func (this *ASTStructNode) DefiningType() IType {
	return NewStructType(this.Name(), this.Members(), this.Location())
}

func (this *ASTStructNode) Accept(visitor IDeclarationVisitor) any {
	return visitor.VisitStructNode(this)
}
