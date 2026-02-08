package models

type NamedType interface {
	IType
	Name() string
	Location() Location
}

type BaseNamedType struct {
	BaseType
	name     string
	location Location
}

func NewBaseNamedType(name string, location Location) *BaseNamedType {
	return &BaseNamedType{
		name:     name,
		location: location,
	}
}

func (t *BaseNamedType) Name() string {
	return t.name
}

func (t *BaseNamedType) Location() Location {
	return t.location
}
