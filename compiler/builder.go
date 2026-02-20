package compiler

import (
	"cbc/loader"
	"cbc/models"
	"cbc/parser"

	"github.com/antlr4-go/antlr/v4"
)

type ASTBuilder struct {
	*parser.BaseCbVisitor
	name        string
	sourcePath  string
	curBaseType models.ITypeRef
}

var _ parser.CbVisitor = (*ASTBuilder)(nil)

func (this *ASTBuilder) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(this)
}

func (this *ASTBuilder) VisitChildren(_ antlr.RuleNode) interface{} {
	return nil
}

func (this *ASTBuilder) VisitTerminal(_ antlr.TerminalNode) interface{} {
	return nil
}

func (this *ASTBuilder) VisitErrorNode(_ antlr.ErrorNode) interface{} {
	return nil
}

func (this *ASTBuilder) VisitProg(ctx *parser.ProgContext) interface{} {
	ctx.ImportStmts().Accept(this)
	decls := ctx.TopDefs().Accept(this).(*models.Declarations)
	return models.NewAST(models.NewLocation(this.sourcePath, ctx.GetStart()), decls)
}

func (this *ASTBuilder) VisitImportStmts(ctx *parser.ImportStmtsContext) interface{} {
	for _, importStmt := range ctx.AllImportStmt() {
		importStmt.Accept(this)
	}
	return this.VisitChildren(ctx)
}

func (this *ASTBuilder) VisitImportStmt(ctx *parser.ImportStmtContext) interface{} {
	path := ctx.Identifier(0).GetText()
	for i := 1; i < len(ctx.AllIdentifier()); i++ {
		path = path + "." + ctx.Identifier(i).GetText()
	}
	loader.LoadLibrary(path)
	return this.VisitChildren(ctx)
}

func (this *ASTBuilder) VisitTopDefs(ctx *parser.TopDefsContext) interface{} {
	decls := models.NewDeclarations()
	for _, defVars := range ctx.AllDefVars() {
		defvars := defVars.Accept(this)
		decls.AddDefvars(defvars.([]*models.DefinedVariable))
	}
	for _, defFun := range ctx.AllDefFunc() {
		f := defFun.Accept(this)
		decls.AddDeffun(f.(*models.DefinedFunction))
	}
	return decls
}

func (this *ASTBuilder) VisitDefVars(ctx *parser.DefVarsContext) interface{} {
	var initialize models.IASTExprNode = nil
	var defs []*models.DefinedVariable
	priv := false
	if ctx.GetPriv() != nil {
		priv = true
	}
	cbType := ctx.GetCbtype().Accept(this).(*models.TypeNode)

	for _, identifier := range ctx.AllIdentifier() {
		initialize = nil
		if ctx.GetHasInit() != nil {
			initialize = ctx.GetInitializer().Accept(this).(models.IASTExprNode)
		}
		dv := models.NewDefinedVariable(priv, cbType, identifier.GetSymbol().GetText(), initialize)
		defs = append(defs, dv)
	}

	return defs
}

func (this *ASTBuilder) VisitDefFunc(ctx *parser.DefFuncContext) interface{} {
	priv := ctx.GetPriv() != nil
	retTypeRef := ctx.GetRetCbtype().Accept(this).(models.ITypeRef)
	name := ctx.Identifier().GetSymbol().GetText()
	params := ctx.Params().Accept(this).(*models.Params)
	body := ctx.Block().Accept(this).(*models.ASTBlockNode)
	funcTypeRef := models.NewFunctionTypeRef(retTypeRef, params.ParametersTypeRef())
	funcTypeNode := models.NewTypeNodeFromRef(funcTypeRef)
	return models.NewDefinedFunction(priv, funcTypeNode, name, params, body)
}

func (this *ASTBuilder) VisitCbType(ctx *parser.CbTypeContext) interface{} {
	typeRef := ctx.CbTypeRef().Accept(this).(models.ITypeRef)
	return models.NewTypeNodeFromRef(typeRef)
}

func (this *ASTBuilder) VisitCbTypeRef(ctx *parser.CbTypeRefContext) interface{} {
	this.curBaseType = ctx.CbTypeRefBase().Accept(this).(models.ITypeRef)
	modifiers := ctx.AllTypeModifier()
	// TODO: add more modifier (sizedArray, FunctionType)
	for i := len(modifiers) - 1; i >= 0; i-- {
		this.curBaseType = modifiers[i].Accept(this).(models.ITypeRef)
	}
	return this.curBaseType
}

