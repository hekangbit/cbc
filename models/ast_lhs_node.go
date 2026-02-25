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

// TODO: can improve? currently use below method to shadow expr node OrigType method
// need runtime panic to detect error
func (this *ASTLHSNode) OrigType() IType {
	panic("ASTLHSNode#OrigType must implemented by concreate lhs node")
}

func (this *ASTLHSNode) Type() IType {
	if this.ty != nil {
		return this.ty
	}
	return this.ASTExprNode._impl.OrigType()
}

func (this *ASTLHSNode) SetType(t IType) {
	this.ty = t
}

func (this *ASTLHSNode) IsLvalue() bool {
	return true
}
