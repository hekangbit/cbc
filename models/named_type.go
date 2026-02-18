package models

type INamedType interface {
	IType
	Name() string
	Location() Location
}

type NamedType struct {
	BaseType
	name     string
	location Location
}

func NewNamedType(name string, location Location) *NamedType {
	return &NamedType{
		name:     name,
		location: location,
	}
}

func (this *NamedType) Name() string {
	return this.name
}

func (this *NamedType) Location() Location {
	return this.location
}
