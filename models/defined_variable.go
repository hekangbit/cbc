package models

import (
	"cbc/asm"
	"strconv"
)

// TODO: cast tmpSeq to int for itoa, any better solution
var tmpSeq int64 = 0

type DefinedVariable struct {
	Variable
	initializer IASTExprNode
	ir          IIRExpr
	sequence    int64
	symbol      asm.ISymbol
}

var _ IVariable = &DefinedVariable{}

func NewDefinedVariable(priv bool, typeNode *ASTTypeNode, name string, init IASTExprNode) *DefinedVariable {
	var p = new(DefinedVariable)
	p._impl = p
	p.isPrivate = priv
	p.name = name
	p.typeNode = typeNode
	p.initializer = init
	p.sequence = -1
	return p
}

func NewTmpNewDefinedVariable(t IType) *DefinedVariable {
	v := NewDefinedVariable(false, NewTypeNodeFromType(t), "@tmp"+strconv.Itoa(int(tmpSeq)), nil)
	tmpSeq++
	return v
}

func (this *DefinedVariable) IsDefined() bool {
	return true
}

func (this *DefinedVariable) SetSequence(seq int64) {
	this.sequence = seq
}

// TODO: currently cast int64 to int, how convert int64 to string
func (this *DefinedVariable) SymbolString() string {
	if this.sequence < 0 {
		return this.name
	}
	return this.name + "." + strconv.Itoa(int(this.sequence))
}

func (this *DefinedVariable) HasInitializer() bool {
	return this.initializer != nil
}

func (this *DefinedVariable) IsInitialized() bool {
	return this.HasInitializer()
}

func (this *DefinedVariable) Initializer() IASTExprNode {
	return this.initializer
}

func (this *DefinedVariable) SetInitializer(expr IASTExprNode) {
	this.initializer = expr
}

func (this *DefinedVariable) SetIR(expr IIRExpr) {
	this.ir = expr
}

func (this *DefinedVariable) IR() IIRExpr {
	return this.ir
}

func (this *DefinedVariable) _Dump(d *Dumper) {
	d.PrintMemberStringNotResolved("name", this.name)
	d.PrintMemberBool("isPrivate", this.isPrivate)
	d.PrintMemberDumpable("typeNode", this.typeNode)
	d.PrintMemberDumpable("initializer", this.initializer)
}

func (this *DefinedVariable) Accept(visitor IEntityVisitor) interface{} {
	return visitor.VisitDefinedVariable(this)
}
