package models

import (
	"cbc/asm"
)

type DefinedVariable struct {
	BaseVariable
	initializer IASTExprNode
	irobj       IRExpr
	sequence    int64
	symbol      asm.ISymbol
}

func NewDefinedVariable(priv bool, typeNode *TypeNode, name string, init IASTExprNode) *DefinedVariable {
	var p = new(DefinedVariable)
	p.isPrivate = priv
	p.name = name
	p.typeNode = typeNode
	p.initializer = init
	p.sequence = -1
	return p
}
