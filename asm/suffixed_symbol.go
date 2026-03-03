package asm

import (
	"cbc/utils"
	"fmt"
)

type SuffixedSymbol struct {
	base   ISymbol
	suffix string
}

var _ ISymbol = &SuffixedSymbol{}

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

func (this *SuffixedSymbol) Cmp(other ILiteral) int {
	switch o := other.(type) {
	case *IntegerLiteral:
		return 1
	case *NamedSymbol:
		return utils.CompareStrings(this.String(), o.String())
	case *UnnamedSymbol:
		return -1
	case *SuffixedSymbol:
		return utils.CompareStrings(this.String(), o.String())
	default:
		panic(fmt.Sprintf("unsupported comparison with %T", other))
	}
}

func (s *SuffixedSymbol) Dump() string {
	return fmt.Sprintf("(SuffixedSymbol %s %s)", s.base.Dump(), utils.DumpString(s.suffix))
}
