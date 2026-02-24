package models

type ASTSuffixOpNode struct {
	*ASTUnaryArithmeticOpNode
}

func NewASTSuffixOpNode(op string, expr IASTExprNode) *ASTSuffixOpNode {
	return &ASTSuffixOpNode{
		ASTUnaryArithmeticOpNode: NewASTUnaryArithmeticOpNode(op, expr),
	}
}

func (this *ASTSuffixOpNode) Accept(visitor IASTVisitor) interface{} {
	return visitor.VisitSuffixOpNode(this)
}
