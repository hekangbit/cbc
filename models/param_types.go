package models

type ParamTypes struct {
	*BaseParamSlots[IType]
}

// NewParams 构造函数
func NewParamTypes(location *Location, typeRefs []IType, vararg bool) *ParamTypes {
	return &ParamTypes{
		BaseParamSlots: NewBaseParamSlots[IType](location, typeRefs, vararg),
	}
}
