package main

import (
	"cbc/compiler"
	"os"
)

func main() {
	compiler.Run(os.Args[1:])
}
