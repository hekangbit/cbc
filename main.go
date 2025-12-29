package main

import (
	"cbc/parser"
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

func main() {
	input := antlr.NewInputStream("3+5*2")
	lexer := parser.NewCbLexer(input)
	for {
		t := lexer.NextToken()
		if t.GetTokenType() == antlr.TokenEOF {
			break
		}
		fmt.Printf("%s (%q)\n", lexer.SymbolicNames[t.GetTokenType()], t.GetText())
	}
}
