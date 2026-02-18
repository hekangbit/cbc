package models

type DefinedFunction struct {
	*Function
	params *Params
	body   *ASTBlockNode
	scope  *LocalScope
	ir     []IRStmt
}

var _ IFunction = (*DefinedFunction)(nil)

func NewDefinedFunction(priv bool, typeNode *TypeNode, name string, params *Params, body *ASTBlockNode) *DefinedFunction {
	return &DefinedFunction{
		Function: NewFunction(priv, typeNode, name),
		params:   params,
		body:     body,
	}
}

func (this *DefinedFunction) IsDefined() bool {
	return true
}

func (this *DefinedFunction) Parameters() []*CBCParameter {
	return this.params.Parameters()
}

func (this *DefinedFunction) Body() *ASTBlockNode {
	return this.body
}

func (this *DefinedFunction) IR() []IRStmt {
	return this.ir
}

func (this *DefinedFunction) SetIR(ir []IRStmt) {
	this.ir = ir
}

func (this *DefinedFunction) SetScope(scope *LocalScope) {
	this.scope = scope
}

func (this *DefinedFunction) LvarScope() *LocalScope {
	return this.body.Scope()
}

func (this *DefinedFunction) LocalVariables() []*DefinedVariable {
	return this.scope.AllLocalVariables()
}

func (this *DefinedFunction) Dump(d *Dumper) {
	d.PrintClass(this, this.Location())
	d.PrintMemberString("name", this.Name(), false)
	d.PrintMemberBool("isPrivate", this.IsPrivate())
	d.PrintMemberDumpable("params", this.params)
	d.PrintMemberDumpable("body", this.body)
}

func (this *DefinedFunction) Accept(visitor IEntityVisitor) interface{} {
	return visitor.VisitDefinedFunction(this)
}
