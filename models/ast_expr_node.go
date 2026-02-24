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
	IsLoadable() bool // TODO: make this method as global function
	IsCallable() bool
	IsPointer() bool
	Accept(visitor IASTVisitor) interface{}
}

type ASTExprNode struct {
	Node
}

// TODO: remove this, cause base struct is abstract class,
// no need implement all interface methods
var _ IASTExprNode = (*ASTExprNode)(nil)

func (n *ASTExprNode) Type() IType {
	panic("Type() must be implemented by concrete type")
}

func (n *ASTExprNode) OrigType() IType {
	panic("OrigType() must be implemented by concrete type")
}

func (n *ASTExprNode) AllocSize() int64 {
	panic("AllocSize() must be implemented by concrete type")
}

func (n *ASTExprNode) IsConstant() bool {
	return false
}

func (n *ASTExprNode) IsParameter() bool {
	return false
}

func (n *ASTExprNode) IsLvalue() bool {
	return false
}

func (n *ASTExprNode) IsAssignable() bool {
	return false
}

func (n *ASTExprNode) IsLoadable() bool {
	return false
}

func (n *ASTExprNode) IsCallable() bool {
	return false
}

func (n *ASTExprNode) IsPointer() bool {
	return false
}

func (n *ASTExprNode) Accept(visitor IASTVisitor) interface{} {
	panic("Accept() must be implemented by concrete type")
}
