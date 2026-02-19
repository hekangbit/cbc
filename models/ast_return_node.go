package models

type ASTReturnNode struct {
	*ASTStmtNode
	expr IASTExprNode
}

var _ IASTStmtNode = (*ASTReturnNode)(nil)

func NewASTReturnNode(loc *Location, expr IASTExprNode) *ASTReturnNode {
	return &ASTReturnNode{
		ASTStmtNode: NewASTStmtNode(loc),
		expr:        expr,
	}
}

func (this *ASTReturnNode) Expr() IASTExprNode {
	return this.expr
}

func (this *ASTReturnNode) SetExpr(expr IASTExprNode) {
	this.expr = expr
}

func (this *ASTReturnNode) Dump(d *Dumper) {
	d.PrintMemberDumpable("expr", this.expr)
}

func (this *ASTReturnNode) Accept(visitor IASTVisitor) interface{} {
	return visitor.VisitReturnNode(this)
}
