package ir

type IR struct {
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
