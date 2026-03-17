package compiler

import (
	"cbc/asm"
	"cbc/models"
	"cbc/utils"
	"container/list"
	"errors"
	"fmt"
)

// record jump label info
type JumpEntry struct {
	label      *asm.Label
	numRefered int64
	isDefined  bool
	location   *models.Location
}

type IRGenerator struct {
	typeTable    *models.TypeTable
	errorHandler *utils.ErrorHandler
	// for CompileFunctionBody start
	stmts         []models.IIRStmt
	scopeStack    *list.List // *models.LocalScope
	breakStack    *list.List // *asm.Label
	continueStack *list.List // *asm.Label
	jumpMap       map[string]*JumpEntry
	exprNestLevel int
	// for CompileFunctionBody end
}

var _ models.IASTVisitor = &IRGenerator{}

func NewIRGenerator(typeTable *models.TypeTable, errorHandler *utils.ErrorHandler) *IRGenerator {
	return &IRGenerator{
		typeTable:    typeTable,
		errorHandler: errorHandler,
	}
}

func (this *IRGenerator) Generate(sem *models.AST) (*models.IR, error) {
	for _, v := range sem.DefinedVariables() {
		if v.HasInitializer() {
			v.SetIR(this.transformExpr(v.Initializer()))
		}
	}
	for _, f := range sem.DefinedFunctions() {
		f.SetIR(this.CompileFunctionBody(f))
	}
	if this.errorHandler.ErrorOccured() {
		return nil, errors.New("IR generation failed.")
	}
	return sem.IR(), nil
}

func (this *IRGenerator) CompileFunctionBody(f *models.DefinedFunction) []models.IIRStmt {
	this.stmts = make([]models.IIRStmt, 0)
	this.scopeStack = list.New()
	this.breakStack = list.New()
	this.continueStack = list.New()
	this.jumpMap = make(map[string]*JumpEntry)
	this.transformStmt(f.Body())
	this.checkJumpLinks(this.jumpMap)
	return this.stmts
}

func (this *IRGenerator) transformStmt(node models.IASTStmtNode) {
	node.Accept(this)
}

func (this *IRGenerator) transformExpr(node models.IASTExprNode) models.IIRExpr {
	this.exprNestLevel++
	expr, _ := node.Accept(this)
	this.exprNestLevel--
	return expr.(models.IIRExpr)
}

func (g *IRGenerator) isStatement() bool {
	return g.exprNestLevel == 0
}

func (g *IRGenerator) assign(loc *models.Location, lhs models.IIRExpr, rhs models.IIRExpr) {
	g.stmts = append(g.stmts, models.NewIRAssign(loc, g.addressOf(lhs), rhs))
}

func (g *IRGenerator) tmpVar(t models.IType) *models.DefinedVariable {
	v, _ := g.scopeStack.Back().Value.(*models.LocalScope).AllocateTmp(t)
	return v
}

func (g *IRGenerator) label(label *asm.Label) {
	g.stmts = append(g.stmts, models.NewIRLabelStmt(nil, label))
}

func (g *IRGenerator) jump(loc *models.Location, target *asm.Label) {
	g.stmts = append(g.stmts, models.NewIRJump(loc, target))
}

func (g *IRGenerator) jumpTo(target *asm.Label) {
	g.jump(nil, target)
}

func (g *IRGenerator) cjump(loc *models.Location, cond models.IIRExpr, thenLabel, elseLabel *asm.Label) {
	g.stmts = append(g.stmts, models.NewIRCJump(loc, cond, thenLabel, elseLabel))
}

func (g *IRGenerator) pushBreak(label *asm.Label) {
	g.breakStack.PushBack(label)
}

func (g *IRGenerator) popBreak() {
	if g.breakStack.Len() == 0 {
		panic("unmatched push/pop for break stack")
	}
	g.breakStack.Remove(g.breakStack.Back())
}

func (g *IRGenerator) currentBreakTarget() (*asm.Label, error) {
	if g.breakStack.Len() == 0 {
		return nil, errors.New("break from out of loop")
	}
	return g.breakStack.Back().Value.(*asm.Label), nil
}

