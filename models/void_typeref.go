package models

// VoidTypeRef void类型引用
type VoidTypeRef struct {
	*BaseTypeRef
}

// NewVoidTypeRef 创建无位置信息的void类型引用
func NewVoidTypeRef() *VoidTypeRef {
	return &VoidTypeRef{
		BaseTypeRef: NewBaseTypeRef(nil),
	}
}

// NewVoidTypeRefWithLocation 创建带位置信息的void类型引用
func NewVoidTypeRefWithLocation(loc *Location) *VoidTypeRef {
	return &VoidTypeRef{
		BaseTypeRef: NewBaseTypeRef(loc),
	}
}

// IsVoid 重写：检查是否为void类型引用
func (v *VoidTypeRef) IsVoid() bool {
	return true
}

// Equals 检查两个VoidTypeRef是否相等
func (v *VoidTypeRef) Equals(other interface{}) bool {
	_, ok := other.(*VoidTypeRef)
	return ok
}

// String 返回字符串表示
func (v *VoidTypeRef) String() string {
	return "void"
}
