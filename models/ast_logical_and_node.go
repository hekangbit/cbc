package models

type ASTLogicalAndNode struct {
	ASTBinaryOpNode
}

var _ IASTExprNode = &ASTLogicalAndNode{}

func NewASTLogicalAndNode(left IASTExprNode, right IASTExprNode) *ASTLogicalAndNode {
	p := &ASTLogicalAndNode{ASTBinaryOpNode: ASTBinaryOpNode{left: left, operator: "&&", right: right}}
	p.ASTBinaryOpNode.ASTExprNode._impl = p
	p.ASTBinaryOpNode.ASTExprNode.Node._impl = p
	return p
}

func (this *ASTLogicalAndNode) Accept(visitor IASTVisitor) (any, error) {
	return visitor.VisitLogicalAndNode(this)
}
