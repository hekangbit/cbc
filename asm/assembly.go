package asm

type IAssembly interface {
	ToSource(table SymbolTable) string
	Dump() string
	IsInstruction() bool
	IsLabel() bool
	IsDirective() bool
	IsComment() bool
	CollectStatistics(stats Statistics)
}

type Assembly struct {
}

var _ IAssembly = (*Assembly)(nil)

func (this *Assembly) ToSource(table SymbolTable) string {
	panic("abstract method: Assembly::ToSource must be implemented by subtype")
}

func (this *Assembly) Dump() string {
	panic("abstract method: Assembly::Dump must be implemented by subtype")
}

func (this *Assembly) IsInstruction() bool {
	return false
}
func (this *Assembly) IsLabel() bool {
	return false
}
func (this *Assembly) IsDirective() bool {
	return false
}
func (this *Assembly) IsComment() bool {
	return false
}

func (this *Assembly) CollectStatistics(stats Statistics) {
	// does nothing by default.
}
