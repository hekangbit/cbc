package models

type IIRStmt interface {
	Dumpable
	Location() *Location
	Accept(visitor IRVisitor) any
	_Dump(d *Dumper)
}

type IRStmt struct {
	_impl    IIRStmt
	location *Location
}

func (this *IRStmt) Location() *Location {
	return this.location
}

func (this *IRStmt) Dump(d *Dumper) {
	d.PrintClass(this._impl, this._impl.Location())
	this._impl._Dump(d)
}
