package asm

type IBaseSymbol interface {
	ISymbol
}

type BaseSymbol struct {
	_impl ISymbol
}

func (this *BaseSymbol) IsZero() bool {
	return false
}

func (this *BaseSymbol) CollectStatistics(stats Statistics) {
	stats.SymbolUsed(this._impl)
}

func (this *BaseSymbol) Plus(n int64) ILiteral {
	panic("must not happen: BaseSymbol.plus called")
}
