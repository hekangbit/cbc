package models

type AST struct {
	name  string
	decls *Declarations
}

func NewAst(name string, decls *Declarations) *AST {
	return &AST{name: name, decls: decls}
}
