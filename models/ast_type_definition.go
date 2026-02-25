package models

type IASTTypeDefinition interface {
	INode
	Name() string
	TypeNode() *ASTTypeNode
	TypeRef() ITypeRef
	Type() IType
	DefiningType() IType
	Accept(visitor IDeclarationVisitor) any
}

type ASTTypeDefinition struct {
	Node
	name     string
	location *Location
	typeNode *ASTTypeNode
}

func (this *ASTTypeDefinition) Name() string {
	return this.name
}

func (this *ASTTypeDefinition) Location() *Location {
	return this.location
}

func (this *ASTTypeDefinition) TypeNode() *ASTTypeNode {
	return this.typeNode
}

func (this *ASTTypeDefinition) TypeRef() ITypeRef {
	return this.typeNode.TypeRef()
}

func (this *ASTTypeDefinition) Type() IType {
	return this.typeNode.Type()
}
