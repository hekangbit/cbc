package models

type IASTLHSNode interface {
	IASTExprNode
	SetType(IType)
}

type ASTLHSNode struct {
	ASTExprNode
	ty     IType
	origTy IType
}

func (this *ASTLHSNode) SetType(t IType) {
	this.ty = t
}

func (this *ASTLHSNode) IsLvalue() bool {
	return true
}
