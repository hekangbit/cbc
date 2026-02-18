package models

type Constant struct {
	*Entity
	value IASTExprNode
}

func NewConstant(typeNode *TypeNode, name string, value IASTExprNode) *Constant {
	return &Constant{
		Entity: NewEntity(true, typeNode, name),
		value:  value,
	}
}

func (this *Constant) IsAssignable() bool {
	return false
}

func (this *Constant) IsDefined() bool {
	return true
}

func (this *Constant) IsInitialized() bool {
	return true
}

func (this *Constant) IsConstant() bool {
	return true
}

func (this *Constant) Value() IASTExprNode {
	return this.value
}

func (this *Constant) Dump(d *Dumper) {
	d.PrintClass(this, this.Location())
	d.PrintMemberString("name", this.name, false)
	d.PrintMemberTypeNode("typeNode", this.typeNode)
	d.PrintMemberDumpable("value", this.value)
}

func (this *Constant) Accept(visitor IEntityVisitor) interface{} {
	return visitor.VisitConstant(this)
}
