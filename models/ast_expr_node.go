package models

type IASTExprNode interface {
	INode
	Type() IType
	OrigType() IType
	AllocSize() int64
	IsConstant() bool
	IsParameter() bool
	IsLvalue() bool
	IsAssignable() bool
	IsLoadable() bool
	IsCallable() bool
	IsPointer() bool
	Accept(visitor ASTVisitor) interface{}
}

type ASTBaseExprNode struct {
	*Node
}

var _ IASTExprNode = (*ASTBaseExprNode)(nil)

func NewBaseExprNode() *ASTBaseExprNode {
	return &ASTBaseExprNode{
		Node: &Node{},
	}
}

func (n *ASTBaseExprNode) Type() IType {
	panic("Type() must be implemented by concrete type")
}

func (n *ASTBaseExprNode) OrigType() IType {
	panic("OrigType() must be implemented by concrete type")
}

func (n *ASTBaseExprNode) AllocSize() int64 {
	panic("AllocSize() must be implemented by concrete type")
}

func (n *ASTBaseExprNode) IsConstant() bool {
	return false
}

func (n *ASTBaseExprNode) IsParameter() bool {
	return false
}

func (n *ASTBaseExprNode) IsLvalue() bool {
	return false
}

func (n *ASTBaseExprNode) IsAssignable() bool {
	return false
}

func (n *ASTBaseExprNode) IsLoadable() bool {
	return false
}

func (n *ASTBaseExprNode) IsCallable() bool {
	return false
}

func (n *ASTBaseExprNode) IsPointer() bool {
	return false
}

func (n *ASTBaseExprNode) Accept(visitor ASTVisitor) interface{} {
	panic("Accept() must be implemented by concrete type")
}
