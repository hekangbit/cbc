package models

type Function interface {
	IEntity
}

type BaseFunction struct {
	BaseEntity
	// callingSymbol Symbol
	// label         Label
}
