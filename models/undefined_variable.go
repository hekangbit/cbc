package models

type UndefinedVariable struct {
	Variable
}

func NewUndefinedVariable(t *TypeNode, name string) *UndefinedVariable {
	return &UndefinedVariable{
		Variable: Variable{Entity: Entity{isPrivate: false, typeNode: t, name: name}},
	}
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

func (this *UndefinedVariable) Dump(d *Dumper) {
	d.PrintClass(this, this.Location())
	d.PrintMemberString("name", this.name, false)
	d.PrintMemberBool("isPrivate", this.IsPrivate())
	d.PrintMemberDumpable("typeNode", this.typeNode)
}

func (this *UndefinedVariable) Accept(visitor IEntityVisitor) interface{} {
	return visitor.VisitUndefinedVariable(this)
}
