package models

type IRAssign struct {
	IRStmt
	lhs IIRExpr
	rhs IIRExpr
}

var _ IIRStmt = &IRAssign{}

func NewIRAssign(loc *Location, lhs IIRExpr, rhs IIRExpr) *IRAssign {
	p := &IRAssign{
		IRStmt: IRStmt{location: loc},
		lhs:    lhs,
		rhs:    rhs,
	}
	p._impl = p
	return p
}

func (this *IRAssign) Lhs() IIRExpr {
	return this.lhs
}

func (this *IRAssign) Rhs() IIRExpr {
	return this.rhs
}

func (this *IRAssign) Accept(visitor IRVisitor) any {
	return visitor.VisitAssign(this)
}

func (this *IRAssign) _Dump(d *Dumper) {
	d.PrintMemberDumpable("lhs", this.lhs)
	d.PrintMemberDumpable("rhs", this.rhs)
}
