package models

type IDeclarationVisitor interface {
	VisitStructNode(*ASTStructNode) any
	VisitUnionNode(*ASTUnionNode) any
	VisitTypedefNode(*ASTTypedefNode) any
}
