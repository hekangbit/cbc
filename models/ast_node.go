package models

type INode interface {
	Dumpable
	Location() *Location
}

// TODO: no need define struct for this interface
// cause No usefull shared method to define
type Node struct {
}

// TODO: all concreate Node struct need has Dump method
// and need call printClass(this, location) first,
// see java code for detail
func (n *Node) Dump(d *Dumper) {
	panic("Node::Dump method must be implemented by concrete node type")
}

func (n *Node) Location() *Location {
	panic("Node::Location method must be implemented by concrete node type")
}
