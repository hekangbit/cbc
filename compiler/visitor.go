package compiler

import "cbc/models"

type IVisitor interface {
	models.IASTVisitor
	visitStmt(models.IASTStmtNode) error
	visitStmts([]models.IASTStmtNode) error
	visitExpr(models.IASTExprNode) error
	visitExprs([]models.IASTExprNode) error
}

type Visitor struct {
	_impl_visitor IVisitor
}

var _ IVisitor = &Visitor{}

func (this *Visitor) visitStmt(stmt models.IASTStmtNode) error {
	_, err := stmt.Accept(this._impl_visitor)
	return err
}

func (this *Visitor) visitStmts(stmts []models.IASTStmtNode) error {
	for _, stms := range stmts {
		this.visitStmt(stms)
	}
	return nil
}

func (this *Visitor) visitExpr(expr models.IASTExprNode) error {
	_, err := expr.Accept(this._impl_visitor)
	return err
}

func (this *Visitor) visitExprs(exprs []models.IASTExprNode) error {
	for _, expr := range exprs {
		this.visitExpr(expr)
	}
	return nil
}

// --- models.IASTVisitor default methods---

func (this *Visitor) VisitBlockNode(node *models.ASTBlockNode) (any, error) {
	for _, v := range node.Variables() {
		if v.HasInitializer() {
			this.visitExpr(v.Initializer())
		}
	}
	this.visitStmts(node.Stmts())
	return nil, nil
}

func (this *Visitor) VisitExprStmtNode(node *models.ASTExprStmtNode) (any, error) {
	this.visitExpr(node.Expr())
	return nil, nil
}

func (this *Visitor) VisitIfNode(n *models.ASTIfNode) (any, error) {
	this.visitExpr(n.Cond())
	this.visitStmt(n.ThenBody())
	if n.ElseBody() != nil {
		this.visitStmt(n.ElseBody())
	}
	return nil, nil
}

func (this *Visitor) VisitSwitchNode(n *models.ASTSwitchNode) (any, error) {
	this.visitExpr(n.Cond())
	buf := make([]models.IASTStmtNode, len(n.Cases()))
	for i, c := range n.Cases() {
		buf[i] = c
	}
	this.visitStmts(buf)
	return nil, nil
}

func (this *Visitor) VisitCaseNode(n *models.ASTCaseNode) (any, error) {
	this.visitExprs(n.Values())
	this.visitStmt(n.Body())
	return nil, nil
}

func (this *Visitor) VisitWhileNode(n *models.ASTWhileNode) (any, error) {
	this.visitExpr(n.Cond())
	this.visitStmt(n.Body())
	return nil, nil
}

func (this *Visitor) VisitDoWhileNode(n *models.ASTDoWhileNode) (any, error) {
	this.visitStmt(n.Body())
	this.visitExpr(n.Cond())
	return nil, nil
}

func (this *Visitor) VisitForNode(n *models.ASTForNode) (any, error) {
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
	return nil, nil
}

func (this *Visitor) VisitBreakNode(*models.ASTBreakNode) (any, error) {
	return nil, nil
}

func (this *Visitor) VisitContinueNode(*models.ASTContinueNode) (any, error) {
	return nil, nil
}

func (this *Visitor) VisitGotoNode(*models.ASTGotoNode) (any, error) {
	return nil, nil
}

func (this *Visitor) VisitLabelNode(n *models.ASTLabelNode) (any, error) {
	this.visitStmt(n.Stmt())
	return nil, nil
}

func (this *Visitor) VisitReturnNode(n *models.ASTReturnNode) (any, error) {
	if n.Expr() != nil {
		this.visitExpr(n.Expr())
	}
	return nil, nil
}

func (this *Visitor) VisitCondExprNode(n *models.ASTCondExprNode) (any, error) {
	this.visitExpr(n.Cond())
	this.visitExpr(n.ThenExpr())
	if n.ElseExpr() != nil {
		this.visitExpr(n.ElseExpr())
	}
	return nil, nil
}

func (this *Visitor) VisitLogicalOrNode(node *models.ASTLogicalOrNode) (any, error) {
	this.visitExpr(node.Left())
	this.visitExpr(node.Right())
	return nil, nil
}

func (this *Visitor) VisitLogicalAndNode(node *models.ASTLogicalAndNode) (any, error) {
	this.visitExpr(node.Left())
	this.visitExpr(node.Right())
	return nil, nil
}

func (this *Visitor) VisitAssignNode(n *models.ASTAssignNode) (any, error) {
	this.visitExpr(n.LHS())
	this.visitExpr(n.RHS())
	return nil, nil
}

func (this *Visitor) VisitOpAssignNode(n *models.ASTOpAssignNode) (any, error) {
	this.visitExpr(n.LHS())
	this.visitExpr(n.RHS())
	return nil, nil
}

func (this *Visitor) VisitBinaryOpNode(n *models.ASTBinaryOpNode) (any, error) {
	this.visitExpr(n.Left())
	this.visitExpr(n.Right())
	return nil, nil
}

func (this *Visitor) VisitUnaryOpNode(node *models.ASTUnaryOpNode) (any, error) {
	this.visitExpr(node.Expr())
	return nil, nil
}

func (this *Visitor) VisitPrefixOpNode(node *models.ASTPrefixOpNode) (any, error) {
	this.visitExpr(node.Expr())
	return nil, nil
}

func (this *Visitor) VisitSuffixOpNode(node *models.ASTSuffixOpNode) (any, error) {
	this.visitExpr(node.Expr())
	return nil, nil
}

func (this *Visitor) VisitFunctionCallNode(node *models.ASTFunctionCallNode) (any, error) {
	this.visitExpr(node.Expr())
	this.visitExprs(node.Args())
	return nil, nil
}

func (this *Visitor) VisitArrayIdxRefNode(node *models.ASTArrayIdxRefNode) (any, error) {
	this.visitExpr(node.Expr())
	this.visitExpr(node.Index())
	return nil, nil
}

func (this *Visitor) VisitMemberNode(node *models.ASTMemberNode) (any, error) {
	this.visitExpr(node.Expr())
	return nil, nil
}

func (this *Visitor) VisitPtrMemberNode(node *models.ASTPtrMemberNode) (any, error) {
	this.visitExpr(node.Expr())
	return nil, nil
}

func (this *Visitor) VisitDereferenceNode(node *models.ASTDereferenceNode) (any, error) {
	this.visitExpr(node.Expr())
	return nil, nil
}

func (this *Visitor) VisitAddressNode(node *models.ASTAddressNode) (any, error) {
	this.visitExpr(node.Expr())
	return nil, nil
}

func (this *Visitor) VisitCastNode(node *models.ASTCastNode) (any, error) {
	this.visitExpr(node.Expr())
	return nil, nil
}

func (this *Visitor) VisitSizeofExprNode(node *models.ASTSizeofExprNode) (any, error) {
	this.visitExpr(node.Expr())
	return nil, nil
}

func (this *Visitor) VisitSizeofTypeNode(*models.ASTSizeofTypeNode) (any, error) {
	return nil, nil
}

func (this *Visitor) VisitVariableNode(*models.ASTVariableNode) (any, error) {
	return nil, nil
}

func (this *Visitor) VisitIntegerLiteralNode(*models.ASTIntegerLiteralNode) (any, error) {
	return nil, nil
}

func (this *Visitor) VisitStringLiteralNode(*models.ASTStringLiteralNode) (any, error) {
	return nil, nil
}
