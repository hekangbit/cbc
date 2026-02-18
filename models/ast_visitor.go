package models

type ASTVisitor interface {
	VisitBlock(*ASTBlockNode) interface{}
	VisitBinaryOp(*ASTBinaryOpNode) interface{}
	VisitReturnNode(*ASTReturnNode) interface{}
	VisitExprStmtNode(*ASTExprStmtNode) interface{}
}
