package models

type ASTSizeofTypeNode struct {
	ASTExprNode
	operand *ASTTypeNode
	ty      *ASTTypeNode
}

var _ IASTExprNode = &ASTSizeofTypeNode{}

func NewASTSizeofTypeNode(operand *ASTTypeNode, tyRef ITypeRef) *ASTSizeofTypeNode {
	p := &ASTSizeofTypeNode{operand: operand, ty: NewTypeNodeFromRef(tyRef)}
	p.ASTExprNode._impl = p
	p.ASTExprNode.Node._impl = p
	return p
}

func (this *ASTSizeofTypeNode) Operand() IType {
	return this.operand.Type()
}

func (this *ASTSizeofTypeNode) OperandTypeNode() *ASTTypeNode {
	return this.operand
}

func (this *ASTSizeofTypeNode) Type() IType {
	return this.ty.Type()
}

func (this *ASTSizeofTypeNode) TypeNode() *ASTTypeNode {
	return this.ty
}

func (this *ASTSizeofTypeNode) Location() *Location {
	return this.operand.Location()
}

func (this *ASTSizeofTypeNode) _Dump(d *Dumper) {
	d.PrintMemberDumpable("operand", this.operand)
}

func (this *ASTSizeofTypeNode) Accept(visitor IASTVisitor) interface{} {
	return visitor.VisitSizeofTypeNode(this)
}
