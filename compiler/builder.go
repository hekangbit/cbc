package compiler

import (
	"cbc/loader"
	"cbc/models"
	"cbc/parser"
	"cbc/utils"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
)

const (
	POSTFIX_OP_KIND_SUFFIX_OP int = iota
	POSTFIX_OP_KIND_ARRAY_INDEX
	POSTFIX_OP_KIND_MEMBER
	POSTFIX_OP_KIND_PTR_MEMBER
	POSTFIX_OP_KIND_FUNCTION_CALL
)

type PostfixOpContent struct {
	kind int
	data interface{}
}

type ASTBuilder struct {
	*parser.BaseCbVisitor
	name        string
	sourcePath  string
	curBaseType models.ITypeRef
}

var _ parser.CbVisitor = &ASTBuilder{}

func (this *ASTBuilder) Loc(token antlr.Token) *models.Location {
	return models.NewLocation(this.sourcePath, token)
}

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
	return models.NewAST(this.Loc(ctx.GetStart()), decls)
}

// TODO: remove visit children
func (this *ASTBuilder) VisitImportStmts(ctx *parser.ImportStmtsContext) interface{} {
	for _, importStmt := range ctx.AllImportStmt() {
		importStmt.Accept(this)
	}
	return this.VisitChildren(ctx)
}

// TODO: remove visit children
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
	cbType := ctx.GetCbtype().Accept(this).(*models.ASTTypeNode)

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
	for i := 0; i < len(modifiers); i++ {
		this.curBaseType = modifiers[i].Accept(this).(models.ITypeRef)
	}
	return this.curBaseType
}

func (this *ASTBuilder) VisitVoidTypeRef(ctx *parser.VoidTypeRefContext) interface{} {
	return models.NewVoidTypeRefWithLocation(this.Loc(ctx.GetStart()))
}

func (this *ASTBuilder) VisitCharTypeRef(ctx *parser.CharTypeRefContext) interface{} {
	return models.NewCharRefWithLocation(this.Loc(ctx.GetStart()))
}

func (this *ASTBuilder) VisitShortTypeRef(ctx *parser.ShortTypeRefContext) interface{} {
	return models.NewShortRefWithLocation(this.Loc(ctx.GetStart()))
}

func (this *ASTBuilder) VisitIntTypeRef(ctx *parser.IntTypeRefContext) interface{} {
	return models.NewIntRefWithLocation(this.Loc(ctx.GetStart()))
}

func (this *ASTBuilder) VisitArrayModifier(ctx *parser.ArrayModifierContext) interface{} {
	return models.NewArrayTypeRef(this.curBaseType)
}

// TODO: convert string to int64, now to int32
func (this *ASTBuilder) VisitSizedArrayModifier(ctx *parser.SizedArrayModifierContext) interface{} {
	length, err := strconv.Atoi(ctx.IntLiteral().GetText())
	if err != nil {
		panic("AST Builder::VisitSizedArrayModifier GetIntConst Fail")
	}
	return models.NewArrayTypeRefWithLen(this.curBaseType, int64(length))
}

func (this *ASTBuilder) VisitFunctionModifier(ctx *parser.FunctionModifierContext) interface{} {
	params := ctx.ParamTypeRefs().Accept(this).(*models.ParamTypeRefs)
	return models.NewFunctionTypeRef(this.curBaseType, params)
}

func (this *ASTBuilder) VisitPointerModifier(ctx *parser.PointerModifierContext) interface{} {
	return models.NewPointerTypeRef(this.curBaseType)
}

func (this *ASTBuilder) VisitParams(ctx *parser.ParamsContext) interface{} {
	voidToken := ctx.GetVoid()
	if voidToken != nil {
		paramDescs := make([]*models.CBCParameter, 0)
		return models.NewParams(this.Loc(voidToken), paramDescs)
	}
	fixedParams := ctx.FixedParams().Accept(this).([]*models.CBCParameter)
	fullParams := models.NewParams(this.Loc(ctx.GetStart()), fixedParams)
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
	typeNode := ctx.CbType().Accept(this).(*models.ASTTypeNode)
	name := ctx.Identifier().GetSymbol().GetText()
	return models.NewCBCParameter(typeNode, name)
}

// TODO
func (this *ASTBuilder) VisitParamTypeRefs(ctx *parser.ParamTypeRefsContext) interface{} {
	return this.VisitChildren(ctx)
}

