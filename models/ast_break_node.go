package models

type ASTBreakNode struct {
	ASTStmtNode
}

var _ IASTStmtNode = &ASTBreakNode{}

func NewASTBreakNode(loc *Location) *ASTBreakNode {
	p := &ASTBreakNode{ASTStmtNode{location: loc}}
	p._impl = p
	return p
}

func (this *ASTBreakNode) _Dump(d *Dumper) {
}

func (this *ASTBreakNode) Accept(visitor IASTVisitor) interface{} {
	return visitor.VisitBreakNode(this)
}
