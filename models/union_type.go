package models

import (
	"cbc/utils"
	"fmt"
)

type UnionType struct {
	CompositeType
}

var _ ICompositeType = &UnionType{}

func NewUnionType(name string, members []*Slot, loc *Location) *UnionType {
	p := new(UnionType)
	p.name = name
	p.location = loc
	p.cachedSize = SizeUnknown
	p.cachedAlign = SizeUnknown
	p.isRecursiveChecked = false
	p.members = members
	p._impl = p
	p._impl_comp_type = p
	return p
}

func (this *UnionType) IsUnion() bool {
	return true
}

func (this *UnionType) IsSameType(other IType) bool {
	if other == nil {
		return false
	}
	return this == other.GetUnionType()
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
