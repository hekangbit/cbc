package compiler

import (
	"cbc/models"
	"cbc/utils"
	"errors"
)

type DereferenceChecker struct {
	Visitor
	typeTable    *models.TypeTable
	errorHandler *utils.ErrorHandler
}

var _ IVisitor = &DereferenceChecker{}

func NewDereferenceChecker(typeTable *models.TypeTable, h *utils.ErrorHandler) *DereferenceChecker {
	p := new(DereferenceChecker)
	p.typeTable = typeTable
	p.errorHandler = h
	p._impl_visitor = p
	return p
}

func (this *DereferenceChecker) semanticError(loc *models.Location, msg string) error {
	this.errorHandler.ErrorWithLoc(loc, msg)
	return errors.New("invalid expr") // TODO: throw new SemanticError("invalid expr");
}

func (this *DereferenceChecker) undereferableError(loc *models.Location) error {
	return this.semanticError(loc, "dereferencing non-pointer expression")
}

func (this *DereferenceChecker) Check(astObj *models.AST) error {
	for _, v := range astObj.DefinedVariables() {
		this.CheckToplevelVariable(v)
	}

	for _, f := range astObj.DefinedFunctions() {
		this.CheckStmtNode(f.Body())
	}

	if this.errorHandler.ErrorOccured() {
		return errors.New("semantic analyze dereference checker failed")
	}

	return nil
}

func (this *DereferenceChecker) handleImplicitAddress(node models.IASTLHSNode) {
	if !node.IsLoadable() {
		t := node.Type()
		if t.IsArray() {
			// int[4] ==> int*
			node.SetType(this.typeTable.PointerTo(t.ElemType()))
		} else {
			// void(int, int) ==> void(int, int)*
			node.SetType(this.typeTable.PointerTo(t))
		}
	}
}

func (this *DereferenceChecker) CheckExprNode(node models.IASTExprNode) error {
	_, err := node.Accept(this)
	return err
}

func (this *DereferenceChecker) CheckStmtNode(node models.IASTStmtNode) error {
	_, err := node.Accept(this)
	return err
}

func (this *DereferenceChecker) CheckToplevelVariable(v *models.DefinedVariable) {
	this.CheckVariable(v)
	if v.HasInitializer() {
		this.CheckConstant(v.Initializer())
	}
}

func (this *DereferenceChecker) CheckVariable(v *models.DefinedVariable) {
	if v.HasInitializer() {
		this.CheckExprNode(v.Initializer())
	}
}

func (this *DereferenceChecker) CheckConstant(expr models.IASTExprNode) {
	if !expr.IsConstant() {
		this.errorHandler.ErrorWithLoc(expr.Location(), "not a constant")
	}
}

func (this *DereferenceChecker) CheckAssignment(node models.IASTAbstractAssignNode) error {
	if !(node.LHS().IsAssignable()) {
		return this.semanticError(node.Location(), "invalid lhs expression")
	}
	return nil
}

func (this *DereferenceChecker) CheckMemberRef(loc *models.Location, t models.IType, memb string) error {
	if !t.IsCompositeType() {
		return this.semanticError(loc, "accessing member `"+memb+"' for non-struct/union: "+t.String())
	}
	if !t.GetCompositeType().HasMember(memb) {
		return this.semanticError(loc, t.String()+" does not have member: "+memb)
	}
	return nil
}

// ---- Implement Visitor ----

func (this *DereferenceChecker) VisitBlockNode(node *models.ASTBlockNode) (any, error) {
	for _, v := range node.Variables() {
		this.CheckVariable(v)
	}
	for _, stmt := range node.Stmts() {
		this.CheckStmtNode(stmt)
	}
	return nil, nil
}

// &a
func (this *DereferenceChecker) VisitAddressNode(node *models.ASTAddressNode) (any, error) {
	var err error
	_, err = this.Visitor.VisitAddressNode(node)
	if err != nil {
		return nil, err
	}
	if !node.Expr().IsLvalue() {
		return nil, this.semanticError(node.Location(), "invalid expression for &")
	}
	baseTy := node.Expr().Type()
	if !node.Expr().IsLoadable() {
		node.SetType(baseTy)
	} else {
		node.SetType(this.typeTable.PointerTo(baseTy))
	}
	return nil, nil
}

