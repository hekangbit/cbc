package models

type IScope interface {
	IsToplevel() bool
	Toplevel() *ToplevelScope
	Parent() IScope
	AddChild(*LocalScope)
	Get(string) (IEntity, error)
}

type Scope struct {
	children []*LocalScope
}

func BaseScope() Scope {
	return Scope{children: make([]*LocalScope, 0)}
}

func (this *Scope) AddChild(s *LocalScope) {
	this.children = append(this.children, s)
}
