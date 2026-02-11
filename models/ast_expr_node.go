package models

// ExprNode 接口定义表达式节点的基本行为
type IASTExprNode interface {
	Node
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

// BaseExprNode 提供表达式节点的默认实现
type BaseASTExprNode struct {
	*BaseNode
}

var _ IASTExprNode = (*BaseASTExprNode)(nil)

// NewBaseExprNode 创建基础表达式节点
func NewBaseExprNode() *BaseASTExprNode {
	return &BaseASTExprNode{
		BaseNode: &BaseNode{},
	}
}

// Type 抽象方法，需要在具体实现中重写
func (n *BaseASTExprNode) Type() IType {
	panic("Type() must be implemented by concrete type")
}

// OrigType 返回原始类型，默认调用Type()
func (n *BaseASTExprNode) OrigType() IType {
	panic("OrigType() must be implemented by concrete type")
}

// AllocSize 返回分配大小
func (n *BaseASTExprNode) AllocSize() int64 {
	panic("AllocSize() must be implemented by concrete type")
}

// IsConstant 默认返回false
func (n *BaseASTExprNode) IsConstant() bool {
	return false
}

// IsParameter 默认返回false
func (n *BaseASTExprNode) IsParameter() bool {
	return false
}

// IsLvalue 默认返回false
func (n *BaseASTExprNode) IsLvalue() bool {
	return false
}

// IsAssignable 默认返回false
func (n *BaseASTExprNode) IsAssignable() bool {
	return false
}

// IsLoadable 默认返回false
func (n *BaseASTExprNode) IsLoadable() bool {
	return false
}

// IsCallable 检查是否可调用
func (n *BaseASTExprNode) IsCallable() bool {
	return false
}

// IsPointer 检查是否为指针类型
func (n *BaseASTExprNode) IsPointer() bool {
	return false
}

// Accept 抽象方法，需要在具体实现中重写
func (n *BaseASTExprNode) Accept(visitor ASTVisitor) interface{} {
	panic("Accept() must be implemented by concrete type")
}
