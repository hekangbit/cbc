package models

import "reflect"

type ParamTypes struct {
	*ParamSlots[IType]
}

func NewParamTypes(location *Location, typeRefs []IType, vararg bool) *ParamTypes {
	return &ParamTypes{
		ParamSlots: NewParamSlotsFull[IType](location, typeRefs, vararg),
	}
}

var _ IParamSlots = (*ParamTypes)(nil)

func (this *ParamTypes) Types() []IType {
	return this.paramDescriptors
}

func (this *ParamTypes) IsSameType(other *ParamTypes) bool {
	if this.vararg != other.vararg {
		return false
	}
	if this.MinArgc() != other.MinArgc() {
		return false
	}
	for i, t := range this.paramDescriptors {
		if !t.IsSameType(other.paramDescriptors[i]) {
			return false
		}
	}
	return true
}

func (this *ParamTypes) Equals(other interface{}) bool {
	if other == nil {
		return false
	}
	otherType, ok := other.(*ParamTypes)
	if !ok {
		return false
	}
	if this.vararg != otherType.vararg {
		return false
	}

	return reflect.DeepEqual(this.paramDescriptors, otherType.paramDescriptors)
}
