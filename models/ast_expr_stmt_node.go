package models

type ASTExprStmtNode struct {
	ASTStmtNode
	expr IASTExprNode
}

var _ IASTStmtNode = &ASTExprStmtNode{}

func NewASTExprStmtNode(loc *Location, expr IASTExprNode) *ASTExprStmtNode {
	p := &ASTExprStmtNode{ASTStmtNode: ASTStmtNode{location: loc}, expr: expr}
	p._impl = p
	return p
}

func (this *ASTExprStmtNode) Expr() IASTExprNode {
	return this.expr
}

func (this *ASTExprStmtNode) SetExpr(expr IASTExprNode) {
	this.expr = expr
}

func (this *ASTExprStmtNode) _Dump(d *Dumper) {
	d.PrintMemberDumpable("expr", this.expr)
}

func (this *ASTExprStmtNode) Accept(visitor IASTVisitor) interface{} {
	return visitor.VisitExprStmtNode(this)
}
