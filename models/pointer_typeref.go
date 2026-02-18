package models

type PointerTypeRef struct {
	*BaseTypeRef
	elemTypeRef ITypeRef
}

var _ ITypeRef = (*PointerTypeRef)(nil)

func NewPointerTypeRef(elemTypeRef ITypeRef) *PointerTypeRef {
	return &PointerTypeRef{
		BaseTypeRef: NewBaseTypeRef(elemTypeRef.Location()),
		elemTypeRef: elemTypeRef,
	}
}

func NewPointerTypeRefEmptyElem() *PointerTypeRef {
	return &PointerTypeRef{
		BaseTypeRef: NewBaseTypeRef(nil),
		elemTypeRef: nil,
	}
}

func (this *PointerTypeRef) IsPointer() bool {
	return true
}

func (this *PointerTypeRef) ElemType() ITypeRef {
	return this.elemTypeRef
}

func (this *PointerTypeRef) Equals(other interface{}) bool {
	otherRef, ok := other.(*PointerTypeRef)
	if !ok {
		return false
	}
	return this.elemTypeRef == otherRef.elemTypeRef
}

func (this *PointerTypeRef) String() string {
	return this.elemTypeRef.String() + "*"
}

func (this *PointerTypeRef) SetElemType(ref ITypeRef) {
	this.elemTypeRef = ref
	this.location = ref.Location()
}
