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
func (this *ASTPtrMemberNode) DereferedCompositeType() ICompositeType {
	pt := GetPointerType(this.expr.Type())
	return GetCompositeType(pt.ElemType()) // catch (ClassCastException err) {	throw new SemanticError(err.getMessage());
}

func (this *ASTPtrMemberNode) DereferedType() IType {
	pt := GetPointerType(this.expr.Type())
	return pt.ElemType() // catch (ClassCastException err) {	throw new SemanticError(err.getMessage());
}

func (this *ASTPtrMemberNode) Expr() IASTExprNode {
	return this.expr
}

func (this *ASTPtrMemberNode) Member() string {
	return this.member
}

// TODO: MemberOffset may return error when not found member
func (this *ASTPtrMemberNode) Offset() int64 {
	return this.DereferedCompositeType().MemberOffset(this.member)
}

// TODO: MemberOffset may return error when not found member
func (this *ASTPtrMemberNode) OrigType() IType {
	return this.DereferedCompositeType().MemberType(this.member)
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

func (this *ASTPtrMemberNode) Accept(visitor IASTVisitor) interface{} {
	return visitor.VisitPtrMemberNode(this)
}
