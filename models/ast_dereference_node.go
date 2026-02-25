package models

type ASTDereferenceNode struct {
	ASTLHSNode
	expr IASTExprNode
}

var _ IASTLHSNode = &ASTDereferenceNode{}

func NewASTDereferenceNode(expr IASTExprNode) *ASTDereferenceNode {
	p := &ASTDereferenceNode{expr: expr}
	p.ASTLHSNode.ASTExprNode._impl = p
	p.ASTLHSNode.ASTExprNode.Node._impl = p
	return p
}

func (this *ASTDereferenceNode) Type() IType {
	if this.ty != nil {
		return this.ty
	}
	return this.OrigType()
}

func (this *ASTDereferenceNode) OrigType() IType {
	return this.expr.Type().ElemType()
}

func (this *ASTDereferenceNode) AllocSize() int64 {
	return this.OrigType().AllocSize()
}

func (this *ASTDereferenceNode) IsAssignable() bool {
	return this.IsLoadable()
}

func (this *ASTDereferenceNode) IsLoadable() bool {
	t := this.OrigType()
	return !t.IsArray() && !t.IsFunction()
}

func (this *ASTDereferenceNode) Expr() IASTExprNode {
	return this.expr
}

func (this *ASTDereferenceNode) SetExpr(expr IASTExprNode) {
	this.expr = expr
}

func (this *ASTDereferenceNode) Location() *Location {
	return this.expr.Location()
}

func (this *ASTDereferenceNode) _Dump(d *Dumper) {
	if this.ty != nil {
		d.PrintMemberType("type", this.ty)
	}
	d.PrintMemberDumpable("expr", this.expr)
}

func (this *ASTDereferenceNode) Accept(visitor IASTVisitor) interface{} {
	return visitor.VisitDereferenceNode(this)
}
