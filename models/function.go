package models

type Function interface {
	IEntity
}

type BaseFunction struct {
	*BaseEntity
	// callingSymbol Symbol
	// label         Label
}

func NewBaseFunction(isPriv bool, typeNode *TypeNode, name string) *BaseFunction {
	return &BaseFunction{
		BaseEntity: NewBaseEntity(isPriv, typeNode, name),
	}
}
