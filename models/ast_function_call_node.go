package models

type ASTFunctionCallNode struct {
	ASTExprNode
	expr IASTExprNode
	args []IASTExprNode
}

var _ IASTExprNode = &ASTFunctionCallNode{}

func NewASTFunctionCallNode(expr IASTExprNode, args []IASTExprNode) *ASTFunctionCallNode {
	p := &ASTFunctionCallNode{expr: expr, args: args}
	p.ASTExprNode._impl = p
	p.ASTExprNode.Node._impl = p
	return p
}

func (this *ASTFunctionCallNode) Expr() IASTExprNode {
	return this.expr
}

/*
* Returns a type of return value of the function which is refered
* by expr.  This method expects expr.type().isCallable() is true.
 */
func (this *ASTFunctionCallNode) Type() IType {
	return this.FunctionType().ReturnType()
}

/*
* Returns a type of function which is refered by expr.
* This method expects expr.type().isCallable() is true.
 */
// TODO: GetPointerType may need return error? cast may fail throw exception in java
func (this *ASTFunctionCallNode) FunctionType() *FunctionType {
	t := GetPointerType(this.expr.Type()).ElemType()
	return GetFunctionType(t)
}

// TODO: int64 -> int ?
func (this *ASTFunctionCallNode) NumArgs() int64 {
	return (int64)(len(this.args))
}

func (this *ASTFunctionCallNode) Args() []IASTExprNode {
	return this.args
}

// called from TypeChecker
func (this *ASTFunctionCallNode) ReplaceArgs(args []IASTExprNode) {
	this.args = args
}

func (this *ASTFunctionCallNode) Location() *Location {
	return this.expr.Location()
}

func (this *ASTFunctionCallNode) _Dump(d *Dumper) {
	d.PrintMemberDumpable("expr", this.expr)
	dumpables := make([]Dumpable, len(this.args))
	for i, tmp := range this.args {
		dumpables[i] = tmp
	}
	d.PrintNodeList("args", dumpables)
}

func (this *ASTFunctionCallNode) Accept(visitor IASTVisitor) interface{} {
	return visitor.VisitFunctionCallNode(this)
}
