package models

import (
	"cbc/utils"
	"fmt"
)

type UnionType struct {
	CompositeType
}

func NewUnionType(name string, members []*Slot, loc *Location) *UnionType {
	p := &UnionType{
		CompositeType: CompositeType{
			NamedType: NamedType{name: name, location: loc},
			members:   members,
		},
	}
	p._impl = p
	return p
}

func (this *UnionType) IsUnion() bool {
	return true
}

func (this *UnionType) IsSameType(other IType) bool {
	if other == nil {
		return false
	}
	otherTy, ok := other.(*UnionType)
	if !ok {
		return false
	}
	return this == otherTy
}

func (this *UnionType) ComputeOffsets() {
	var maxSize int64 = 0
	var maxAlign int64 = 1
	for _, s := range this.Members() {
		s.SetOffset(0)
		if size := s.AllocSize(); size > maxSize {
			maxSize = size
		}
		if align := s.Alignment(); align > maxAlign {
			maxAlign = align
		}
	}
	this.cachedSize = utils.Align(maxSize, maxAlign)
	this.cachedAlign = maxAlign
}

func (u *UnionType) String() string {
	return fmt.Sprintf("union %s", u.name)
}
