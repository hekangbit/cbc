package models

type IRExprStmt struct {
	IRStmt
	expr IIRExpr
}

var _ IIRStmt = &IRExprStmt{}

func NewIRExprStmt(loc *Location, expr IIRExpr) *IRExprStmt {
	p := &IRExprStmt{
		IRStmt: IRStmt{location: loc},
		expr:   expr,
	}
	p._impl = p
	return p
}

func (this *IRExprStmt) Expr() IIRExpr {
	return this.expr
}

func (this *IRExprStmt) Accept(visitor IRVisitor) interface{} {
	return visitor.VisitExprStmt(this)
}

func (this *IRExprStmt) _Dump(d *Dumper) {
	d.PrintMemberDumpable("expr", this.expr)
}
