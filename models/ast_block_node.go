package models

type ASTBlockNode struct {
	*ASTStmtNode
	variables []*DefinedVariable
	stmts     []IASTStmtNode
	scope     *LocalScope
}

func NewASTBlockNode(loc *Location, vars []*DefinedVariable, stmts []IASTStmtNode) *ASTBlockNode {
	return &ASTBlockNode{
		ASTStmtNode: NewASTStmtNode(loc),
		variables:   vars,
		stmts:       stmts,
	}
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

func (this *ASTBlockNode) Dump(d *Dumper) {
	d.PrintClass(this, this.Location())
}

func (this *ASTBlockNode) Accept(visitor IASTVisitor) interface{} {
	return visitor.VisitBlock(this)
}
