package models

// StmtNode 接口定义
type IStmtNode interface {
	Node
	Location() *Location
	Accept(visitor ASTVisitor) interface{}
}

// StmtNodeBase 作为具体语句节点的嵌入基础结构
type BaseStmtNode struct {
	location *Location
}

// NewStmtNodeBase 创建基础语句节点
func NewBaseStmtNode(loc *Location) *BaseStmtNode {
	return &BaseStmtNode{
		location: loc,
	}
}

// Location 返回位置信息
func (s *BaseStmtNode) Location() *Location {
	return s.location
}
