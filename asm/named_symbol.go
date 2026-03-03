package asm

import (
	"cbc/utils"
	"fmt"
)

type NamedSymbol struct {
	BaseSymbol
	name string
}

var _ IBaseSymbol = &NamedSymbol{}

func NewNamedSymbol(name string) *NamedSymbol {
	p := &NamedSymbol{name: name}
	p._impl = p
	return p
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

func (this *NamedSymbol) Cmp(other ILiteral) int {
	switch o := other.(type) {
	case *IntegerLiteral:
		return 1
	case *NamedSymbol:
		return utils.CompareStrings(this.name, o.name)
	case *UnnamedSymbol:
		return -1
	case *SuffixedSymbol:
		return utils.CompareStrings(this.String(), o.String())
	default:
		panic(fmt.Sprintf("unsupported comparison with %T", other))
	}
}

func (this *NamedSymbol) Dump() string {
	return fmt.Sprintf("(NamedSymbol %s)", utils.DumpString(this.name))
}
