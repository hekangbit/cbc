package models

import "cbc/asm"

type IRLabelStmt struct {
	IRStmt
	label *asm.Label
}

var _ IIRStmt = &IRLabelStmt{}

func NewIRLabelStmt(loc *Location, label *asm.Label) *IRLabelStmt {
	p := &IRLabelStmt{
		IRStmt: IRStmt{location: loc},
		label:  label,
	}
	p._impl = p

	return p
}

func (this *IRLabelStmt) Label() *asm.Label {
	return this.label
}

func (this *IRLabelStmt) Accept(visitor IRVisitor) any {
	return visitor.VisitLabelStmt(this)
}

func (this *IRLabelStmt) _Dump(d *Dumper) {
	d.PrintMemberLabel("label", this.label)
}
