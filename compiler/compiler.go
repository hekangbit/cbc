package compiler

import (
	"cbc/models"
	"cbc/parser"
	"cbc/sysdep/x86"
	"cbc/utils"
	"errors"
	"fmt"
	"os"

	"github.com/antlr4-go/antlr/v4"
)

var CompilerProgramName string = "cbc"
var CompilerVersion string = "1.0.0"
var errorHandler *utils.ErrorHandler = utils.NewErrorHandler(CompilerProgramName)

type CustomErrorListener struct {
	*antlr.DefaultErrorListener
	hasError   bool
	sourcePath string
	count      int
}

func NewCustomErrorListener(sourcePath string) *CustomErrorListener {
	return &CustomErrorListener{
		DefaultErrorListener: &antlr.DefaultErrorListener{},
		hasError:             false,
		sourcePath:           sourcePath,
		count:                0,
	}
}

func (l *CustomErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	fmt.Fprintf(os.Stderr, "%s:%d:%d error: %s\n", l.sourcePath, line, column, msg)
	l.hasError = true
	l.count++
}

func DebugDump(path string) {
	src, err := os.ReadFile(path)
	if err != nil {
		os.Exit(64)
	}
	input := antlr.NewInputStream(string(src))
	cbLexer := parser.NewCbLexer(input)
	for {
		t := cbLexer.NextToken()
		if t.GetTokenType() == antlr.TokenEOF {
			break
		}
		fmt.Printf("%s (%q)\n", cbLexer.SymbolicNames[t.GetTokenType()], t.GetText())
	}
	cbLexer.Reset()
	tokenStream := antlr.NewCommonTokenStream(cbLexer, antlr.TokenDefaultChannel)
	cbParser := parser.NewCbParser(tokenStream)
	tree := cbParser.Prog()
	fmt.Println(tree.ToStringTree(cbParser.RuleNames, cbParser))
	builder := &ASTBuilder{BaseCbVisitor: &parser.BaseCbVisitor{}}
	program := tree.Accept(builder) // builder.Visit(tree)
	fmt.Println(program.(*models.AST))
}

func DumpAST(ast *models.AST, mode CompilerMode) bool {
	switch mode {
	case COMPILER_MODE_DumpTokens:
		ast.DumpTokens(os.Stdout)
		return true
	case COMPILER_MODE_DumpAST:
		ast.DumpNode()
		return true
	case COMPILER_MODE_DumpStmt:
		return true
	case COMPILER_MODE_DumpExpr:
		return true
	default:
		return false
	}
}

func DumpSemant(ast *models.AST, mode CompilerMode) bool {
	switch mode {
	case COMPILER_MODE_DumpReference, COMPILER_MODE_DumpSemantic:
		ast.DumpNode()
		return true
	default:
		return false
	}
}

func DumpIR(ir *models.IR, mode CompilerMode) bool {
	switch mode {
	case COMPILER_MODE_DumpIR:
		ir.Dump()
		return true
	default:
		return false
	}
}

func DumpAsm(asmObj *x86.AssemblyCode, mode CompilerMode) bool {
	return false
}

func PrintAsm(asmObj *x86.AssemblyCode, mode CompilerMode) bool {
	return false
}

func GenerateExecutable(opts *Options) {
}

func GenerateSharedLibrary(opts *Options) {
}

func ParseFile(path string, opts *Options) *models.AST {
	src, err := os.ReadFile(path)
	if err != nil {
		os.Exit(64)
	}
	input := antlr.NewInputStream(string(src))
	cbLexer := parser.NewCbLexer(input)

	cbLexer.RemoveErrorListeners() // can keep default listener
	errorListener := NewCustomErrorListener(path)
	cbLexer.AddErrorListener(errorListener)

	tokenStream := antlr.NewCommonTokenStream(cbLexer, antlr.TokenDefaultChannel)
	cbParser := parser.NewCbParser(tokenStream)

	cbParser.RemoveErrorListeners()
	cbParser.AddErrorListener(errorListener)
	tree := cbParser.Prog()

	if errorListener.hasError {
		fmt.Printf("stop compile %s\n", path)
		fmt.Printf("totaly %d errors\n", errorListener.count)
		os.Exit(1)
	}

	builder := &ASTBuilder{BaseCbVisitor: &parser.BaseCbVisitor{}, sourcePath: path}
	program := tree.Accept(builder)
	cbAST := program.(*models.AST)
	cbAST.SetStream(cbLexer, tokenStream)
	return cbAST
}

