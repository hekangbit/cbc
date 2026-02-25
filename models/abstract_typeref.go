package models

import (
	"cbc/util"
	"fmt"
)

type ITypeRef interface {
	fmt.Stringer
	Location() *Location
}

// TODO: check hash algorithm, this is not the real object ptr
// make this as a global functino
func ITypeRefHashCode(t ITypeRef) int32 {
	return util.HashCode(util.ToString(t))
}

type BaseTypeRef struct {
	location *Location
}

var _ ITypeRef = (*BaseTypeRef)(nil)

func (this *BaseTypeRef) Location() *Location {
	return this.location
}

func (this *BaseTypeRef) String() string {
	panic("String method must be implemented by concrete type")
}
