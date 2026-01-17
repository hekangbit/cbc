package main

import (
	"cbc/parser"
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

type CbVisitorImpl struct {
	*parser.BaseCbVisitor
	name string
}

func (v *CbVisitorImpl) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

func (v *CbVisitorImpl) VisitChildren(_ antlr.RuleNode) interface{} {
	return nil
}
func (v *CbVisitorImpl) VisitTerminal(_ antlr.TerminalNode) interface{} {
	return nil
}
func (v *CbVisitorImpl) VisitErrorNode(_ antlr.ErrorNode) interface{} {
	return nil
}

func (v *CbVisitorImpl) VisitProg(ctx *parser.ProgContext) interface{} {
	fmt.Println("CbVisitorImpl VisitProg")
	return v.BaseCbVisitor.VisitProg(ctx)
}

func (v *CbVisitorImpl) VisitImportStmts(ctx *parser.ImportStmtsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *CbVisitorImpl) VisitImportStmt(ctx *parser.ImportStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *CbVisitorImpl) VisitTopDefs(ctx *parser.TopDefsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *CbVisitorImpl) VisitDefVars(ctx *parser.DefVarsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *CbVisitorImpl) VisitDefFunc(ctx *parser.DefFuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *CbVisitorImpl) VisitTypedef(ctx *parser.TypedefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *CbVisitorImpl) VisitBlock(ctx *parser.BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *CbVisitorImpl) VisitStmt(ctx *parser.StmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *CbVisitorImpl) VisitIfStmt(ctx *parser.IfStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *CbVisitorImpl) VisitWhileStmt(ctx *parser.WhileStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *CbVisitorImpl) VisitForStmt(ctx *parser.ForStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *CbVisitorImpl) VisitBreakStmt(ctx *parser.BreakStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *CbVisitorImpl) VisitContinueStmt(ctx *parser.ContinueStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *CbVisitorImpl) VisitGotoStmt(ctx *parser.GotoStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *CbVisitorImpl) VisitReturnStmt(ctx *parser.ReturnStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *CbVisitorImpl) VisitType(ctx *parser.TypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *CbVisitorImpl) VisitTyperef(ctx *parser.TyperefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *CbVisitorImpl) VisitTyperefBase(ctx *parser.TyperefBaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *CbVisitorImpl) VisitParams(ctx *parser.ParamsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *CbVisitorImpl) VisitFixedParams(ctx *parser.FixedParamsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *CbVisitorImpl) VisitParam(ctx *parser.ParamContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *CbVisitorImpl) VisitParamTyperefs(ctx *parser.ParamTyperefsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *CbVisitorImpl) VisitFixedparamTyperefs(ctx *parser.FixedparamTyperefsContext) interface{} {
	return v.VisitChildren(ctx)
}
