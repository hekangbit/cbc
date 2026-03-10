package models

import (
	"cbc/utils"
	"fmt"
)

const (
	markChecking = 1
	markChecked  = 2
)

type TypeTable struct {
	intSize     int
	longSize    int
	pointerSize int
	table       map[ITypeRef]IType
}

func NewTypeTable(intSize, longSize, pointerSize int) *TypeTable {
	return &TypeTable{
		intSize:     intSize,
		longSize:    longSize,
		pointerSize: pointerSize,
		table:       make(map[ITypeRef]IType),
	}
}

func ILP32() *TypeTable { return newTable(1, 2, 4, 4, 4) }
func ILP64() *TypeTable { return newTable(1, 2, 8, 8, 8) }
func LP64() *TypeTable  { return newTable(1, 2, 4, 8, 8) }
func LLP64() *TypeTable { return newTable(1, 2, 4, 4, 8) }

func newTable(charSize, shortSize, intSize, longSize, ptrSize int) *TypeTable {
	tt := NewTypeTable(intSize, longSize, ptrSize)
	tt.Put(NewVoidTypeRef(), &VoidType{})
	tt.Put(NewCharRef(), &IntegerType{size: int64(charSize), isSigned: true, name: "char"})
	tt.Put(NewShortRef(), &IntegerType{size: int64(shortSize), isSigned: true, name: "short"})
	tt.Put(NewIntRef(), &IntegerType{size: int64(intSize), isSigned: true, name: "int"})
	tt.Put(NewLongRef(), &IntegerType{size: int64(longSize), isSigned: true, name: "long"})
	tt.Put(NewUCharRef(), &IntegerType{size: int64(charSize), isSigned: false, name: "unsigned char"})
	tt.Put(NewUShortRef(), &IntegerType{size: int64(shortSize), isSigned: false, name: "unsigned short"})
	tt.Put(NewUIntRef(), &IntegerType{size: int64(intSize), isSigned: false, name: "unsigned int"})
	tt.Put(NewULongRef(), &IntegerType{size: int64(longSize), isSigned: false, name: "unsigned long"})
	return tt
}

func (tt *TypeTable) Put(ref ITypeRef, t IType) {
	if _, ok := tt.table[ref]; ok {
		panic(fmt.Sprintf("duplicated type definition: %v", ref))
	}
	tt.table[ref] = t
}

func (tt *TypeTable) IsDefined(ref ITypeRef) bool {
	_, ok := tt.table[ref]
	return ok
}

func (tt *TypeTable) Get(ref ITypeRef) IType {
	if t, ok := tt.table[ref]; ok {
		return t
	}
	switch r := ref.(type) {
	case *UserTypeRef:
		panic(fmt.Sprintf("undefined type: %s", r.Name()))
	case *PointerTypeRef:
		base := tt.Get(r.ElemType())
		t := NewPointerType(int64(tt.pointerSize), base) // &PointerType{size: int64(tt.pointerSize), elemType: base}
		tt.table[ref] = t
		return t
	case *ArrayTypeRef:
		base := tt.Get(r.ElemType())
		t := NewArrayTypeWithLen(base, r.Length(), int64(tt.pointerSize)) // &ArrayType{BaseType: base, Length: r.Length}
		tt.table[ref] = t
		return t
	case *FunctionTypeRef:
		ret := tt.Get(r.ReturnType())
		t := NewFunctionType(ret, r.Params().InternTypes(tt)) // &FunctionType{ReturnType: ret, ParamTypes: params}
		tt.table[ref] = t
		return t
	default:
		// TODO: java throw new Error("unregistered type: " + ref.toString()); so golang need return error
		panic(fmt.Sprintf("unregistered type: %v", ref))
	}
}

func (tt *TypeTable) GetParamType(ref ITypeRef) IType {
	t := tt.Get(ref)
	if arr, ok := t.(*ArrayType); ok {
		return tt.PointerTo(arr.ElemType())
	}
	return t
}

func (tt *TypeTable) PointerTo(elemType IType) *PointerType {
	return NewPointerType(int64(tt.pointerSize), elemType)
}

