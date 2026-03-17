package models

import "cbc/asm"

type IRCall struct {
	IRExpr
	expr IIRExpr
	args []IIRExpr
}

var _ IIRExpr = &IRCall{}

func NewIRCall(t asm.Type, expr IIRExpr, args []IIRExpr) *IRCall {
	p := &IRCall{
		IRExpr: IRExpr{typ: t},
		expr:   expr,
		args:   args,
	}
	p._impl = p
	return p
}

func (this *IRCall) Expr() IIRExpr {
	return this.expr
}

func (this *IRCall) Args() []IIRExpr {
	return this.args
}

func (this *IRCall) NumArgs() int {
	return len(this.args)
}

func (this *IRCall) IsStaticCall() bool {
	ent := this.expr.GetEntityForce()
	if ent == nil {
		return false
	}
	_, ok := ent.(IFunction)
	return ok
}

func (this *IRCall) Function() IFunction {
	ent := this.expr.GetEntityForce()
	if ent == nil {
		panic("not a static funcall")
	}
	fn, ok := ent.(IFunction)
	if !ok {
		panic("entity is not a Function")
	}
	return fn
}

func (this *IRCall) Accept(visitor IRVisitor) any {
	return visitor.VisitCall(this)
}

func (this *IRCall) _Dump(d *Dumper) {
	d.PrintMemberDumpable("expr", this.expr)

	buf := make([]Dumpable, len(this.args))
	for i, tmp := range this.args {
		buf[i] = tmp
	}
	d.PrintNodeList("args", buf)
}
