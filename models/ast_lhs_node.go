package models

type IASTLHSNode interface {
	IASTExprNode
	SetType(IType)
}

type ASTLHSNode struct {
	ASTExprNode
	ty     IType
	origTy IType // TODO: seem can remove this field
}

func (this *ASTLHSNode) Type() IType {
	if this.ty != nil {
		return this.ty
	}
	return this._impl.OrigType()
}

func (this *ASTLHSNode) OrigType() IType {
	panic("ASTLHSNode#OrigType must implemented by concreate lhs node")
}

func (this *ASTLHSNode) SetType(t IType) {
	this.ty = t
}

func (this *ASTLHSNode) IsLvalue() bool {
	return true
}

func (this *ASTLHSNode) IsAssignable() bool {
	return this._impl.IsLoadable()
}

func (this *ASTLHSNode) IsLoadable() bool {
	t := this._impl.OrigType()
	return !t.IsArray() && !t.IsFunction()
}
