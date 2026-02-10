package models

// IntegerTypeRef 整数类型引用
type IntegerTypeRef struct {
	*BaseTypeRef
	name string
}

// 以下是静态工厂方法对应的包级函数

// CharRef 创建无位置信息的char类型引用
func NewCharRef() *IntegerTypeRef {
	return NewIntegerTypeRef("char")
}

// CharRefWithLocation 创建带位置信息的char类型引用
func NewCharRefWithLocation(loc *Location) *IntegerTypeRef {
	return NewIntegerTypeRefWithLocation("char", loc)
}

// ShortRef 创建无位置信息的short类型引用
func NewShortRef() *IntegerTypeRef {
	return NewIntegerTypeRef("short")
}

// ShortRefWithLocation 创建带位置信息的short类型引用
func NewShortRefWithLocation(loc *Location) *IntegerTypeRef {
	return NewIntegerTypeRefWithLocation("short", loc)
}

// IntRef 创建无位置信息的int类型引用
func NewIntRef() *IntegerTypeRef {
	return NewIntegerTypeRef("int")
}

// IntRefWithLocation 创建带位置信息的int类型引用
func NewIntRefWithLocation(loc *Location) *IntegerTypeRef {
	return NewIntegerTypeRefWithLocation("int", loc)
}

// LongRef 创建无位置信息的long类型引用
func NewLongRef() *IntegerTypeRef {
	return NewIntegerTypeRef("long")
}

// LongRefWithLocation 创建带位置信息的long类型引用
func NewLongRefWithLocation(loc *Location) *IntegerTypeRef {
	return NewIntegerTypeRefWithLocation("long", loc)
}

// UcharRef 创建无位置信息的unsigned char类型引用
func NewUcharRef() *IntegerTypeRef {
	return NewIntegerTypeRef("unsigned char")
}

// UcharRefWithLocation 创建带位置信息的unsigned char类型引用
func NewUcharRefWithLocation(loc *Location) *IntegerTypeRef {
	return NewIntegerTypeRefWithLocation("unsigned char", loc)
}

// UshortRef 创建无位置信息的unsigned short类型引用
func NewUshortRef() *IntegerTypeRef {
	return NewIntegerTypeRef("unsigned short")
}

// UshortRefWithLocation 创建带位置信息的unsigned short类型引用
func NewUshortRefWithLocation(loc *Location) *IntegerTypeRef {
	return NewIntegerTypeRefWithLocation("unsigned short", loc)
}

// UintRef 创建无位置信息的unsigned int类型引用
func NewUintRef() *IntegerTypeRef {
	return NewIntegerTypeRef("unsigned int")
}

// UintRefWithLocation 创建带位置信息的unsigned int类型引用
func NewUintRefWithLocation(loc *Location) *IntegerTypeRef {
	return NewIntegerTypeRefWithLocation("unsigned int", loc)
}

// UlongRef 创建无位置信息的unsigned long类型引用
func NewUlongRef() *IntegerTypeRef {
	return NewIntegerTypeRef("unsigned long")
}

// UlongRefWithLocation 创建带位置信息的unsigned long类型引用
func NewUlongRefWithLocation(loc *Location) *IntegerTypeRef {
	return NewIntegerTypeRefWithLocation("unsigned long", loc)
}

// 构造函数

// NewIntegerTypeRef 创建整数类型引用（不带位置信息）
func NewIntegerTypeRef(name string) *IntegerTypeRef {
	return NewIntegerTypeRefWithLocation(name, nil)
}

// NewIntegerTypeRefWithLocation 创建整数类型引用（带位置信息）
func NewIntegerTypeRefWithLocation(name string, loc *Location) *IntegerTypeRef {
	return &IntegerTypeRef{
		BaseTypeRef: NewBaseTypeRef(loc),
		name:        name,
	}
}

// Name 返回类型名称
func (i *IntegerTypeRef) Name() string {
	return i.name
}

// Equals 检查两个整数类型引用是否相等
func (i *IntegerTypeRef) Equals(other interface{}) bool {
	otherRef, ok := other.(*IntegerTypeRef)
	if !ok {
		return false
	}
	return i.name == otherRef.name
}

// String 返回字符串表示
func (i *IntegerTypeRef) String() string {
	return i.name
}
