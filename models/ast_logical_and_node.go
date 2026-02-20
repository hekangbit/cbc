package models

type ASTLogicalAndNode struct {
	*ASTBinaryOpNode
}

func NewASTLogicalAndNode(left IASTExprNode, right IASTExprNode) *ASTLogicalAndNode {
	return &ASTLogicalAndNode{
		ASTBinaryOpNode: NewASTBinaryOpNode(left, "||", right),
	}
}

func (this *ASTLogicalAndNode) Accept(visitor IASTVisitor) interface{} {
	return visitor.VisitLogicalAndNode(this)
}
