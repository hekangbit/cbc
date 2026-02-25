package models

type CBCParameter struct {
	DefinedVariable
}

var _ Dumpable = &CBCParameter{}

func NewCBCParameter(typenode *ASTTypeNode, name string) *CBCParameter {
	var p = new(CBCParameter)
	p.isPrivate = false
	p.typeNode = typenode
	return p

}
func (param *CBCParameter) IsParameter() bool {
	return true
}

func (param *CBCParameter) _Dump(d *Dumper) {
	d.PrintMemberStringNotResolved("name", param.name)
	d.PrintMemberDumpable("typeNode", param.typeNode)
}
