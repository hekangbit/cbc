package models

type ASTSwitchNode struct {
	ASTStmtNode
	cond  IASTExprNode
	cases []*ASTCaseNode
}

var _ IASTStmtNode = &ASTSwitchNode{}

func NewASTSwitchNode(loc *Location, cond IASTExprNode, cases []*ASTCaseNode) *ASTSwitchNode {
	p := &ASTSwitchNode{
		ASTStmtNode: ASTStmtNode{location: loc},
		cond:        cond,
		cases:       cases,
	}
	p._impl = p
	return p
}

func (this *ASTSwitchNode) Cond() IASTExprNode {
	return this.cond
}

func (this *ASTSwitchNode) Cases() []*ASTCaseNode {
	return this.cases
}

func (this *ASTSwitchNode) Accept(visitor IASTVisitor) any {
	return visitor.VisitSwitchNode(this)
}

func (this *ASTSwitchNode) _Dump(d *Dumper) {
	d.PrintMemberDumpable("cond", this.cond)
	buf := make([]Dumpable, len(this.cases))
	for i, tmp := range this.cases {
		buf[i] = tmp
	}
	d.PrintNodeList("cases", buf)
}
