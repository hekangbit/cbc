package models

type IASTStmtNode interface {
	INode
	Location() *Location
	Accept(visitor IASTVisitor) interface{}
}

type ASTStmtNode struct {
	Node
	location *Location
}

func NewASTStmtNode(loc *Location) *ASTStmtNode {
	return &ASTStmtNode{
		location: loc,
	}
}

func (this *ASTStmtNode) Location() *Location {
	return this.location
}
