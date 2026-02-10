package models

import (
	"cbc/util"
	"fmt"
)

type ITypeRef interface {
	fmt.Stringer
	Location() *Location
	HashCode() int32
}

type BaseTypeRef struct {
	location *Location
}

var _ ITypeRef = (*BaseTypeRef)(nil)

func NewBaseTypeRef(loc *Location) *BaseTypeRef {
	return &BaseTypeRef{location: loc}

}

func (p *BaseTypeRef) Location() *Location {
	return p.location
}

func (p *BaseTypeRef) HashCode() int32 {
	return util.HashCode(util.ToString(p))
}

func (p *BaseTypeRef) String() string {
	panic("String method must be implemented by concrete type")
}
