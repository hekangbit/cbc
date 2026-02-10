package models

import "github.com/antlr4-go/antlr/v4"

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
