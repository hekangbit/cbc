package models

type IRReturn struct {
	IRStmt
	expr IIRExpr
}

func NewIRReturn(loc *Location, expr IIRExpr) *IRReturn {
	p := &IRReturn{
		IRStmt: IRStmt{location: loc},
		expr:   expr,
	}
	p._impl = p
	return p
}

func (this *IRReturn) Expr() IIRExpr {
	return this.expr
}

func (this *IRReturn) Accept(visitor IRVisitor) any {
	return visitor.VisitReturn(this)
}

func (this *IRReturn) _Dump(d *Dumper) {
	d.PrintMemberDumpable("expr", this.expr)
}
