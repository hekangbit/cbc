package models

type IASTTypeDefinition interface {
	INode
	Name() string
	TypeNode() *TypeNode
	TypeRef() ITypeRef
	Type() IType
	DefiningType() IType
	Accept(visitor IDeclarationVisitor) any
}

type ASTTypeDefinition struct {
	Node
	name     string
	location *Location
	typeNode *TypeNode
}

// TODO: remove interface check, ASTTypeDefinition is abstract klass, no need panic method
var _ IASTTypeDefinition = &ASTTypeDefinition{}

func NewASTTypeDefinition(loc *Location, ref ITypeRef, name string) *ASTTypeDefinition {
	return &ASTTypeDefinition{
		location: loc,
		name:     name,
		typeNode: NewTypeNodeFromRef(ref),
	}
}

func (this *ASTTypeDefinition) Name() string {
	return this.name
}

func (this *ASTTypeDefinition) Location() *Location {
	return this.location
}

func (this *ASTTypeDefinition) TypeNode() *TypeNode {
	return this.typeNode
}

func (this *ASTTypeDefinition) TypeRef() ITypeRef {
	return this.typeNode.TypeRef()
}

func (this *ASTTypeDefinition) Type() IType {
	return this.typeNode.Type()
}

func (this *ASTTypeDefinition) DefiningType() IType {
	panic("abstract method: DefiningType must be implemented by subtype")
}

func (this *ASTTypeDefinition) Accept(visitor IDeclarationVisitor) any {
	panic("abstract method: Accept must be implemented by subtype")
}
