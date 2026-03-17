package models

type ASTLabelNode struct {
	ASTStmtNode
	name string
	stmt IASTStmtNode
}

var _ IASTStmtNode = &ASTLabelNode{}

func NewASTLabelNode(loc *Location, name string, stmt IASTStmtNode) *ASTLabelNode {
	p := &ASTLabelNode{
		ASTStmtNode: ASTStmtNode{location: loc},
		name:        name,
		stmt:        stmt,
	}
	p._impl = p
	return p
}

func (this *ASTLabelNode) Name() string {
	return this.name
}

func (this *ASTLabelNode) Stmt() IASTStmtNode {
	return this.stmt
}

func (this *ASTLabelNode) Accept(visitor IASTVisitor) (any, error) {
	return visitor.VisitLabelNode(this)
}

func (this *ASTLabelNode) _Dump(d *Dumper) {
	d.PrintMemberStringNotResolved("name", this.name)
	d.PrintMemberDumpable("stmt", this.stmt)
}
