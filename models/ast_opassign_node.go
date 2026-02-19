package models

type ASTOpAssignNode struct {
	ASTAbstractAssignNode
	operator string
}

func NewASTOpAssignNode(lhs IASTExprNode, op string, rhs IASTExprNode) *ASTOpAssignNode {
	return &ASTOpAssignNode{
		ASTAbstractAssignNode: ASTAbstractAssignNode{
			lhs: lhs,
			rhs: rhs,
		},
		operator: op,
	}
}

func (this *ASTOpAssignNode) Accept(visitor IASTVisitor) interface{} {
	return visitor.VisitOpAssignNode(this)
}
