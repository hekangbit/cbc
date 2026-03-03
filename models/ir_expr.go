package models

import (
	"cbc/asm"
	"fmt"
)

type IIRExpr interface {
	Dumpable
	Type() asm.Type
	IsVar() bool
	IsAddr() bool
	IsConstant() bool
	AsmValue() *asm.ImmediateValue
	Address() asm.IOperand
	Memref() asm.IMemoryReference
	AddressNode(asm.Type) IIRExpr
	GetEntityForce() IEntity
	Accept(IRVisitor) any
	_Dump(*Dumper)
}

type IRExpr struct {
	_impl IIRExpr
	typ   asm.Type
}

func (this *IRExpr) Type() asm.Type { return this.typ }

func (this *IRExpr) IsVar() bool      { return false }
func (this *IRExpr) IsAddr() bool     { return false }
func (this *IRExpr) IsConstant() bool { return false }

func (this *IRExpr) AsmValue() *asm.ImmediateValue {
	panic("Expr#AsmValue called")
}

func (this *IRExpr) Address() asm.IOperand {
	panic("Expr#Address called")
}

// TODO: may need remove abstract method
func (this *IRExpr) Memref() asm.IMemoryReference {
	panic("Expr#Memref called")
}

// TODO: java throw new Error("unexpected node for LHS: " + getClass());
func (this *IRExpr) AddressNode(typ asm.Type) IIRExpr {
	panic(fmt.Sprintf("unexpected node for LHS: %T", this))
}

func (this *IRExpr) GetEntityForce() IEntity {
	return nil
}

// TODO: print asm.Type as int?
func (this *IRExpr) Dump(d *Dumper) {
	d.PrintClassNoLoc(this._impl)
	d.PrintMemberInt("type", int(this.typ))
	this._impl._Dump(d)
}
