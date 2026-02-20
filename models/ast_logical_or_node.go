package models

type ASTLogicalOrNode struct {
	*ASTBinaryOpNode
}

func NewASTLogicalOrNode(left IASTExprNode, right IASTExprNode) *ASTLogicalOrNode {
	return &ASTLogicalOrNode{
		ASTBinaryOpNode: NewASTBinaryOpNode(left, "||", right),
	}
}

func (this *ASTLogicalOrNode) Accept(visitor IASTVisitor) interface{} {
	return visitor.VisitLogicalOrNode(this)
}
