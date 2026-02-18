package models

import "cbc/asm"

type ConstantEntry struct {
	value   string
	symbol  asm.ISymbol
	memref  *asm.MemoryReference
	address *asm.ImmediateValue
}

func NewConstantEntry(val string) *ConstantEntry {
	return &ConstantEntry{value: val}
}

func (this *ConstantEntry) Value() string {
	return this.value
}

func (this *ConstantEntry) SetSymbol(sym asm.ISymbol) {
	this.symbol = sym
}

func (this *ConstantEntry) Symbol() asm.ISymbol {
	if this.symbol == nil {
		panic("must not happen: symbol == nil")
	}
	return this.symbol
}

func (this *ConstantEntry) SetMemref(mem *asm.MemoryReference) {
	this.memref = mem
}

func (this *ConstantEntry) Memref() *asm.MemoryReference {
	if this.memref == nil {
		panic("must not happen: memref == nil")
	}
	return this.memref
}

func (this *ConstantEntry) SetAddress(imm *asm.ImmediateValue) {
	this.address = imm
}

func (this *ConstantEntry) Address() *asm.ImmediateValue {
	return this.address
}
