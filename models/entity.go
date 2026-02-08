package models

import (
	"cbc/asm"
)

type IEntity interface {
	Dumpable
	Name() string
	SymbolString() string
	IsDefined() bool
	IsInitialized() bool
	IsConstant() bool
	Value() ExprNode
	IsParameter() bool
	IsPrivate() bool
	TypeNode() *TypeNode
	Type() IType
	AllocSize() int64
	Alignment() int64
	Refered()
	IsRefered() bool
	SetMemref(mem asm.IMemoryReference)
	Memref() asm.IMemoryReference
	SetAddressMem(mem asm.IMemoryReference)
	SetAddressImm(imm asm.IImmediateValue)
	Address() asm.IOperand
	CheckAddress()
	Location() *Location
	Accept() interface{}
	Dump(*Dumper)
	Dump_(*Dumper)
}

// Entity 基础结构体（对应 Java 的 abstract class Entity）
// 包含所有字段和已实现的方法；抽象方法默认 panic，由嵌入者覆盖
type BaseEntity struct {
	name      string
	isPrivate bool
	typeNode  *TypeNode
	nRefered  int64
	memref    asm.IMemoryReference
	address   asm.IOperand
}

// EntityVisitor 访问者接口
// 注：Go 不支持泛型方法，因此 Accept 返回 interface{}，由调用方类型断言
// type EntityVisitor interface {
// 	VisitUndefined(e *UndefinedEntity) interface{}
// 	VisitVariable(e *VariableEntity) interface{}
// 	VisitFunction(e *FunctionEntity) interface{}
// 	// ... 根据实际子类添加其他 Visit 方法
// }

// NewEntity 构造函数（对应 Java 的 public Entity(...)）
// 注：Go 没有 protected，通过包内可见性控制
func NewBaseEntity(priv bool, typ *TypeNode, name string) *BaseEntity {
	return &BaseEntity{
		name:      name,
		isPrivate: priv,
		typeNode:  typ,
		nRefered:  0,
	}
}

// Name 返回名称
func (e *BaseEntity) Name() string {
	return e.name
}

// SymbolString 返回符号字符串
func (e *BaseEntity) SymbolString() string {
	return e.Name()
}

// ========== 抽象方法（基类默认 panic，子类必须覆盖） ==========

// IsDefined 抽象方法
func (e *BaseEntity) IsDefined() bool {
	panic("abstract method: IsDefined")
}

// IsInitialized 抽象方法
func (e *BaseEntity) IsInitialized() bool {
	panic("abstract method: IsInitialized")
}

// ========== 已实现的具体方法 ==========

// IsConstant 默认返回 false，可变实体可覆盖
func (e *BaseEntity) IsConstant() bool {
	return false
}

// Value 默认抛出错误（panic）
func (e *BaseEntity) Value() ExprNode {
	panic("Entity#value")
}

// IsParameter 默认返回 false
func (e *BaseEntity) IsParameter() bool {
	return false
}

// IsPrivate 返回是否为私有
func (e *BaseEntity) IsPrivate() bool {
	return e.isPrivate
}

// TypeNode 返回类型节点
func (e *BaseEntity) TypeNode() *TypeNode {
	return e.typeNode
}

// Type 返回类型
func (e *BaseEntity) Type() IType {
	return e.typeNode.Type()
}

// AllocSize 返回分配大小
func (e *BaseEntity) AllocSize() int64 {
	return e.Type().AllocSize()
}

// Alignment 返回对齐要求
func (e *BaseEntity) Alignment() int64 {
	return e.Type().Alignment()
}

// Refered 增加引用计数（保持 Java 原拼写 refered）
func (e *BaseEntity) Refered() {
	e.nRefered++
}

// IsRefered 检查是否被引用过
func (e *BaseEntity) IsRefered() bool {
	return e.nRefered > 0
}

// SetMemref 设置内存引用
func (e *BaseEntity) SetMemref(mem asm.IMemoryReference) {
	e.memref = mem
}

// Memref 获取内存引用（会先检查地址是否已解析）
func (e *BaseEntity) Memref() asm.IMemoryReference {
	e.CheckAddress()
	return e.memref
}

// SetAddressMem 设置地址（MemoryReference 版本）
// 对应 Java 的 setAddress(MemoryReference mem)
func (e *BaseEntity) SetAddressMem(mem asm.IMemoryReference) {
	e.address = mem
}

// SetAddressImm 设置地址（ImmediateValue 版本）
// 对应 Java 的 setAddress(ImmediateValue imm)
func (e *BaseEntity) SetAddressImm(imm asm.IImmediateValue) {
	e.address = imm
}

// Address 获取地址（会检查地址是否已解析）
func (e *BaseEntity) Address() asm.IOperand {
	e.CheckAddress()
	return e.address
}

// CheckAddress 检查地址是否已解析（保护方法，包内可见）
func (e *BaseEntity) CheckAddress() {
	if e.memref == nil && e.address == nil {
		panic("address did not resolved: " + e.name)
	}
}

// Location 返回代码位置
func (e *BaseEntity) Location() *Location {
	return e.typeNode.Location()
}

// Dump 实现 Dumpable 接口
// 对应 Java 的 public void dump(Dumper d)
func (e *BaseEntity) Dump(d *Dumper) {
}
