package asm

import "reflect"

type IImmediateValue interface {
	IOperand
}

type ImmediateValue struct {
	Operand
	expr ILiteral
}

func NewImmediateValue(expr ILiteral) *ImmediateValue {
	return &ImmediateValue{expr: expr}
}

func NewImmediateValueInt(n int64) *ImmediateValue {
	// TODO:
	panic("TODO: NewImmediateValueInt")
}

func (this *ImmediateValue) Expr() ILiteral {
	return this.expr
}

func (this *ImmediateValue) Equal(other any) bool {
	otherImm, ok := other.(*ImmediateValue)
	if !ok {
		return false
	}
	return reflect.DeepEqual(this.expr, otherImm.expr)
}

func (iv *ImmediateValue) ToSource(table SymbolTable) string {
	return "$" + iv.expr.ToSourceWithTable(table)
}

func (iv *ImmediateValue) Dump() string {
	return "(ImmediateValue " + iv.expr.Dump() + ")"
}

func (iv *ImmediateValue) CollectStatistics(stats Statistics) {
	// does nothing
}
