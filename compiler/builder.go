package compiler

import (
	"cbc/loader"
	"cbc/models"
	"cbc/parser"
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

type ASTBuilder struct {
	*parser.BaseCbVisitor
	name        string
	sourcePath  string
	curBaseType models.ITypeRef
}

var _ parser.CbVisitor = (*ASTBuilder)(nil)

func (v *ASTBuilder) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

func (v *ASTBuilder) VisitChildren(_ antlr.RuleNode) interface{} {
	return nil
}

func (v *ASTBuilder) VisitTerminal(_ antlr.TerminalNode) interface{} {
	return nil
}

func (v *ASTBuilder) VisitErrorNode(_ antlr.ErrorNode) interface{} {
	return nil
}

func (v *ASTBuilder) VisitProg(ctx *parser.ProgContext) interface{} {
	fmt.Println("ASTBuilder VisitProg")
	ctx.ImportStmts().Accept(v)
	decls := ctx.TopDefs().Accept(v).(*models.Declarations)
	return models.NewAst("cbc program", decls)
}

func (v *ASTBuilder) VisitImportStmts(ctx *parser.ImportStmtsContext) interface{} {
	fmt.Println("ASTBuilder VisitImportStmts")
	for _, importStmt := range ctx.AllImportStmt() {
		importStmt.Accept(v)
	}
	return v.VisitChildren(ctx)
}

func (v *ASTBuilder) VisitImportStmt(ctx *parser.ImportStmtContext) interface{} {
	fmt.Println("ASTBuilder VisitImportStmt")
	path := ctx.Identifier(0).GetText()
	for i := 1; i < len(ctx.AllIdentifier()); i++ {
		path = path + "." + ctx.Identifier(i).GetText()
	}
	loader.LoadLibrary(path)
	return v.VisitChildren(ctx)
}

func (v *ASTBuilder) VisitTopDefs(ctx *parser.TopDefsContext) interface{} {
	fmt.Println("ASTBuilder VisitTopDefs")
	decls := models.NewDeclarations()
	for _, defVars := range ctx.AllDefVars() {
		defvars := defVars.Accept(v)
		decls.AddDefvars(defvars.([]*models.DefinedVariable))
	}
	for _, defFun := range ctx.AllDefFunc() {
		f := defFun.Accept(v)
		decls.AddDefFunc(f.(*models.DefinedFunction))
	}
	return decls
}

func (v *ASTBuilder) VisitDefVars(ctx *parser.DefVarsContext) interface{} {
	var initialize *models.ExprNode = nil
	var defs []*models.DefinedVariable
	priv := false
	if ctx.GetPriv() != nil {
		priv = true
	}
	cbType := ctx.GetCbtype().Accept(v).(*models.TypeNode)

	for _, identifier := range ctx.AllIdentifier() {
		initialize = nil
		if ctx.GetHasInit() != nil {
			initialize = ctx.GetInitializer().Accept(v).(*models.ExprNode)
		}
		dv := models.NewDefinedVariable(priv, cbType, identifier.GetSymbol().GetText(), initialize)
		defs = append(defs, dv)
	}

	return defs
}

func (v *ASTBuilder) VisitDefFunc(ctx *parser.DefFuncContext) interface{} {
	priv := false
	if ctx.GetPriv() != nil {
		priv = true
	}
	retType := ctx.GetRetCbtype().Accept(v).(*models.TypeNode)
	name := ctx.Identifier().GetSymbol().GetText()
	params := ctx.Params().Accept(v).(*models.Params)
	body := ctx.Block().Accept(v).(*models.ASTBlockNode)

	// DefinedFunction var7;
	// try {
	// 		boolean priv = this.storage();
	// 		TypeRef ret = this.typeref();
	// 		String n = this.name();
	// 		this.jj_consume_token(46);
	// 		Params ps = this.params();
	// 		this.jj_consume_token(51);
	// 		BlockNode body = this.block();
	// 		TypeRef t = new FunctionTypeRef(ret, ps.parametersTypeRef());
	// 		if ("" == null) {
	// 			throw new Error("Missing return statement in function");
	// 		}

	// 		var7 = new DefinedFunction(priv, new TypeNode(t), n, ps, body);
	// } finally {
	// 		this.trace_return("defun");
	// }

	return models.NewDefinedFunction(priv, retType, name, params, body)
}

func (v *ASTBuilder) VisitParams(ctx *parser.ParamsContext) interface{} {
	voidToken := ctx.GetVoid()
	if voidToken != nil {
		paramDescs := make([]*models.CBCParameter, 0)
		return models.NewParams(models.NewLocation(v.sourcePath, voidToken), paramDescs)
	}
	params := ctx.FixedParams().Accept(v).(*models.Params)
	if ctx.GetHasVararg() != nil {
		params.AcceptVarargs()
	}
	return params
}

func (v *ASTBuilder) VisitBlock(ctx *parser.BlockContext) interface{} {
	return nil
}

func (v *ASTBuilder) VisitCondExpr(ctx *parser.CondExprContext) interface{} {
	return &models.ExprNode{}
}

func (v *ASTBuilder) VisitCbType(ctx *parser.CbTypeContext) interface{} {
	return nil
}

func (v *ASTBuilder) VisitCbTypeRef(ctx *parser.CbTypeRefContext) interface{} {
	v.curBaseType = ctx.CbTypeRefBase().Accept(v).(models.ITypeRef)
	modifiers := ctx.AllTypeModifier()
	// TODO: add more modifier (sizedArray, FunctionType)
	for i := len(modifiers) - 1; i >= 0; i-- {
		v.curBaseType = modifiers[i].Accept(v).(models.ITypeRef)
	}
	return v.curBaseType
}

func (v *ASTBuilder) VisitVoidTypeRef(ctx *parser.VoidTypeRefContext) interface{} {
	return models.NewVoidTypeRefWithLocation(models.NewLocation(v.sourcePath, ctx.GetStart()))
}

func (v *ASTBuilder) VisitCharTypeRef(ctx *parser.CharTypeRefContext) interface{} {
	return models.NewCharRefWithLocation(models.NewLocation(v.sourcePath, ctx.GetStart()))
}

func (v *ASTBuilder) VisitShortTypeRef(ctx *parser.ShortTypeRefContext) interface{} {
	return models.NewShortRefWithLocation(models.NewLocation(v.sourcePath, ctx.GetStart()))
}

func (v *ASTBuilder) VisitIntTypeRef(ctx *parser.IntTypeRefContext) interface{} {
	return models.NewIntRefWithLocation(models.NewLocation(v.sourcePath, ctx.GetStart()))
}

func (this *ASTBuilder) VisitArrayModifier(ctx *parser.ArrayModifierContext) interface{} {
	return models.NewArrayTypeRef(this.curBaseType)
}

func (this *ASTBuilder) VisitPointerModifier(ctx *parser.PointerModifierContext) interface{} {
	return models.NewPointerTypeRef(this.curBaseType)
}
