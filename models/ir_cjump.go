package models

import "cbc/asm"

type IRCJump struct {
	IRStmt
	cond      IIRExpr
	thenLabel *asm.Label
	elseLabel *asm.Label
}

var _ IIRStmt = &IRCJump{}

func NewIRCJump(loc *Location, cond IIRExpr, thenLabel *asm.Label, elseLabel *asm.Label) *IRCJump {
	p := &IRCJump{
		IRStmt:    IRStmt{location: loc},
		cond:      cond,
		thenLabel: thenLabel,
		elseLabel: elseLabel,
	}
	p._impl = p
	return p
}

func (this *IRCJump) Cond() IIRExpr {
	return this.cond
}

func (this *IRCJump) ThenLabel() *asm.Label {
	return this.thenLabel
}

func (this *IRCJump) ElseLabel() *asm.Label {
	return this.elseLabel
}

func (this *IRCJump) Accept(visitor IRVisitor) any {
	return visitor.VisitCJump(this)
}

func (this *IRCJump) _Dump(d *Dumper) {
	d.PrintMemberDumpable("cond", this.cond)
	d.PrintMemberLabel("thenLabel", this.thenLabel)
	d.PrintMemberLabel("elseLabel", this.elseLabel)
}
