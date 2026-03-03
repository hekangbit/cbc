package models

import "cbc/asm"

type IRVar struct {
	IRExpr
	ent IEntity
}

var _ IIRExpr = &IRVar{}

func NewIRVar(typ asm.Type, ent IEntity) *IRVar {
	p := &IRVar{
		IRExpr: IRExpr{typ: typ},
		ent:    ent,
	}
	p._impl = p
	return p
}

func (this *IRVar) IsVar() bool { return true }

// TODO: java enum is class, but here in golang it is int
func (this *IRVar) Type() asm.Type {
	// if this.IRExpr.Type() == nil {
	// 	panic("IRVar is too big to load by 1 insn")
	// }
	return this.IRExpr.Type()
}

func (this *IRVar) Name() string { return this.ent.Name() }

func (this *IRVar) Entity() IEntity { return this.ent }

func (this *IRVar) Address() asm.IOperand { return this.ent.Address() }

func (this *IRVar) Memref() asm.IMemoryReference { return this.ent.Memref() }

func (this *IRVar) AddressNode(typ asm.Type) IIRExpr {
	return NewIRAddr(typ, this.ent)
}

func (this *IRVar) GetEntityForce() IEntity {
	return this.ent
}

func (this *IRVar) Accept(visitor IRVisitor) interface{} {
	return visitor.VisitVar(this)
}

func (this *IRVar) _Dump(d *Dumper) {
	d.PrintMemberStringNotResolved("entity", this.ent.Name())
}