func (g *IRGenerator) pushContinue(label *asm.Label) {
	g.continueStack.PushBack(label)
}

func (g *IRGenerator) popContinue() {
	if g.continueStack.Len() == 0 {
		panic("unmatched push/pop for continue stack")
	}
	g.continueStack.Remove(g.continueStack.Back())
}

func (g *IRGenerator) currentContinueTarget() (*asm.Label, error) {
	if g.continueStack.Len() == 0 {
		return nil, errors.New("continue from out of loop")
	}
	return g.continueStack.Back().Value.(*asm.Label), nil
}

func (g *IRGenerator) VisitBlockNode(node *models.ASTBlockNode) (any, error) {
	g.scopeStack.PushBack(node.Scope())
	for _, v := range node.Variables() {
		if v.HasInitializer() {
			if v.IsPrivate() {
				v.SetIR(g.transformExpr(v.Initializer()))
			} else {
				g.assign(v.Location(), g.ref(v), g.transformExpr(v.Initializer()))
			}
		}
	}
	for _, s := range node.Stmts() {
		g.transformStmt(s)
	}
	g.scopeStack.Remove(g.scopeStack.Back())
	return nil, nil
}

// TODO
func (g *IRGenerator) VisitExprStmtNode(node *models.ASTExprStmtNode) (any, error) {
	_, err := node.Expr().Accept(g)
	if err != nil {
		g.errorHandler.WarnWithLoc(node.Location(), "useless expression")
	}
	return nil, nil
}

func (g *IRGenerator) VisitIfNode(node *models.ASTIfNode) (any, error) {
	thenLabel := asm.NewLabelUnnamed()
	elseLabel := asm.NewLabelUnnamed()
	endLabel := asm.NewLabelUnnamed()
	cond := g.transformExpr(node.Cond())
	if node.ElseBody() == nil {
		g.cjump(node.Location(), cond, thenLabel, endLabel)
		g.label(thenLabel)
		g.transformStmt(node.ThenBody())
		g.label(endLabel)
	} else {
		g.cjump(node.Location(), cond, thenLabel, elseLabel)
		g.label(thenLabel)
		g.transformStmt(node.ThenBody())
		g.jumpTo(endLabel)
		g.label(elseLabel)
		g.transformStmt(node.ElseBody())
		g.label(endLabel)
	}
	return nil, nil
}

func (g *IRGenerator) VisitSwitchNode(node *models.ASTSwitchNode) (any, error) {
	cases := make([]*models.IRCase, 0)
	endLabel := asm.NewLabelUnnamed()
	defaultLabel := endLabel

	cond := g.transformExpr(node.Cond())
	for _, c := range node.Cases() {
		if c.IsDefault() {
			defaultLabel = c.Label()
		} else {
			for _, val := range c.Values() {
				v := g.transformExpr(val).(*models.IRInt)
				cases = append(cases, models.NewIRCase(v.Value(), c.Label()))
			}
		}
	}
	g.stmts = append(g.stmts, models.NewIRSwitch(node.Location(), cond, cases, defaultLabel, endLabel))
	g.pushBreak(endLabel)
	for _, c := range node.Cases() {
		g.label(c.Label())
		g.transformStmt(c.Body())
	}
	g.popBreak()
	g.label(endLabel)
	return nil, nil
}

func (g *IRGenerator) VisitCaseNode(node *models.ASTCaseNode) (any, error) {
	panic("VisitCaseNode must not happen")
}

func (g *IRGenerator) VisitWhileNode(node *models.ASTWhileNode) (any, error) {
	begLabel := asm.NewLabelUnnamed()
	bodyLabel := asm.NewLabelUnnamed()
	endLabel := asm.NewLabelUnnamed()

	g.label(begLabel)
	g.cjump(node.Location(), g.transformExpr(node.Cond()), bodyLabel, endLabel)
	g.label(bodyLabel)
	g.pushContinue(begLabel)
	g.pushBreak(endLabel)
	g.transformStmt(node.Body())
	g.popBreak()
	g.popContinue()
	g.jumpTo(begLabel)
	g.label(endLabel)
	return nil, nil
}

