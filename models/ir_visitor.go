package models

type IRVisitor interface {
	// Statements
	VisitExprStmt(*IRExprStmt) any
	VisitAssign(*IRAssign) any
	VisitCJump(*IRCJump) any
	VisitJump(*IRJump) any
	VisitSwitch(*IRSwitch) any
	VisitLabelStmt(*IRLabelStmt) any
	VisitReturn(*IRReturn) any

	// Expressions
	VisitUni(*IRUni) any
	VisitBin(*IRBin) any
	VisitCall(*IRCall) any
	VisitAddr(*IRAddr) any
	VisitMem(*IRMem) any
	VisitVar(*IRVar) any
	VisitInt(*IRInt) any
	VisitStr(*IRStr) any
}
