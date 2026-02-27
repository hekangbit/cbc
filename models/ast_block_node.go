package models

type ASTBlockNode struct {
	ASTStmtNode
	variables []*DefinedVariable
	stmts     []IASTStmtNode
	scope     *LocalScope
}

var _ IASTStmtNode = &ASTBlockNode{}

func NewASTBlockNode(loc *Location, vars []*DefinedVariable, stmts []IASTStmtNode) *ASTBlockNode {
	p := &ASTBlockNode{ASTStmtNode: ASTStmtNode{location: loc}, variables: vars, stmts: stmts}
	p._impl = p
	return p
}

func (this *ASTBlockNode) Variables() []*DefinedVariable {
	return this.variables
}

func (this *ASTBlockNode) Stmts() []IASTStmtNode {
	return this.stmts
}

func (this *ASTBlockNode) TailStmt() IASTStmtNode {
	if len(this.stmts) == 0 {
		return nil
	}
	return this.stmts[len(this.stmts)-1]
}

func (this *ASTBlockNode) Scope() *LocalScope {
	return this.scope
}

func (this *ASTBlockNode) SetScope(scope *LocalScope) {
	this.scope = scope
}

func (this *ASTBlockNode) _Dump(d *Dumper) {
	ivars := make([]Dumpable, len(this.variables))
	for i, tmp := range this.variables {
		ivars[i] = tmp
	}
	d.PrintNodeList("variables", ivars)

	istmts := make([]Dumpable, len(this.stmts))
	for i, tmp := range this.stmts {
		istmts[i] = tmp
	}
	d.PrintNodeList("stmts", istmts)
}

func (this *ASTBlockNode) Accept(visitor IASTVisitor) interface{} {
	return visitor.VisitBlockNode(this)
}
