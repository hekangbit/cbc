package models

type IScope interface {
	IsToplevel() bool
	Toplevel() *ToplevelScope
	Parent() IScope
	AddChild(*LocalScope)
	Get(string) IEntity
}

type Scope struct {
	children []*LocalScope
}

func NewScope() *Scope {
	return &Scope{
		children: make([]*LocalScope, 0),
	}
}

func (this *Scope) AddChild(s *LocalScope) {
	this.children = append(this.children, s)
}

func (this *Scope) IsToplevel() bool {
	panic("Scope::IsToplevel need implenmented by concreate struct")
}

func (this *Scope) Toplevel() *ToplevelScope {
	panic("Scope::Toplevel need implenmented by concreate struct")
}

func (this *Scope) Parent() IScope {
	panic("Scope::Parent need implenmented by concreate struct")
}

func (this *Scope) Get(name string) IEntity {
	panic("Scope::Get need implenmented by concreate struct")
}
