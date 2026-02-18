package models

type IParamSlots interface {
	Argc() int
	MinArgc() int
	AcceptVarargs()
	IsVararg() bool
	Location() *Location
}

type ParamSlots[T any] struct {
	location         *Location
	paramDescriptors []T
	vararg           bool
}

func NewParamSlots[T any](paramDescs []T) *ParamSlots[T] {
	return &ParamSlots[T]{
		location:         nil,
		paramDescriptors: paramDescs,
		vararg:           false,
	}
}

func NewParamSlotsWithLoc[T any](loc *Location, paramDescs []T) *ParamSlots[T] {
	return &ParamSlots[T]{
		location:         loc,
		paramDescriptors: paramDescs,
		vararg:           false,
	}
}

func NewParamSlotsFull[T any](loc *Location, paramDescs []T, vararg bool) *ParamSlots[T] {
	return &ParamSlots[T]{
		location:         loc,
		paramDescriptors: paramDescs,
		vararg:           vararg,
	}
}

func (p *ParamSlots[T]) Argc() int {
	if p.vararg {
		//  throw new Error("must not happen: Param#argc for vararg");
	}
	return len(p.paramDescriptors)
}

func (p *ParamSlots[T]) MinArgc() int {
	return len(p.paramDescriptors)
}

func (p *ParamSlots[T]) AcceptVarargs() {
	p.vararg = true
}

func (p *ParamSlots[T]) IsVararg() bool {
	return p.vararg
}

func (p *ParamSlots[T]) Location() *Location {
	return p.location
}
