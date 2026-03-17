package models

import "cbc/asm"

type ASTCaseNode struct {
	ASTStmtNode
	label  *asm.Label
	values []IASTExprNode
	body   *ASTBlockNode
}

var _ IASTStmtNode = &ASTCaseNode{}

func NewASTCaseNode(loc *Location, values []IASTExprNode, body *ASTBlockNode) *ASTCaseNode {
	p := &ASTCaseNode{
		ASTStmtNode: ASTStmtNode{location: loc},
		label:       asm.NewLabelUnnamed(),
		values:      values,
		body:        body,
	}
	p._impl = p
	return p
}

func (this *ASTCaseNode) Values() []IASTExprNode {
	return this.values
}

func (this *ASTCaseNode) IsDefault() bool {
	return len(this.values) == 0
}

func (this *ASTCaseNode) Body() *ASTBlockNode {
	return this.body
}

func (this *ASTCaseNode) Label() *asm.Label {
	return this.label
}

func (this *ASTCaseNode) Accept(visitor IASTVisitor) (any, error) {
	return visitor.VisitCaseNode(this)
}

func (this *ASTCaseNode) _Dump(d *Dumper) {
	buf := make([]Dumpable, len(this.values))
	for i, tmp := range this.values {
		buf[i] = tmp
	}
	d.PrintNodeList("values", buf)
	d.PrintMemberDumpable("body", this.body)
}
