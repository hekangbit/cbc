package models

// ASTVisitor 接口（需要根据实际情况定义）
type ASTVisitor interface {
	// 具体的访问方法...
	VisitBlock(*ASTBlockNode) interface{}
	VisitBinaryOp(*ASTBinaryOpNode) interface{}
	VisitReturnNode(*ASTReturnNode) interface{}
}
