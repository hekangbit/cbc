package asm

import (
	"fmt"
	"strconv"
)

type IntegerLiteral struct {
	value int64
}

var _ ILiteral = &IntegerLiteral{}

func NewIntegerLiteral(n int64) *IntegerLiteral {
	return &IntegerLiteral{value: n}
}

func (this *IntegerLiteral) Value() int64 {
	return this.value
}

func (this *IntegerLiteral) IsZero() bool {
	return this.value == 0
}

func (this *IntegerLiteral) Plus(diff int64) ILiteral {
	return NewIntegerLiteral(this.value + diff)
}

func (this *IntegerLiteral) ToSource() string {
	return strconv.FormatInt(this.value, 10)
}

func (this *IntegerLiteral) ToSourceWithTable(table SymbolTable) string {
	return this.ToSource()
}

func (this *IntegerLiteral) Dump() string {
	return "(IntegerLiteral " + this.ToSource() + ")"
}

func (this *IntegerLiteral) CollectStatistics(stats Statistics) {
	// nothing
}

func (this *IntegerLiteral) CompareTo(lit ILiteral) int {
	return -(lit.Cmp(this))
}

func (this *IntegerLiteral) Cmp(other ILiteral) int {
	switch o := other.(type) {
	case *IntegerLiteral:
		if this.value < o.value {
			return -1
		} else if this.value > o.value {
			return 1
		}
		return 0
	case *NamedSymbol, *UnnamedSymbol, *SuffixedSymbol:
		return -1
	default:
		panic(fmt.Sprintf("unsupported comparison with %T", other))
	}
}

func (this *IntegerLiteral) Equals(other ILiteral) bool {
	if o, ok := other.(*IntegerLiteral); ok {
		return this.value == o.value
	}
	return false
}

func (this *IntegerLiteral) String() string {
	return this.ToSource()
}
