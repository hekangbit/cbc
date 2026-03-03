package models

type ASTCastNode struct {
	ASTExprNode
	typeNode *ASTTypeNode
	expr     IASTExprNode
}

var _ IASTExprNode = &ASTCastNode{}

func NewASTCastNode(t IType, expr IASTExprNode) *ASTCastNode {
	return NewASTCastNodeWithTypeNode(NewTypeNodeFromType(t), expr)
}

func NewASTCastNodeWithTypeNode(typeNode *ASTTypeNode, expr IASTExprNode) *ASTCastNode {
	p := &ASTCastNode{
		typeNode: typeNode,
		expr:     expr,
	}
	p.ASTExprNode._impl = p
	p.ASTExprNode.Node._impl = p
	return p
}

func (this *ASTCastNode) Type() IType {
	return this.typeNode.Type()
}

func (this *ASTCastNode) TypeNode() *ASTTypeNode {
	return this.typeNode
}

func (this *ASTCastNode) Expr() IASTExprNode {
	return this.expr
}

func (this *ASTCastNode) IsLvalue() bool {
	return this.expr.IsLvalue()
}

func (this *ASTCastNode) IsAssignable() bool {
	return this.expr.IsAssignable()
}

func (this *ASTCastNode) IsEffectiveCast() bool {
	return this.Type().Size() > this.expr.Type().Size()
}

func (this *ASTCastNode) Location() *Location {
	return this.typeNode.Location()
}

func (this *ASTCastNode) _Dump(d *Dumper) {
	d.PrintMemberDumpable("typeNode", this.typeNode)
	d.PrintMemberDumpable("expr", this.expr)
}

func (this *ASTCastNode) Accept(visitor IASTVisitor) interface{} {
	return visitor.VisitCastNode(this)
}
