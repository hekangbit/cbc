package asm

import "fmt"

type ISymbol interface {
	ILiteral
	fmt.Stringer
	Name() string
}
