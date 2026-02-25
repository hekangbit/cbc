package models

type ASTSuffixOpNode struct {
	ASTUnaryArithmeticOpNode
}

var _ IASTExprNode = &ASTSuffixOpNode{}

func NewASTSuffixOpNode(op string, expr IASTExprNode) *ASTSuffixOpNode {
	p := &ASTSuffixOpNode{
		ASTUnaryArithmeticOpNode: ASTUnaryArithmeticOpNode{
			ASTUnaryOpNode: ASTUnaryOpNode{operator: op, expr: expr},
			amount:         1,
		},
	}
	p.ASTUnaryArithmeticOpNode.ASTUnaryOpNode.ASTExprNode._impl = p
	p.ASTUnaryArithmeticOpNode.ASTUnaryOpNode.ASTExprNode.Node._impl = p
	return p
}

func (this *ASTSuffixOpNode) Accept(visitor IASTVisitor) interface{} {
	return visitor.VisitSuffixOpNode(this)
}