func (g *IRGenerator) VisitDoWhileNode(node *models.ASTDoWhileNode) (any, error) {
	begLabel := asm.NewLabelUnnamed()
	contLabel := asm.NewLabelUnnamed()
	endLabel := asm.NewLabelUnnamed()

	g.pushContinue(contLabel)
	g.pushBreak(endLabel)
	g.label(begLabel)
	g.transformStmt(node.Body())
	g.popBreak()
	g.popContinue()
	g.label(contLabel)
	g.cjump(node.Location(), g.transformExpr(node.Cond()), begLabel, endLabel)
	g.label(endLabel)
	return nil, nil
}

func (g *IRGenerator) VisitForNode(node *models.ASTForNode) (any, error) {
	begLabel := asm.NewLabelUnnamed()
	bodyLabel := asm.NewLabelUnnamed()
	contLabel := asm.NewLabelUnnamed()
	endLabel := asm.NewLabelUnnamed()
	if node.Init() != nil {
		g.transformStmt(node.Init())
	}
	g.label(begLabel)
	g.cjump(node.Location(), g.transformExpr(node.Cond()), bodyLabel, endLabel)
	g.label(bodyLabel)
	g.pushContinue(contLabel)
	g.pushBreak(endLabel)
	g.transformStmt(node.Body())
	g.popBreak()
	g.popContinue()
	g.label(contLabel)
	if node.Incr() != nil {
		g.transformStmt(node.Incr())
	}
	g.jumpTo(begLabel)
	g.label(endLabel)
	return nil, nil
}

func (g *IRGenerator) VisitBreakNode(node *models.ASTBreakNode) (any, error) {
	label, err := g.currentBreakTarget()
	if err != nil {
		g.error(node, err.Error())
	}
	g.jump(node.Location(), label)
	return nil, nil
}

func (g *IRGenerator) VisitContinueNode(node *models.ASTContinueNode) (any, error) {
	label, err := g.currentContinueTarget()
	if err != nil {
		g.error(node, err.Error())
	}
	g.jump(node.Location(), label)
	return nil, nil
}

func (g *IRGenerator) VisitLabelNode(node *models.ASTLabelNode) (any, error) {
	label, err := g.defineLabel(node.Name(), node.Location())
	if err != nil {
		g.error(node, err.Error())
		return nil, nil
	}
	g.stmts = append(g.stmts, models.NewIRLabelStmt(node.Location(), label))
	if node.Stmt() != nil {
		g.transformStmt(node.Stmt())
	}
	return nil, nil
}

func (g *IRGenerator) VisitGotoNode(node *models.ASTGotoNode) (any, error) {
	g.jump(node.Location(), g.referLabel(node.Target()))
	return nil, nil
}

func (g *IRGenerator) VisitReturnNode(node *models.ASTReturnNode) (any, error) {
	var expr models.IIRExpr
	if node.Expr() != nil {
		expr = g.transformExpr(node.Expr())
	}
	g.stmts = append(g.stmts, models.NewIRReturn(node.Location(), expr))
	return nil, nil
}

func (g *IRGenerator) VisitCondExprNode(node *models.ASTCondExprNode) (any, error) {
	thenLabel := asm.NewLabelUnnamed()
	elseLabel := asm.NewLabelUnnamed()
	endLabel := asm.NewLabelUnnamed()
	v := g.tmpVar(node.Type())

	cond := g.transformExpr(node.Cond())
	g.cjump(node.Location(), cond, thenLabel, elseLabel)
	g.label(thenLabel)
	g.assign(node.ThenExpr().Location(), g.ref(v), g.transformExpr(node.ThenExpr()))
	g.jumpTo(endLabel)
	g.label(elseLabel)
	g.assign(node.ElseExpr().Location(), g.ref(v), g.transformExpr(node.ElseExpr()))
	g.jumpTo(endLabel)
	g.label(endLabel)
	if g.isStatement() {
		return nil, nil
	}
	return g.ref(v), nil
}

