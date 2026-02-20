package models

type ASTCondExprNode struct {
	ASTExprNode
	cond     IASTExprNode
	thenExpr IASTExprNode
	elseExpr IASTExprNode
}

func NewASTCondExprNode(c IASTExprNode, t IASTExprNode, e IASTExprNode) *ASTCondExprNode {
	return &ASTCondExprNode{cond: c, thenExpr: t, elseExpr: e}
}

func (this *ASTCondExprNode) Type() IType {
	return this.thenExpr.Type()
}

func (this *ASTCondExprNode) Cond() IASTExprNode {
	return this.cond
}

func (this *ASTCondExprNode) ThenExpr() IASTExprNode {
	return this.thenExpr
}

func (this *ASTCondExprNode) SetThenExpr(expr IASTExprNode) {
	this.thenExpr = expr
}

func (this *ASTCondExprNode) ElseExpr() IASTExprNode {
	return this.elseExpr
}

func (this *ASTCondExprNode) SetElseExpr(expr IASTExprNode) {
	this.elseExpr = expr
}

func (this *ASTCondExprNode) Location() *Location {
	return this.cond.Location()
}

func (this *ASTCondExprNode) Dump(d *Dumper) {
	d.PrintMemberDumpable("cond", this.cond)
	d.PrintMemberDumpable("thenExpr", this.thenExpr)
	d.PrintMemberDumpable("elseExpr", this.elseExpr)
}

func (this *ASTCondExprNode) Accept(visitor IASTVisitor) interface{} {
	return visitor.VisitCondExprNode(this)
}
