package models

// TODO: finish this node
type ASTPtrMemberNode struct {
	ASTLHSNode
	expr   IASTExprNode
	member string
}

var _ IASTLHSNode = &ASTPtrMemberNode{}

func NewASTPtrMemberNode(expr IASTExprNode, member string) *ASTPtrMemberNode {
	p := &ASTPtrMemberNode{
		expr:   expr,
		member: member,
	}
	p.ASTLHSNode.ASTExprNode._impl = p
	p.ASTLHSNode.ASTExprNode.Node._impl = p
	return p
}

// TODO: cast may need return error
// catch (ClassCastException err) {	throw new SemanticError(err.getMessage());
func (this *ASTPtrMemberNode) DereferedCompositeType() ICompositeType {
	pt := this.expr.Type().GetPointerType()
	return pt.ElemType().GetCompositeType()
}

// TODO: check
// catch (ClassCastException err) {	throw new SemanticError(err.getMessage());
func (this *ASTPtrMemberNode) DereferedType() IType {
	pt := this.expr.Type().GetPointerType()
	return pt.ElemType()
}

func (this *ASTPtrMemberNode) Expr() IASTExprNode {
	return this.expr
}

func (this *ASTPtrMemberNode) Member() string {
	return this.member
}

// TODO: MemberOffset may return error when not found member
func (this *ASTPtrMemberNode) Offset() int64 {
	offset, _ := this.DereferedCompositeType().MemberOffset(this.member)
	return offset
}

// TODO: MemberOffset may return error when not found member
func (this *ASTPtrMemberNode) OrigType() IType {
	t, _ := this.DereferedCompositeType().MemberType(this.member)
	return t
}

func (this *ASTPtrMemberNode) Location() *Location {
	return this.expr.Location()
}

func (this *ASTPtrMemberNode) _Dump(d *Dumper) {
	if this.ty != nil {
		d.PrintMemberType("type", this.ty)
	}
	d.PrintMemberDumpable("expr", this.expr)
	d.PrintMemberStringNotResolved("member", this.member)
}

func (this *ASTPtrMemberNode) Accept(visitor IASTVisitor) (any, error) {
	return visitor.VisitPtrMemberNode(this)
}
