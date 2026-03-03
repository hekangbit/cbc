package models

type ASTForNode struct {
	ASTStmtNode
	init IASTStmtNode
	cond IASTExprNode
	incr IASTStmtNode
	body IASTStmtNode
}

var _ IASTStmtNode = &ASTForNode{}

func NewASTForNode(loc *Location, init IASTExprNode, cond IASTExprNode, incr IASTExprNode, body IASTStmtNode) *ASTForNode {
	p := &ASTForNode{
		ASTStmtNode: ASTStmtNode{location: loc},
		body:        body,
	}
	p._impl = p

	if init != nil {
		p.init = NewASTExprStmtNode(init.Location(), init)
	} else {
		p.init = nil
	}

	if cond != nil {
		p.cond = cond
	} else {
		p.cond = NewASTIntegerLiteralNode(nil, NewIntRef(), 1) // always true
	}

	if incr != nil {
		p.incr = NewASTExprStmtNode(incr.Location(), incr)
	} else {
		p.incr = nil
	}

	return p
}

func (this *ASTForNode) Init() IASTStmtNode {
	return this.init
}

func (this *ASTForNode) Cond() IASTExprNode {
	return this.cond
}

func (this *ASTForNode) Incr() IASTStmtNode {
	return this.incr
}

func (this *ASTForNode) Body() IASTStmtNode {
	return this.body
}

func (this *ASTForNode) Accept(visitor IASTVisitor) any {
	return visitor.VisitForNode(this)
}

func (this *ASTForNode) _Dump(d *Dumper) {
	d.PrintMemberDumpable("init", this.init)
	d.PrintMemberDumpable("cond", this.cond)
	d.PrintMemberDumpable("incr", this.incr)
	d.PrintMemberDumpable("body", this.body)
}
