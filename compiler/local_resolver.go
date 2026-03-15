package compiler

import (
	"cbc/models"
	"cbc/utils"
	"errors"
)

type LocalResolver struct {
	Visitor
	scopeStack    []models.IScope
	constantTable models.ConstantTable
	errorHandler  *utils.ErrorHandler
}

var _ IVisitor = &LocalResolver{}

func NewLocalResolver(h *utils.ErrorHandler) *LocalResolver {
	p := new(LocalResolver)
	p.errorHandler = h
	p._impl_visitor = p
	return p
}

func (this *LocalResolver) ResolveStmt(n models.IASTStmtNode) (any, error) {
	return n.Accept(this), nil
}

func (this *LocalResolver) ResolveExpr(n models.IASTExprNode) (any, error) {
	return n.Accept(this), nil
}

func (this *LocalResolver) Resolve(astObj *models.AST) error {
	toplevel := models.NewToplevelScope()
	this.scopeStack = append(this.scopeStack, toplevel)

	for _, decl := range astObj.Declarations() {
		err := toplevel.DeclareEntity(decl)
		if err != nil {
			this.errorHandler.ErrorWithLoc(decl.Location(), err.Error())
		}
	}
	for _, ent := range astObj.Definitions() {
		err := toplevel.DefineEntity(ent)
		if err != nil {
			this.errorHandler.ErrorWithLoc(ent.Location(), err.Error())
		}
	}

	this.ResolveGvarInitializers(astObj.DefinedVariables())
	this.ResolveConstantValues(astObj.Constants())
	this.ResolveFunctions(astObj.DefinedFunctions())

	toplevel.CheckReferences(this.errorHandler)

	if this.errorHandler.ErrorOccured() {
		return errors.New("semantic analyze local resolve failed")
	}

	err := astObj.SetScope(toplevel)
	if err != nil {
		return err
	}
	astObj.SetConstantTable(&this.constantTable)
	return nil
}

func (this *LocalResolver) ResolveGvarInitializers(gvars []*models.DefinedVariable) {
	for _, gvar := range gvars {
		if gvar.HasInitializer() {
			this.ResolveExpr(gvar.Initializer())
		}
	}
}

func (this *LocalResolver) ResolveConstantValues(consts []*models.Constant) {
	for _, c := range consts {
		this.ResolveExpr(c.Value())
	}
}

func (this *LocalResolver) ResolveFunctions(functions []*models.DefinedFunction) {
	for _, function := range functions {
		paras := make([]models.IDefinedVariable, len(function.Parameters()))
		for i, p := range function.Parameters() {
			paras[i] = p
		}
		this.PushScope(paras)
		this.ResolveStmt(function.Body())
		function.SetScope(this.PopScope())
	}
}

func (this *LocalResolver) CurrentScope() models.IScope {
	return this.scopeStack[len(this.scopeStack)-1]
}

func (this *LocalResolver) PushScope(vars []models.IDefinedVariable) {
	scope := models.NewLocalScope(this.CurrentScope())
	for _, v := range vars {
		err := scope.DefineVariable(v)
		if err != nil {
			this.errorHandler.ErrorWithLoc(v.Location(), err.Error())
		}
	}
	this.scopeStack = append(this.scopeStack, scope)
}

func (this *LocalResolver) PopScope() *models.LocalScope {
	scope := this.scopeStack[len(this.scopeStack)-1]
	this.scopeStack = this.scopeStack[:len(this.scopeStack)-1]
	return scope.(*models.LocalScope) // push pop pair, always return localScope, no need check cast
}

func (this *LocalResolver) VisitBlockNode(node *models.ASTBlockNode) any {
	vars := make([]models.IDefinedVariable, len(node.Variables()))
	for i, v := range node.Variables() {
		vars[i] = v
	}
	this.PushScope(vars)
	this.Visitor.VisitBlockNode(node)
	node.SetScope(this.PopScope())
	return nil
}

func (this *LocalResolver) VisitStringLiteralNode(node *models.ASTStringLiteralNode) any {
	node.SetEntry(this.constantTable.Intern(node.Value()))
	return nil
}

func (this *LocalResolver) VisitVariableNode(node *models.ASTVariableNode) any {
	ent, err := this.CurrentScope().Get(node.Name())
	if err != nil {
		this.errorHandler.ErrorWithLoc(node.Location(), err.Error())
		return nil
	}
	ent.Refered()
	node.SetEntity(ent)
	return nil
}