func (g *IRGenerator) VisitLogicalAndNode(node *models.ASTLogicalAndNode) (any, error) {
	rightLabel := asm.NewLabelUnnamed()
	endLabel := asm.NewLabelUnnamed()
	v := g.tmpVar(node.Type())

	g.assign(node.Left().Location(), g.ref(v), g.transformExpr(node.Left()))
	g.cjump(node.Location(), g.ref(v), rightLabel, endLabel)
	g.label(rightLabel)
	g.assign(node.Right().Location(), g.ref(v), g.transformExpr(node.Right()))
	g.label(endLabel)
	if g.isStatement() {
		return nil, nil
	}
	return g.ref(v), nil
}

func (g *IRGenerator) VisitLogicalOrNode(node *models.ASTLogicalOrNode) (any, error) {
	rightLabel := asm.NewLabelUnnamed()
	endLabel := asm.NewLabelUnnamed()
	v := g.tmpVar(node.Type())

	g.assign(node.Left().Location(), g.ref(v), g.transformExpr(node.Left()))
	g.cjump(node.Location(), g.ref(v), endLabel, rightLabel)
	g.label(rightLabel)
	g.assign(node.Right().Location(), g.ref(v), g.transformExpr(node.Right()))
	g.label(endLabel)
	if g.isStatement() {
		return nil, nil
	}
	return g.ref(v), nil
}

func (g *IRGenerator) VisitAssignNode(node *models.ASTAssignNode) (any, error) {
	lloc := node.LHS().Location()
	rloc := node.RHS().Location()
	if g.isStatement() {
		rhs := g.transformExpr(node.RHS())
		g.assign(lloc, g.transformExpr(node.LHS()), rhs)
		return nil, nil
	} else {
		tmp := g.tmpVar(node.RHS().Type())
		g.assign(rloc, g.ref(tmp), g.transformExpr(node.RHS()))
		g.assign(lloc, g.transformExpr(node.LHS()), g.ref(tmp))
		return g.ref(tmp), nil
	}
}

func (g *IRGenerator) VisitOpAssignNode(node *models.ASTOpAssignNode) (any, error) {
	rhs := g.transformExpr(node.RHS())
	lhs := g.transformExpr(node.LHS())
	t := node.LHS().Type()
	op := models.InternBinary(node.Operator(), t.IsSigned())
	return g.transformOpAssign(node.Location(), op, t, lhs, rhs), nil
}

func (g *IRGenerator) VisitPrefixOpNode(node *models.ASTPrefixOpNode) (any, error) {
	t := node.Expr().Type()
	return g.transformOpAssign(node.Location(),
		g.binOp(node.Operator()), t,
		g.transformExpr(node.Expr()), g.imm(t, 1)), nil
}

func (g *IRGenerator) VisitSuffixOpNode(node *models.ASTSuffixOpNode) (any, error) {
	expr := g.transformExpr(node.Expr())
	t := node.Expr().Type()
	op := g.binOp(node.Operator())
	loc := node.Location()

	if g.isStatement() {
		g.transformOpAssign(loc, op, t, expr, g.imm(t, 1))
		return nil, nil
	} else if expr.IsVar() {
		v := g.tmpVar(t)
		g.assign(loc, g.ref(v), expr)
		g.assign(loc, expr, g.bin(op, t, g.ref(v), g.imm(t, 1)))
		return g.ref(v), nil
	} else {
		a := g.tmpVar(g.pointerTo(t))
		v := g.tmpVar(t)
		g.assign(loc, g.ref(a), g.addressOf(expr))
		g.assign(loc, g.ref(v), g.memOfEntity(a))
		g.assign(loc, g.memOfEntity(a), g.bin(op, t, g.memOfEntity(a), g.imm(t, 1)))
		return g.ref(v), nil
	}
}

