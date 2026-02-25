package models

type ASTArrayIdxRefNode struct {
	ASTLHSNode
	expr  IASTExprNode
	index IASTExprNode
}

var _ IASTLHSNode = &ASTArrayIdxRefNode{}

func NewASTArrayIdxRefNode(expr IASTExprNode, index IASTExprNode) *ASTArrayIdxRefNode {
	p := &ASTArrayIdxRefNode{expr: expr, index: index}
	p.ASTLHSNode.ASTExprNode._impl = p
	p.ASTLHSNode.ASTExprNode.Node._impl = p
	return p
}

func (this *ASTArrayIdxRefNode) Expr() IASTExprNode {
	return this.expr
}

func (this *ASTArrayIdxRefNode) Index() IASTExprNode {
	return this.index
}

// isMultiDimension a[x][y][z] = true.
// isMultiDimension a[x][y] = true.
// isMultiDimension a[x] = false.
func (this *ASTArrayIdxRefNode) IsMultiDimension() bool {
	p, ok := this.expr.(*ASTArrayIdxRefNode)
	if !ok {
		return false
	}
	return !(p.OrigType().IsPointer())
}

// Returns base expression of (multi-dimension) array.
// e.g.  baseExpr of a[x][y][z] is a.
func (this *ASTArrayIdxRefNode) BaseExpr() IASTExprNode {
	if this.IsMultiDimension() {
		p := this.expr.(*ASTArrayIdxRefNode)
		return p.BaseExpr()
	}
	return this.expr
}

// element size of this (multi-dimension) array
func (this *ASTArrayIdxRefNode) ElementSize() int64 {
	return this.OrigType().AllocSize()
}

func (this *ASTArrayIdxRefNode) Length() int64 {
	origTy := (this.expr.OrigType()).(*ArrayType)
	return origTy.Length()
}

func (this *ASTArrayIdxRefNode) OrigType() IType {
	return this.expr.OrigType().ElemType()
}

func (this *ASTArrayIdxRefNode) Location() *Location {
	return this.expr.Location()
}

func (this *ASTArrayIdxRefNode) _Dump(d *Dumper) {
	if this.ty != nil {
		d.PrintMemberType("type", this.ty)
	}
	d.PrintMemberDumpable("expr", this.expr)
	d.PrintMemberDumpable("expr", this.index)
}

func (this *ASTArrayIdxRefNode) Accept(visitor IASTVisitor) interface{} {
	return visitor.VisitArrayIdxRefNode(this)
}
