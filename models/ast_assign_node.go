package models

type ASTAssignNode struct {
	ASTAbstractAssignNode
}

func NewASTAssignNode(lhs IASTExprNode, rhs IASTExprNode) *ASTAssignNode {
	return &ASTAssignNode{
		ASTAbstractAssignNode: ASTAbstractAssignNode{
			lhs: lhs,
			rhs: rhs,
		},
	}
}

func (this *ASTAssignNode) Accept(visitor IASTVisitor) interface{} {
	return visitor.VisitAssignNode(this)
}
