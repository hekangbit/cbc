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
	name string
}

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
	retType := ctx.GetCbtype().Accept(v).(*models.TypeNode)
	name := ctx.Identifier().GetSymbol().GetText()
	// ps := ctx.Params().Accept(v)
	// body := ctx.Block().Accept(v)

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

	return models.NewDefinedFunction(priv, retType, name)
}

func (v *ASTBuilder) VisitBlock() interface{} {
	return nil
}

func (v *ASTBuilder) VisitCondExpr(ctx *parser.CondExprContext) interface{} {
	return &models.ExprNode{}
}

func (V *ASTBuilder) VisitCbType(ctx *parser.CbTypeContext) interface{} {
	return &models.TypeNode{}
}
