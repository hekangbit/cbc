package models

// StmtNode 接口定义
type IASTStmtNode interface {
	Node
	Location() *Location
	Accept(visitor ASTVisitor) interface{}
}

// StmtNodeBase 作为具体语句节点的嵌入基础结构
type BaseASTStmtNode struct {
	location *Location
}

// NewStmtNodeBase 创建基础语句节点
func NewBaseASTStmtNode(loc *Location) *BaseASTStmtNode {
	return &BaseASTStmtNode{
		location: loc,
	}
}

// Location 返回位置信息
func (s *BaseASTStmtNode) Location() *Location {
	return s.location
}
