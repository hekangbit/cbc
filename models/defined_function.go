package models

type DefinedFunction struct {
	Function
	params *Params
	body   *ASTBlockNode
	scope  *LocalScope
	ir     []IRStmt
}

var _ IFunction = &DefinedFunction{}

func NewDefinedFunction(priv bool, t *ASTTypeNode, name string, params *Params, body *ASTBlockNode) *DefinedFunction {
	var p = new(DefinedFunction)
	p._impl = p
	p.isPrivate = priv
	p.name = name
	p.typeNode = t
	p.params = params
	p.body = body
	return p
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

func (this *DefinedFunction) _Dump(d *Dumper) {
	d.PrintMemberString("name", this.Name(), false)
	d.PrintMemberBool("isPrivate", this.IsPrivate())
	d.PrintMemberDumpable("params", this.params)
	d.PrintMemberDumpable("body", this.body)
}

func (this *DefinedFunction) Accept(visitor IEntityVisitor) interface{} {
	return visitor.VisitDefinedFunction(this)
}
