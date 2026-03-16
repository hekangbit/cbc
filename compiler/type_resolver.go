package compiler

import (
	"cbc/models"
	"cbc/utils"
	"errors"
)

type TypeResolver struct {
	Visitor
	typeTable    *models.TypeTable
	errorHandler *utils.ErrorHandler
}

var _ IVisitor = &TypeResolver{}
var _ models.IDeclarationVisitor = &TypeResolver{}
var _ models.IEntityVisitor = &TypeResolver{}

func NewTypeResolver(typeTable *models.TypeTable, errorHandler *utils.ErrorHandler) *TypeResolver {
	p := new(TypeResolver)
	p.typeTable = typeTable
	p.errorHandler = errorHandler
	p._impl_visitor = p
	return p
}

func (this *TypeResolver) Resolve(astObj *models.AST) error {
	this.defineTypes(astObj.Types())
	for _, node := range astObj.Types() {
		node.Accept(this)
	}
	for _, ent := range astObj.Entities() {
		ent.Accept(this)
	}
	if this.errorHandler.ErrorOccured() {
		return errors.New("semantic analyze type resolve failed")
	}
	return nil
}

func (this *TypeResolver) defineTypes(typeDefNodes []models.IASTAbstractTypeDefinitionNode) {
	for _, typeDefNode := range typeDefNodes {
		if this.typeTable.IsDefined(typeDefNode.TypeRef()) {
			this.errorHandler.ErrorWithLoc(typeDefNode.Location(), "duplicated type definition: "+typeDefNode.TypeRef().String())
			continue
		}
		this.typeTable.Put(typeDefNode.TypeRef(), typeDefNode.DefiningType())
	}
}

func (this *TypeResolver) bindType(n *models.ASTTypeNode) {
	if n.IsResolved() {
		return
	}
	n.SetType(this.typeTable.Get(n.TypeRef()))
}

func (this *TypeResolver) resolveCompositeType(def models.IASTAbstractCompositeTypeDefinitionNode) error {
	t := this.typeTable.Get(def.TypeNode().TypeRef())
	compType, ok := t.(models.ICompositeType)
	if !ok {
		panic("cannot intern struct/union: " + def.Name())
	}
	for _, slot := range compType.Members() {
		this.bindType(slot.TypeNode())
	}
	return nil
}

func (this *TypeResolver) resolveFunctionHeader(f models.IFunction) {
	this.bindType(f.TypeNode())
	for _, param := range f.Parameters() {
		// notice: array to pointer
		t := this.typeTable.GetParamType(param.TypeNode().TypeRef())
		param.TypeNode().SetType(t)
	}
}

// --- Implement DeclarationVisitor interface---

func (this *TypeResolver) VisitStructNode(structNode *models.ASTStructNode) any {
	this.resolveCompositeType(structNode)
	return nil
}

func (this *TypeResolver) VisitUnionNode(unionNode *models.ASTUnionNode) any {
	this.resolveCompositeType(unionNode)
	return nil
}

func (this *TypeResolver) VisitTypedefNode(typedefNode *models.ASTTypedefNode) any {
	this.bindType(typedefNode.TypeNode())
	this.bindType(typedefNode.RealTypeNode())
	return nil
}

// --- Implement EntityVisitor interface---

func (this *TypeResolver) VisitConstant(c *models.Constant) any {
	this.bindType(c.TypeNode())
	this.visitExpr(c.Value())
	return nil
}

func (this *TypeResolver) VisitDefinedVariable(varNode *models.DefinedVariable) any {
	this.bindType(varNode.TypeNode())
	if varNode.HasInitializer() {
		this.visitExpr(varNode.Initializer())
	}
	return nil
}

func (this *TypeResolver) VisitUndefinedVariable(varNode *models.UndefinedVariable) any {
	this.bindType(varNode.TypeNode())
	return nil
}

func (this *TypeResolver) VisitDefinedFunction(funcNode *models.DefinedFunction) any {
	this.resolveFunctionHeader(funcNode)
	this.visitStmt(funcNode.Body())
	return nil
}

func (this *TypeResolver) VisitUndefinedFunction(funcNode *models.UndefinedFunction) any {
	this.resolveFunctionHeader(funcNode)
	return nil
}

// --- Implement ASTVisitor interface ---

func (this *TypeResolver) VisitBlockNode(node *models.ASTBlockNode) any {
	for _, v := range node.Variables() {
		v.Accept(this)
	}
	this.visitStmts(node.Stmts())
	return nil
}

func (this *TypeResolver) VisitCastNode(node *models.ASTCastNode) any {
	this.bindType(node.TypeNode())
	this.Visitor.VisitCastNode(node)
	return nil
}

func (this *TypeResolver) VisitSizeofExprNode(node *models.ASTSizeofExprNode) any {
	this.bindType(node.TypeNode())
	this.visitExpr(node.Expr())
	return nil
}

func (this *TypeResolver) VisitSizeofTypeNode(node *models.ASTSizeofTypeNode) any {
	this.bindType(node.OperandTypeNode())
	this.bindType(node.TypeNode())
	return nil
}

func (this *TypeResolver) VisitIntegerLiteralNode(node *models.ASTIntegerLiteralNode) any {
	this.bindType(node.TypeNode())
	return nil
}

func (this *TypeResolver) VisitStringLiteralNode(node *models.ASTStringLiteralNode) any {
	this.bindType(node.TypeNode())
	return nil
}
