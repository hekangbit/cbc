package models

type ParamTypeRefs struct {
	*BaseParamSlots[ITypeRef]
}

// NewParams 构造函数
func NewParamTypeRefs(location *Location, typeRefs []ITypeRef, vararg bool) *ParamTypeRefs {
	return &ParamTypeRefs{
		BaseParamSlots: NewBaseParamSlots[ITypeRef](location, typeRefs, vararg),
	}
}
