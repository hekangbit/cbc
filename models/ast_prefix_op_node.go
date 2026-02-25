package models

type ASTPrefixOpNode struct {
	ASTUnaryArithmeticOpNode
}

var _ IASTExprNode = &ASTPrefixOpNode{}

func NewASTPrefixOpNode(op string, expr IASTExprNode) *ASTPrefixOpNode {
	p := &ASTPrefixOpNode{
		ASTUnaryArithmeticOpNode: ASTUnaryArithmeticOpNode{
			ASTUnaryOpNode: ASTUnaryOpNode{operator: op, expr: expr},
			amount:         1,
		},
	}
	p.ASTUnaryArithmeticOpNode.ASTUnaryOpNode.ASTExprNode._impl = p
	p.ASTUnaryArithmeticOpNode.ASTUnaryOpNode.ASTExprNode.Node._impl = p
	return p
}

func (this *ASTPrefixOpNode) Accept(visitor IASTVisitor) interface{} {
	return visitor.VisitPrefixOpNode(this)
}
