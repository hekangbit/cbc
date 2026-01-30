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
	ctx.TopDefs().Accept(v)
	return models.NewAst("cbc program")
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
	return v.VisitChildren(ctx)
}

func (v *ASTBuilder) VisitDefVars(ctx *parser.DefVarsContext) interface{} {
	var defs []*models.DefinedVariable
	isStatic := false
	priv := ctx.GetPriv()
	if priv != nil {
		isStatic = true
	}
	ctx.GetCbtype()
	cbType := models.TypeNode{}

	for i, identifier := range ctx.AllIdentifier() {
		init := ctx.Expr(i).Accept(v)
		dv := models.NewDefinedVariable(isStatic, cbType, identifier.GetSymbol().GetText(), *(init.(*models.ExprNode)))
		defs = append(defs, dv)
	}

	return defs
}

func (v *ASTBuilder) VisitCondExpr(ctx *parser.CondExprContext) interface{} {
	return &models.ExprNode{}
}
