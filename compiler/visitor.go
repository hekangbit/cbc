package compiler

import "cbc/models"

type IVisitor interface {
	models.IASTVisitor
}

type Visitor struct {
	_impl_visitor IVisitor
}

var _ IVisitor = &Visitor{}

func (this *Visitor) visitStmt(stmt models.IASTStmtNode) {
	stmt.Accept(this._impl_visitor)
}

func (this *Visitor) visitStmts(stmts []models.IASTStmtNode) {
	for _, stms := range stmts {
		this.visitStmt(stms)
	}
}

func (this *Visitor) visitExpr(expr models.IASTExprNode) {
	expr.Accept(this._impl_visitor)
}

func (this *Visitor) visitExprs(exprs []models.IASTExprNode) {
	for _, expr := range exprs {
		this.visitExpr(expr)
	}
}

// --- ASTVisitor default methods---

func (this *Visitor) VisitBlockNode(node *models.ASTBlockNode) interface{} {
	for _, v := range node.Variables() {
		if v.HasInitializer() {
			this.visitExpr(v.Initializer())
		}
	}
	this.visitStmts(node.Stmts())
	return nil
}

func (this *Visitor) VisitExprStmtNode(node *models.ASTExprStmtNode) interface{} {
	this.visitExpr(node.Expr())
	return nil
}

func (this *Visitor) VisitIfNode(n *models.ASTIfNode) interface{} {
	this.visitExpr(n.Cond())
	this.visitStmt(n.ThenBody())
	if n.ElseBody() != nil {
		this.visitStmt(n.ElseBody())
	}
	return nil
}

func (this *Visitor) VisitSwitchNode(n *models.ASTSwitchNode) interface{} {
	this.visitExpr(n.Cond())
	buf := make([]models.IASTStmtNode, len(n.Cases()))
	for i, c := range n.Cases() {
		buf[i] = c
	}
	this.visitStmts(buf)
	return nil
}

func (this *Visitor) VisitCaseNode(n *models.ASTCaseNode) interface{} {
	this.visitExprs(n.Values())
	this.visitStmt(n.Body())
	return nil
}

func (this *Visitor) VisitWhileNode(n *models.ASTWhileNode) interface{} {
	this.visitExpr(n.Cond())
	this.visitStmt(n.Body())
	return nil
}

func (this *Visitor) VisitDoWhileNode(n *models.ASTDoWhileNode) interface{} {
	this.visitStmt(n.Body())
	this.visitExpr(n.Cond())
	return nil
}

func (this *Visitor) VisitForNode(n *models.ASTForNode) interface{} {
	if n.Init() != nil {
		this.visitStmt(n.Init())
	}
	if n.Cond() != nil {
		this.visitExpr(n.Cond())
	}
	if n.Incr() != nil {
		this.visitStmt(n.Incr())
	}
	this.visitStmt(n.Body())
	return nil
}

func (this *Visitor) VisitBreakNode(*models.ASTBreakNode) interface{}       { return nil }
func (this *Visitor) VisitContinueNode(*models.ASTContinueNode) interface{} { return nil }
func (this *Visitor) VisitGotoNode(*models.ASTGotoNode) interface{}         { return nil }

func (this *Visitor) VisitLabelNode(n *models.ASTLabelNode) interface{} {
	this.visitStmt(n.Stmt())
	return nil
}

func (this *Visitor) VisitReturnNode(n *models.ASTReturnNode) interface{} {
	if n.Expr() != nil {
		this.visitExpr(n.Expr())
	}
	return nil
}

func (this *Visitor) VisitCondExprNode(n *models.ASTCondExprNode) interface{} {
	this.visitExpr(n.Cond())
	this.visitExpr(n.ThenExpr())
	if n.ElseExpr() != nil {
		this.visitExpr(n.ElseExpr())
	}
	return nil
}

func (this *Visitor) VisitLogicalOrNode(node *models.ASTLogicalOrNode) interface{} {
	this.visitExpr(node.Left())
	this.visitExpr(node.Right())
	return nil
}

func (this *Visitor) VisitLogicalAndNode(node *models.ASTLogicalAndNode) interface{} {
	this.visitExpr(node.Left())
	this.visitExpr(node.Right())
	return nil
}

func (this *Visitor) VisitAssignNode(n *models.ASTAssignNode) interface{} {
	this.visitExpr(n.LHS())
	this.visitExpr(n.RHS())
	return nil
}

func (this *Visitor) VisitOpAssignNode(n *models.ASTOpAssignNode) interface{} {
	this.visitExpr(n.LHS())
	this.visitExpr(n.RHS())
	return nil
}

func (this *Visitor) VisitBinaryOpNode(n *models.ASTBinaryOpNode) interface{} {
	this.visitExpr(n.Left())
	this.visitExpr(n.Right())
	return nil
}

func (this *Visitor) VisitUnaryOpNode(node *models.ASTUnaryOpNode) interface{} {
	this.visitExpr(node.Expr())
	return nil
}

func (this *Visitor) VisitPrefixOpNode(node *models.ASTPrefixOpNode) interface{} {
	this.visitExpr(node.Expr())
	return nil
}

func (this *Visitor) VisitSuffixOpNode(node *models.ASTSuffixOpNode) interface{} {
	this.visitExpr(node.Expr())
	return nil
}

func (this *Visitor) VisitFunctionCallNode(node *models.ASTFunctionCallNode) interface{} {
	this.visitExpr(node.Expr())
	this.visitExprs(node.Args())
	return nil
}

func (this *Visitor) VisitArrayIdxRefNode(node *models.ASTArrayIdxRefNode) interface{} {
	this.visitExpr(node.Expr())
	this.visitExpr(node.Index())
	return nil
}

func (this *Visitor) VisitMemberNode(node *models.ASTMemberNode) interface{} {
	this.visitExpr(node.Expr())
	return nil
}

func (this *Visitor) VisitPtrMemberNode(node *models.ASTPtrMemberNode) interface{} {
	this.visitExpr(node.Expr())
	return nil
}

func (this *Visitor) VisitDereferenceNode(node *models.ASTDereferenceNode) interface{} {
	this.visitExpr(node.Expr())
	return nil
}

func (this *Visitor) VisitAddressNode(node *models.ASTAddressNode) interface{} {
	this.visitExpr(node.Expr())
	return nil
}

func (this *Visitor) VisitCastNode(node *models.ASTCastNode) interface{} {
	this.visitExpr(node.Expr())
	return nil
}

func (this *Visitor) VisitSizeofExprNode(node *models.ASTSizeofExprNode) interface{} {
	this.visitExpr(node.Expr())
	return nil
}

func (this *Visitor) VisitSizeofTypeNode(*models.ASTSizeofTypeNode) interface{}         { return nil }
func (this *Visitor) VisitVariableNode(*models.ASTVariableNode) interface{}             { return nil }
func (this *Visitor) VisitIntegerLiteralNode(*models.ASTIntegerLiteralNode) interface{} { return nil }
func (this *Visitor) VisitStringLiteralNode(*models.ASTStringLiteralNode) interface{}   { return nil }
