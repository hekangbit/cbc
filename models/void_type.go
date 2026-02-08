package models

type VoidType struct {
	BaseType
}

// NewVoidType 创建空类型
func NewVoidType() *VoidType {
	return &VoidType{}
}

// 实现Type接口的方法

// Size 返回类型大小
func (v *VoidType) Size() int64 {
	return 1 // 根据Java原代码，void类型的大小为1
}

// AllocSize 分配大小
func (v *VoidType) AllocSize() int64 {
	return v.Size()
}

// Alignment 对齐大小
func (v *VoidType) Alignment() int64 {
	return 1
}

// IsSameType 检查类型是否相同
func (v *VoidType) IsSameType(other IType) bool {
	return other.IsVoid()
}

// IsCompatible 检查类型是否兼容
func (v *VoidType) IsCompatible(other IType) bool {
	return other.IsVoid()
}

// IsCastableTo 检查是否可以强制转换到目标类型
func (v *VoidType) IsCastableTo(target IType) bool {
	return target.IsVoid()
}

// 重写类型检查方法
func (v *VoidType) IsVoid() bool {
	return true
}

// Equals 检查两个VoidType是否相等
func (v *VoidType) Equals(other interface{}) bool {
	_, ok := other.(*VoidType)
	return ok
}

// String 返回字符串表示
func (v *VoidType) String() string {
	return "void"
}
