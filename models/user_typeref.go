package models

type UserTypeRef struct {
	BaseTypeRef
	name string
}

func NewUserTypeRef(name string) *UserTypeRef {
	return &UserTypeRef{
		BaseTypeRef: BaseTypeRef{location: nil},
		name:        name,
	}
}

func NewUserTypeRefWithLoc(loc *Location, name string) *UserTypeRef {
	return &UserTypeRef{
		BaseTypeRef: BaseTypeRef{location: loc},
		name:        name,
	}
}

func (this *UserTypeRef) IsUserType() bool {
	return true
}

func (this *UserTypeRef) Name() string {
	return this.name
}

func (this *UserTypeRef) Equals(other any) bool {
	otherRef, ok := other.(*UserTypeRef)
	if !ok {
		return false
	}
	return this.name == otherRef.name
}

func (this *UserTypeRef) String() string {
	return this.name
}
