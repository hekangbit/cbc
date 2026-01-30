package models

type Declarations struct {
	defvars map[*DefinedVariable]struct{}
}

func NewDeclarations() *Declarations {
	return &Declarations{defvars: make(map[*DefinedVariable]struct{})}
}

func (decls *Declarations) Add(new_decls *Declarations) {
	// merge de
	for k, _ := range new_decls.defvars {
		decls.defvars[k] = struct{}{}
	}
}

func (decls *Declarations) AddDefvars(vars []*DefinedVariable) {
	for _, pdv := range vars {
		decls.defvars[pdv] = struct{}{}
	}
}
