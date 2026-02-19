package models

import (
	"cbc/parser"
	"io"

	"github.com/antlr4-go/antlr/v4"
)

type AST struct {
	Node
	source        *Location
	decls         *Declarations
	scope         *ToplevelScope
	constantTable *ConstantTable
	stream        *antlr.CommonTokenStream
	cbLexer       *parser.CbLexer
}

func NewAST(source *Location, declarations *Declarations) *AST {
	return &AST{
		source: source,
		decls:  declarations,
	}
}

func (this *AST) Location() *Location {
	return this.source
}

func (this *AST) Types() []IASTTypeDefinition {
	result := []IASTTypeDefinition{}
	for _, t := range this.decls.Defstructs() {
		result = append(result, t)
	}
	for _, t := range this.decls.Defunions() {
		result = append(result, t)
	}
	for _, t := range this.decls.Typedefs() {
		result = append(result, t)
	}
	return result
}

func (this *AST) Entities() []IEntity {
	result := []IEntity{}
	for _, f := range this.decls.Funcdecls() {
		result = append(result, f)
	}
	for _, v := range this.decls.Vardecls() {
		result = append(result, v)
	}
	for _, v := range this.decls.Defvars() {
		result = append(result, v)
	}
	for _, f := range this.decls.Deffuns() {
		result = append(result, f)
	}
	for _, c := range this.decls.Constants() {
		result = append(result, c)
	}
	return result
}

func (this *AST) Declarations() []IEntity {
	result := []IEntity{}
	for _, f := range this.decls.Funcdecls() {
		result = append(result, f)
	}
	for _, v := range this.decls.Vardecls() {
		result = append(result, v)
	}
	return result
}

func (this *AST) Definitions() []IEntity {
	result := []IEntity{}
	for _, v := range this.decls.Defvars() {
		result = append(result, v)
	}
	for _, f := range this.decls.Deffuns() {
		result = append(result, f)
	}
	for _, c := range this.decls.Constants() {
		result = append(result, c)
	}
	return result
}

func (this *AST) Constants() []*Constant {
	return this.decls.Constants()
}

func (this *AST) DefinedVariables() []*DefinedVariable {
	return this.decls.Defvars()
}

func (this *AST) DefinedFunctions() []*DefinedFunction {
	return this.decls.Deffuns()
}

func (this *AST) SetScope(scope *ToplevelScope) {
	if this.scope != nil {
		panic("must not happen: ToplevelScope set twice")
	}
	this.scope = scope
}

func (this *AST) Scope() *ToplevelScope {
	if this.scope == nil {
		panic("must not happen: AST.scope is nil")
	}
	return this.scope
}

func (this *AST) SetConstantTable(table *ConstantTable) {
	if this.constantTable != nil {
		panic("must not happen: ConstantTable set twice")
	}
	this.constantTable = table
}

func (this *AST) ConstantTable() *ConstantTable {
	if this.constantTable == nil {
		panic("must not happen: AST.constantTable is null")
	}
	return this.constantTable
}

func (this *AST) IR() *IR {
	panic("TODO: AST::IR")
}

func (this *AST) DumpDefinedVariables(d *Dumper, buf []*DefinedVariable) {
	dumpables := make([]Dumpable, len(buf))
	for i, p := range buf {
		dumpables[i] = p
	}
	d.PrintNodeList("variables", dumpables)
}

func (this *AST) DumpDefinedFunctions(d *Dumper, buf []*DefinedFunction) {
	dumpables := make([]Dumpable, len(buf))
	for i, p := range buf {
		dumpables[i] = p
	}
	d.PrintNodeList("functions", dumpables)
}

func (this *AST) Dump(d *Dumper) {
	d.PrintClass(this, this.Location())
	this.DumpDefinedVariables(d, this.DefinedVariables())
	this.DumpDefinedFunctions(d, this.DefinedFunctions())
}

const numLeftColumns = 24

func (this *AST) SetStream(cbLexer *parser.CbLexer, stream *antlr.CommonTokenStream) {
	this.cbLexer = cbLexer
	this.stream = stream
}

func (this *AST) DumpTokens(w io.Writer) {
	for _, t := range this.stream.GetAllTokens() {
		if t.GetTokenType() != antlr.TokenEOF {
			this.PrintPair(this.cbLexer.SymbolicNames[t.GetTokenType()], t.GetText(), w)
		}
	}
}

func (this *AST) PrintPair(key string, value string, w io.Writer) {
	io.WriteString(w, key)
	for n := numLeftColumns - len(key); n > 0; n-- {
		io.WriteString(w, " ")
	}
	io.WriteString(w, value)
	io.WriteString(w, "\n")
}

// GetSingleMainStmt return the first statement in cb main function
func (this *AST) GetSingleMainStmt() IASTStmtNode {
	for _, f := range this.DefinedFunctions() {
		if f.Name() == "main" {
			stmts := f.Body().Stmts()
			if len(stmts) == 0 {
				return nil
			}
			return stmts[0]
		}
	}
	return nil
}

// GetSingleMainExpr return the expr of signle statement in main function
func (this *AST) GetSingleMainExpr() IASTExprNode {
	stmt := this.GetSingleMainStmt()
	if stmt == nil {
		return nil
	}
	switch s := stmt.(type) {
	case *ASTExprStmtNode:
		return s.Expr()
	case *ASTReturnNode:
		return s.Expr()
	default:
		return nil
	}
}
