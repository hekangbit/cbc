package models

type Constant struct {
	Entity
	value IASTExprNode
}

var _ IEntity = &Constant{}

func NewConstant(typeNode *ASTTypeNode, name string, value IASTExprNode) *Constant {
	p := &Constant{
		Entity: Entity{name: name, isPrivate: true, typeNode: typeNode, nRefered: 0},
		value:  value,
	}
	p._impl = p
	return p
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

func (this *Constant) _Dump(d *Dumper) {
	d.PrintMemberString("name", this.name, false)
	d.PrintMemberTypeNode("typeNode", this.typeNode)
	d.PrintMemberDumpable("value", this.value)
}

func (this *Constant) Accept(visitor IEntityVisitor) interface{} {
	return visitor.VisitConstant(this)
}