// TODO
func (this *ASTBuilder) VisitFixedparamTypeRefs(ctx *parser.FixedparamTypeRefsContext) interface{} {
	return this.VisitChildren(ctx)
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

	return models.NewASTBlockNode(this.Loc(ctx.GetStart()), defLocalVars, stmts)
}

func (this *ASTBuilder) VisitExprStatement(ctx *parser.ExprStatementContext) interface{} {
	expr := ctx.Expr().Accept(this).(models.IASTExprNode)
	return models.NewASTExprStmtNode(this.Loc(ctx.GetStart()), expr)
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
	return models.NewASTReturnNode(this.Loc(ctx.GetStart()), exprNode)
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
	l := ctx.Expr8(0).Accept(this).(models.IASTExprNode)
	for i := 1; i < len(ctx.AllExpr8()); i++ {
		r := ctx.Expr8(i).Accept(this).(models.IASTExprNode)
		l = models.NewASTLogicalOrNode(l, r)
	}
	return l
}

func (this *ASTBuilder) VisitExpr8(ctx *parser.Expr8Context) interface{} {
	l := ctx.Expr7(0).Accept(this).(models.IASTExprNode)
	for i := 1; i < len(ctx.AllExpr7()); i++ {
		r := ctx.Expr7(i).Accept(this).(models.IASTExprNode)
		l = models.NewASTLogicalAndNode(l, r)
	}
	return l
}

func (this *ASTBuilder) VisitExpr7(ctx *parser.Expr7Context) interface{} {
	l := ctx.Expr6(0).Accept(this).(models.IASTExprNode)
	children := ctx.GetChildren()
	for i := 1; i < len(ctx.AllExpr6()); i++ {
		r := ctx.Expr6(i).Accept(this).(models.IASTExprNode)
		opIndex := 2*i - 1
		opNode := children[opIndex].(antlr.TerminalNode)
		opText := opNode.GetText()
		l = models.NewASTBinaryOpNode(l, opText, r)
	}
	return l
}

func (this *ASTBuilder) VisitExpr6(ctx *parser.Expr6Context) interface{} {
	l := ctx.Expr5(0).Accept(this).(models.IASTExprNode)
	for i := 1; i < len(ctx.AllExpr5()); i++ {
		r := ctx.Expr5(i).Accept(this).(models.IASTExprNode)
		l = models.NewASTBinaryOpNode(l, "|", r)
	}
	return l
}

func (this *ASTBuilder) VisitExpr5(ctx *parser.Expr5Context) interface{} {
	l := ctx.Expr4(0).Accept(this).(models.IASTExprNode)
	for i := 1; i < len(ctx.AllExpr4()); i++ {
		r := ctx.Expr4(i).Accept(this).(models.IASTExprNode)
		l = models.NewASTBinaryOpNode(l, "^", r)
	}
	return l
}

func (this *ASTBuilder) VisitExpr4(ctx *parser.Expr4Context) interface{} {
	l := ctx.Expr3(0).Accept(this).(models.IASTExprNode)
	for i := 1; i < len(ctx.AllExpr3()); i++ {
		r := ctx.Expr3(i).Accept(this).(models.IASTExprNode)
		l = models.NewASTBinaryOpNode(l, "&", r)
	}
	return l
}

func (this *ASTBuilder) VisitExpr3(ctx *parser.Expr3Context) interface{} {
	l := ctx.Expr2(0).Accept(this).(models.IASTExprNode)
	children := ctx.GetChildren()
	for i := 1; i < len(ctx.AllExpr2()); i++ {
		r := ctx.Expr2(i).Accept(this).(models.IASTExprNode)
		opIndex := 2*i - 1
		opNode := children[opIndex].(antlr.TerminalNode)
		opText := opNode.GetText()
		l = models.NewASTBinaryOpNode(l, opText, r)
	}
	return l
}

func (this *ASTBuilder) VisitExpr2(ctx *parser.Expr2Context) interface{} {
	l := ctx.Expr1(0).Accept(this).(models.IASTExprNode)
	children := ctx.GetChildren()
	for i := 1; i < len(ctx.AllExpr1()); i++ {
		r := ctx.Expr1(i).Accept(this).(models.IASTExprNode)
		opIndex := 2*i - 1
		opNode := children[opIndex].(antlr.TerminalNode)
		opText := opNode.GetText()
		l = models.NewASTBinaryOpNode(l, opText, r)
	}
	return l
}

