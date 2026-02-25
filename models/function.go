package models

import "cbc/asm"

type IFunction interface {
	IEntity
	Parameters() []*CBCParameter
	ReturnType() IType
	IsVoid() bool
	SetCallingSymbol(asm.ISymbol)
	CallingSymbol() asm.ISymbol
	Label() *asm.Label
}

type Function struct {
	Entity
	callingSymbol asm.ISymbol
	label         *asm.Label
}

func (this *Function) IsInitialized() bool {
	return true
}

func (this *Function) Parameters() []*CBCParameter {
	panic("Function::Parameters abstract method")
}

// TODO: GetFunctionType may return error, like java throw cast exception
func (this *Function) ReturnType() IType {
	t := GetFunctionType(this.Type())
	return t.ReturnType()
}

func (this *Function) IsVoid() bool {
	return this.ReturnType().IsVoid()
}

func (this *Function) SetCallingSymbol(sym asm.ISymbol) {
	if this.callingSymbol != nil {
		//TODO: java use throw new Error
		panic("must not happen: Function#callingSymbol was set again")
	}
	this.callingSymbol = sym
}

func (this *Function) CallingSymbol() asm.ISymbol {
	if this.callingSymbol == nil {
		//TODO: java use throw new Error
		panic("must not happen: Function#callingSymbol called but null")
	}
	return this.callingSymbol
}

func (this *Function) Label() *asm.Label {
	if this.label != nil {
		return this.label
	}
	return asm.NewLabel(this.CallingSymbol())
}
