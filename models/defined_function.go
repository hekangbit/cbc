package models

type DefinedFunction struct {
	*BaseFunction
	params *Params
	body   *ASTBlockNode
	scope  LocalScope
	ir     []IRStmt
}

func NewDefinedFunction(priv bool, typeNode *TypeNode, name string, params *Params, body *ASTBlockNode) *DefinedFunction {
	return &DefinedFunction{
		BaseFunction: NewBaseFunction(priv, typeNode, name),
		params:       params,
		body:         body,
	}
}
