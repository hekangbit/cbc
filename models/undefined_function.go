package models

type UndefinedFunction struct {
	Function
	params *Params
}

func NewUndefinedFunction(t *TypeNode, name string, params *Params) *UndefinedFunction {
	return &UndefinedFunction{
		Function: Function{Entity: Entity{isPrivate: false, typeNode: t, name: name}},
		params:   params,
	}
}

func (this *UndefinedFunction) Parameters() []*CBCParameter {
	return this.params.Parameters()
}

func (this *UndefinedFunction) IsDefined() bool {
	return false
}

func (this *UndefinedFunction) Dump(d *Dumper) {
	d.PrintMemberStringNotResolved("name", this.name)
	d.PrintMemberBool("isPrivate", this.IsPrivate())
	d.PrintMemberDumpable("typeNode", this.typeNode)
	d.PrintMemberDumpable("params", this.params)
}

func (this *UndefinedFunction) Accept(visitor IEntityVisitor) interface{} {
	return visitor.VisitUndefinedFunction(this)
}
