package models

import (
	"fmt"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

type Location struct {
	sourceName string
	token      antlr.Token
}

func NewLocation(source string, token antlr.Token) *Location {
	return &Location{
		sourceName: source,
		token:      token,
	}
}

func GetLineText(input antlr.CharStream, line int) string {
	whole := input.GetTextFromInterval(antlr.NewInterval(0, input.Size()-1))
	lines := strings.Split(whole, "\n")
	if line-1 < 0 || line-1 >= len(lines) {
		return ""
	}
	return lines[line-1]
}

func (this *Location) SourceName() string {
	return this.sourceName
}

func (this *Location) Token() antlr.Token {
	return this.token
}

func (this *Location) Lineno() int {
	return this.token.GetLine()
}

func (this *Location) Column() int {
	return this.token.GetColumn()
}

func (this *Location) Line() string {
	return GetLineText(this.token.GetInputStream(), this.token.GetLine())
}

func (this *Location) NumberedLine() string {
	return fmt.Sprintf("line %d: %s", this.Lineno(), this.Line())
}

func (this *Location) String() string {
	return fmt.Sprintf("%s:%d", this.sourceName, this.Lineno())
}
