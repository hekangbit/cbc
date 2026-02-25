package models

type Dumpable interface {
	Dump(*Dumper)
}
