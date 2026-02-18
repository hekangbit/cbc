package models

import "reflect"

type ParamTypeRefs struct {
	*ParamSlots[ITypeRef]
}

// NewParams 构造函数
func NewParamTypeRefs(location *Location, paramDescs []ITypeRef, vararg bool) *ParamTypeRefs {
	return &ParamTypeRefs{
		ParamSlots: NewParamSlotsFull(location, paramDescs, vararg),
	}
}

func (this *ParamTypeRefs) TypeRefs() []ITypeRef {
	return this.paramDescriptors
}

func (this *ParamTypeRefs) InternTypes(table TypeTable) *ParamTypes {
	cbtypes := make([]IType, 0)
	for _, ref := range this.paramDescriptors {
		cbtypes = append(cbtypes, table.GetParamType(ref))
	}
	return NewParamTypes(this.location, cbtypes, this.vararg)
}

func (this *ParamTypeRefs) Equals(other interface{}) bool {
	otherRef, ok := other.(*ParamTypeRefs)
	if !ok {
		return false
	}

	if this.vararg != otherRef.vararg {
		return false
	}

	if len(this.paramDescriptors) != len(otherRef.paramDescriptors) {
		return false
	}
	// TODO: need check the deep equal correctness
	return reflect.DeepEqual(this.paramDescriptors, otherRef.paramDescriptors)
}
