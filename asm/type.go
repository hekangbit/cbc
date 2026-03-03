package asm

import "fmt"

// TODO: java Type is enum, which is class, how about golang
type Type int

const (
	INT8 Type = iota
	INT16
	INT32
	INT64
)

func GetType(size int64) Type {
	switch size {
	case 1:
		return INT8
	case 2:
		return INT16
	case 4:
		return INT32
	case 8:
		return INT64
	default:
		panic(fmt.Sprintf("unsupported asm type size: %d", size))
	}
}

func (t Type) Size() int {
	switch t {
	case INT8:
		return 1
	case INT16:
		return 2
	case INT32:
		return 4
	case INT64:
		return 8
	default:
		panic("must not happen")
	}
}
