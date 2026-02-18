package models

import "cbc/util"

type Declarations struct {
	defvars    *util.LinkedHashSet[*DefinedVariable]
	vardecls   *util.LinkedHashSet[*UndefinedVariable]
	deffuns    *util.LinkedHashSet[*DefinedFunction]
	funcdecls  *util.LinkedHashSet[*UndefinedFunction]
	constants  *util.LinkedHashSet[*Constant]
	defstructs *util.LinkedHashSet[*ASTStructNode]
	defunions  *util.LinkedHashSet[*ASTUnionNode]
	typedefs   *util.LinkedHashSet[*ASTTypedefNode]
}

func NewDeclarations() *Declarations {
	return &Declarations{
		defvars:    util.NewLinkedHashSet[*DefinedVariable](),
		vardecls:   util.NewLinkedHashSet[*UndefinedVariable](),
		deffuns:    util.NewLinkedHashSet[*DefinedFunction](),
		funcdecls:  util.NewLinkedHashSet[*UndefinedFunction](),
		constants:  util.NewLinkedHashSet[*Constant](),
		defstructs: util.NewLinkedHashSet[*ASTStructNode](),
		defunions:  util.NewLinkedHashSet[*ASTUnionNode](),
		typedefs:   util.NewLinkedHashSet[*ASTTypedefNode](),
	}
}

func (this *Declarations) Add(other *Declarations) {
	this.defvars.Merge(other.defvars)
	this.vardecls.Merge(other.vardecls)
	this.deffuns.Merge(other.deffuns)
	this.funcdecls.Merge(other.funcdecls)
	this.constants.Merge(other.constants)
	this.defstructs.Merge(other.defstructs)
	this.defunions.Merge(other.defunions)
	this.typedefs.Merge(other.typedefs)
}

func (this *Declarations) AddDefvar(v *DefinedVariable) {
	this.defvars.Add(v)
}

func (this *Declarations) AddDefvars(vars []*DefinedVariable) {
	this.defvars.AddAll(vars)
}

func (this *Declarations) Defvars() []*DefinedVariable {
	return this.defvars.ToSlice()
}

func (this *Declarations) AddVardecl(v *UndefinedVariable) {
	this.vardecls.Add(v)
}

func (this *Declarations) Vardecls() []*UndefinedVariable {
	return this.vardecls.ToSlice()
}

func (this *Declarations) AddConstant(c *Constant) {
	this.constants.Add(c)
}

func (this *Declarations) Constants() []*Constant {
	return this.constants.ToSlice()
}

func (this *Declarations) AddDeffun(f *DefinedFunction) {
	this.deffuns.Add(f)
}

func (this *Declarations) Deffuns() []*DefinedFunction {
	return this.deffuns.ToSlice()
}

func (this *Declarations) AddFuncdecl(f *UndefinedFunction) {
	this.funcdecls.Add(f)
}

func (this *Declarations) Funcdecls() []*UndefinedFunction {
	return this.funcdecls.ToSlice()
}

func (this *Declarations) AddDefstruct(n *ASTStructNode) {
	this.defstructs.Add(n)
}

func (this *Declarations) Defstructs() []*ASTStructNode {
	return this.defstructs.ToSlice()
}

func (this *Declarations) AddDefunion(n *ASTUnionNode) {
	this.defunions.Add(n)
}

func (this *Declarations) Defunions() []*ASTUnionNode {
	return this.defunions.ToSlice()
}

func (this *Declarations) AddTypedef(n *ASTTypedefNode) {
	this.typedefs.Add(n)
}

func (this *Declarations) Typedefs() []*ASTTypedefNode {
	return this.typedefs.ToSlice()
}
