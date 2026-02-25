package models

type ASTAddressNode struct {
	ASTExprNode
	ty   IType
	expr IASTExprNode
}

var _ IASTExprNode = &ASTAddressNode{}

func NewASTAddressNode(expr IASTExprNode) *ASTAddressNode {
	p := &ASTAddressNode{expr: expr}
	p.ASTExprNode._impl = p
	p.ASTExprNode.Node._impl = p
	return p
}

func (this *ASTAddressNode) Expr() IASTExprNode {
	return this.expr
}

func (this *ASTAddressNode) Type() IType {
	if this.ty == nil {
		panic("type is nil")
	}
	return this.ty
}

func (this *ASTAddressNode) SetType(t IType) {
	if this.ty != nil {
		panic("type set twice")
	}
	this.ty = t
}

func (this *ASTAddressNode) Location() *Location {
	return this.expr.Location()
}

func (this *ASTAddressNode) _Dump(d *Dumper) {
	if this.ty != nil {
		d.PrintMemberType("type", this.ty)
	}
	d.PrintMemberDumpable("expr", this.expr)
}

func (this *ASTAddressNode) Accept(visitor IASTVisitor) interface{} {
	return visitor.VisitAddressNode(this)
}
