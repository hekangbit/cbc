package models

import (
	"cbc/asm"
	"cbc/utils"
)

type IRStr struct {
	IRExpr
	entry *ConstantEntry
}

var _ IIRExpr = &IRStr{}

func NewIRStr(t asm.Type, entry *ConstantEntry) *IRStr {
	p := &IRStr{
		IRExpr: IRExpr{typ: t},
		entry:  entry,
	}
	p._impl = p
	return p
}

func (this *IRStr) Entry() *ConstantEntry {
	return this.entry
}

func (this *IRStr) Symbol() asm.ISymbol {
	return this.entry.Symbol()
}

func (this *IRStr) IsConstant() bool {
	return true
}

func (this *IRStr) Memref() asm.IMemoryReference {
	return this.entry.Memref()
}

func (this *IRStr) Address() asm.IOperand {
	return this.entry.Address()
}

func (this *IRStr) AsmValue() *asm.ImmediateValue {
	return this.entry.Address()
}

func (this *IRStr) Accept(visitor IRVisitor) interface{} {
	return visitor.VisitStr(this)
}

// TODO: java call toString, means call Object class method
func (this *IRStr) _Dump(d *Dumper) {
	d.PrintMemberStringNotResolved("entry", utils.ToString(this.entry)) // TODO
}
