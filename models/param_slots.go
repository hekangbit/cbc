package models

type ParamSlots interface {
	Argc() int
	MinArgc() int
	AcceptVarargs()
	IsVararg() bool
	Location() Location
}

type BaseParamSlots[T any] struct {
	location         *Location
	paramDescriptors []T
	vararg           bool
}

// 构造函数
func NewBaseParamSlots[T any](loc *Location, paramDescs []T, vararg bool) *BaseParamSlots[T] {
	return &BaseParamSlots[T]{
		location:         loc,
		paramDescriptors: paramDescs,
		vararg:           vararg,
	}
}

func NewBaseParamSlotsNoVararg[T any](loc *Location, paramDescs []T) *BaseParamSlots[T] {
	return &BaseParamSlots[T]{
		location:         loc,
		paramDescriptors: paramDescs,
		vararg:           false,
	}
}

func NewBaseParamSlotsNoLocation[T any](loc *Location, paramDescs []T) *BaseParamSlots[T] {
	return &BaseParamSlots[T]{
		location:         nil,
		paramDescriptors: paramDescs,
		vararg:           false,
	}
}

func (p *BaseParamSlots[T]) Argc() int {
	if p.vararg {
		//  throw new Error("must not happen: Param#argc for vararg");
	}
	return len(p.paramDescriptors)
}

func (p *BaseParamSlots[T]) MinArgc() int {
	return len(p.paramDescriptors)
}

func (p *BaseParamSlots[T]) AcceptVarargs() {
	p.vararg = true
}

func (p *BaseParamSlots[T]) IsVararg() bool {
	return p.vararg
}

func (p *BaseParamSlots[T]) Location() *Location {
	return p.location
}
