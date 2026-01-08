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
	// for _, child := range ctx.AllStatement() {
	// 	v.Visit(child.(antlr.ParseTree))
	// }
	return v.BaseCbVisitor.VisitProg(ctx)
}

func (v *CbVisitorImpl) VisitStatement(ctx *parser.StatementContext) interface{} {
	fmt.Println("CbVisitorImpl VisitStatement")
	return nil
}

func (v *CbVisitorImpl) VisitExpr(ctx *parser.ExprContext) interface{} {
	fmt.Println("CbVisitorImpl VisitExpr")
	return nil
}
