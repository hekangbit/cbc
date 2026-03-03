package models

import "cbc/asm"

type IRCase struct {
	Dumpable
	value int64
	label *asm.Label
}

func NewIRCase(value int64, label *asm.Label) *IRCase {
	p := &IRCase{
		value: value,
		label: label,
	}
	return p
}

func (this *IRCase) Dump(d *Dumper) {
	d.PrintClassNoLoc(this)
	d.PrintMemberInt64("value", this.value)
	d.PrintMemberLabel("label", this.label)
}
