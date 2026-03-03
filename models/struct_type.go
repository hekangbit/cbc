package models

import (
	"cbc/utils"
	"fmt"
)

type StructType struct {
	CompositeType
}

func NewStructType(name string, membs []*Slot, loc *Location) *StructType {
	p := &StructType{
		CompositeType: CompositeType{
			NamedType:          NamedType{name: name, location: loc},
			cachedSize:         SizeUnknown,
			cachedAlign:        SizeUnknown,
			isRecursiveChecked: false,
			members:            membs,
		},
	}
	p._impl = p
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

func (this *StructType) ComputeOffset() {
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
