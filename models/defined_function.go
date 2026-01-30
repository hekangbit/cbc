package models

type DefinedFunction struct {
	body   ASTBlockNode
	scope  LocalScope
	params Params
	ir     []IRStmt
}
