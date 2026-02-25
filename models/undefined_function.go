package models

type UndefinedFunction struct {
	Function
	params *Params
}

var _ IFunction = &UndefinedFunction{}

func NewUndefinedFunction(t *ASTTypeNode, name string, params *Params) *UndefinedFunction {
	var p = new(UndefinedFunction)
	p._impl = p
	p.isPrivate = false
	p.name = name
	p.typeNode = t
	p.params = params
	return p
}

func (this *UndefinedFunction) Parameters() []*CBCParameter {
	return this.params.Parameters()
}

func (this *UndefinedFunction) IsDefined() bool {
	return false
}

func (this *UndefinedFunction) _Dump(d *Dumper) {
	d.PrintMemberStringNotResolved("name", this.name)
	d.PrintMemberBool("isPrivate", this.IsPrivate())
	d.PrintMemberDumpable("typeNode", this.typeNode)
	d.PrintMemberDumpable("params", this.params)
}

func (this *UndefinedFunction) Accept(visitor IEntityVisitor) interface{} {
	return visitor.VisitUndefinedFunction(this)
}
