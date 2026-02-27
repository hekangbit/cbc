package models

import "cbc/utils"

type Declarations struct {
	defvars    *utils.LinkedHashSet[*DefinedVariable]
	vardecls   *utils.LinkedHashSet[*UndefinedVariable]
	deffuns    *utils.LinkedHashSet[*DefinedFunction]
	funcdecls  *utils.LinkedHashSet[*UndefinedFunction]
	constants  *utils.LinkedHashSet[*Constant]
	defstructs *utils.LinkedHashSet[*ASTStructNode]
	defunions  *utils.LinkedHashSet[*ASTUnionNode]
	typedefs   *utils.LinkedHashSet[*ASTTypedefNode]
}

func NewDeclarations() *Declarations {
	return &Declarations{
		defvars:    utils.NewLinkedHashSet[*DefinedVariable](),
		vardecls:   utils.NewLinkedHashSet[*UndefinedVariable](),
		deffuns:    utils.NewLinkedHashSet[*DefinedFunction](),
		funcdecls:  utils.NewLinkedHashSet[*UndefinedFunction](),
		constants:  utils.NewLinkedHashSet[*Constant](),
		defstructs: utils.NewLinkedHashSet[*ASTStructNode](),
		defunions:  utils.NewLinkedHashSet[*ASTUnionNode](),
		typedefs:   utils.NewLinkedHashSet[*ASTTypedefNode](),
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
