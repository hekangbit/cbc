package asm

import (
	"fmt"
)

type UnnamedSymbol struct {
	BaseSymbol
}

var _ IBaseSymbol = &UnnamedSymbol{}

func NewUnnamedSymbol() *UnnamedSymbol {
	p := &UnnamedSymbol{}
	p._impl = p
	return p
}

func (this *UnnamedSymbol) Name() string {
	panic("unnamed symbol")
}

func (this *UnnamedSymbol) ToSource() string {
	panic("UnnamedSymbol#ToSource() called")
}

func (this *UnnamedSymbol) ToSourceWithTable(table SymbolTable) string {
	return table.SymbolString(this)
}

func (this *UnnamedSymbol) String() string {
	return this._impl.String()
}

func (this *UnnamedSymbol) CompareTo(lit ILiteral) int {
	return -(lit.CompareTo(this))
}

func (this *UnnamedSymbol) Cmp(other ILiteral) int {
	switch o := other.(type) {
	case *IntegerLiteral:
		return 1
	case *NamedSymbol:
		return 1
	case *UnnamedSymbol:
		// TODO: string compare logic
		str := this.String()
		symStr := o.String()
		if str < symStr {
			return -1
		} else if str > symStr {
			return 1
		}
		return 0
	case *SuffixedSymbol:
		return 1
	default:
		panic(fmt.Sprintf("unsupported comparison with %T", other))
	}
}

func (this *UnnamedSymbol) Dump() string {
	// TODO: hashCode, and string format
	addr := fmt.Sprintf("%p", this)
	hexPart := addr
	if len(addr) > 2 && addr[:2] == "0x" {
		hexPart = addr[2:]
	}
	return fmt.Sprintf("(UnnamedSymbol @%s)", hexPart)
}
