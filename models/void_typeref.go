package models

type VoidTypeRef struct {
	BaseTypeRef
}

func NewVoidTypeRef() *VoidTypeRef {
	return &VoidTypeRef{
		BaseTypeRef: BaseTypeRef{location: nil},
	}
}

func NewVoidTypeRefWithLocation(loc *Location) *VoidTypeRef {
	return &VoidTypeRef{
		BaseTypeRef: BaseTypeRef{location: loc},
	}
}

func (v *VoidTypeRef) IsVoid() bool {
	return true
}

func (v *VoidTypeRef) String() string {
	return "void"
}
