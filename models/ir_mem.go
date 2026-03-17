package models

import "cbc/asm"

type IRMem struct {
	IRExpr
	expr IIRExpr
}

var _ IIRExpr = &IRMem{}

func NewIRMem(t asm.Type, expr IIRExpr) *IRMem {
	p := &IRMem{
		IRExpr: IRExpr{typ: t},
		expr:   expr,
	}
	p._impl = p
	return p
}

func (this *IRMem) Expr() IIRExpr {
	return this.expr
}

func (this *IRMem) AddressNode(typ asm.Type) IIRExpr {
	return this.expr
}

func (this *IRMem) Accept(visitor IRVisitor) any {
	return visitor.VisitMem(this)
}

func (this *IRMem) _Dump(d *Dumper) {
	d.PrintMemberDumpable("expr", this.expr)
}
