package models

type IASTLiteralNode interface {
	IASTExprNode
	TypeNode() *ASTTypeNode
}

type ASTLiteralNode struct {
	ASTExprNode
	location *Location
	typeNode *ASTTypeNode
}

func (this *ASTLiteralNode) Location() *Location {
	return this.location
}

func (this *ASTLiteralNode) Type() IType {
	return this.typeNode.Type()
}

func (this *ASTLiteralNode) TypeNode() *ASTTypeNode {
	return this.typeNode
}

func (this *ASTLiteralNode) IsConstant() bool {
	return true
}
