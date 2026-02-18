package models

import (
	"errors"
	"fmt"
)

type ICompositeType interface {
	INamedType
	Members() []*Slot
	MemberTypes() []IType
	HasMember(name string) bool
	MemberType(name string) (IType, error)
	MemberOffset(name string) (int64, error)
	GetMember(name string) (*Slot, error)
}

type CompositeType struct {
	NamedType
	members            []*Slot
	cachedSize         int64
	cachedAlign        int64
	isRecursiveChecked bool
}

func NewCompositeType(name string, members []*Slot, location Location) *CompositeType {
	return &CompositeType{
		NamedType:          *NewNamedType(name, location),
		members:            members,
		cachedSize:         SizeUnknown,
		cachedAlign:        SizeUnknown,
		isRecursiveChecked: false,
	}
}

func (this *CompositeType) IsCompositeType() bool {
	return true
}

func (this *CompositeType) IsSameType(other IType) bool {
	return this.CompareMemberTypes(other, "IsSameType")
}

func (this *CompositeType) IsCompatible(target IType) bool {
	return this.CompareMemberTypes(target, "IsCompatible")
}

func (this *CompositeType) IsCastableTo(target IType) bool {
	return this.CompareMemberTypes(target, "IsCastableTo")
}

func (this *CompositeType) CompareMemberTypes(other IType, cmpType string) bool {
	if this.IsStruct() && !other.IsStruct() {
		return false
	}
	if this.IsUnion() && !other.IsUnion() {
		return false
	}

	otherComposite := other.GetCompositeType()

	otherMembers := otherComposite.Members()

	// check member number
	if len(this.members) != int(other.Size()) {
		return false
	}

	// check each member type
	for i, member := range this.members {
		otherMember := otherMembers[i]

		var cmpResult bool
		switch cmpType {
		case "IsSameType":
			cmpResult = member.Type().IsSameType(otherMember.Type())
		case "IsCompatible":
			cmpResult = member.Type().IsCompatible(otherMember.Type())
		case "IsCastableTo":
			cmpResult = member.Type().IsCastableTo(otherMember.Type())
		default:
			return false
		}

		if !cmpResult {
			return false
		}
	}

	return true
}

func (this *CompositeType) Size() int64 {
	if this.cachedSize == SizeUnknown {
		this.ComputeOffsets()
	}
	return this.cachedSize
}

func (this *CompositeType) AllocSize() int64 {
	return this.Size()
}

func (this *CompositeType) Alignment() int64 {
	if this.cachedAlign == SizeUnknown {
		this.ComputeOffsets()
	}
	return this.cachedAlign
}

func (this *CompositeType) Members() ([]*Slot, error) {
	return this.members, nil
}

func (this *CompositeType) MemberTypes() []IType {
	result := make([]IType, len(this.members))
	for i, slot := range this.members {
		result[i] = slot.Type()
	}
	return result
}

func (this *CompositeType) HasMember(name string) bool {
	_, err := this.GetMember(name)
	return err == nil
}

func (this *CompositeType) MemberType(name string) (IType, error) {
	slot, err := this.FetchMember(name)
	if err != nil {
		return nil, err
	}
	return slot.Type(), nil
}

func (this *CompositeType) MemberOffset(name string) (int64, error) {
	slot, err := this.FetchMember(name)
	if err != nil {
		return 0, err
	}

	if slot.offset == SizeUnknown {
		this.ComputeOffsets()
	}
	return slot.offset, nil
}

func (this *CompositeType) ComputeOffsets() {
	panic("ComputeOffsets must be implemented by concrete composite type")
}

func (this *CompositeType) FetchMember(name string) (*Slot, error) {
	slot, err := this.GetMember(name)
	if err != nil {
		return nil, errors.New("no such member in " + this.Name() + ": " + name)
	}
	return slot, nil
}

func (this *CompositeType) GetMember(name string) (*Slot, error) {
	for _, slot := range this.members {
		if slot.name == name {
			return slot, nil
		}
	}
	return nil, errors.New("member not found")
}

func (this *CompositeType) String() string {
	memberCount := len(this.members)
	if memberCount == 0 {
		return this.Name() + " {}"
	}

	return fmt.Sprintf("%s {%d members}", this.Name(), memberCount)
}
