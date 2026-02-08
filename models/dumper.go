package models

import (
	"cbc/util"
	"fmt"
	"io"
	"strings"
)

const indentString = "    "

// Dumper 工具接口（假设定义，需根据项目补充）
type Dumper struct {
	nIndent int
	stream  io.Writer
}

// NewDumper 创建一个新的Dumper实例
func NewDumper(s io.Writer) *Dumper {
	return &Dumper{
		stream:  s,
		nIndent: 0,
	}
}

// PrintClass 打印类信息
func (d *Dumper) PrintClass(obj interface{}, loc *Location) {
	d.printIndent()
	fmt.Fprintf(d.stream, "<<%T>> (%v)\n", obj, loc)
}

// PrintNodeList 打印节点列表
func (d *Dumper) PrintNodeList(name string, nodes []Dumpable) {
	d.printIndent()
	fmt.Fprintf(d.stream, "%s:\n", name)
	d.indent()
	for _, n := range nodes {
		n.Dump(d)
	}
	d.unindent()
}

// PrintMemberInt 打印int类型成员
func (d *Dumper) PrintMemberInt(name string, n int) {
	d.printPair(name, fmt.Sprintf("%d", n))
}

// PrintMemberInt64 打印long类型成员
func (d *Dumper) PrintMemberInt64(name string, n int64) {
	d.printPair(name, fmt.Sprintf("%d", n))
}

// PrintMemberBool 打印bool类型成员
func (d *Dumper) PrintMemberBool(name string, b bool) {
	d.printPair(name, fmt.Sprintf("%t", b))
}

// PrintMemberTypeRef 打印TypeRef类型成员
func (d *Dumper) PrintMemberTypeRef(name string, ref ITypeRef) {
	d.printPair(name, ref.String())
}

// PrintMemberType 打印Type类型成员
func (d *Dumper) PrintMemberType(name string, t IType) {
	if t == nil {
		d.printPair(name, "null")
	} else {
		d.printPair(name, t.String())
	}
}

// PrintMemberString 打印字符串成员（带解析状态）
func (d *Dumper) PrintMemberString(name string, str string, isResolved bool) {
	resolved := ""
	if isResolved {
		resolved = " (resolved)"
	}
	d.printPair(name, util.DumpString(str)+resolved)
}

// PrintMemberStringSimple 打印简单字符串成员
func (d *Dumper) PrintMemberStringSimple(name string, str string) {
	d.PrintMemberString(name, str, false)
}

func (d *Dumper) printPair(name, value string) {
	d.printIndent()
	fmt.Fprintf(d.stream, "%s: %s\n", name, value)
}

// PrintMemberTypeNode 打印TypeNode成员
func (d *Dumper) PrintMemberTypeNode(name string, n *TypeNode) {
	d.printIndent()
	resolved := ""
	if n.IsResolved() {
		resolved = " (resolved)"
	}
	fmt.Fprintf(d.stream, "%s: %v%s\n", name, n.TypeRef(), resolved)
}

// PrintMemberDumpable 打印Dumpable成员
func (d *Dumper) PrintMemberDumpable(name string, n Dumpable) {
	d.printIndent()
	if n == nil {
		fmt.Fprintf(d.stream, "%s: null\n", name)
	} else {
		fmt.Fprintf(d.stream, "%s:\n", name)
		d.indent()
		n.Dump(d)
		d.unindent()
	}
}

func (d *Dumper) indent() {
	d.nIndent++
}

func (d *Dumper) unindent() {
	d.nIndent--
}

func (d *Dumper) printIndent() {
	d.stream.Write([]byte(strings.Repeat(indentString, d.nIndent)))
}
