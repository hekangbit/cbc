package asm

type ILiteral interface {
	CompareTo(other ILiteral) int
	ToSource() string
	ToSourceWithTable(table SymbolTable) string
	Dump() string
	CollectStatistics(stats Statistics)
	IsZero() bool
	Plus(diff int64) ILiteral
	CmpIntegerLiteral(i *IntegerLiteral) int
	CmpNamedSymbol(sym *NamedSymbol) int
	CmpUnnamedSymbol(sym *UnnamedSymbol) int
	CmpSuffixedSymbol(sym *SuffixedSymbol) int
}
