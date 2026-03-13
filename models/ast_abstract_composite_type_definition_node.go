package models

type IASTAbstractCompositeTypeDefinitionNode interface {
	IASTAbstractTypeDefinitionNode
	Kind() string
	Members() []*Slot
	IsCompositeType() bool
}

type ASTAbstractCompositeTypeDefinitionNode struct {
	ASTAbstractTypeDefinitionNode
	members []*Slot
}

func (this *ASTAbstractCompositeTypeDefinitionNode) IsCompositeType() bool {
	return true
}

func (this *ASTAbstractCompositeTypeDefinitionNode) Members() []*Slot {
	return this.members
}

func (this *ASTAbstractCompositeTypeDefinitionNode) _Dump(d *Dumper) {
	d.PrintMemberStringNotResolved("name", this.name)
	buf := make([]Dumpable, len(this.members))
	for i, tmp := range this.members {
		buf[i] = tmp
	}
	d.PrintNodeList("members", buf)
}
