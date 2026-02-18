package models

type IntegerTypeRef struct {
	*BaseTypeRef
	name string
}

func NewCharRef() *IntegerTypeRef {
	return NewIntegerTypeRef("char")
}

func NewCharRefWithLocation(loc *Location) *IntegerTypeRef {
	return NewIntegerTypeRefWithLocation("char", loc)
}

func NewShortRef() *IntegerTypeRef {
	return NewIntegerTypeRef("short")
}

func NewShortRefWithLocation(loc *Location) *IntegerTypeRef {
	return NewIntegerTypeRefWithLocation("short", loc)
}

func NewIntRef() *IntegerTypeRef {
	return NewIntegerTypeRef("int")
}

func NewIntRefWithLocation(loc *Location) *IntegerTypeRef {
	return NewIntegerTypeRefWithLocation("int", loc)
}

func NewLongRef() *IntegerTypeRef {
	return NewIntegerTypeRef("long")
}

func NewLongRefWithLocation(loc *Location) *IntegerTypeRef {
	return NewIntegerTypeRefWithLocation("long", loc)
}

func NewUcharRef() *IntegerTypeRef {
	return NewIntegerTypeRef("unsigned char")
}

func NewUcharRefWithLocation(loc *Location) *IntegerTypeRef {
	return NewIntegerTypeRefWithLocation("unsigned char", loc)
}

func NewUshortRef() *IntegerTypeRef {
	return NewIntegerTypeRef("unsigned short")
}

func NewUshortRefWithLocation(loc *Location) *IntegerTypeRef {
	return NewIntegerTypeRefWithLocation("unsigned short", loc)
}

func NewUintRef() *IntegerTypeRef {
	return NewIntegerTypeRef("unsigned int")
}

func NewUintRefWithLocation(loc *Location) *IntegerTypeRef {
	return NewIntegerTypeRefWithLocation("unsigned int", loc)
}

func NewUlongRef() *IntegerTypeRef {
	return NewIntegerTypeRef("unsigned long")
}

func NewUlongRefWithLocation(loc *Location) *IntegerTypeRef {
	return NewIntegerTypeRefWithLocation("unsigned long", loc)
}

func NewIntegerTypeRef(name string) *IntegerTypeRef {
	return NewIntegerTypeRefWithLocation(name, nil)
}

func NewIntegerTypeRefWithLocation(name string, loc *Location) *IntegerTypeRef {
	return &IntegerTypeRef{
		BaseTypeRef: NewBaseTypeRef(loc),
		name:        name,
	}
}

func (i *IntegerTypeRef) Name() string {
	return i.name
}

func (i *IntegerTypeRef) Equals(other interface{}) bool {
	otherRef, ok := other.(*IntegerTypeRef)
	if !ok {
		return false
	}
	return i.name == otherRef.name
}

func (i *IntegerTypeRef) String() string {
	return i.name
}
