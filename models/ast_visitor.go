package models

type IASTVisitor interface {
	VisitBlock(*ASTBlockNode) interface{}
	VisitBinaryOp(*ASTBinaryOpNode) interface{}
	VisitReturnNode(*ASTReturnNode) interface{}
	VisitExprStmtNode(*ASTExprStmtNode) interface{}
	VisitAssignNode(*ASTAssignNode) interface{}
	VisitOpAssignNode(*ASTOpAssignNode) interface{}
	VisitCondExprNode(*ASTCondExprNode) interface{}
	VisitLogicalOrNode(*ASTLogicalOrNode) interface{}
	VisitLogicalAndNode(*ASTLogicalAndNode) interface{}
	VisitUnaryOpNode(*ASTUnaryOpNode) interface{}
	VisitPrefixOpNode(*ASTPrefixOpNode) interface{}
}
