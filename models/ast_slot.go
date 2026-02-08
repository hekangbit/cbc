package models

import "fmt"

// Slot 结构体成员槽位
type Slot struct {
	*BaseNode
	typeNode *TypeNode
	name     string
	offset   int64
}

// NewSlot 创建新的槽位
func NewSlot(typeNode *TypeNode, name string) *Slot {
	return &Slot{
		BaseNode: &BaseNode{},
		typeNode: typeNode,
		name:     name,
		offset:   SizeUnknown,
	}
}

// TypeNode 返回类型节点
func (s *Slot) TypeNode() *TypeNode {
	return s.typeNode
}

// TypeRef 返回类型引用
func (s *Slot) TypeRef() ITypeRef {
	return s.typeNode.TypeRef()
}

// Type 返回类型
func (s *Slot) Type() IType {
	return s.typeNode.Type()
}

// Name 返回名称
func (s *Slot) Name() string {
	return s.name
}

// Size 返回大小
func (s *Slot) Size() int64 {
	return s.Type().Size()
}

// AllocSize 返回分配大小
func (s *Slot) AllocSize() int64 {
	return s.Type().AllocSize()
}

// Alignment 返回对齐大小
func (s *Slot) Alignment() int64 {
	return s.Type().Alignment()
}

// Offset 返回偏移量
func (s *Slot) Offset() int64 {
	return s.offset
}

// SetOffset 设置偏移量
func (s *Slot) SetOffset(offset int64) {
	s.offset = offset
}

// Location 返回位置信息
func (s *Slot) Location() *Location {
	return s.typeNode.Location()
}

// dump 实现内部转储方法
func (s *Slot) Dump(d *Dumper) {
	// d.PrintField("name", s.name)
	// d.PrintField("typeNode", s.typeNode)
}

// String 返回字符串表示
func (s *Slot) String() string {
	return fmt.Sprintf("Slot{name: %s, type: %v, offset: %d}", s.name, s.Type(), s.offset)
}
