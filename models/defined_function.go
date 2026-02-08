package models

type DefinedFunction struct {
	body   ASTBlockNode
	scope  LocalScope
	params Params
	ir     []IRStmt
}

func NewDefinedFunction(prov bool, typeNode *TypeNode, name string) *DefinedFunction {
	var p = new(DefinedFunction)
	return p
}
