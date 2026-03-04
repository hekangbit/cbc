package models

type IASTCompositeTypeDefinition interface {
	IASTTypeDefinition
}

type ASTCompositeTypeDefinition struct {
	ASTTypeDefinition
	members []*Slot
}

func NewASTCompositeTypeDefinition(loc *Location, ref ITypeRef, name string, members []*Slot) *ASTCompositeTypeDefinition {
	p := &ASTCompositeTypeDefinition{
		ASTTypeDefinition: ASTTypeDefinition{
			name:     name,
			location: loc,
			typeNode: NewASTTypeNodeFromRef(ref),
		},
		members: members,
	}
	p._impl = p
	return p
}

func (this *ASTCompositeTypeDefinition) IsCompositeType() bool {
	return true
}

func (this *ASTCompositeTypeDefinition) Members() []*Slot {
	return this.members
}

func (this *ASTCompositeTypeDefinition) _Dump(d *Dumper) {
	d.PrintMemberStringNotResolved("name", this.name)
	buf := make([]Dumpable, len(this.members))
	for i, tmp := range this.members {
		buf[i] = tmp
	}
	d.PrintNodeList("members", buf)
}
