package models

type ASTAssignNode struct {
	ASTAbstractAssignNode
}

var _ IASTAbstractAssignNode = &ASTAssignNode{}

func NewASTAssignNode(lhs IASTExprNode, rhs IASTExprNode) *ASTAssignNode {
	p := &ASTAssignNode{ASTAbstractAssignNode: ASTAbstractAssignNode{lhs: lhs, rhs: rhs}}
	p.ASTAbstractAssignNode.ASTExprNode._impl = p
	p.ASTAbstractAssignNode.ASTExprNode.Node._impl = p
	return p
}

func (this *ASTAssignNode) Accept(visitor IASTVisitor) interface{} {
	return visitor.VisitAssignNode(this)
}
