package models

import (
	"fmt"
)

type ICompositeType interface {
	INamedType
	HasMember(name string) bool
	Members() []*Slot
	MemberType(name string) (IType, error)
	MemberTypes() []IType
	MemberOffset(name string) (int64, error)
	FetchMember(name string) (*Slot, error)
	GetMember(name string) (*Slot, error)
	CompareMemberTypes(other IType, cmpType string) bool
	ComputeOffsets()
}

type CompositeType struct {
	NamedType
	members            []*Slot
	cachedSize         int64
	cachedAlign        int64
	isRecursiveChecked bool
	_impl_comp_type    ICompositeType
}

func (this *CompositeType) Size() int64 {
	if this.cachedSize == SizeUnknown {
		this._impl_comp_type.ComputeOffsets()
	}
	return this.cachedSize
}

func (this *CompositeType) AllocSize() int64 {
	return this.Size()
}

func (this *CompositeType) Alignment() int64 {
	if this.cachedAlign == SizeUnknown {
		this._impl_comp_type.ComputeOffsets()
	}
	return this.cachedAlign
}

func (this *CompositeType) IsCompositeType() bool {
	return true
}

func (this *CompositeType) IsSameType(other IType) bool {
	return this._impl_comp_type.CompareMemberTypes(other, "IsSameType")
}

func (this *CompositeType) IsCompatible(target IType) bool {
	return this._impl_comp_type.CompareMemberTypes(target, "IsCompatible")
}

func (this *CompositeType) IsCastableTo(target IType) bool {
	return this._impl_comp_type.CompareMemberTypes(target, "IsCastableTo")
}

func (this *CompositeType) HasMember(name string) bool {
	_, err := this.GetMember(name)
	return err == nil
}

func (this *CompositeType) Members() []*Slot {
	return this.members
}

func (this *CompositeType) MemberType(name string) (IType, error) {
	slot, err := this.FetchMember(name)
	if err != nil {
		return nil, err
	}
	return slot.Type(), nil
}

func (this *CompositeType) MemberTypes() []IType {
	result := make([]IType, len(this.members))
	for i, slot := range this.members {
		result[i] = slot.Type()
	}
	return result
}

func (this *CompositeType) MemberOffset(name string) (int64, error) {
	slot, err := this.FetchMember(name)
	if err != nil {
		return 0, err
	}

	if slot.offset == SizeUnknown {
		this._impl_comp_type.ComputeOffsets()
	}
	return slot.offset, nil
}

func (this *CompositeType) FetchMember(name string) (*Slot, error) {
	slot, err := this.GetMember(name)
	if err != nil {
		return nil, err
	}
	return slot, nil
}

func (this *CompositeType) GetMember(name string) (*Slot, error) {
	for _, slot := range this.members {
		if slot.name == name {
			return slot, nil
		}
	}
	return nil, fmt.Errorf("member <%s> not found in %s", name, this.Name())
}

// TODO: how to use type size to compare composite, members size means array length?
func (this *CompositeType) CompareMemberTypes(other IType, cmpType string) bool {
	if this.IsStruct() && !other.IsStruct() {
		return false
	}
	if this.IsUnion() && !other.IsUnion() {
		return false
	}
	if len(this.members) != int(other.Size()) {
		return false
	}

	// check each member type
	otherComposite := other.GetCompositeType()
	otherMembers := otherComposite.Members()
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