func SemanticAnalyzePhase1(astNode *models.AST, typeTable *models.TypeTable, opts *Options) error {
	localResolver := NewLocalResolver(errorHandler)
	err := localResolver.Resolve(astNode)
	if err != nil {
		return errors.New("SemanticAnalyzePhase1 Local Resolver fail")
	}
	// typeResolver := NewTypeResolver(typeTable, errorHandler)
	// typeResolver.Resolve(astNode)
	return nil
}

func SemanticAnalyzePhase2(astNode *models.AST, typeTable *models.TypeTable, opts *Options) error {
	return nil
}

func GenerateIR(sem *models.AST, typeTable *models.TypeTable) *models.IR {
	generator := NewIRGenerator(typeTable, errorHandler)
	ir, err := generator.Generate(sem)
	if err != nil {
		panic("generate ir fail")
	}
	return ir
}

func GenerateAssembly(ir *models.IR, opts *Options) *x86.AssemblyCode {
	return &x86.AssemblyCode{}
}

func WriteFile(path string, content string) {
}

// parse file
// semantic analyze
// generate ir
// generate asm
// write file
func CompileToAsm(srcPath string, dstPath string, opts *Options) error {
	var err error
	fmt.Println("CompileToAsm: " + srcPath + " to " + dstPath)
	typeTable := opts.TypeTable()
	astObj := ParseFile(srcPath, opts)
	if DumpAST(astObj, opts.Mode()) {
		return nil
	}
	err = SemanticAnalyzePhase1(astObj, typeTable, opts)
	if err != nil {
		return err
	}
	if DumpSemant(astObj, opts.Mode()) {
		return nil
	}
	err = SemanticAnalyzePhase2(astObj, typeTable, opts)
	if err != nil {
		return err
	}
	// ir := GenerateIR(semObj, typeTable)
	// if DumpIR(ir, opts.Mode()) {
	// 	return
	// }
	// asmObj := GenerateAssembly(ir, opts)
	// if DumpAsm(asmObj, opts.Mode()) {
	// 	return
	// }
	// if PrintAsm(asmObj, opts.Mode()) {
	// 	return
	// }
	// WriteFile(dstPath, asmObj.String())
	return nil
}

func AssembleToObj(srcPath string, dstPath string, opts *Options) {
	fmt.Println("AssembleToObj: " + srcPath + " to " + dstPath)
}

func Link(opts *Options) {
	fmt.Println("Link")
	if !opts.IsGeneratingSharedLibrary() {
		GenerateExecutable(opts)
	} else {
		GenerateSharedLibrary(opts)
	}
}

func CompileFile(src SourceFile, opts *Options) error {
	var err error
	// .cb -> .s
	if src.IsCbSource() {
		destPath := opts.AsmFileNameOf(src)
		err = CompileToAsm(src.Path(), destPath, opts)
		if err != nil {
			return err
		}
		src.SetCurrentName(destPath)
	}
	if !opts.IsAssembleRequired() {
		return nil
	}
	// .s -> .o
	if src.IsAssemblySource() {
		destPath := opts.ObjFileNameOf(src)
		AssembleToObj(src.Path(), destPath, opts)
		src.SetCurrentName(destPath)
	}
	return nil
}

func CompileFiles(srcs []SourceFile, opts *Options) error {
	var err error
	for _, src := range srcs {
		err = CompileFile(src, opts)
		if err != nil {
			return err
		}
	}
	return nil
}

func Build(srcs []SourceFile, opts *Options) error {
	var err error

	err = CompileFiles(srcs, opts)
	if err != nil {
		return err
	}

	if !opts.IsLinkRequired() {
		return nil
	}

	Link(opts)

	return nil
}

func Run(args []string) {
	opts, err := ParseOptions(args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %v\n", err)
		os.Exit(1)
	}
	srcs := opts.SourceFiles()
	Build(srcs, opts)
}
