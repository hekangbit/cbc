package models

type IDeclarationVisitor interface {
	VisitStructNode(*ASTStructNode) (any, error)
	VisitUnionNode(*ASTUnionNode) (any, error)
	VisitTypedefNode(*ASTTypedefNode) (any, error)
}
