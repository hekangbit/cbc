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

func (this *BaseTypeRef) Location() *Location {
	return this.location
}

func (this *BaseTypeRef) HashCode() int32 {
	return util.HashCode(util.ToString(this))
}

func (this *BaseTypeRef) String() string {
	panic("String method must be implemented by concrete type")
}
