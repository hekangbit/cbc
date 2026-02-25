package models

type UndefinedVariable struct {
	Variable
}

func NewUndefinedVariable(t *ASTTypeNode, name string) *UndefinedVariable {
	var p = new(UndefinedVariable)
	p._impl = p
	p.isPrivate = false
	p.name = name
	p.typeNode = t
	return p
}

func (this *UndefinedVariable) IsDefined() bool {
	return false
}

func (this *UndefinedVariable) IsPrivate() bool {
	return false
}

func (this *UndefinedVariable) IsInitialized() bool {
	return false
}

func (this *UndefinedVariable) _Dump(d *Dumper) {
	d.PrintMemberString("name", this.name, false)
	d.PrintMemberBool("isPrivate", this.IsPrivate())
	d.PrintMemberDumpable("typeNode", this.typeNode)
}

func (this *UndefinedVariable) Accept(visitor IEntityVisitor) interface{} {
	return visitor.VisitUndefinedVariable(this)
}
