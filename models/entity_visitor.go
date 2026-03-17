package models

// TODO
type IEntityVisitor interface {
	VisitDefinedVariable(*DefinedVariable) (any, error)
	VisitUndefinedVariable(*UndefinedVariable) (any, error)
	VisitDefinedFunction(*DefinedFunction) (any, error)
	VisitUndefinedFunction(*UndefinedFunction) (any, error)
	VisitConstant(*Constant) (any, error)
}
