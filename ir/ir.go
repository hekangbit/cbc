package ast

type IR interface {
}

type IRExpr interface {
	IR
}

type IRStmt interface {
}

type IRAssign struct {
}

type IRCJump struct {
}

type IRJump struct {
}

type IRBin struct {
}
