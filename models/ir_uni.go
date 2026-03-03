package models

import (
	"cbc/asm"
	"strconv"
)

type IRUni struct {
	IRExpr
	op   Op
	expr IIRExpr
}

func NewIRUni(typ asm.Type, op Op, expr IIRExpr) *IRUni {
	p := &IRUni{
		IRExpr: IRExpr{typ: typ},
		op:     op,
		expr:   expr,
	}
	p._impl = p
	return p
}

func (this *IRUni) Op() Op {
	return this.op
}

func (this *IRUni) Expr() IIRExpr {
	return this.expr
}

func (this *IRUni) Accept(visitor IRVisitor) interface{} {
	return visitor.VisitUni(this)
}

func (this *IRUni) _Dump(d *Dumper) {
	d.PrintMemberStringNotResolved("op", strconv.Itoa(int(this.op)))
	d.PrintMemberDumpable("expr", this.expr)
}
