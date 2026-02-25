package models

import (
	"cbc/asm"
)

type IEntity interface {
	Dumpable
	Name() string
	SymbolString() string
	IsDefined() bool
	IsInitialized() bool
	IsConstant() bool
	Value() IASTExprNode
	IsParameter() bool
	IsPrivate() bool
	TypeNode() *ASTTypeNode
	Type() IType
	AllocSize() int64
	Alignment() int64
	Refered()
	IsRefered() bool
	SetMemref(asm.IMemoryReference)
	Memref() asm.IMemoryReference
	SetAddressMem(asm.IMemoryReference)
	SetAddressImm(asm.IImmediateValue)
	Address() asm.IOperand
	CheckAddress()
	Location() *Location
	Accept(IEntityVisitor) interface{}
	_Dump(*Dumper)
}

type Entity struct {
	_impl     IEntity
	name      string
	isPrivate bool
	typeNode  *ASTTypeNode
	nRefered  int64
	memref    asm.IMemoryReference
	address   asm.IOperand
}

func (e *Entity) Name() string {
	return e.name
}

func (e *Entity) SymbolString() string {
	return e.Name()
}

func (e *Entity) IsConstant() bool {
	return false
}

func (e *Entity) Value() IASTExprNode {
	panic("Entity#value")
}

func (e *Entity) IsParameter() bool {
	return false
}

func (e *Entity) IsPrivate() bool {
	return e.isPrivate
}

func (e *Entity) TypeNode() *ASTTypeNode {
	return e.typeNode
}

func (e *Entity) Type() IType {
	return e.typeNode.Type()
}

func (e *Entity) AllocSize() int64 {
	return e.Type().AllocSize()
}

func (e *Entity) Alignment() int64 {
	return e.Type().Alignment()
}

func (e *Entity) Refered() {
	e.nRefered++
}

func (e *Entity) IsRefered() bool {
	return e.nRefered > 0
}

func (e *Entity) SetMemref(mem asm.IMemoryReference) {
	e.memref = mem
}

func (e *Entity) Memref() asm.IMemoryReference {
	e.CheckAddress()
	return e.memref
}

func (e *Entity) SetAddressMem(mem asm.IMemoryReference) {
	e.address = mem
}

func (e *Entity) SetAddressImm(imm asm.IImmediateValue) {
	e.address = imm
}

func (e *Entity) Address() asm.IOperand {
	e.CheckAddress()
	return e.address
}

func (e *Entity) CheckAddress() {
	if e.memref == nil && e.address == nil {
		panic("address did not resolved: " + e.name)
	}
}

func (e *Entity) Location() *Location {
	return e.typeNode.Location()
}

func (e *Entity) Dump(d *Dumper) {
	d.PrintClass(e._impl, e.Location())
	e._impl._Dump(d)
}
