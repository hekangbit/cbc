package models

type IDeclarationVisitor interface {
	VisitStructNode(*ASTStructNode) interface{}
	VisitUnionNode(*ASTUnionNode) interface{}
	VisitTypedefNode(*ASTTypedefNode) interface{}
}
