package asm

type Label struct {
	Assembly
	symbol ISymbol
}

var _ IAssembly = (*Label)(nil)

func NewLabel(sym ISymbol) *Label {
	return &Label{symbol: sym}
}

func NewLabelUnnamed() *Label {
	return &Label{symbol: NewUnnamedSymbol()}
}

func (this *Label) Symbol() ISymbol {
	return this.symbol
}

func (this *Label) IsLabel() bool {
	return true
}

func (this *Label) ToSource(table SymbolTable) string {
	return this.symbol.ToSourceWithTable(table) + ":"
}

func (this *Label) Dump() string {
	return "(Label " + this.symbol.Dump() + ")"
}