func (tt *TypeTable) voidType() *VoidType         { return (tt.Get(NewVoidTypeRef())).(*VoidType) }
func (tt *TypeTable) signedChar() *IntegerType    { return (tt.Get(NewCharRef())).(*IntegerType) }
func (tt *TypeTable) signedShort() *IntegerType   { return (tt.Get(NewShortRef())).(*IntegerType) }
func (tt *TypeTable) signedInt() *IntegerType     { return (tt.Get(NewIntRef())).(*IntegerType) }
func (tt *TypeTable) signedLong() *IntegerType    { return (tt.Get(NewLongRef())).(*IntegerType) }
func (tt *TypeTable) unsignedChar() *IntegerType  { return (tt.Get(NewUCharRef())).(*IntegerType) }
func (tt *TypeTable) unsignedShort() *IntegerType { return (tt.Get(NewUShortRef())).(*IntegerType) }
func (tt *TypeTable) unsignedInt() *IntegerType   { return (tt.Get(NewUIntRef())).(*IntegerType) }
func (tt *TypeTable) unsignedLong() *IntegerType  { return (tt.Get(NewULongRef())).(*IntegerType) }

func (tt *TypeTable) IntSize() int                    { return tt.intSize }
func (tt *TypeTable) LongSize() int                   { return tt.longSize }
func (tt *TypeTable) PointerSize() int                { return tt.pointerSize }
func (tt *TypeTable) MaxIntSize() int                 { return tt.pointerSize }
func (tt *TypeTable) SignedStackType() *IntegerType   { return tt.signedLong() }
func (tt *TypeTable) UnsignedStackType() *IntegerType { return tt.unsignedLong() }

func (tt *TypeTable) PtrDiffType() IType {
	return tt.Get(tt.PtrDiffTypeRef())
}

func (tt *TypeTable) PtrDiffTypeRef() ITypeRef {
	name := tt.ptrDiffTypeName()
	return NewIntegerTypeRef(name)
}

func (tt *TypeTable) ptrDiffTypeName() string {
	if tt.signedLong().size == int64(tt.pointerSize) {
		return "long"
	}
	if tt.signedInt().size == int64(tt.pointerSize) {
		return "int"
	}
	if tt.signedShort().size == int64(tt.pointerSize) {
		return "short"
	}
	panic("must not happen: integer.size != pointer.size")
}

func (tt *TypeTable) Types() []IType {
	types := make([]IType, 0, len(tt.table))
	for _, t := range tt.table {
		types = append(types, t)
	}
	return types
}

func (tt *TypeTable) SemanticCheck(h utils.ErrorHandler) {
	for _, t := range tt.Types() {
		if ct, ok := t.(*CompositeType); ok {
			tt.checkVoidMembersComposite(ct, h) // check void field in struct or union
			tt.checkDuplicatedMembers(ct, h)    // check dupliate field name
		} else if at, ok := t.(*ArrayType); ok {
			tt.checkVoidMembersArray(at, h) // check void element type of array type
		}
		tt.checkRecursiveDefinition(t, h)
	}
}

func (tt *TypeTable) checkVoidMembersArray(at *ArrayType, h utils.ErrorHandler) {
	if _, ok := at.ElemType().(*VoidType); ok {
		h.Error("array cannot contain void")
	}
}

func (tt *TypeTable) checkVoidMembersComposite(ct *CompositeType, h utils.ErrorHandler) {
	for _, s := range ct.Members() {
		if _, ok := s.Type().(*VoidType); ok {
			h.Error("struct/union cannot contain void")
		}
	}
}

func (tt *TypeTable) checkDuplicatedMembers(ct *CompositeType, h utils.ErrorHandler) {
	seen := make(map[string]bool)
	for _, s := range ct.Members() {
		if seen[s.Name()] {
			h.ErrorWithLoc(ct.Location(), ct.String()+" has duplicated member: "+s.Name())
		}
		seen[s.Name()] = true
	}
}

func (tt *TypeTable) checkRecursiveDefinition(t IType, h utils.ErrorHandler) {
	marks := make(map[IType]int)
	tt.checkRecursiveDefinitionInternal(t, marks, h)
}

func (tt *TypeTable) checkRecursiveDefinitionInternal(t IType, marks map[IType]int, h utils.ErrorHandler) {
	if marks[t] == markChecking {
		nameTy := t.(INamedType)
		h.ErrorWithLoc(nameTy.Location(), fmt.Sprintf("recursive type definition: %s", t.String()))
		return
	}
	if marks[t] == markChecked {
		return
	}
	marks[t] = markChecking
	switch typ := t.(type) {
	case *CompositeType:
		for _, s := range typ.Members() {
			tt.checkRecursiveDefinitionInternal(s.Type(), marks, h)
		}
	case *ArrayType:
		tt.checkRecursiveDefinitionInternal(typ.ElemType(), marks, h)
	case *UserType:
		tt.checkRecursiveDefinitionInternal(typ.RealType(), marks, h)
	}
	marks[t] = markChecked
}
