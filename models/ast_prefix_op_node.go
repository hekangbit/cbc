package models

type ASTPrefixOpNode struct {
	*ASTUnaryArithmeticOpNode
}

func NewASTPrefixOpNode(op string, expr IASTExprNode) *ASTPrefixOpNode {
	return &ASTPrefixOpNode{
		ASTUnaryArithmeticOpNode: NewASTUnaryArithmeticOpNode(op, expr),
	}
}

func (this *ASTPrefixOpNode) Accept(visitor IASTVisitor) interface{} {
	return visitor.VisitPrefixOpNode(this)
}
