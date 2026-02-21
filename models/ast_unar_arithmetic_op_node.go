package models

type ASTUnaryArithmeticOpNode struct {
	*ASTUnaryOpNode
	amount int64
}

func NewASTUnaryArithmeticOpNode(op string, expr IASTExprNode) *ASTUnaryArithmeticOpNode {
	return &ASTUnaryArithmeticOpNode{
		ASTUnaryOpNode: NewASTUnaryOpNode(op, expr),
		amount:         1,
	}
}

func (this *ASTUnaryArithmeticOpNode) SetExpr(expr IASTExprNode) {
	this.expr = expr
}

func (this *ASTUnaryArithmeticOpNode) Amount() int64 {
	return this.amount
}

func (this *ASTUnaryArithmeticOpNode) SetAmount(amount int64) {
	this.amount = amount
}
