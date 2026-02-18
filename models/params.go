package models

var _ Dumpable = (*Params)(nil)

type Params struct {
	*ParamSlots[*CBCParameter]
	paramDescriptors []*CBCParameter
}

func NewParams(loc *Location, paramDescs []*CBCParameter) *Params {
	return &Params{
		ParamSlots: NewParamSlotsFull(loc, paramDescs, false),
	}
}

func (p *Params) Parameters() []*CBCParameter {
	return p.paramDescriptors
}

func (p *Params) ParametersTypeRef() *ParamTypeRefs {
	typeRefs := make([]ITypeRef, 0, len(p.paramDescriptors))

	for _, param := range p.paramDescriptors {
		if param.TypeNode() != nil {
			typeRefs = append(typeRefs, param.TypeNode().TypeRef())
		}
	}

	return NewParamTypeRefs(p.location, typeRefs, p.vararg)
}

func (p *Params) Equals(other interface{}) bool {
	if otherParams, ok := other.(*Params); ok {
		return p.EqualsParams(otherParams)
	}
	return false
}

func (p *Params) EqualsParams(other *Params) bool {
	if other == nil {
		return false
	}

	if other.vararg != p.vararg {
		return false
	}

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

func (p *Params) compareCBCParameter(a, b *CBCParameter) bool {
	if a.Name() != b.Name() {
		return false
	}

	return a.Type().IsSameType(b.Type())
}

func (p *Params) Dump(d *Dumper) {
	dumpables := make([]Dumpable, len(p.Parameters()))
	for i, p := range p.Parameters() {
		dumpables[i] = p
	}
	d.PrintNodeList("parameters", dumpables)
}
