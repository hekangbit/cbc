package models

import (
	"io"
	"os"
)

type INode interface {
	Dumpable
	DumpNode()
	DumpByStream(io.Writer)
	_Dump(d *Dumper)
	Location() *Location
}

type Node struct {
	_impl INode
}

func (this *Node) DumpNode() {
	this._impl.DumpByStream(os.Stdout)
}

func (this *Node) DumpByStream(s io.Writer) {
	this._impl.Dump(NewDumper(s))
}

func (this *Node) Dump(d *Dumper) {
	// TODO: java dump this class, but how about golang
	d.PrintClass(this._impl, this._impl.Location())
	this._impl._Dump(d)
}
