package models

type ASTContinueNode struct {
	ASTStmtNode
}

var _ IASTStmtNode = &ASTContinueNode{}

func NewASTContinueNode(loc *Location) *ASTContinueNode {
	p := &ASTContinueNode{ASTStmtNode: ASTStmtNode{location: loc}}
	p._impl = p
	return p
}

func (this *ASTContinueNode) _Dump(d *Dumper) {
}

func (this *ASTContinueNode) Accept(visitor IASTVisitor) interface{} {
	return visitor.VisitContinueNode(this)
}
