package models

type ASTGotoNode struct {
	ASTStmtNode
	target string
}

var _ IASTStmtNode = &ASTGotoNode{}

func NewASTGotoNode(loc *Location, target string) *ASTGotoNode {
	p := &ASTGotoNode{
		ASTStmtNode: ASTStmtNode{location: loc},
		target:      target,
	}
	p._impl = p
	return p
}

func (this *ASTGotoNode) Target() string {
	return this.target
}

func (this *ASTGotoNode) Accept(visitor IASTVisitor) (any, error) {
	return visitor.VisitGotoNode(this)
}

func (this *ASTGotoNode) _Dump(d *Dumper) {
	d.PrintMemberStringNotResolved("target", this.target)
}
