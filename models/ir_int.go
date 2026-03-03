package models

import "cbc/asm"

type IRInt struct {
	IRExpr
	value int64
}

var _ IIRExpr = &IRInt{}

func NewIRInt(t asm.Type, value int64) *IRInt {
	p := &IRInt{
		IRExpr: IRExpr{typ: t},
		value:  value,
	}
	p._impl = p
	return p
}

func (this *IRInt) Value() int64 {
	return this.value
}

func (this *IRInt) IsConstant() bool {
	return true
}

func (this *IRInt) AsmValue() *asm.ImmediateValue {
	return asm.NewImmediateValue(asm.NewIntegerLiteral(this.value))
}

// TODO: java throw new Error("must not happen: IntValue#memref");
func (this *IRInt) Memref() asm.IMemoryReference {
	panic("must not happen: IntValue#memref")
}

func (this *IRInt) Accept(visitor IRVisitor) interface{} {
	return visitor.VisitInt(this)
}

func (this *IRInt) _Dump(d *Dumper) {
	d.PrintMemberInt64("value", this.value)
}
