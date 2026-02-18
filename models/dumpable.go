package models

type Dumpable interface {
	Dump(d *Dumper)
}
