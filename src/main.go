package main

import (
	"cbc/parser"
	"fmt"
	"os"

	"github.com/antlr4-go/antlr/v4"
)

func main() {
	argNum := len(os.Args)
	fmt.Println("argNum: ", argNum)

	for i := 1; i < argNum; i++ {
		arg := os.Args[i]

		src, err := os.ReadFile(arg)
		if err != nil {
			os.Exit(64)
		}
		input := antlr.NewInputStream(string(src))
		cbLexer := parser.NewCbLexer(input)
		tokenStream := antlr.NewCommonTokenStream(cbLexer, antlr.TokenDefaultChannel)
		cbParser := parser.NewCbParser(tokenStream)
		tree := cbParser.Prog()
		fmt.Println(tree.ToStringTree(cbParser.RuleNames, cbParser))

		visitor := &CbVisitorImpl{BaseCbVisitor: &parser.BaseCbVisitor{}}
		visitor.Visit(tree) // tree.Accept(visitor)

		for {
			t := cbLexer.NextToken()
			if t.GetTokenType() == antlr.TokenEOF {
				break
			}
			fmt.Printf("%s (%q)\n", cbLexer.SymbolicNames[t.GetTokenType()], t.GetText())
		}
	}
}
