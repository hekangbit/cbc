package models

type ASTDoWhileNode struct {
	ASTStmtNode
	body IASTStmtNode
	cond IASTExprNode
}

var _ IASTStmtNode = &ASTDoWhileNode{}

func NewASTDoWhileNode(loc *Location, body IASTStmtNode, cond IASTExprNode) *ASTDoWhileNode {
	p := &ASTDoWhileNode{ASTStmtNode: ASTStmtNode{location: loc}, body: body, cond: cond}
	p._impl = p
	return p
}

func (this *ASTDoWhileNode) Cond() IASTExprNode {
	return this.cond
}

func (this *ASTDoWhileNode) Body() IASTStmtNode {
	return this.body
}

func (this *ASTDoWhileNode) _Dump(d *Dumper) {
	d.PrintMemberDumpable("cond", this.cond)
	d.PrintMemberDumpable("body", this.body)
}

func (this *ASTDoWhileNode) Accept(visitor IASTVisitor) interface{} {
	return visitor.VisitDoWhileNode(this)
}