func (this *ASTBuilder) VisitExpr1(ctx *parser.Expr1Context) interface{} {
	l := ctx.Term(0).Accept(this).(models.IASTExprNode)
	children := ctx.GetChildren()
	for i := 1; i < len(ctx.AllTerm()); i++ {
		r := ctx.Term(i).Accept(this).(models.IASTExprNode)
		opIndex := 2*i - 1
		opNode := children[opIndex].(antlr.TerminalNode)
		opText := opNode.GetText()
		l = models.NewASTBinaryOpNode(l, opText, r)
	}
	return l
}

func (this *ASTBuilder) VisitTermCast(ctx *parser.TermCastContext) interface{} {
	return nil
}

func (this *ASTBuilder) VisitTermUnary(ctx *parser.TermUnaryContext) interface{} {
	return ctx.Unary().Accept(this)
}

func (this *ASTBuilder) VisitCastExpr(ctx *parser.CastExprContext) interface{} {
	return nil
}

func (this *ASTBuilder) VisitUnaryPrefixIncrement(ctx *parser.UnaryPrefixIncrementContext) interface{} {
	n := ctx.Unary().Accept(this).(models.IASTExprNode)
	return models.NewASTPrefixOpNode("++", n)
}

func (this *ASTBuilder) VisitUnaryPrefixDecrement(ctx *parser.UnaryPrefixDecrementContext) interface{} {
	n := ctx.Unary().Accept(this).(models.IASTExprNode)
	return models.NewASTPrefixOpNode("--", n)
}

func (this *ASTBuilder) VisitUnaryPrefixPlus(ctx *parser.UnaryPrefixPlusContext) interface{} {
	n := ctx.Unary().Accept(this).(models.IASTExprNode)
	return models.NewASTUnaryOpNode("+", n)
}

func (this *ASTBuilder) VisitUnaryPrefixMinus(ctx *parser.UnaryPrefixMinusContext) interface{} {
	n := ctx.Unary().Accept(this).(models.IASTExprNode)
	return models.NewASTUnaryOpNode("-", n)
}

func (this *ASTBuilder) VisitUnaryPrefixLogicalNot(ctx *parser.UnaryPrefixLogicalNotContext) interface{} {
	n := ctx.Unary().Accept(this).(models.IASTExprNode)
	return models.NewASTUnaryOpNode("!", n)
}

func (this *ASTBuilder) VisitUnaryPrefixfixBitwiseNot(ctx *parser.UnaryPrefixfixBitwiseNotContext) interface{} {
	n := ctx.Unary().Accept(this).(models.IASTExprNode)
	return models.NewASTUnaryOpNode("~", n)
}

func (this *ASTBuilder) VisitUnaryPrefixDereference(ctx *parser.UnaryPrefixDereferenceContext) interface{} {
	n := ctx.Unary().Accept(this).(models.IASTExprNode)
	return models.NewASTDereferenceNode(n)
}

func (this *ASTBuilder) VisitUnaryPrefixAddress(ctx *parser.UnaryPrefixAddressContext) interface{} {
	n := ctx.Unary().Accept(this).(models.IASTExprNode)
	return models.NewASTAddressNode(n)
}

func (this *ASTBuilder) VisitUnaryPrefixSizeofType(ctx *parser.UnaryPrefixSizeofTypeContext) interface{} {
	t := ctx.CbType().Accept(this).(*models.ASTTypeNode)
	return models.NewASTSizeofTypeNode(t, models.NewUlongRef())
}

func (this *ASTBuilder) VisitUnaryPrefixSizeof(ctx *parser.UnaryPrefixSizeofContext) interface{} {
	e := ctx.Unary().Accept(this).(models.IASTExprNode)
	return models.NewASTSizeofExprNode(e, models.NewUlongRef())
}

func (this *ASTBuilder) VisitUnaryPostfix(ctx *parser.UnaryPostfixContext) interface{} {
	return ctx.Postfix().Accept(this)
}

