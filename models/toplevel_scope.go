package models

import (
	"cbc/utils"
	"fmt"
)

type ToplevelScope struct {
	Scope
	entities             map[string]IEntity
	staticLocalVariables []IDefinedVariable
}

var _ IScope = &ToplevelScope{}

func NewToplevelScope() *ToplevelScope {
	p := new(ToplevelScope)
	p.Scope = BaseScope()
	p.entities = make(map[string]IEntity)
	p.staticLocalVariables = nil
	return p
}

func (this *ToplevelScope) IsToplevel() bool {
	return true
}

func (this *ToplevelScope) Parent() IScope {
	return nil
}

func (this *ToplevelScope) Toplevel() *ToplevelScope {
	return this
}

func (this *ToplevelScope) DeclareEntity(ent IEntity) error {
	e, ok := this.entities[ent.Name()]
	if ok {
		msg := fmt.Sprintf("duplicated declaration: %s: %s and %s", ent.Name(), e.Location().String(), ent.Location().String())
		return fmt.Errorf("%s", msg)
	}
	this.entities[ent.Name()] = ent
	return nil
}

func (this *ToplevelScope) DefineEntity(ent IEntity) error {
	e, ok := this.entities[ent.Name()]
	if ok && e.IsDefined() {
		msg := fmt.Sprintf("duplicated definition: %s: %s and %s", ent.Name(), e.Location().String(), ent.Location().String())
		return fmt.Errorf("%s", msg)
	}
	this.entities[ent.Name()] = ent
	return nil
}

func (this *ToplevelScope) Get(name string) (IEntity, error) {
	ent, ok := this.entities[name]
	if ok {
		return ent, nil
	}
	return nil, fmt.Errorf("%s", fmt.Sprintf("unresolved reference: %s", name))
}

func (this *ToplevelScope) AllGlobalVariables() []IVariable {
	result := make([]IVariable, 0)
	for _, ent := range this.entities {
		v, ok := ent.(IVariable)
		if ok {
			result = append(result, v)
		}
	}
	buf := make([]IVariable, 0)
	for _, tmp := range this.StaticLocalVariables() {
		buf = append(buf, tmp.(IVariable))
	}
	result = append(result, buf...)
	return result
}

func (this *ToplevelScope) DefinedGlobalScopeVariables() []IDefinedVariable {
	result := make([]IDefinedVariable, 0)
	for _, ent := range this.entities {
		v, ok := ent.(IDefinedVariable)
		if ok {
			result = append(result, v)
		}
	}
	result = append(result, this.StaticLocalVariables()...)
	return result
}

func (this *ToplevelScope) StaticLocalVariables() []IDefinedVariable {
	if this.staticLocalVariables == nil {
		this.staticLocalVariables = make([]IDefinedVariable, 0)
		for _, s := range this.children {
			this.staticLocalVariables = append(this.staticLocalVariables, s.StaticLocalVariables()...)
		}
		seqTable := make(map[string]int64)
		for _, v := range this.staticLocalVariables {
			seq, ok := seqTable[v.Name()]
			if !ok {
				v.SetSequence(0)
				seqTable[v.Name()] = 1
			} else {
				v.SetSequence(seq)
				seqTable[v.Name()] = seq + 1
			}
		}
	}
	return this.staticLocalVariables
}

func (this *ToplevelScope) CheckReferences(h *utils.ErrorHandler) {
	for _, ent := range this.entities {
		if ent.IsDefined() && ent.IsPrivate() && !ent.IsConstant() && !ent.IsRefered() {
			h.WarnWithLoc(ent.Location(), fmt.Sprintf("unused top level variable: : %s", ent.Name()))
		}
	}
	for _, funcScope := range this.children {
		for _, funcBodyScope := range funcScope.children {
			funcBodyScope.CheckReferences(h)
		}
	}
}
