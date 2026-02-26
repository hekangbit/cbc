package models

type ASTIntegerLiteralNode struct {
	ASTLiteralNode
	value int64
}

var _ IASTLiteralNode = &ASTIntegerLiteralNode{}

// TODO: access parent field directly, how to improve
func NewASTIntegerLiteralNode(loc *Location, ref ITypeRef, value int64) *ASTIntegerLiteralNode {
	var p = new(ASTIntegerLiteralNode)
	p.ASTLiteralNode.ASTExprNode._impl = p
	p.ASTLiteralNode.ASTExprNode.Node._impl = p
	p.location = loc
	p.typeNode = NewTypeNodeFromRef(ref)
	p.value = value
	return p
}

func (this *ASTIntegerLiteralNode) Value() int64 {
	return this.value
}

// TODO: int64 to int may error when print member char
func (this *ASTIntegerLiteralNode) _Dump(d *Dumper) {
	d.PrintMemberDumpable("typeNode", this.typeNode)
	if this.typeNode.typeRef.(*IntegerTypeRef).name == "char" {
		d.PrintMemberChar("value", int(this.value))
	} else {
		d.PrintMemberInt64("value", this.value)
	}
}

func (this *ASTIntegerLiteralNode) Accept(visitor IASTVisitor) interface{} {
	return visitor.VisitIntegerLiteralNode(this)
}
