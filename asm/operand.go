package asm

type IOperand interface {
	IOperandPattern
	ToSourceWithTable(SymbolTable) string
	Dump() string
	IsRegister() bool
	IsMemoryReference() bool
	IntegerLiteral() IntegerLiteral
	CollectStatistics(Statistics)
}

type Operand struct {
}

func (this *Operand) ToSourceWithTable(table SymbolTable) string {
	panic("abstract method: Operand::ToSourceWithTable")
}

func (this *Operand) Dump() string {
	panic("abstract method: Operand::Dump")
}

func (this *Operand) isRegister() bool {
	return false
}

func (this *Operand) isMemoryReference() bool {
	return false
}

func (this *Operand) IntegerLiteral() *IntegerLiteral {
	return nil
}

func (this *Operand) CollectStatistics(Statistics) {
	panic("abstract method: Operand::CollectStatistics")
}
func (this *Operand) Match(operand *Operand) bool {
	return this == operand
}