func (g *IRGenerator) transformOpAssign(loc *models.Location, op models.Op, lhsType models.IType, lhs models.IIRExpr, rhs models.IIRExpr) any {
	if lhs.IsVar() {
		g.assign(loc, lhs, g.bin(op, lhsType, lhs, rhs))
		if g.isStatement() {
			return nil
		}
		return lhs
	} else {
		a := g.tmpVar(g.pointerTo(lhsType))
		g.assign(loc, g.ref(a), g.addressOf(lhs))
		g.assign(loc, g.memOfEntity(a), g.bin(op, lhsType, g.memOfEntity(a), rhs))
		if g.isStatement() {
			return nil
		}
		return g.memOfEntity(a)
	}
}

func (g *IRGenerator) bin(op models.Op, leftType models.IType, left, right models.IIRExpr) *models.IRBin {
	if g.isPointerArithmetic(op, leftType) {
		return models.NewIRBin(left.Type(), op, left,
			models.NewIRBin(right.Type(), models.OpMUL, right, g.ptrBaseSize(leftType)))
	} else {
		return models.NewIRBin(left.Type(), op, left, right)
	}
}

func (g *IRGenerator) VisitFunctionCallNode(node *models.ASTFunctionCallNode) (any, error) {
	args := make([]models.IIRExpr, 0, len(node.Args()))
	for _, arg := range node.Args() {
		args = append(args, g.transformExpr(arg))
	}
	call := models.NewIRCall(g.asmType(node.Type()), g.transformExpr(node.Expr()), args)
	if g.isStatement() {
		g.stmts = append(g.stmts, models.NewIRExprStmt(node.Location(), call))
		return nil, nil
	} else {
		tmp := g.tmpVar(node.Type())
		g.assign(node.Location(), g.ref(tmp), call)
		return g.ref(tmp), nil
	}
}

func (g *IRGenerator) VisitBinaryOpNode(node *models.ASTBinaryOpNode) (any, error) {
	right := g.transformExpr(node.Right())
	left := g.transformExpr(node.Left())
	op := models.InternBinary(node.Operator(), node.Type().IsSigned())
	t := node.Type()
	r := node.Right().Type()
	l := node.Left().Type()

	var irBin *models.IRBin
	if g.isPointerDiff(op, l, r) {
		tmp := models.NewIRBin(g.asmType(t), op, left, right)
		irBin = models.NewIRBin(g.asmType(t), models.OpSDIV, tmp, g.ptrBaseSize(l))
	} else if g.isPointerArithmetic(op, l) {
		irBin = models.NewIRBin(g.asmType(t), op, left, models.NewIRBin(g.asmType(r), models.OpMUL, right, g.ptrBaseSize(l)))
	} else if g.isPointerArithmetic(op, r) {
		irBin = models.NewIRBin(g.asmType(t), op, models.NewIRBin(g.asmType(l), models.OpMUL, left, g.ptrBaseSize(r)), right)
	} else {
		irBin = models.NewIRBin(g.asmType(t), op, left, right)
	}
	return irBin, nil
}

func (g *IRGenerator) VisitUnaryOpNode(node *models.ASTUnaryOpNode) (any, error) {
	if node.Operator() == "+" {
		return g.transformExpr(node.Expr()), nil
	} else {
		return models.NewIRUni(g.asmType(node.Type()),
			models.InternUnary(node.Operator()),
			g.transformExpr(node.Expr())), nil
	}
}

func (g *IRGenerator) VisitArrayIdxRefNode(node *models.ASTArrayIdxRefNode) (any, error) {
	expr := g.transformExpr(node.BaseExpr())
	offset := models.NewIRBin(g.ptrdiff_t(), models.OpMUL, g.size(node.ElementSize()), g.transformIndex(node))
	addr := models.NewIRBin(g.ptr_t(), models.OpADD, expr, offset)
	return g.memOfExpr(addr, node.Type()), nil
}

