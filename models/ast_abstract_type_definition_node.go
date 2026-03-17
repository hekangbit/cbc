package models

type IASTAbstractTypeDefinitionNode interface {
	INode
	Name() string
	TypeNode() *ASTTypeNode
	TypeRef() ITypeRef
	Type() IType
	DefiningType() IType
	Accept(visitor IDeclarationVisitor) (any, error)
}

type ASTAbstractTypeDefinitionNode struct {
	Node
	name     string
	location *Location
	typeNode *ASTTypeNode
}

func (this *ASTAbstractTypeDefinitionNode) Name() string {
	return this.name
}

func (this *ASTAbstractTypeDefinitionNode) Location() *Location {
	return this.location
}

func (this *ASTAbstractTypeDefinitionNode) TypeNode() *ASTTypeNode {
	return this.typeNode
}

func (this *ASTAbstractTypeDefinitionNode) TypeRef() ITypeRef {
	return this.typeNode.TypeRef()
}

func (this *ASTAbstractTypeDefinitionNode) Type() IType {
	return this.typeNode.Type()
}
