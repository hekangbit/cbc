package models

import (
	"cbc/asm"
)

type Entity struct {
	name      string
	isPrivate bool
	typeNode  TypeNode
	nRefered  int64
	memref    asm.MemoryReference
	address   asm.Operand
}
