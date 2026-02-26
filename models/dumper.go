package models

import (
	"cbc/util"
	"fmt"
	"io"
	"strings"
)

const indentString = "    "

type Dumper struct {
	nIndent int
	stream  io.Writer
}

func NewDumper(s io.Writer) *Dumper {
	return &Dumper{
		stream:  s,
		nIndent: 0,
	}
}

func (this *Dumper) PrintClass(obj interface{}, loc *Location) {
	this.printIndent()
	fmt.Fprintf(this.stream, "<<%T>> (%v)\n", obj, loc)
}

func (this *Dumper) PrintNodeList(name string, nodes []Dumpable) {
	this.printIndent()
	fmt.Fprintf(this.stream, "%s:\n", name)
	this.indent()
	for _, n := range nodes {
		n.Dump(this)
	}
	this.unindent()
}

func (this *Dumper) PrintMemberChar(name string, n int) {
	this.printPair(name, fmt.Sprintf("%d, (%c)", n, n))
}

func (this *Dumper) PrintMemberInt(name string, n int) {
	this.printPair(name, fmt.Sprintf("%d", n))
}

func (this *Dumper) PrintMemberInt64(name string, n int64) {
	this.printPair(name, fmt.Sprintf("%d", n))
}

func (this *Dumper) PrintMemberBool(name string, b bool) {
	this.printPair(name, fmt.Sprintf("%t", b))
}

func (this *Dumper) PrintMemberTypeRef(name string, ref ITypeRef) {
	this.printPair(name, ref.String())
}

func (this *Dumper) PrintMemberType(name string, t IType) {
	if t == nil {
		this.printPair(name, "null")
	} else {
		this.printPair(name, t.String())
	}
}

func (this *Dumper) PrintMemberString(name string, str string, isResolved bool) {
	resolved := ""
	if isResolved {
		resolved = " (resolved)"
	}
	this.printPair(name, util.DumpString(str)+resolved)
}

func (this *Dumper) PrintMemberStringNotResolved(name string, str string) {
	this.PrintMemberString(name, str, false)
}

func (this *Dumper) printPair(name, value string) {
	this.printIndent()
	fmt.Fprintf(this.stream, "%s: %s\n", name, value)
}

func (this *Dumper) PrintMemberTypeNode(name string, n *ASTTypeNode) {
	this.printIndent()
	resolved := ""
	if n.IsResolved() {
		resolved = " (resolved)"
	}
	fmt.Fprintf(this.stream, "%s: %v%s\n", name, n.TypeRef(), resolved)
}

func (this *Dumper) PrintMemberDumpable(name string, n Dumpable) {
	this.printIndent()
	if n == nil {
		fmt.Fprintf(this.stream, "%s: null\n", name)
	} else {
		fmt.Fprintf(this.stream, "%s:\n", name)
		this.indent()
		n.Dump(this)
		this.unindent()
	}
}

func (this *Dumper) indent() {
	this.nIndent++
}

func (this *Dumper) unindent() {
	this.nIndent--
}

func (this *Dumper) printIndent() {
	this.stream.Write([]byte(strings.Repeat(indentString, this.nIndent)))
}
