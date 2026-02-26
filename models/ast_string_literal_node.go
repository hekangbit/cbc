package models

type ASTStringLiteralNode struct {
	ASTLiteralNode
	value string
	entry *ConstantEntry
}

var _ IASTLiteralNode = &ASTStringLiteralNode{}

func NewASTStringLiteralNode(loc *Location, ref ITypeRef, value string) *ASTStringLiteralNode {
	var p = new(ASTStringLiteralNode)
	p.ASTLiteralNode.ASTExprNode._impl = p
	p.ASTLiteralNode.ASTExprNode.Node._impl = p
	p.location = loc
	p.typeNode = NewTypeNodeFromRef(ref)
	p.value = value
	return p
}

func (this *ASTStringLiteralNode) Value() string {
	return this.value
}

func (this *ASTStringLiteralNode) Entry() *ConstantEntry {
	return this.entry
}

func (this *ASTStringLiteralNode) SetEntry(ent *ConstantEntry) {
	this.entry = ent
}

func (this *ASTStringLiteralNode) _Dump(d *Dumper) {
	d.PrintMemberStringNotResolved("value", this.value)
}

func (this *ASTStringLiteralNode) Accept(visitor IASTVisitor) interface{} {
	return visitor.VisitStringLiteralNode(this)
}
