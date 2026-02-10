package models

type PointerTypeRef struct {
	*BaseTypeRef
	elemTypeRef ITypeRef
}

var _ ITypeRef = (*PointerTypeRef)(nil)

// NewPointerTypeRef 创建指针类型引用
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

// IsPointer 检查是否为指针类型引用
func (this *PointerTypeRef) IsPointer() bool {
	return true
}

// BaseType 返回基类型
func (this *PointerTypeRef) ElemType() ITypeRef {
	return this.elemTypeRef
}

// Equals 检查两个指针类型引用是否相等
func (this *PointerTypeRef) Equals(other interface{}) bool {
	otherRef, ok := other.(*PointerTypeRef)
	if !ok {
		return false
	}
	return this.elemTypeRef == otherRef.elemTypeRef
}

// String 返回字符串表示
func (this *PointerTypeRef) String() string {
	return this.elemTypeRef.String() + "*"
}

func (this *PointerTypeRef) SetElemType(ref ITypeRef) {
	this.elemTypeRef = ref
	this.location = ref.Location()
}
