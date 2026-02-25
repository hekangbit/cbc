package models

import "fmt"

type Slot struct {
	Node
	typeNode *ASTTypeNode
	name     string
	offset   int64
}

var _ INode = &Slot{}

func NewSlot(typeNode *ASTTypeNode, name string) *Slot {
	p := &Slot{typeNode: typeNode, name: name, offset: SizeUnknown}
	p._impl = p
	return p
}

func (s *Slot) TypeNode() *ASTTypeNode {
	return s.typeNode
}

func (s *Slot) TypeRef() ITypeRef {
	return s.typeNode.TypeRef()
}

func (s *Slot) Type() IType {
	return s.typeNode.Type()
}

func (s *Slot) Name() string {
	return s.name
}

func (s *Slot) Size() int64 {
	return s.Type().Size()
}

func (s *Slot) AllocSize() int64 {
	return s.Type().AllocSize()
}

func (s *Slot) Alignment() int64 {
	return s.Type().Alignment()
}

func (s *Slot) Offset() int64 {
	return s.offset
}

func (s *Slot) SetOffset(offset int64) {
	s.offset = offset
}

func (s *Slot) Location() *Location {
	return s.typeNode.Location()
}

func (s *Slot) _Dump(d *Dumper) {
	d.PrintMemberStringNotResolved("name", s.name)
	d.PrintMemberDumpable("typeNode", s.typeNode)
}

func (s *Slot) String() string {
	return fmt.Sprintf("Slot{name: %s, type: %v, offset: %d}", s.name, s.Type(), s.offset)
}