func (this *ASTBuilder) VisitVoidTypeRef(ctx *parser.VoidTypeRefContext) interface{} {
	return models.NewVoidTypeRefWithLocation(models.NewLocation(this.sourcePath, ctx.GetStart()))
}

func (this *ASTBuilder) VisitCharTypeRef(ctx *parser.CharTypeRefContext) interface{} {
	return models.NewCharRefWithLocation(models.NewLocation(this.sourcePath, ctx.GetStart()))
}

func (this *ASTBuilder) VisitShortTypeRef(ctx *parser.ShortTypeRefContext) interface{} {
	return models.NewShortRefWithLocation(models.NewLocation(this.sourcePath, ctx.GetStart()))
}

func (this *ASTBuilder) VisitIntTypeRef(ctx *parser.IntTypeRefContext) interface{} {
	return models.NewIntRefWithLocation(models.NewLocation(this.sourcePath, ctx.GetStart()))
}

func (this *ASTBuilder) VisitArrayModifier(ctx *parser.ArrayModifierContext) interface{} {
	return models.NewArrayTypeRef(this.curBaseType)
}

func (this *ASTBuilder) VisitPointerModifier(ctx *parser.PointerModifierContext) interface{} {
	return models.NewPointerTypeRef(this.curBaseType)
}

func (this *ASTBuilder) VisitParams(ctx *parser.ParamsContext) interface{} {
	voidToken := ctx.GetVoid()
	if voidToken != nil {
		paramDescs := make([]*models.CBCParameter, 0)
		return models.NewParams(models.NewLocation(this.sourcePath, voidToken), paramDescs)
	}
	fixedParams := ctx.FixedParams().Accept(this).([]*models.CBCParameter)
	fullParams := models.NewParams(models.NewLocation(this.sourcePath, ctx.GetStart()), fixedParams)
	if ctx.GetHasVararg() != nil {
		fullParams.AcceptVarargs()
	}
	return fullParams
}

func (this *ASTBuilder) VisitFixedParams(ctx *parser.FixedParamsContext) interface{} {
	params := make([]*models.CBCParameter, 0)
	for _, paramCtx := range ctx.AllParam() {
		param := paramCtx.Accept(this).(*models.CBCParameter)
		params = append(params, param)
	}
	return params
}

func (this *ASTBuilder) VisitParam(ctx *parser.ParamContext) interface{} {
	typeNode := ctx.CbType().Accept(this).(*models.TypeNode)
	name := ctx.Identifier().GetSymbol().GetText()
	return models.NewCBCParameter(typeNode, name)
}

func (this *ASTBuilder) VisitBlock(ctx *parser.BlockContext) interface{} {
	defLocalVars := make([]*models.DefinedVariable, 0)
	for _, defVarsCtx := range ctx.AllDefVars() {
		vars := defVarsCtx.Accept(this).([]*models.DefinedVariable)
		defLocalVars = append(defLocalVars, vars...)
	}

	stmts := make([]models.IASTStmtNode, 0)
	for _, stmtsCtx := range ctx.AllStmt() {
		stmt := stmtsCtx.Accept(this).(models.IASTStmtNode)
		stmts = append(stmts, stmt)
	}

	return models.NewASTBlockNode(models.NewLocation(this.sourcePath, ctx.GetStart()), defLocalVars, stmts)
}

func (this *ASTBuilder) VisitExprStatement(ctx *parser.ExprStatementContext) interface{} {
	expr := ctx.Expr().Accept(this).(models.IASTExprNode)
	return models.NewASTExprStmtNode(models.NewLocation(this.sourcePath, ctx.GetStart()), expr)
}

func (this *ASTBuilder) VisitBlockStatement(ctx *parser.BlockStatementContext) interface{} {
	return ctx.Block().Accept(this)
}

func (this *ASTBuilder) VisitIfStatement(ctx *parser.IfStatementContext) interface{} {
	return ctx.IfStmt().Accept(this)
}

func (this *ASTBuilder) VisitIfStmt(ctx *parser.IfStmtContext) interface{} {
	return nil
}

func (this *ASTBuilder) VisitWhileStatement(ctx *parser.WhileStatementContext) interface{} {
	return ctx.WhileStmt().Accept(this)
}

func (this *ASTBuilder) VisitWhileStmt(cgtx *parser.WhileStmtContext) interface{} {
	return nil
}

