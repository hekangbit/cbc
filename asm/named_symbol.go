package asm

import (
	"cbc/util"
	"fmt"
)

type NamedSymbol struct {
	BaseSymbol
	name string
}

func NewNamedSymbol(name string) *NamedSymbol {
	return &NamedSymbol{name: name}
}

func (this *NamedSymbol) Name() string {
	return this.name
}

func (this *NamedSymbol) ToSource() string {
	return this.name
}

func (this *NamedSymbol) ToSourceWithTable(table SymbolTable) string {
	return this.name
}

func (this *NamedSymbol) String() string {
	return "#" + this.name
}

func (this *NamedSymbol) CompareTo(lit ILiteral) int {
	return -lit.CompareTo(this)
}

func (this *NamedSymbol) CmpIntegerLiteral(i *IntegerLiteral) int {
	return 1
}

func (this *NamedSymbol) CmpNamedSymbol(sym *NamedSymbol) int {
	return util.CompareStrings(this.name, sym.name)
}

func (this *NamedSymbol) CmpUnnamedSymbol(sym *UnnamedSymbol) int {
	return -1
}

func (this *NamedSymbol) CmpSuffixedSymbol(sym *SuffixedSymbol) int {
	return util.CompareStrings(this.String(), sym.String())
}

func (this *NamedSymbol) Dump() string {
	return fmt.Sprintf("(NamedSymbol %s)", util.DumpString(this.name))
}
