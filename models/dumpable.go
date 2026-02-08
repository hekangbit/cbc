package models

// Dumpable 接口（原 Java 的 Dumpable）
type Dumpable interface {
	Dump(d *Dumper)
}
