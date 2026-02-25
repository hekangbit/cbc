package models

type IASTStmtNode interface {
	INode
	Accept(visitor IASTVisitor) interface{}
}

type ASTStmtNode struct {
	Node
	location *Location
}

func (this *ASTStmtNode) Location() *Location {
	return this.location
}
