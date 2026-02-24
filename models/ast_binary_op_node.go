package models

type ASTBinaryOpNode struct {
	ASTExprNode
	operator string
	left     IASTExprNode
	right    IASTExprNode
	type_    IType
}

var _ IASTExprNode = &ASTBinaryOpNode{}

func NewASTBinaryOpNode(left IASTExprNode, op string, right IASTExprNode) *ASTBinaryOpNode {
	return &ASTBinaryOpNode{
		operator: op,
		left:     left,
		right:    right,
		type_:    nil,
	}
}

func NewASTBinaryOpNodeWithType(ty IType, left IASTExprNode, op string, right IASTExprNode) *ASTBinaryOpNode {
	return &ASTBinaryOpNode{
		operator: op,
		left:     left,
		right:    right,
		type_:    ty,
	}
}

func (this *ASTBinaryOpNode) Operator() string {
	return this.operator
}

func (this *ASTBinaryOpNode) Type() IType {
	if this.type_ == nil {
		return this.left.Type()
	}
	return this.type_
}

func (this *ASTBinaryOpNode) SetType(ty IType) {
	if this.type_ != nil {
		panic("BinaryOp#setType called twice")
	}
	this.type_ = ty
}

func (this *ASTBinaryOpNode) Left() IASTExprNode {
	return this.left
}

func (this *ASTBinaryOpNode) SetLeft(left IASTExprNode) {
	this.left = left
}

func (this *ASTBinaryOpNode) Right() IASTExprNode {
	return this.right
}

func (this *ASTBinaryOpNode) SetRight(right IASTExprNode) {
	this.right = right
}

func (this *ASTBinaryOpNode) Location() *Location {
	return this.left.Location()
}

func (this *ASTBinaryOpNode) Dump(d *Dumper) {
	d.PrintMemberString("operator", this.operator, false)
	d.PrintMemberDumpable("left", this.left)
	d.PrintMemberDumpable("right", this.right)
}

func (this *ASTBinaryOpNode) Accept(visitor IASTVisitor) interface{} {
	return visitor.VisitBinaryOp(this)
}
