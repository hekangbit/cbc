package main

import (
	"cbc/ast"
	"cbc/parser"
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
	builder := &ast.ASTBuilder{BaseCbVisitor: &parser.BaseCbVisitor{}}
	program := tree.Accept(builder) // builder.Visit(tree)
	fmt.Println(program)
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

func Assemble(srcPath string, dstPath string, opts *Options) {
	fmt.Println("Assemble " + srcPath + " to " + dstPath)
}

func Compile(srcPath string, dstPath string, opts *Options) {
	fmt.Println("Compile " + srcPath + " to " + dstPath)
	// parse file
	// semantic analyze
	// generate ir
	// generate asm
	// write file

	// test code, dump all tokens, visit tree
	DebugDump(srcPath)
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

func main() {
	opts, err := ParseOptions(os.Args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Options: ", opts)
	srcs := opts.SourceFiles()
	Build(srcs, opts)
}
