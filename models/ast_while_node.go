package models

type ASTWhileNode struct {
	ASTStmtNode
	body IASTStmtNode
	cond IASTExprNode
}

var _ IASTStmtNode = &ASTWhileNode{}

func NewASTWhileNode(loc *Location, cond IASTExprNode, body IASTStmtNode) *ASTWhileNode {
	p := &ASTWhileNode{ASTStmtNode: ASTStmtNode{location: loc}, cond: cond, body: body}
	p._impl = p
	return p
}

func (this *ASTWhileNode) Cond() IASTExprNode {
	return this.cond
}

func (this *ASTWhileNode) Body() IASTStmtNode {
	return this.body
}

func (this *ASTWhileNode) _Dump(d *Dumper) {
	d.PrintMemberDumpable("cond", this.cond)
	d.PrintMemberDumpable("body", this.body)
}

func (this *ASTWhileNode) Accept(visitor IASTVisitor) (any, error) {
	return visitor.VisitWhileNode(this)
}
