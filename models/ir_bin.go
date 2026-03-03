package models

import (
	"cbc/asm"
	"strconv"
)

type IRBin struct {
	IRExpr
	op    Op
	left  IIRExpr
	right IIRExpr
}

var _ IIRExpr = &IRBin{}

func NewIRBin(t asm.Type, op Op, left IIRExpr, right IIRExpr) *IRBin {
	p := &IRBin{
		IRExpr: IRExpr{typ: t},
		op:     op,
		left:   left,
		right:  right,
	}
	p._impl = p
	return p
}

func (this *IRBin) Accept(visitor IRVisitor) any {
	return visitor.VisitBin(this)
}

func (this *IRBin) _Dump(d *Dumper) {
	d.PrintMemberStringNotResolved("op", strconv.Itoa(int(this.op)))
	d.PrintMemberDumpable("left", this.left)
	d.PrintMemberDumpable("right", this.right)
}
