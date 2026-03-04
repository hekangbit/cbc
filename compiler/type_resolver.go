package compiler

import (
	"cbc/models"
	"cbc/utils"
)

type TypeResolver struct {
	Visitor
	typeTable    *models.TypeTable
	errorHandler *utils.ErrorHandler
}

// need support 3 interface
var _ IVisitor = &TypeResolver{}
var _ models.IDeclarationVisitor = &TypeResolver{}
var _ models.IEntityVisitor = &TypeResolver{}

func NewTypeResolver(typeTable *models.TypeTable, errorHandler *utils.ErrorHandler) *TypeResolver {
	p := &TypeResolver{
		typeTable:    typeTable,
		errorHandler: errorHandler,
	}
	p.Visitor._impl_visitor = p
	return p
}

func (this *TypeResolver) Resolve(astObj *models.AST) {
	this.defineTypes(astObj.Types())
	for _, tdef := range astObj.Types() {
		tdef.Accept(this)
	}
	for _, e := range astObj.Entities() {
		e.Accept(this)
	}
}

func (this *TypeResolver) defineTypes(deftypes []models.IASTTypeDefinition) {
	for _, def := range deftypes {
		if this.typeTable.IsDefined(def.TypeRef()) {
			this.error(def, "duplicated type definition: "+def.TypeRef().String())
		} else {
			this.typeTable.Put(def.TypeRef(), def.DefiningType())
		}
	}
}

func (this *TypeResolver) bindType(n *models.ASTTypeNode) {
	if n.IsResolved() {
		return
	}
	n.SetType(this.typeTable.Get(n.TypeRef()))
}

func (this *TypeResolver) error(node models.INode, msg string) {
	this.errorHandler.Error(node.Location().String(), msg)
}

// --- Implement DeclarationVisitor ---

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

func (this *TypeResolver) resolveCompositeType(def models.IASTCompositeTypeDefinition) {
	ct, ok := this.typeTable.Get(def.TypeNode().TypeRef()).(models.ICompositeType)
	if !ok {
		// TODO: java throw new Error("cannot intern struct/union: " + def.name());
		panic("cannot intern struct/union: " + def.Name())
	}
	for _, slot := range ct.Members() {
		this.bindType(slot.TypeNode())
	}
}

// --- implement EntityVisitor ---

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

func (this *TypeResolver) VisitConstant(c *models.Constant) any {
	this.bindType(c.TypeNode())
	this.visitExpr(c.Value())
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

func (this *TypeResolver) resolveFunctionHeader(f models.IFunction) {
	this.bindType(f.TypeNode())
	for _, param := range f.Parameters() {
		// notice: array to pointer
		t := this.typeTable.GetParamType(param.TypeNode().TypeRef())
		param.TypeNode().SetType(t)
	}
}

// --- 覆盖 ASTVisitor 中的部分方法，以进行类型绑定 ---

func (this *TypeResolver) VisitBlockNode(node *models.ASTBlockNode) any {
	for _, v := range node.Variables() {
		v.Accept(this) // 通过EntityVisitor访问
	}
	this.visitStmts(node.Stmts())
	return nil
}

func (this *TypeResolver) VisitCastNode(node *models.ASTCastNode) any {
	this.bindType(node.TypeNode())
	return this.Visitor.VisitCastNode(node)
}

// TODO: remove duplicate
// func (this *TypeResolver) VisitCastNode(node *models.ASTCastNode) any {
// 	this.bindType(node.TypeNode())
// 	this.visitExpr(node.Expr())
// 	return nil
// }

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
