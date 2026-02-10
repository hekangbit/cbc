package models

type DefinedFunction struct {
	*BaseFunction
	params *Params
	body   *ASTBlockNode
	scope  LocalScope
	ir     []IRStmt
}

func NewDefinedFunction(priv bool, typeNode *TypeNode, name string, params *Params, body *ASTBlockNode) *DefinedFunction {
	var p = new(DefinedFunction)
	p.isPrivate = priv
	p.params = params
	p.body = body
	return p
}
