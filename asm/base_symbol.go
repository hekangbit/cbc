package asm

type IBaseSymbol interface {
	ISymbol
}

type BaseSymbol struct {
	Symbol
}

var _ IBaseSymbol = (*BaseSymbol)(nil)

func (this *BaseSymbol) Name() string {
	panic("")
}
func (this *BaseSymbol) String() string {
	panic("")
}

func (this *BaseSymbol) CompareTo(other ILiteral) int {
	panic("")
}

func (this *BaseSymbol) Dump() string {
	panic("")
}

func (this *BaseSymbol) ToSource() string {
	panic("")
}

func (this *BaseSymbol) ToSourceWithTable(table SymbolTable) string {
	panic("")
}

func (this *BaseSymbol) IsZero() bool {
	return false
}

func (this *BaseSymbol) CollectStatistics(stats Statistics) {
	stats.SymbolUsed(this)
}

func (this *BaseSymbol) Plus(n int64) ILiteral {
	panic("must not happen: BaseSymbol.plus called")
}
