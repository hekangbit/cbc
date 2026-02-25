package models

// TODO
type IEntityVisitor interface {
	VisitDefinedVariable(*DefinedVariable) interface{}
	VisitUndefinedVariable(*UndefinedVariable) interface{}
	VisitDefinedFunction(*DefinedFunction) interface{}
	VisitUndefinedFunction(*UndefinedFunction) interface{}
	VisitConstant(*Constant) interface{}
}
