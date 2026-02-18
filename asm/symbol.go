package asm

import "fmt"

type ISymbol interface {
	ILiteral
	fmt.Stringer
	Name() string
}

type Symbol struct {
}

var _ ISymbol = (*Symbol)(nil)

func (this *Symbol) Name() string {
	panic("")
}
func (this *Symbol) String() string {
	panic("")
}

func (this *Symbol) CompareTo(other ILiteral) int {
	panic("")
}

func (this *Symbol) Dump() string {
	panic("")
}

func (this *Symbol) ToSource() string {
	panic("")
}

func (this *Symbol) ToSourceWithTable(table SymbolTable) string {
	panic("")
}

func (this *Symbol) IsZero() bool {
	return false
}

func (this *Symbol) CollectStatistics(stats Statistics) {
	stats.SymbolUsed(this)
}

func (this *Symbol) Plus(n int64) ILiteral {
	panic("must not happen: BaseSymbol.plus called")
}

func (this *Symbol) CmpIntegerLiteral(i *IntegerLiteral) int {
	panic("abstractmethod: Symbol::CmpIntegerLiteral")
}
func (this *Symbol) CmpNamedSymbol(sym *NamedSymbol) int {
	panic("abstractmethod: Symbol::CmpNamedSymbol")
}

func (this *Symbol) CmpUnnamedSymbol(sym *UnnamedSymbol) int {
	panic("abstractmethod: Symbol::CmpUnnamedSymbol")
}

func (this *Symbol) CmpSuffixedSymbol(sym *SuffixedSymbol) int {
	panic("abstractmethod: Symbol::CmpSuffixedSymbol")
}
