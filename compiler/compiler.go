package compiler

import (
	"cbc/models"
	"cbc/parser"
	"cbc/sysdep/x86"
	"cbc/utils"
	"fmt"
	"os"

	"github.com/antlr4-go/antlr/v4"
)

var CompilerProgramName string = "cbc"
var CompilerVersion string = "1.0.0"

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

func ParseFile(path string, opts *Options, h *utils.ErrorHandler) *models.AST {
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

	builder := NewASTBuilder(path, h)
	program := tree.Accept(builder)
	cbAST := program.(*models.AST)
	cbAST.SetStream(cbLexer, tokenStream)
	return cbAST
}

func SemanticAnalyzeResolvePhase(astNode *models.AST, typeTable *models.TypeTable, opts *Options, errorHandler *utils.ErrorHandler) error {
	var err error
	localResolver := NewLocalResolver(errorHandler)
	err = localResolver.Resolve(astNode)
	if err != nil {
		return err
	}
	typeResolver := NewTypeResolver(typeTable, errorHandler)
	err = typeResolver.Resolve(astNode)
	if err != nil {
		return err
	}
	return nil
}

func SemanticAnalyzeCheckPhase(astNode *models.AST, typeTable *models.TypeTable, opts *Options, errorHandler *utils.ErrorHandler) error {
	var err error
	err = typeTable.SemanticCheck(errorHandler)
	if err != nil {
		return err
	}
	return nil
}

func GenerateIR(sem *models.AST, typeTable *models.TypeTable, errorHandler *utils.ErrorHandler) *models.IR {
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

func CompileToAsm(srcPath string, dstPath string, opts *Options, errorHandler *utils.ErrorHandler) error {
	var err error
	fmt.Println("CompileToAsm: " + srcPath + " to " + dstPath)
	typeTable := opts.TypeTable()

	// 1. parse file
	astObj := ParseFile(srcPath, opts, errorHandler)
	if errorHandler.ErrorOccured() {
		return fmt.Errorf("compile parse file <%s> fail", srcPath)
	}
	if DumpAST(astObj, opts.Mode()) {
		return nil
	}

	// 2. semantic analyze
	err = SemanticAnalyzeResolvePhase(astObj, typeTable, opts, errorHandler)
	if err != nil {
		return err
	}
	if DumpSemant(astObj, opts.Mode()) {
		return nil
	}
	err = SemanticAnalyzeCheckPhase(astObj, typeTable, opts, errorHandler)
	if err != nil {
		return err
	}

	// 3. generate ir
	// ir := GenerateIR(semObj, typeTable, errorHandler)
	// if DumpIR(ir, opts.Mode()) {
	// 	return
	// }

	// 4. generate asm
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

func CompileFile(src SourceFile, opts *Options, errorHandler *utils.ErrorHandler) error {
	var err error

	// 1. .cb -> .s
	if src.IsCbSource() {
		destPath := opts.AsmFileNameOf(src)
		err = CompileToAsm(src.Path(), destPath, opts, errorHandler)
		if err != nil {
			return err
		}
		src.SetCurrentName(destPath)
	}

	// 2. .s -> .o
	if !opts.IsAssembleRequired() {
		return nil
	}
	if src.IsAssemblySource() {
		destPath := opts.ObjFileNameOf(src)
		AssembleToObj(src.Path(), destPath, opts)
		src.SetCurrentName(destPath)
	}

	return nil
}

func CompileFiles(srcs []SourceFile, opts *Options, errorHandler *utils.ErrorHandler) error {
	var err error
	for _, src := range srcs {
		err = CompileFile(src, opts, errorHandler)
		if err != nil {
			return err
		}
	}
	return nil
}

func GenerateExecutable(opts *Options) {
}

func GenerateSharedLibrary(opts *Options) {
}

func Link(opts *Options) {
	fmt.Println("Link")
	if !opts.IsGeneratingSharedLibrary() {
		GenerateExecutable(opts)
	} else {
		GenerateSharedLibrary(opts)
	}
}

func Build(srcs []SourceFile, opts *Options, errorHandler *utils.ErrorHandler) error {
	var err error

	err = CompileFiles(srcs, opts, errorHandler)
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
	var errorHandler = utils.NewErrorHandler(CompilerProgramName)
	opts, err := ParseOptions(args)
	if err != nil {
		errorHandler.Error(err.Error())
		fmt.Println("Try \"cbc --help\" for usage")
		os.Exit(1)
	}
	srcs := opts.SourceFiles()
	Build(srcs, opts, errorHandler)
	if errorHandler.IssueOccured() {
		errorHandler.DumpTotal()
	}
}
