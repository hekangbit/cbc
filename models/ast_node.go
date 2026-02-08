package models

type Node interface {
	Dumpable
	Location() *Location
}

type BaseNode struct {
}

// // Dump 默认转储到标准输出
// func (n *BaseNode) DumpDefault() {
// 	n.DumpTo(os.Stdout)
// }

// // DumpTo 转储到指定的输出流
// func (n *BaseNode) DumpTo(out io.Writer) {
// 	dumper := NewDumper(out)
// 	n.DumpWith(dumper)
// }

// // DumpWith 使用指定的转储器进行转储
// func (n *BaseNode) DumpWith(d *Dumper) {
// 	// 获取实际节点的类型和位置
// 	// 由于 BaseNode 不知道具体实现，这里需要在子类中重写
// 	d.PrintClass(n, nil)
// 	n.dump(d)
// }

// dump 内部转储方法（由子类实现）
func (n *BaseNode) Dump(d *Dumper) {
	// 由具体节点类型实现
	panic("dump method must be implemented by concrete node type")
}

// Location 获取位置信息（由子类实现）
func (n *BaseNode) Location() *Location {
	// 由具体节点类型实现
	panic("Location method must be implemented by concrete node type")
}
