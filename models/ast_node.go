package models

// TODO: dump overload

type INode interface {
	Dumpable
	Location() *Location
}

type Node struct {
}

func (n *Node) Dump(d *Dumper) {
	panic("dump method must be implemented by concrete node type")
}

func (n *Node) Location() *Location {
	panic("Location method must be implemented by concrete node type")
}
