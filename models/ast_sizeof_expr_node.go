package models

type ASTSizeofExprNode struct {
	ASTExprNode
	expr IASTExprNode
	ty   *ASTTypeNode
}

var _ IASTExprNode = &ASTSizeofExprNode{}

func NewASTSizeofExprNode(expr IASTExprNode, tyRef ITypeRef) *ASTSizeofExprNode {
	p := &ASTSizeofExprNode{expr: expr, ty: NewTypeNodeFromRef(tyRef)}
	p.ASTExprNode._impl = p
	p.ASTExprNode.Node._impl = p
	return p
}

func (this *ASTSizeofExprNode) Expr() IASTExprNode {
	return this.expr
}

func (this *ASTSizeofExprNode) SetExpr(expr IASTExprNode) {
	this.expr = expr
}

func (this *ASTSizeofExprNode) Type() IType {
	return this.ty.Type()
}

func (this *ASTSizeofExprNode) TypeNode() *ASTTypeNode {
	return this.ty
}

func (this *ASTSizeofExprNode) Location() *Location {
	return this.expr.Location()
}

func (this *ASTSizeofExprNode) _Dump(d *Dumper) {
	d.PrintMemberDumpable("expr", this.expr)
}

func (this *ASTSizeofExprNode) Accept(visitor IASTVisitor) interface{} {
	return visitor.VisitSizeofExprNode(this)
}
