package models

import (
	"cbc/asm"
	"strconv"
)

var tmpSeq int = 0

type DefinedVariable struct {
	Variable
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

func NewTmpNewDefinedVariable(t IType) *DefinedVariable {
	v := NewDefinedVariable(false, NewTypeNodeFromType(t), "@tmp"+strconv.Itoa(tmpSeq), nil)
	tmpSeq++
	return v
}

func (this *DefinedVariable) Accept(visitor IEntityVisitor) interface{} {
	return visitor.VisitDefinedVariable(this)
}
