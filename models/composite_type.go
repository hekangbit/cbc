package models

import (
	"errors"
	"fmt"
)

type ICompositeType interface {
	NamedType
	Members() []*Slot
	MemberTypes() []IType
	HasMember(name string) bool
	MemberType(name string) (IType, error)
	MemberOffset(name string) (int64, error)
	GetMember(name string) (*Slot, error)
}

// BaseCompositeType 复合类型基础结构体
type BaseCompositeType struct {
	BaseNamedType
	members            []*Slot
	cachedSize         int64
	cachedAlign        int64
	isRecursiveChecked bool
}

// 构造函数
func NewBaseCompositeType(name string, members []*Slot, location Location) *BaseCompositeType {
	return &BaseCompositeType{
		BaseNamedType:      *NewBaseNamedType(name, location),
		members:            members,
		cachedSize:         SizeUnknown,
		cachedAlign:        SizeUnknown,
		isRecursiveChecked: false,
	}
}

// 实现Type接口的方法
func (c *BaseCompositeType) IsCompositeType() bool {
	return true
}

// IsSameType 检查类型是否相同
func (c *BaseCompositeType) IsSameType(other IType) bool {
	return c.CompareMemberTypes(other, "IsSameType")
}

// IsCompatible 检查类型是否兼容
func (c *BaseCompositeType) IsCompatible(target IType) bool {
	return c.CompareMemberTypes(target, "IsCompatible")
}

// IsCastableTo 检查是否可以强制转换到目标类型
func (c *BaseCompositeType) IsCastableTo(target IType) bool {
	return c.CompareMemberTypes(target, "IsCastableTo")
}

// CompareMemberTypes 比较成员类型（Go版本不使用反射）
func (c *BaseCompositeType) CompareMemberTypes(other IType, cmpType string) bool {
	if c.IsStruct() && !other.IsStruct() {
		return false
	}
	if c.IsUnion() && !other.IsUnion() {
		return false
	}

	otherComposite := other.GetCompositeType()

	// 获取其他复合类型的成员
	otherMembers := otherComposite.Members()

	// 检查成员数量是否相同
	if len(c.members) != int(other.Size()) {
		return false
	}

	// 比较每个成员的类型
	for i, member := range c.members {
		otherMember := otherMembers[i]

		var cmpResult bool
		switch cmpType {
		case "IsSameType":
			cmpResult = member.Type().IsSameType(otherMember.Type())
		case "IsCompatible":
			cmpResult = member.Type().IsCompatible(otherMember.Type())
		case "IsCastableTo":
			cmpResult = member.Type().IsCastableTo(otherMember.Type())
		default:
			return false
		}

		if !cmpResult {
			return false
		}
	}

	return true
}

// Size 返回类型大小
func (c *BaseCompositeType) Size() int64 {
	if c.cachedSize == SizeUnknown {
		c.ComputeOffsets()
	}
	return c.cachedSize
}

// AllocSize 分配大小
func (c *BaseCompositeType) AllocSize() int64 {
	return c.Size()
}

// Alignment 对齐大小
func (c *BaseCompositeType) Alignment() int64 {
	if c.cachedAlign == SizeUnknown {
		c.ComputeOffsets()
	}
	return c.cachedAlign
}

// Members 获取成员列表
func (c *BaseCompositeType) Members() ([]*Slot, error) {
	return c.members, nil
}

// MemberTypes 获取成员类型列表
func (c *BaseCompositeType) MemberTypes() []IType {
	result := make([]IType, len(c.members))
	for i, slot := range c.members {
		result[i] = slot.Type()
	}
	return result
}

// HasMember 检查是否有指定成员
func (c *BaseCompositeType) HasMember(name string) bool {
	_, err := c.GetMember(name)
	return err == nil
}

// MemberType 获取成员类型
func (c *BaseCompositeType) MemberType(name string) (IType, error) {
	slot, err := c.FetchMember(name)
	if err != nil {
		return nil, err
	}
	return slot.Type(), nil
}

// MemberOffset 获取成员偏移量
func (c *BaseCompositeType) MemberOffset(name string) (int64, error) {
	slot, err := c.FetchMember(name)
	if err != nil {
		return 0, err
	}

	if slot.offset == SizeUnknown {
		c.ComputeOffsets()
	}
	return slot.offset, nil
}

// ComputeOffsets 计算偏移量（由子类实现）
func (c *BaseCompositeType) ComputeOffsets() {
	// 这是一个模板方法，由具体子类实现
	panic("ComputeOffsets must be implemented by concrete composite type")
}

// FetchMember 获取成员（内部方法）
func (c *BaseCompositeType) FetchMember(name string) (*Slot, error) {
	slot, err := c.GetMember(name)
	if err != nil {
		return nil, errors.New("no such member in " + c.Name() + ": " + name)
	}
	return slot, nil
}

// GetMember 获取成员
func (c *BaseCompositeType) GetMember(name string) (*Slot, error) {
	for _, slot := range c.members {
		if slot.name == name {
			return slot, nil
		}
	}
	return nil, errors.New("member not found")
}

// String 返回字符串表示
func (c *BaseCompositeType) String() string {
	memberCount := len(c.members)
	if memberCount == 0 {
		return c.Name() + " {}"
	}

	return fmt.Sprintf("%s {%d members}", c.Name(), memberCount)
}
