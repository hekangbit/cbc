package asm

import "fmt"

type SymbolTable struct {
	base string
	m    map[*UnnamedSymbol]string
	seq  int64
}

const dummySymbolBase = "L"

var dummy = NewSymbolTable(dummySymbolBase)

func Dummy() *SymbolTable {
	return dummy
}

func NewSymbolTable(base string) *SymbolTable {
	return &SymbolTable{
		base: base,
		m:    make(map[*UnnamedSymbol]string),
		seq:  0,
	}
}

func (st *SymbolTable) newString() string {
	s := fmt.Sprintf("%s%d", st.base, st.seq)
	st.seq++
	return s
}

func (st *SymbolTable) NewSymbol() *NamedSymbol {
	return NewNamedSymbol(st.newString())
}

func (st *SymbolTable) SymbolString(sym *UnnamedSymbol) string {
	if str, ok := st.m[sym]; ok {
		return str
	}
	newStr := st.newString()
	st.m[sym] = newStr
	return newStr
}
