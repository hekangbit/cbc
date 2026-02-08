package models

import "fmt"

type CBCParameter struct {
	DefinedVariable
}

var _ Dumpable = (*CBCParameter)(nil)

func NewCBCParameter(typenode *TypeNode, name string) *CBCParameter {
	var p = new(CBCParameter)
	p.isPrivate = false
	p.typeNode = typenode
	return p

}
func (param *CBCParameter) IsParameter() bool {
	return true
}

func (param *CBCParameter) Dump(d *Dumper) {
	fmt.Println("name ", param.name)         // TODO
	fmt.Println("typeNode ", param.typeNode) // TODO
}