func (this *ASTBuilder) VisitPostfix(ctx *parser.PostfixContext) interface{} {
	expr := ctx.Primary().Accept(this).(models.IASTExprNode)
	for _, tmpCtx := range ctx.AllPostfixOp() {
		content := tmpCtx.Accept(this).(PostfixOpContent)
		switch content.kind {
		case POSTFIX_OP_KIND_SUFFIX_OP:
			expr = models.NewASTSuffixOpNode(content.data.(string), expr)
		case POSTFIX_OP_KIND_ARRAY_INDEX:
			expr = models.NewASTArrayIdxRefNode(expr, content.data.(models.IASTExprNode))
		case POSTFIX_OP_KIND_MEMBER:
			expr = models.NewASTMemberNode(expr, content.data.(string))
		case POSTFIX_OP_KIND_PTR_MEMBER:
			expr = models.NewASTPtrMemberNode(expr, content.data.(string))
		case POSTFIX_OP_KIND_FUNCTION_CALL:
			expr = models.NewASTFunctionCallNode(expr, content.data.([]models.IASTExprNode))
		default:
			panic("ASTBuilder#VisitPostfix invalid postfix op kind")
		}
	}
	return expr
}

func (this *ASTBuilder) VisitPostInc(ctx *parser.PostIncContext) interface{} {
	return PostfixOpContent{kind: POSTFIX_OP_KIND_SUFFIX_OP, data: "++"}
}

func (this *ASTBuilder) VisitPostDec(ctx *parser.PostDecContext) interface{} {
	return PostfixOpContent{kind: POSTFIX_OP_KIND_SUFFIX_OP, data: "--"}
}

func (this *ASTBuilder) VisitPostArrayIndex(ctx *parser.PostArrayIndexContext) interface{} {
	idx := ctx.Expr().Accept(this).(models.IASTExprNode)
	return PostfixOpContent{kind: POSTFIX_OP_KIND_ARRAY_INDEX, data: idx}
}

func (this *ASTBuilder) VisitPostMember(ctx *parser.PostMemberContext) interface{} {
	memb := ctx.Identifier().Accept(this).(string)
	return PostfixOpContent{kind: POSTFIX_OP_KIND_MEMBER, data: memb}
}

func (this *ASTBuilder) VisitPostPtrMember(ctx *parser.PostPtrMemberContext) interface{} {
	memb := ctx.Identifier().Accept(this).(string)
	return PostfixOpContent{kind: POSTFIX_OP_KIND_PTR_MEMBER, data: memb}
}

func (this *ASTBuilder) VisitFuncCall(ctx *parser.FuncCallContext) interface{} {
	args := ctx.Args().Accept(this).([]models.IASTExprNode)
	return PostfixOpContent{kind: POSTFIX_OP_KIND_FUNCTION_CALL, data: args}
}

func (this *ASTBuilder) VisitArgs(ctx *parser.ArgsContext) interface{} {
	var args []models.IASTExprNode = make([]models.IASTExprNode, 0)
	for _, tmpCtx := range ctx.AllExpr() {
		args = append(args, tmpCtx.Accept(this).(models.IASTExprNode))
	}
	return args
}

// TODO: How to get int64 const value
// TODO: java IntegerLiteralNode integerNode(Location loc, String image)
// can handle UL, L, U suffix
func (this *ASTBuilder) VisitIntConst(ctx *parser.IntConstContext) interface{} {
	val, err := strconv.Atoi(ctx.IntLiteral().GetText())
	if err != nil {
		panic("AST Builder: GetIntConst Fail")
	}
	p := models.NewASTIntegerLiteralNode(this.Loc(ctx.GetStart()), models.NewLongRef(), int64(val))
	return p
}

// TODO: cast int to int64, and need check effect
func (this *ASTBuilder) VisitCharConst(ctx *parser.CharConstContext) interface{} {
	r := []rune(ctx.Character().GetText())[1]
	val := int(r)
	p := models.NewASTIntegerLiteralNode(this.Loc(ctx.GetStart()), models.NewCharRef(), int64(val))
	return p
}

func (this *ASTBuilder) VisitStringConst(ctx *parser.StringConstContext) interface{} {
	s, err := utils.StringValue(ctx.GetText())
	if err != nil {
		panic("ASTBuilder#VisitStringConst string literal invalid")
	}
	return models.NewASTStringLiteralNode(this.Loc(ctx.GetStart()), models.NewPointerTypeRef(models.NewCharRef()), s)
}

func (this *ASTBuilder) VisitIdentifier(ctx *parser.IdentifierContext) interface{} {
	return models.NewASTVariableNode(this.Loc(ctx.GetStart()), ctx.Identifier().GetText())
}

func (this *ASTBuilder) VisitParenExpr(ctx *parser.ParenExprContext) interface{} {
	return ctx.Expr().Accept(this)
}
