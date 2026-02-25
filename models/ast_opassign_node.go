package models

type ASTOpAssignNode struct {
	ASTAbstractAssignNode
	operator string
}

var _ IASTAbstractAssignNode = &ASTOpAssignNode{}

func NewASTOpAssignNode(lhs IASTExprNode, op string, rhs IASTExprNode) *ASTOpAssignNode {
	p := &ASTOpAssignNode{ASTAbstractAssignNode: ASTAbstractAssignNode{lhs: lhs, rhs: rhs}, operator: op}
	p.ASTAbstractAssignNode.ASTExprNode._impl = p
	p.ASTAbstractAssignNode.ASTExprNode.Node._impl = p
	return p
}

func (this *ASTOpAssignNode) Accept(visitor IASTVisitor) interface{} {
	return visitor.VisitOpAssignNode(this)
}
