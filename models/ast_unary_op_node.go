package models

type ASTUnaryOpNode struct {
	ASTExprNode
	operator string
	expr     IASTExprNode
	opType   IType
}

var _ IASTExprNode = &ASTUnaryOpNode{}

func NewASTUnaryOpNode(op string, expr IASTExprNode) *ASTUnaryOpNode {
	p := &ASTUnaryOpNode{operator: op, expr: expr}
	p.ASTExprNode._impl = p
	p.ASTExprNode.Node._impl = p
	return p
}

func (this *ASTUnaryOpNode) Operator() string {
	return this.operator
}

func (this *ASTUnaryOpNode) Type() IType {
	return this.expr.Type()
}

func (this *ASTUnaryOpNode) SetOpType(t IType) {
	this.opType = t
}

func (this *ASTUnaryOpNode) OpType() IType {
	return this.opType
}

func (this *ASTUnaryOpNode) Expr() IASTExprNode {
	return this.expr
}

func (this *ASTUnaryOpNode) SetExpr(expr IASTExprNode) {
	this.expr = expr
}

func (this *ASTUnaryOpNode) Location() *Location {
	return this.expr.Location()
}

func (this *ASTUnaryOpNode) _Dump(d *Dumper) {
	d.PrintMemberStringNotResolved("operator", this.operator)
	d.PrintMemberDumpable("expr", this.expr)
}

func (this *ASTUnaryOpNode) Accept(visitor IASTVisitor) interface{} {
	return visitor.VisitUnaryOpNode(this)
}
