package models

import "cbc/utils"

type LocalScope struct {
	Scope
	parent    IScope
	variables map[string]*DefinedVariable
}

func NewLocalScope(parent IScope) *LocalScope {
	scope := &LocalScope{
		parent: parent,
	}
	parent.AddChild(scope)
	scope.variables = make(map[string]*DefinedVariable)
	return scope
}

func (this *LocalScope) IsTopLevel() bool {
	return false
}

func (this *LocalScope) TopLevel() *ToplevelScope {
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

func (this *LocalScope) DefineVariable(v *DefinedVariable) {
	if _, ok := this.variables[v.Name()]; ok {
		panic("duplicated variable: " + v.Name())
	}
	this.variables[v.Name()] = v
}

func (this *LocalScope) AllocateTmp(t IType) {
	v := NewTmpNewDefinedVariable(t)
	this.DefineVariable(v)
}

func (this *LocalScope) Get(name string) IEntity {
	v, ok := this.variables[name]
	if ok {
		return v
	}
	return this.parent.Get(name)
}

func (this *LocalScope) AllLocalVariables() []*DefinedVariable {
	result := make([]*DefinedVariable, 0)
	for _, s := range this.AllLocalScopes() {
		result = append(result, s.LocalVariables()...)
	}
	return result
}

func (this *LocalScope) LocalVariables() []*DefinedVariable {
	result := make([]*DefinedVariable, 0)
	for _, v := range this.variables {
		if !v.IsPrivate() {
			result = append(result, v)
		}
	}
	return result
}

func (this *LocalScope) StaticLocalVariables() []*DefinedVariable {
	result := make([]*DefinedVariable, 0)
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

func (this *LocalScope) CheckReferences(h utils.ErrorHandler) {
	for _, v := range this.variables {
		if !v.IsRefered() {
			// TODO: localscope.java can report warning, h.warn(...)
		}
	}
	for _, child := range this.children {
		child.CheckReferences(h)
	}
}
