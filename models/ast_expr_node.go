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
	_impl IASTExprNode
}

func (this *ASTExprNode) OrigType() IType {
	return this._impl.Type()
}

func (this *ASTExprNode) AllocSize() int64 {
	return this._impl.Type().AllocSize()
}

func (this *ASTExprNode) IsConstant() bool {
	return false
}

func (this *ASTExprNode) IsParameter() bool {
	return false
}

func (this *ASTExprNode) IsLvalue() bool {
	return false
}

func (this *ASTExprNode) IsAssignable() bool {
	return false
}

func (this *ASTExprNode) IsLoadable() bool {
	return false
}

// TODO: Java use try catch, when SemanticError return false
func (this *ASTExprNode) IsCallable() bool {
	return this._impl.Type().IsCallable()
}

// TODO: Java use try catch, when SemanticError return false
func (this *ASTExprNode) IsPointer() bool {
	return this._impl.Type().IsPointer()
}
