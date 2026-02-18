package models

type ASTExprStmtNode struct {
	*ASTStmtNode
	expr IASTExprNode
}

func NewASTExprStmtNode(loc *Location, expr IASTExprNode) *ASTExprStmtNode {
	return &ASTExprStmtNode{
		ASTStmtNode: NewASTStmtNode(loc),
		expr:        expr,
	}

}

func (this *ASTExprStmtNode) Expr() IASTExprNode {
	return this.expr
}

func (this *ASTExprStmtNode) SetExpr(expr IASTExprNode) {
	this.expr = expr
}

func (this *ASTExprStmtNode) Dump(d *Dumper) {
	d.PrintMemberDumpable("expr", this.expr)
}

func (this *ASTExprStmtNode) Accept(visitor ASTVisitor) interface{} {
	return visitor.VisitExprStmtNode(this)
}