func (g *IRGenerator) transformIndex(node *models.ASTArrayIdxRefNode) models.IIRExpr {
	if node.IsMultiDimension() {
		return models.NewIRBin(g.int_t(), models.OpADD, g.transformExpr(node.Index()),
			models.NewIRBin(g.int_t(), models.OpMUL, models.NewIRInt(g.int_t(), node.Length()),
				g.transformIndex(node.Expr().(*models.ASTArrayIdxRefNode))))
	} else {
		return g.transformExpr(node.Index())
	}
}

// TODO: Offset return error when member not exist, need catch
func (g *IRGenerator) VisitMemberNode(node *models.ASTMemberNode) (any, error) {
	expr := g.addressOf(g.transformExpr(node.Expr()))
	offset := g.ptrdiff(node.Offset())
	addr := models.NewIRBin(g.ptr_t(), models.OpADD, expr, offset)
	if node.IsLoadable() {
		return g.memOfExpr(addr, node.Type()), nil
	} else {
		return addr, nil
	}
}

// TODO: Offset return error when member not exist, need catch
func (g *IRGenerator) VisitPtrMemberNode(node *models.ASTPtrMemberNode) (any, error) {
	expr := g.transformExpr(node.Expr())
	offset := g.ptrdiff(node.Offset())
	addr := models.NewIRBin(g.ptr_t(), models.OpADD, expr, offset)
	if node.IsLoadable() {
		return g.memOfExpr(addr, node.Type()), nil
	} else {
		return addr, nil
	}
}

func (g *IRGenerator) VisitDereferenceNode(node *models.ASTDereferenceNode) (any, error) {
	addr := g.transformExpr(node.Expr())
	if node.IsLoadable() {
		return g.memOfExpr(addr, node.Type()), nil
	} else {
		return addr, nil
	}
}

func (g *IRGenerator) VisitAddressNode(node *models.ASTAddressNode) (any, error) {
	e := g.transformExpr(node.Expr())
	if node.Expr().IsLoadable() {
		return g.addressOf(e), nil
	} else {
		return e, nil
	}
}

func (g *IRGenerator) VisitCastNode(node *models.ASTCastNode) (any, error) {
	if node.IsEffectiveCast() {
		op := models.OpUCAST
		if node.Expr().Type().IsSigned() {
			op = models.OpSCAST
		}
		return models.NewIRUni(g.asmType(node.Type()), op, g.transformExpr(node.Expr())), nil
	} else if g.isStatement() {
		g.transformStmt(node.Expr())
		return nil, nil
	} else {
		return g.transformExpr(node.Expr()), nil
	}
}

func (g *IRGenerator) VisitSizeofExprNode(node *models.ASTSizeofExprNode) (any, error) {
	return models.NewIRInt(g.size_t(), node.Expr().AllocSize()), nil
}

func (g *IRGenerator) VisitSizeofTypeNode(node *models.ASTSizeofTypeNode) (any, error) {
	return models.NewIRInt(g.size_t(), node.Operand().AllocSize()), nil
}

func (g *IRGenerator) VisitVariableNode(node *models.ASTVariableNode) (any, error) {
	if node.Entity().IsConstant() {
		return g.transformExpr(node.Entity().Value()), nil
	}
	varVar := g.ref(node.Entity())
	if node.IsLoadable() {
		return varVar, nil
	} else {
		return g.addressOf(varVar), nil
	}
}

func (g *IRGenerator) VisitIntegerLiteralNode(node *models.ASTIntegerLiteralNode) (any, error) {
	return models.NewIRInt(g.asmType(node.Type()), node.Value()), nil
}

func (g *IRGenerator) VisitStringLiteralNode(node *models.ASTStringLiteralNode) (any, error) {
	return models.NewIRStr(g.asmType(node.Type()), node.Entry()), nil
}

func (g *IRGenerator) isPointerDiff(op models.Op, l, r models.IType) bool {
	return op == models.OpSUB && l.IsPointer() && r.IsPointer()
}

func (g *IRGenerator) isPointerArithmetic(op models.Op, operandType models.IType) bool {
	switch op {
	case models.OpADD, models.OpSUB:
		return operandType.IsPointer()
	default:
		return false
	}
}

