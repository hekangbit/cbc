package compiler

import (
	"cbc/ast"
	"cbc/loader"
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
	ctx.TopDefs().Accept(v)
	return &ast.AST{Name: "cbc program"}
}

func (v *ASTBuilder) VisitImportStmts(ctx *parser.ImportStmtsContext) interface{} {
	fmt.Println("ASTBuilder VisitImportStmts")
	for _, importStmt := range ctx.AllImportStmt() {
		importStmt.Accept(v)
	}
	return v.VisitChildren(ctx)
}

func (v *ASTBuilder) VisitImportStmt(ctx *parser.ImportStmtContext) interface{} {
	fmt.Println("ASTBuilder VisitImportStmt ")
	path := ctx.Identifier(0).GetText()
	for i := 1; i < len(ctx.AllIdentifier()); i++ {
		path = path + "." + ctx.Identifier(i).GetText()
	}
	loader.LoadLibrary(path)
	return v.VisitChildren(ctx)
}

func (v *ASTBuilder) VisitTopDefs(ctx *parser.TopDefsContext) interface{} {
	fmt.Println("ASTBuilder VisitTopDefs")
	return v.VisitChildren(ctx)
}

func (v *ASTBuilder) VisitDefVars(ctx *parser.DefVarsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ASTBuilder) VisitDefFunc(ctx *parser.DefFuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ASTBuilder) VisitTypedef(ctx *parser.TypedefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ASTBuilder) VisitBlock(ctx *parser.BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ASTBuilder) VisitStmt(ctx *parser.StmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ASTBuilder) VisitIfStmt(ctx *parser.IfStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ASTBuilder) VisitWhileStmt(ctx *parser.WhileStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ASTBuilder) VisitForStmt(ctx *parser.ForStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ASTBuilder) VisitBreakStmt(ctx *parser.BreakStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ASTBuilder) VisitContinueStmt(ctx *parser.ContinueStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ASTBuilder) VisitGotoStmt(ctx *parser.GotoStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ASTBuilder) VisitReturnStmt(ctx *parser.ReturnStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ASTBuilder) VisitType(ctx *parser.TypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ASTBuilder) VisitTyperef(ctx *parser.TyperefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ASTBuilder) VisitTyperefBase(ctx *parser.TyperefBaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ASTBuilder) VisitParams(ctx *parser.ParamsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ASTBuilder) VisitFixedParams(ctx *parser.FixedParamsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ASTBuilder) VisitParam(ctx *parser.ParamContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ASTBuilder) VisitParamTyperefs(ctx *parser.ParamTyperefsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ASTBuilder) VisitFixedparamTyperefs(ctx *parser.FixedparamTyperefsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *ASTBuilder) VisitCharConst(ctx *parser.CharConstContext) interface{} {
	return ctx.GetText()[0]
}
