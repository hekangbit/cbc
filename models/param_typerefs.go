package models

type ParamTypeRefs struct {
	*ParamSlots[ITypeRef]
}

func NewParamTypeRefs(location *Location, paramDescs []ITypeRef, vararg bool) *ParamTypeRefs {
	return &ParamTypeRefs{
		ParamSlots: NewParamSlotsFull(location, paramDescs, vararg),
	}
}

func (this *ParamTypeRefs) TypeRefs() []ITypeRef {
	return this.paramDescriptors
}

func (this *ParamTypeRefs) InternTypes(table *TypeTable) *ParamTypes {
	cbtypes := make([]IType, 0)
	for _, ref := range this.paramDescriptors {
		cbtypes = append(cbtypes, table.GetParamType(ref))
	}
	return NewParamTypes(this.location, cbtypes, this.vararg)
}
