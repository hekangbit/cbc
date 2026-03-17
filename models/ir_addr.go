package models

import "cbc/asm"

type IRAddr struct {
	IRExpr
	entity IEntity
}

var _ IIRExpr = &IRAddr{}

func NewIRAddr(typ asm.Type, entity IEntity) *IRAddr {
	p := &IRAddr{
		IRExpr: IRExpr{typ: typ},
		entity: entity,
	}
	p._impl = p
	return p
}

func (this *IRAddr) IsAddr() bool {
	return true
}

func (this *IRAddr) Entity() IEntity {
	return this.entity
}

func (this *IRAddr) Address() asm.IOperand {
	return this.entity.Address()
}

func (this *IRAddr) Memref() asm.IMemoryReference {
	return this.entity.Memref()
}

func (this *IRAddr) GetEntityForce() IEntity {
	return this.entity
}

func (this *IRAddr) Accept(visitor IRVisitor) any {
	return visitor.VisitAddr(this)
}

func (this *IRAddr) _Dump(d *Dumper) {
	d.PrintMemberStringNotResolved("entity", this.entity.Name())
}
