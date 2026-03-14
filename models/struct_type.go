package models

import (
	"cbc/utils"
	"fmt"
)

type StructType struct {
	CompositeType
}

var _ ICompositeType = &StructType{}

func NewStructType(name string, members []*Slot, loc *Location) *StructType {
	p := new(StructType)
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

func (this *StructType) IsStruct() bool {
	return true
}

func (this *StructType) IsSameType(other IType) bool {
	if !other.IsStruct() {
		return false
	}
	return this == other.GetStructType() // TODO: is this correct?
}

func (this *StructType) ComputeOffsets() {
	var offset int64 = 0
	var maxAlign int64 = 1
	for _, slot := range this.Members() {
		offset = utils.Align(offset, slot.AllocSize())
		slot.SetOffset(offset)
		offset += slot.AllocSize()
		if maxAlign < slot.Alignment() {
			maxAlign = slot.Alignment()
		}
	}
	this.cachedSize = utils.Align(offset, maxAlign)
	this.cachedAlign = maxAlign
}

func (this *StructType) String() string {
	return fmt.Sprintf("struct %s", this.name)
}
