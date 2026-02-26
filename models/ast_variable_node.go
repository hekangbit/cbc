package models

type ASTVariableNode struct {
	ASTLHSNode
	location *Location
	name     string
	entity   IEntity
}

var _ IASTLHSNode = &ASTVariableNode{}

func NewASTVariableNode(loc *Location, name string) *ASTVariableNode {
	var p = new(ASTVariableNode)
	p.ASTLHSNode.ASTExprNode._impl = p
	p.ASTLHSNode.ASTExprNode.Node._impl = p
	p.location = loc
	p.name = name
	return p
}

func NewASTVariableNodeWithVar(v *DefinedVariable) *ASTVariableNode {
	var p = new(ASTVariableNode)
	p.ASTLHSNode.ASTExprNode._impl = p
	p.ASTLHSNode.ASTExprNode.Node._impl = p
	p.entity = v
	p.name = v.Name()
	return p
}

func (this *ASTVariableNode) Name() string {
	return this.name
}

func (this *ASTVariableNode) IsResolved() bool {
	return this.entity != nil
}

// TODO: java throw new Error("VariableNode.entity == null"); when nil
func (this *ASTVariableNode) Entity() IEntity {
	if this.entity == nil {
		panic("VariableNode.entity == null")
	}
	return this.entity
}

func (this *ASTVariableNode) SetEntity(ent IEntity) {
	this.entity = ent
}

// TODO: fix constant entity bug, what does this bug mean
func (this *ASTVariableNode) IsLvalue() bool {
	if this.entity.IsConstant() {
		return false
	}
	return true
}

// fix constant entity bug
func (this *ASTVariableNode) IsAssignable() bool {
	if this.entity.IsConstant() {
		return false
	}
	return this.IsLoadable()
}

func (this *ASTVariableNode) TypeNode() *ASTTypeNode {
	return this.Entity().TypeNode()
}

func (this *ASTVariableNode) IsParameter() bool {
	return this.Entity().IsParameter()
}

func (this *ASTVariableNode) OrigType() IType {
	return this.Entity().Type()
}

func (this *ASTVariableNode) Location() *Location {
	return this.location
}

func (this *ASTVariableNode) _Dump(d *Dumper) {
	if this.ty != nil {
		d.PrintMemberType("typpe", this.ty)
	}
	d.PrintMemberString("name", this.name, this.IsResolved())
}

func (this *ASTVariableNode) Accept(visitor IASTVisitor) interface{} {
	return visitor.VisitVariableNode(this)
}
