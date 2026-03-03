package models

type ASTIfNode struct {
	ASTStmtNode
	cond     IASTExprNode
	thenBody IASTStmtNode
	elseBody IASTStmtNode
}

var _ IASTStmtNode = &ASTIfNode{}

func NewASTIfNode(loc *Location, c IASTExprNode, t IASTStmtNode, e IASTStmtNode) *ASTIfNode {
	p := &ASTIfNode{
		ASTStmtNode: ASTStmtNode{location: loc},
		cond:        c,
		thenBody:    t,
		elseBody:    e,
	}
	p._impl = p
	return p
}

func (this *ASTIfNode) Cond() IASTExprNode {
	return this.cond
}

func (this *ASTIfNode) ThenBody() IASTStmtNode {
	return this.thenBody
}

func (this *ASTIfNode) ElseBody() IASTStmtNode {
	return this.elseBody
}

func (this *ASTIfNode) _Dump(d *Dumper) {
	d.PrintMemberDumpable("cond", this.cond)
	d.PrintMemberDumpable("thenBody", this.thenBody)
	d.PrintMemberDumpable("elseBody", this.elseBody)
}

func (this *ASTIfNode) Accept(visitor IASTVisitor) interface{} {
	return visitor.VisitIfNode(this)
}
