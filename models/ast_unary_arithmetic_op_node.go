package models

type ASTUnaryArithmeticOpNode struct {
	ASTUnaryOpNode
	amount int64
}

var _ IASTExprNode = &ASTUnaryArithmeticOpNode{}

func NewASTUnaryArithmeticOpNode(op string, expr IASTExprNode) *ASTUnaryArithmeticOpNode {
	p := &ASTUnaryArithmeticOpNode{
		ASTUnaryOpNode: ASTUnaryOpNode{operator: op, expr: expr},
		amount:         1,
	}
	p.ASTUnaryOpNode.ASTExprNode._impl = p
	p.ASTUnaryOpNode.ASTExprNode.Node._impl = p
	return p
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
