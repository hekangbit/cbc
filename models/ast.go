package models

type AST struct {
	name string
}

func NewAst(name string) *AST {
  return &AST{name: name}
}
