package asm

type ILiteral interface {
	CompareTo(ILiteral) int
	ToSource() string
	ToSourceWithTable(SymbolTable) string
	Dump() string
	CollectStatistics(Statistics)
	IsZero() bool
	Plus(int64) ILiteral
	Cmp(ILiteral) int
}
