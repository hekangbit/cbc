package models

// TypeNode 类型节点
type TypeNode struct {
	*BaseNode
	typeRef ITypeRef
	typ     IType
}

// NewTypeNodeFromRef 从类型引用创建类型节点
func NewTypeNodeFromRef(ref ITypeRef) *TypeNode {
	return &TypeNode{
		BaseNode: &BaseNode{},
		typeRef:  ref,
		typ:      nil,
	}
}

// NewTypeNodeFromType 从类型创建类型节点
func NewTypeNodeFromType(typ IType) *TypeNode {
	return &TypeNode{
		BaseNode: &BaseNode{},
		typeRef:  nil,
		typ:      typ,
	}
}

// TypeRef 返回类型引用
func (tn *TypeNode) TypeRef() ITypeRef {
	return tn.typeRef
}

// Type 返回类型（如果未解析则返回错误）
func (tn *TypeNode) Type() IType {
	if tn.typ == nil {
		panic("TypeNode not resolved and no typeRef available")
	}
	return tn.typ
}

// IsResolved 检查类型是否已解析
func (tn *TypeNode) IsResolved() bool {
	return tn.typ != nil
}

// SetType 设置类型（只能调用一次）
func (tn *TypeNode) SetType(typ IType) {
	if tn.typ != nil {
		panic("TypeNode#SetType called twice")
	}
	tn.typ = typ
}

// Location 返回位置信息
func (tn *TypeNode) Location() *Location {
	if tn.typeRef != nil {
		return tn.typeRef.Location()
	}
	return nil
}

// dump 实现内部转储方法
func (tn *TypeNode) Dump(d *Dumper) {
	// d.PrintField("typeRef", tn.typeRef)
	// d.PrintField("type", tn.typ)
}
