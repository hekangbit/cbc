package models

// BlockNode 表示代码块节点
type ASTBlockNode struct {
	*BaseASTStmtNode
	variables []*DefinedVariable
	stmts     []IASTStmtNode
	scope     *LocalScope
}

// NewBlockNode 创建新的代码块节点
func NewASTBlockNode(loc *Location, vars []*DefinedVariable, stmts []IASTStmtNode) *ASTBlockNode {
	return &ASTBlockNode{
		BaseASTStmtNode: NewBaseASTStmtNode(loc),
		variables:       vars,
		stmts:           stmts,
	}
}

// Variables 返回变量列表
func (this *ASTBlockNode) Variables() []*DefinedVariable {
	return this.variables
}

// Stmts 返回语句列表
func (this *ASTBlockNode) Stmts() []IASTStmtNode {
	return this.stmts
}

// TailStmt 返回最后一个语句
func (this *ASTBlockNode) TailStmt() IASTStmtNode {
	if len(this.stmts) == 0 {
		return nil
	}
	return this.stmts[len(this.stmts)-1]
}

// Scope 返回局部作用域
func (this *ASTBlockNode) Scope() *LocalScope {
	return this.scope
}

// SetScope 设置局部作用域
func (this *ASTBlockNode) SetScope(scope *LocalScope) {
	this.scope = scope
}

func (this *ASTBlockNode) Dump(d *Dumper) {
	d.PrintClass(this, this.Location())
}

// Accept 实现访问者模式
func (this *ASTBlockNode) Accept(visitor ASTVisitor) interface{} {
	return visitor.VisitBlock(this)
}
