package asm

import "fmt"

type UnnamedSymbol struct {
	BaseSymbol
}

var _ IBaseSymbol = (*UnnamedSymbol)(nil)

func NewUnnamedSymbol() *UnnamedSymbol {
	return &UnnamedSymbol{}
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
	return this.BaseSymbol.String()
}

func (this *UnnamedSymbol) CompareTo(lit ILiteral) int {
	return -(lit.CompareTo(this))
}

func (this *UnnamedSymbol) CmpIntegerLiteral(i *IntegerLiteral) int {
	return 1
}

func (this *UnnamedSymbol) CmpNamedSymbol(sym *NamedSymbol) int {
	return 1
}

func (this *UnnamedSymbol) CmpUnnamedSymbol(sym *UnnamedSymbol) int {
	// TODO: string compare logic
	str := this.String()
	symStr := sym.String()
	if str < symStr {
		return -1
	} else if str > symStr {
		return 1
	}
	return 0
}

func (this *UnnamedSymbol) CmpSuffixedSymbol(sym *SuffixedSymbol) int {
	return 1
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
