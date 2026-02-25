package models

type ASTLogicalOrNode struct {
	ASTBinaryOpNode
}

var _ IASTExprNode = &ASTLogicalOrNode{}

func NewASTLogicalOrNode(left IASTExprNode, right IASTExprNode) *ASTLogicalOrNode {
	p := &ASTLogicalOrNode{ASTBinaryOpNode: ASTBinaryOpNode{left: left, operator: "||", right: right}}
	p.ASTBinaryOpNode.ASTExprNode._impl = p
	p.ASTBinaryOpNode.ASTExprNode.Node._impl = p
	return p
}

func (this *ASTLogicalOrNode) Accept(visitor IASTVisitor) interface{} {
	return visitor.VisitLogicalOrNode(this)
}
