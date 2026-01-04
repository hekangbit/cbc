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
		lexer := parser.NewCbLexer(input)
		tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
		parser := parser.NewCbParser(tokenStream)
		tree := parser.Prog()
		fmt.Println(tree.ToStringTree(parser.RuleNames, parser))

		for {
			t := lexer.NextToken()
			if t.GetTokenType() == antlr.TokenEOF {
				break
			}
			fmt.Printf("%s (%q)\n", lexer.SymbolicNames[t.GetTokenType()], t.GetText())
		}
	}
}
