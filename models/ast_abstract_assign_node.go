package models

type IASTAbstractAssignNode interface {
	IASTExprNode
	LHS() IASTExprNode
	RHS() IASTExprNode
	SetRHS(IASTExprNode)
}

type ASTAbstractAssignNode struct {
	ASTExprNode
	lhs IASTExprNode
	rhs IASTExprNode
}

func (this *ASTAbstractAssignNode) Type() IType {
	return this.lhs.Type()
}

func (this *ASTAbstractAssignNode) LHS() IASTExprNode {
	return this.lhs
}

func (this *ASTAbstractAssignNode) RHS() IASTExprNode {
	return this.rhs
}

func (this *ASTAbstractAssignNode) SetRHS(expr IASTExprNode) {
	this.rhs = expr
}

func (this *ASTAbstractAssignNode) Location() *Location {
	return this.LHS().Location()
}

func (this *ASTAbstractAssignNode) _Dump(d *Dumper) {
	d.PrintMemberDumpable("lhs", this.lhs)
	d.PrintMemberDumpable("rhs", this.rhs)
}