// *a
func (this *DereferenceChecker) VisitDereferenceNode(node *models.ASTDereferenceNode) (any, error) {
	var err error
	_, err = this.Visitor.VisitDereferenceNode(node)
	if err != nil {
		return nil, err
	}
	if !node.Expr().IsPointer() {
		return this.undereferableError(node.Location()), nil
	}

	// case 1:
	// int[4]* p; p is pointer, *p is dereferenceNode
	// Type() is int[4], then will set type of ASTDereferenceNode as int*
	// case 2:
	// void(int,int)* p; p is pointer, *p is dereferenceNode
	// Type() is void(int,int), then will set type of ASTDereferenceNode as void(int,int)*
	this.handleImplicitAddress(node)
	return nil, nil
}

func (this *DereferenceChecker) VisitAssignNode(node *models.ASTAssignNode) (any, error) {
	var err error
	_, err = this.Visitor.VisitAssignNode(node)
	if err != nil {
		return nil, err
	}
	err = this.CheckAssignment(node)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (this *DereferenceChecker) VisitOpAssignNode(node *models.ASTOpAssignNode) (any, error) {
	var err error
	_, err = this.Visitor.VisitOpAssignNode(node)
	if err != nil {
		return nil, err
	}
	err = this.CheckAssignment(node)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (this *DereferenceChecker) VisitPrefixOpNode(node *models.ASTPrefixOpNode) (any, error) {
	var err error
	_, err = this.Visitor.VisitPrefixOpNode(node)
	if err != nil {
		return nil, err
	}
	if !(node.Expr().IsAssignable()) {
		return nil, this.semanticError(node.Location(), "cannot prefix increment/decrement")
	}
	return nil, nil
}

func (this *DereferenceChecker) VisitSuffixOpNode(node *models.ASTSuffixOpNode) (any, error) {
	var err error
	_, err = this.Visitor.VisitSuffixOpNode(node)
	if err != nil {
		return nil, err
	}
	if !(node.Expr().IsAssignable()) {
		return nil, this.semanticError(node.Location(), "cannot suffix increment/decrement")
	}
	return nil, nil
}

func (this *DereferenceChecker) VisitMemberNode(node *models.ASTMemberNode) (any, error) {
	var err error
	_, err = this.Visitor.VisitMemberNode(node)
	if err != nil {
		return nil, err
	}
	err = this.CheckMemberRef(node.Location(), node.Expr().Type(), node.Member())
	if err != nil {
		return nil, err
	}
	this.handleImplicitAddress(node)
	return nil, nil
}

func (this *DereferenceChecker) VisitPtrMemberNode(node *models.ASTPtrMemberNode) (any, error) {
	var err error
	_, err = this.Visitor.VisitPtrMemberNode(node)
	if err != nil {
		return nil, err
	}
	err = this.CheckMemberRef(node.Location(), node.DereferedType(), node.Member())
	if err != nil {
		return nil, err
	}
	this.handleImplicitAddress(node)
	return nil, nil
}

func (this *DereferenceChecker) VisitFunctionCallNode(node *models.ASTFunctionCallNode) (any, error) {
	var err error
	_, err = this.Visitor.VisitFunctionCallNode(node)
	if err != nil {
		return nil, err
	}
	if node.Expr().IsCallable() {
		return nil, this.semanticError(node.Location(), "calling object is not a function")
	}
	return nil, nil
}

func (this *DereferenceChecker) VisitArrayIdxRefNode(node *models.ASTArrayIdxRefNode) (any, error) {
	var err error
	_, err = this.Visitor.VisitArrayIdxRefNode(node)
	if err != nil {
		return nil, err
	}
	if !node.Expr().IsPointer() {
		return nil, this.semanticError(node.Location(), "indexing non-array/pointer expression")
	}
	this.handleImplicitAddress(node)
	return nil, nil
}

func (this *DereferenceChecker) VisitVariableNode(node *models.ASTVariableNode) (any, error) {
	var err error
	_, err = this.Visitor.VisitVariableNode(node)
	if err != nil {
		return nil, err
	}
	if node.Entity().IsConstant() {
		this.CheckConstant(node.Entity().Value())
	}
	this.handleImplicitAddress(node)
	return nil, nil
}

func (this *DereferenceChecker) VisitCastNode(node *models.ASTCastNode) (any, error) {
	var err error
	_, err = this.Visitor.VisitCastNode(node)
	if err != nil {
		return nil, err
	}
	if node.Type().IsArray() {
		return nil, this.semanticError(node.Location(), "cast specifies array type")
	}
	return nil, nil
}
