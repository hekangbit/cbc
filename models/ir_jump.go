package models

import "cbc/asm"

type IRJump struct {
	IRStmt
	label *asm.Label
}

var _ IIRStmt = &IRCJump{}

func NewIRJump(loc *Location, label *asm.Label) *IRJump {
	p := &IRJump{
		IRStmt: IRStmt{location: loc},
		label:  label,
	}
	p._impl = p
	return p
}

func (this *IRJump) Label() *asm.Label {
	return this.label
}

func (this *IRJump) Accept(visitor IRVisitor) any {
	return visitor.VisitJump(this)
}

func (this *IRJump) _Dump(d *Dumper) {
	d.PrintMemberLabel("label", this.label)
}
