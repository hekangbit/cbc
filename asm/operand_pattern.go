package asm

type IOperandPattern interface {
	Match(*Operand) bool
}
