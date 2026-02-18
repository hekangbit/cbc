package asm

type IMemoryReference interface {
	IOperand
	CompareTo(other *MemoryReference) int
	FixOffset(int64)
	CmpDirectMemRef(*DirectMemoryReference) int
	CmpIndirectMemRef(*IndirectMemoryReference) int
}

type MemoryReference struct {
	Operand
}

func (this *MemoryReference) IsMemoryReference() bool {
	return true
}

func (this *MemoryReference) CompareTo(other *MemoryReference) int {
	panic("abstract method: MemoryReference::CompareTo")
}
func (this *MemoryReference) FixOffset(int64) {
	panic("abstract method: MemoryReference::FixOffset")
}
func (this *MemoryReference) CmpDirectMemRef(*DirectMemoryReference) int {
	panic("abstract method: MemoryReference::CmpDirectMemRef")
}
func (this *MemoryReference) CmpIndirectMemRef(*IndirectMemoryReference) int {
	panic("abstract method: MemoryReference::CmpIndirectMemRef")
}
