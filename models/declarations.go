package models

type Declarations struct {
	defvars map[*DefinedVariable]struct{}
	deffuns map[*DefinedFunction]struct{}
}

func NewDeclarations() *Declarations {
	return &Declarations{
		defvars: make(map[*DefinedVariable]struct{}),
		deffuns: make(map[*DefinedFunction]struct{}),
	}
}

func (decls *Declarations) Add(new_decls *Declarations) {
	for k := range new_decls.defvars {
		decls.defvars[k] = struct{}{}
	}
	for k := range new_decls.deffuns {
		decls.deffuns[k] = struct{}{}
	}
}

func (decls *Declarations) AddDefvars(vars []*DefinedVariable) {
	for _, pdv := range vars {
		decls.defvars[pdv] = struct{}{}
	}
}

func (decls *Declarations) AddDefFunc(function *DefinedFunction) {
	decls.deffuns[function] = struct{}{}
}
