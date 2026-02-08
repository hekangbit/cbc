package models

var _ Dumpable = (*Params)(nil)

type Params struct {
	*BaseParamSlots[*CBCParameter]
	// 可以添加Params特有的字段
	paramDescriptors []*CBCParameter
}

// 构造函数
func NewParams(loc *Location, paramDescs []*CBCParameter) *Params {
	return &Params{
		BaseParamSlots: NewBaseParamSlots(loc, paramDescs, false),
	}
}

func (p *Params) Parameters() []*CBCParameter {
	return p.paramDescriptors
}

// ParametersTypeRef 方法
func (p *Params) ParametersTypeRef() *ParamTypeRefs {
	typeRefs := make([]ITypeRef, 0, len(p.paramDescriptors))

	for _, param := range p.paramDescriptors {
		if param.TypeNode() != nil {
			typeRefs = append(typeRefs, param.TypeNode().TypeRef())
		}
	}

	return NewParamTypeRefs(p.location, typeRefs, p.vararg)
}

// Equals 方法 - 对应Java中的equals(Object)
func (p *Params) Equals(other interface{}) bool {
	if otherParams, ok := other.(*Params); ok {
		return p.EqualsParams(otherParams)
	}
	return false
}

// EqualsParams 方法 - 对应Java中的equals(Params)
func (p *Params) EqualsParams(other *Params) bool {
	if other == nil {
		return false
	}

	// 比较vararg
	if other.vararg != p.vararg {
		return false
	}

	// 比较paramDescriptors
	if len(other.paramDescriptors) != len(p.paramDescriptors) {
		return false
	}

	for i := range p.paramDescriptors {
		if !p.compareCBCParameter(p.paramDescriptors[i], other.paramDescriptors[i]) {
			return false
		}
	}

	return true
}

// 辅助方法：比较两个CBCParameter
func (p *Params) compareCBCParameter(a, b *CBCParameter) bool {
	// 比较名称
	if a.Name() != b.Name() {
		return false
	}

	// 比较类型节点
	return a.Type().IsSameType(b.Type())
}

// Dump 方法 - 实现Dumpable接口
func (p *Params) Dump(d *Dumper) {
	dumpables := make([]Dumpable, len(p.Parameters()))
	for i, p := range p.Parameters() {
		dumpables[i] = p
	}
	d.PrintNodeList("parameters", dumpables)
}
