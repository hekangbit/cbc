package models

import (
	"fmt"
)

type ITypeRef interface {
	fmt.Stringer
	Location() *Location
}

type BaseTypeRef struct {
	location *Location
}

func (this *BaseTypeRef) Location() *Location {
	return this.location
}
