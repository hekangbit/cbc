package models

type ASTSizeofTypeNode struct {
	ASTExprNode
	operand *TypeNode
	ty      *TypeNode
}

var _ IASTExprNode = &ASTSizeofTypeNode{}

func NewASTSizeofTypeNode(operand *TypeNode, tyRef ITypeRef) *ASTSizeofTypeNode {
	return &ASTSizeofTypeNode{operand: operand, ty: NewTypeNodeFromRef(tyRef)}
}

func (this *ASTSizeofTypeNode) Operand() IType {
	return this.operand.Type()
}

func (this *ASTSizeofTypeNode) OperandTypeNode() *TypeNode {
	return this.operand
}

func (this *ASTSizeofTypeNode) Type() IType {
	return this.ty.Type()
}

func (this *ASTSizeofTypeNode) TypeNode() *TypeNode {
	return this.ty
}

func (this *ASTSizeofTypeNode) Location() *Location {
	return this.operand.Location()
}

func (this *ASTSizeofTypeNode) Dump(d *Dumper) {
	d.PrintMemberDumpable("operand", this.operand)
}

func (this *ASTSizeofTypeNode) Accept(visitor IASTVisitor) interface{} {
	return visitor.VisitSizeofTypeNode(this)
}