func (g *IRGenerator) ptrBaseSize(t models.IType) models.IIRExpr {
	return models.NewIRInt(g.ptrdiff_t(), t.ElemType().Size())
}

func (g *IRGenerator) binOp(uniOp string) models.Op {
	if uniOp == "++" {
		return models.OpADD
	} else {
		return models.OpSUB
	}
}

func (g *IRGenerator) addressOf(expr models.IIRExpr) models.IIRExpr {
	return expr.AddressNode(g.ptr_t())
}

func (g *IRGenerator) ref(ent models.IEntity) *models.IRVar {
	return models.NewIRVar(g.varType(ent.Type()), ent)
}

func (g *IRGenerator) memOfEntity(ent models.IEntity) *models.IRMem {
	return models.NewIRMem(g.asmType(ent.Type().ElemType()), g.ref(ent))
}

func (g *IRGenerator) memOfExpr(expr models.IIRExpr, t models.IType) *models.IRMem {
	return models.NewIRMem(g.asmType(t), expr)
}

func (g *IRGenerator) ptrdiff(n int64) *models.IRInt {
	return models.NewIRInt(g.ptrdiff_t(), n)
}

func (g *IRGenerator) size(n int64) *models.IRInt {
	return models.NewIRInt(g.size_t(), n)
}

func (g *IRGenerator) imm(operandType models.IType, n int64) *models.IRInt {
	if operandType.IsPointer() {
		return models.NewIRInt(g.ptrdiff_t(), n)
	} else {
		return models.NewIRInt(g.int_t(), n)
	}
}

func (g *IRGenerator) pointerTo(t models.IType) models.IType {
	return g.typeTable.PointerTo(t)
}

func (g *IRGenerator) asmType(t models.IType) asm.Type {
	if t.IsVoid() {
		return g.int_t()
	}
	return asm.GetType(t.Size())
}

func (g *IRGenerator) varType(t models.IType) asm.Type {
	if !t.IsScalar() {
		// TODO: java return null, 		return nil
		// here panic
		panic("type is not scalar")
	}
	return asm.GetType(t.Size())
}

func (g *IRGenerator) int_t() asm.Type {
	return asm.GetType(int64(g.typeTable.IntSize()))
}

func (g *IRGenerator) size_t() asm.Type {
	return asm.GetType(int64(g.typeTable.LongSize()))
}

func (g *IRGenerator) ptr_t() asm.Type {
	return asm.GetType(int64(g.typeTable.PointerSize()))
}

func (g *IRGenerator) ptrdiff_t() asm.Type {
	return asm.GetType(int64(g.typeTable.LongSize()))
}

func (g *IRGenerator) error(node models.INode, msg string) {
	g.errorHandler.ErrorWithLoc(node.Location(), msg)
}

func (g *IRGenerator) defineLabel(name string, loc *models.Location) (*asm.Label, error) {
	ent := g.getJumpEntry(name)
	if ent.isDefined {
		return nil, fmt.Errorf("duplicated jump labels in %s(): %s", name, name)
	}
	ent.isDefined = true
	ent.location = loc
	return ent.label, nil
}

func (g *IRGenerator) referLabel(name string) *asm.Label {
	ent := g.getJumpEntry(name)
	ent.numRefered++
	return ent.label
}

func (g *IRGenerator) getJumpEntry(name string) *JumpEntry {
	if ent, ok := g.jumpMap[name]; ok {
		return ent
	}
	ent := &JumpEntry{
		label: asm.NewLabelUnnamed(),
	}
	g.jumpMap[name] = ent
	return ent
}

func (g *IRGenerator) checkJumpLinks(jumpMap map[string]*JumpEntry) {
	for name, jump := range jumpMap {
		if !jump.isDefined {
			g.errorHandler.ErrorWithLoc(jump.location, "undefined label: "+name)
		}
		if jump.numRefered == 0 {
			g.errorHandler.WarnWithLoc(jump.location, "useless label: "+name)
		}
	}
}
