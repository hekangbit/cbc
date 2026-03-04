package models

// TODO
type IEntityVisitor interface {
	VisitDefinedVariable(*DefinedVariable) any
	VisitUndefinedVariable(*UndefinedVariable) any
	VisitDefinedFunction(*DefinedFunction) any
	VisitUndefinedFunction(*UndefinedFunction) any
	VisitConstant(*Constant) any
}
