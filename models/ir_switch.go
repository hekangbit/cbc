package models

import "cbc/asm"

type IRSwitch struct {
	IRStmt
	cond         IIRExpr
	cases        []*IRCase
	defaultLabel *asm.Label
	endLabel     *asm.Label
}

var _ IIRStmt = &IRSwitch{}

func NewIRSwitch(loc *Location, cond IIRExpr, cases []*IRCase, defaultLabel, endLabel *asm.Label) *IRSwitch {
	p := &IRSwitch{
		IRStmt:       IRStmt{location: loc},
		cond:         cond,
		cases:        cases,
		defaultLabel: defaultLabel,
		endLabel:     endLabel,
	}
	p._impl = p
	return p
}

func (this *IRSwitch) Cond() IIRExpr {
	return this.cond
}

func (this *IRSwitch) Cases() []*IRCase {
	return this.cases
}

func (this *IRSwitch) DefaultLabel() *asm.Label {
	return this.defaultLabel
}

func (this *IRSwitch) EndLabel() *asm.Label {
	return this.endLabel
}

func (this *IRSwitch) Accept(visitor IRVisitor) any {
	return visitor.VisitSwitch(this)
}

func (this *IRSwitch) _Dump(d *Dumper) {
	d.PrintMemberDumpable("cond", this.cond)
	buf := make([]Dumpable, len(this.cases))
	for i, tmp := range this.cases {
		buf[i] = tmp
	}
	d.PrintNodeList("cases", buf)
	d.PrintMemberLabel("defaultLabel", this.defaultLabel)
	d.PrintMemberLabel("endLabel", this.endLabel)
}
