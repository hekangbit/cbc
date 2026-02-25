package models

type ASTReturnNode struct {
	ASTStmtNode
	expr IASTExprNode
}

var _ IASTStmtNode = &ASTReturnNode{}

func NewASTReturnNode(loc *Location, expr IASTExprNode) *ASTReturnNode {
	p := &ASTReturnNode{ASTStmtNode: ASTStmtNode{location: loc}, expr: expr}
	p._impl = p
	return p
}

func (this *ASTReturnNode) Expr() IASTExprNode {
	return this.expr
}

func (this *ASTReturnNode) SetExpr(expr IASTExprNode) {
	this.expr = expr
}

func (this *ASTReturnNode) _Dump(d *Dumper) {
	d.PrintMemberDumpable("expr", this.expr)
}

func (this *ASTReturnNode) Accept(visitor IASTVisitor) interface{} {
	return visitor.VisitReturnNode(this)
}
