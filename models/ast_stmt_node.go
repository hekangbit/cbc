package models

type IASTStmtNode interface {
	INode
	Accept(visitor IASTVisitor) (any, error)
}

type ASTStmtNode struct {
	Node
	location *Location
}

func (this *ASTStmtNode) Location() *Location {
	return this.location
}