func (this *ASTBuilder) VisitForStatement(ctx *parser.ForStatementContext) interface{} {
	return ctx.ForStmt().Accept(this)
}

func (this *ASTBuilder) VisitForStmt(ctx *parser.ForStmtContext) interface{} {
	return nil
}

func (this *ASTBuilder) VisitBreakStatement(ctx *parser.BreakStatementContext) interface{} {
	return ctx.BreakStmt().Accept(this)
}

func (this *ASTBuilder) VisitBreakStmt(ctx *parser.BreakStmtContext) interface{} {
	return nil
}

func (this *ASTBuilder) VisitContinueStatement(ctx *parser.ContinueStatementContext) interface{} {
	return ctx.ContinueStmt().Accept(this)
}

func (this *ASTBuilder) VisitContinueStmt(ctx *parser.ContinueStmtContext) interface{} {
	return nil
}

func (this *ASTBuilder) VisitGotoStatement(ctx *parser.GotoStatementContext) interface{} {
	return ctx.GotoStmt().Accept(this)
}

func (this *ASTBuilder) VisitGotoStmt(ctx *parser.GotoStmtContext) interface{} {
	return nil
}

func (this *ASTBuilder) VisitReturnStatement(ctx *parser.ReturnStatementContext) interface{} {
	return ctx.ReturnStmt().Accept(this)
}

func (this *ASTBuilder) VisitReturnStmt(ctx *parser.ReturnStmtContext) interface{} {
	var exprNode models.IASTExprNode = nil
	if ctx.Expr() != nil {
		exprNode = ctx.Expr().Accept(this).(models.IASTExprNode)
	}
	return models.NewASTReturnNode(models.NewLocation(this.sourcePath, ctx.GetStart()), exprNode)
}

func (this *ASTBuilder) VisitAssignOp(ctx *parser.AssignOpContext) interface{} {
	return ctx.GetText()
}

func (this *ASTBuilder) VisitAssignExpr(ctx *parser.AssignExprContext) interface{} {
	lhs := ctx.Term().Accept(this).(models.IASTExprNode)
	op := ctx.AssignOp().Accept(this).(string)
	rhs := ctx.Expr().Accept(this).(models.IASTExprNode)
	if op == "=" {
		return models.NewASTAssignNode(lhs, rhs)
	}
	return models.NewASTOpAssignNode(lhs, op, rhs)
}

func (this *ASTBuilder) VisitNoneAssignExpr(ctx *parser.NoneAssignExprContext) interface{} {
	return ctx.Expr10().Accept(this)
}

func (this *ASTBuilder) VisitExpr10(ctx *parser.Expr10Context) interface{} {
	c := ctx.Expr9().Accept(this).(models.IASTExprNode)
	condThenExprCtx := ctx.Expr()
	if condThenExprCtx != nil {
		t := condThenExprCtx.Accept(this).(models.IASTExprNode)
		e := ctx.Expr10().Accept(this).(models.IASTExprNode)
		return models.NewASTCondExprNode(c, t, e)
	}
	return c
}

func (this *ASTBuilder) VisitExpr9(ctx *parser.Expr9Context) interface{} {
	return &models.ASTExprNode{}
}

func (this *ASTBuilder) VisitExpr8(ctx *parser.Expr8Context) interface{} {
	return this.VisitChildren(ctx)
}

func (this *ASTBuilder) VisitExpr7(ctx *parser.Expr7Context) interface{} {
	return this.VisitChildren(ctx)
}

func (this *ASTBuilder) VisitExpr6(ctx *parser.Expr6Context) interface{} {
	return this.VisitChildren(ctx)
}

func (this *ASTBuilder) VisitExpr5(ctx *parser.Expr5Context) interface{} {
	return this.VisitChildren(ctx)
}

func (this *ASTBuilder) VisitExpr4(ctx *parser.Expr4Context) interface{} {
	return this.VisitChildren(ctx)
}

func (this *ASTBuilder) VisitExpr3(ctx *parser.Expr3Context) interface{} {
	return this.VisitChildren(ctx)
}

func (this *ASTBuilder) VisitExpr2(ctx *parser.Expr2Context) interface{} {
	return this.VisitChildren(ctx)
}

func (this *ASTBuilder) VisitExpr1(ctx *parser.Expr1Context) interface{} {
	return this.VisitChildren(ctx)
}

func (this *ASTBuilder) VisitTerm(ctx *parser.TermContext) interface{} {
	return this.VisitChildren(ctx)
}
