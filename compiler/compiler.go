package compiler

import (
	"cbc/models"
	"cbc/parser"
	"cbc/sysdep/x86"
	"fmt"
	"os"

	"github.com/antlr4-go/antlr/v4"
)

var CompilerProgramName string = "cbc"
var CompilerVersion string = "1.0.0"

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
	return program.(*models.AST)
}

func SemanticAnalyze(astNode *models.AST, typeTable models.TypeTable, opts *Options) *models.AST {
	return astNode
}

func GenerateIR(astNode *models.AST, typeTable models.TypeTable) *models.IR {
	return &models.IR{}
}

func GenerateAssembly(ir *models.IR, opts *Options) x86.AssemblyCode {
	return x86.AssemblyCode{}
}

func WriteFile(path string, content string) {
}

func Compile(srcPath string, dstPath string, opts *Options) {
	var tree *models.AST
	fmt.Println("Compile " + srcPath + " to " + dstPath)
	// parse file
	// semantic analyze
	// generate ir
	// generate asm
	// write file

	typeTable := opts.GetTypeTable()
	tree = ParseFile(srcPath, opts)
	// if (DumpAST(tree, opts.mode())) {
	// 	return
	// }
	tree = SemanticAnalyze(tree, typeTable, opts)
	// if (DumpSemant(tree, opts.mode())) {
	// 	return
	// }
	ir := GenerateIR(tree, typeTable)
	// if (DumpIR(ir, opts.mode())) {
	// 	return
	// }
	asm := GenerateAssembly(ir, opts)
	// if dumpAsm(asm, opts.mode()) {
	// 	return
	// }
	// if printAsm(asm, opts.mode()) {
	// 	return
	// }
	WriteFile(dstPath, asm.String())
}

func Assemble(srcPath string, dstPath string, opts *Options) {
	fmt.Println("Assemble " + srcPath + " to " + dstPath)
}

func Link(opts *Options) {
	fmt.Println("Link")
	if !opts.IsGeneratingSharedLibrary() {
		GenerateExecutable(opts)
	} else {
		GenerateSharedLibrary(opts)
	}
}

func Build(srcs []SourceFile, opts *Options) {
	// compile all source files
	for _, src := range srcs {
		// .cb -> .s
		if src.IsCbSource() {
			destPath := opts.AsmFileNameOf(src)
			Compile(src.Path(), destPath, opts)
			src.SetCurrentName(destPath)
		}
		if !opts.IsAssembleRequired() {
			continue
		}
		// .s -> .o
		if src.IsAssemblySource() {
			destPath := opts.ObjFileNameOf(src)
			Assemble(src.Path(), destPath, opts)
			src.SetCurrentName(destPath)
		}
	}

	// link
	if !opts.IsLinkRequired() {
		return
	}
	Link(opts)
}

func Run(args []string) {
	opts, err := ParseOptions(args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Options: ", opts)
	srcs := opts.SourceFiles()
	Build(srcs, opts)
}
