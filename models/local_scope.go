package models

import (
	"cbc/utils"
	"fmt"
)

type LocalScope struct {
	Scope
	parent    IScope
	variables map[string]IDefinedVariable
}

var _ IScope = &LocalScope{}

func NewLocalScope(parent IScope) *LocalScope {
	scope := &LocalScope{
		parent: parent,
	}
	parent.AddChild(scope)
	scope.variables = make(map[string]IDefinedVariable)
	return scope
}

func (this *LocalScope) IsToplevel() bool {
	return false
}

func (this *LocalScope) Toplevel() *ToplevelScope {
	return this.parent.Toplevel()
}

func (this *LocalScope) Parent() IScope {
	return this.parent
}

func (this *LocalScope) Children() []*LocalScope {
	return this.children
}

func (this *LocalScope) IsDefinedLocally(name string) bool {
	if _, ok := this.variables[name]; ok {
		return true
	}
	return false
}

func (this *LocalScope) DefineVariable(v IDefinedVariable) {
	if _, ok := this.variables[v.Name()]; ok {
		panic("duplicated variable: " + v.Name())
	}
	this.variables[v.Name()] = v
}

func (this *LocalScope) AllocateTmp(t IType) *DefinedVariable {
	v := NewTmpNewDefinedVariable(t)
	this.DefineVariable(v)
	return v
}

func (this *LocalScope) Get(name string) (IEntity, error) {
	v, ok := this.variables[name]
	if ok {
		return v, nil
	}
	return this.parent.Get(name)
}

func (this *LocalScope) AllLocalVariables() []IDefinedVariable {
	result := make([]IDefinedVariable, 0)
	for _, s := range this.AllLocalScopes() {
		result = append(result, s.LocalVariables()...)
	}
	return result
}

func (this *LocalScope) LocalVariables() []IDefinedVariable {
	result := make([]IDefinedVariable, 0)
	for _, v := range this.variables {
		if !v.IsPrivate() {
			result = append(result, v)
		}
	}
	return result
}

func (this *LocalScope) StaticLocalVariables() []IDefinedVariable {
	result := make([]IDefinedVariable, 0)
	for _, s := range this.AllLocalScopes() {
		for _, v := range s.variables {
			if v.IsPrivate() {
				result = append(result, v)
			}
		}
	}
	return result
}

func (this *LocalScope) AllLocalScopes() []*LocalScope {
	result := make([]*LocalScope, 0)
	result = this.CollectScope(result)
	return result
}

func (this *LocalScope) CollectScope(buf []*LocalScope) []*LocalScope {
	buf = append(buf, this)
	for _, s := range this.children {
		buf = s.CollectScope(buf)
	}
	return buf
}

func (this *LocalScope) CheckReferences(h *utils.ErrorHandler) {
	for _, v := range this.variables {
		if !v.IsRefered() {
			h.WarnWithLoc(v.Location(), fmt.Sprintf("unused variable: %s", v.Name()))
		}
	}
	for _, child := range this.children {
		child.CheckReferences(h)
	}
}
