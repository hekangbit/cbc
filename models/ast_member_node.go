package models

type ASTMemberNode struct {
	ASTLHSNode
	expr   IASTExprNode
	member string
}

var _ IASTLHSNode = &ASTMemberNode{}

func NewASTMemberNode(expr IASTExprNode, member string) *ASTMemberNode {
	return &ASTMemberNode{
		expr:   expr,
		member: member,
	}
}

func (this *ASTMemberNode) ElemType() ICompositeType {
	// TODO: cbc java use try catch
	// java throw new SemanticError(err.getMessage()); when cast fail
	// design static shared method, to cast IType to CompositeType
	return GetCompositeType(this.expr.Type())
}

func (this *ASTMemberNode) Expr() IASTExprNode {
	return this.expr
}

func (this *ASTMemberNode) Member() string {
	return this.member
}

func (this *ASTMemberNode) Offset() int64 {
	return this.ElemType().MemberOffset(this.member)
}

// TODO: MemberType need return error when Member string not exist
// java throw Exception, but golang can't
func (this *ASTMemberNode) OrigType() IType {
	return this.ElemType().MemberType(this.member)
}

func (this *ASTMemberNode) Location() *Location {
	return this.expr.Location()
}

func (this *ASTMemberNode) Dump(d *Dumper) {
	if this.ty != nil {
		d.PrintMemberType("type", this.ty)
	}
	d.PrintMemberDumpable("expr", this.expr)
	d.PrintMemberStringNotResolved("member", this.member)
}

func (this *ASTMemberNode) Accept(visitor IASTVisitor) interface{} {
	return visitor.VisitMemberNodeNode(this)
}
