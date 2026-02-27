package asm

import (
	"cbc/utils"
	"fmt"
)

type SuffixedSymbol struct {
	Symbol
	base   ISymbol
	suffix string
}

var _ ISymbol = (*SuffixedSymbol)(nil)

func NewSuffixedSymbol(base ISymbol, suffix string) *SuffixedSymbol {
	return &SuffixedSymbol{
		base:   base,
		suffix: suffix,
	}
}

func (s *SuffixedSymbol) IsZero() bool {
	return false
}

func (s *SuffixedSymbol) CollectStatistics(stats Statistics) {
	s.base.CollectStatistics(stats)
}

func (s *SuffixedSymbol) Plus(diff int64) ILiteral {
	panic("must not happen: SuffixedSymbol.Plus called")
}

func (s *SuffixedSymbol) Name() string {
	return s.base.Name()
}

func (s *SuffixedSymbol) ToSource() string {
	return s.base.ToSource() + s.suffix
}

func (s *SuffixedSymbol) ToSourceWithTable(table SymbolTable) string {
	return s.base.ToSourceWithTable(table) + s.suffix
}

func (s *SuffixedSymbol) String() string {
	return s.base.String() + s.suffix
}

func (s *SuffixedSymbol) CompareTo(lit ILiteral) int {
	return -lit.CompareTo(s)
}

func (s *SuffixedSymbol) CmpIntegerLiteral(i *IntegerLiteral) int {
	return 1
}

func (s *SuffixedSymbol) CmpNamedSymbol(sym *NamedSymbol) int {
	return utils.CompareStrings(s.String(), sym.String())
}

func (s *SuffixedSymbol) CmpUnnamedSymbol(sym *UnnamedSymbol) int {
	return -1
}

func (s *SuffixedSymbol) CmpSuffixedSymbol(sym *SuffixedSymbol) int {
	return utils.CompareStrings(s.String(), sym.String())
}

func (s *SuffixedSymbol) Dump() string {
	return fmt.Sprintf("(SuffixedSymbol %s %s)", s.base.Dump(), utils.DumpString(s.suffix))
}
