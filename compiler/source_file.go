package compiler

import (
	"path/filepath"
	"strings"
)

const (
	EXT_CFLAT_SOURCE    = ".cb"
	EXT_ASSEMBLY_SOURCE = ".s"
	EXT_OBJECT_FILE     = ".o"
	EXT_STATIC_LIBRARY  = ".a"
	EXT_SHARED_LIBRARY  = ".so"
	EXT_EXECUTABLE_FILE = ""
)

var KNOWN_EXTENSIONS []string = []string{
	EXT_CFLAT_SOURCE,
	EXT_ASSEMBLY_SOURCE,
	EXT_OBJECT_FILE,
	EXT_STATIC_LIBRARY,
	EXT_SHARED_LIBRARY,
	EXT_EXECUTABLE_FILE,
}

type SourceFile struct {
	originalName string
	currentName  string
}

func (src *SourceFile) IsSourceFile() bool {
	return true
}

func (src *SourceFile) String() string {
	return src.currentName
}

func (src *SourceFile) Path() string {
	return src.currentName
}

func (src *SourceFile) CurrentName() string {
	return src.currentName
}

func (src *SourceFile) SetCurrentName(name string) {
	src.currentName = name
}

func (src *SourceFile) IsknownFileType() bool {
	ext := src.ExtName(src.originalName)
	for _, e := range KNOWN_EXTENSIONS {
		if e == ext {
			return true
		}
	}
	return false
}

func (src *SourceFile) IsCbSource() bool {
	return strings.HasSuffix(src.currentName, EXT_CFLAT_SOURCE)
}

func (src *SourceFile) IsAssemblySource() bool {
	return strings.HasSuffix(src.currentName, EXT_ASSEMBLY_SOURCE)
}

func (src *SourceFile) IsObjectFile() bool {
	return strings.HasSuffix(src.currentName, EXT_OBJECT_FILE)
}

func (src *SourceFile) IsSharedLibrary() bool {
	return strings.HasSuffix(src.currentName, EXT_STATIC_LIBRARY)
}

func (src *SourceFile) IsStaticLibrary() bool {
	return strings.HasSuffix(src.currentName, EXT_SHARED_LIBRARY)
}

func (src *SourceFile) IsExecutable() bool {
	return strings.HasSuffix(src.currentName, EXT_EXECUTABLE_FILE)
}

func (src *SourceFile) AsmFileName() string {
	return src.ReplaceExt(EXT_ASSEMBLY_SOURCE)
}

func (src *SourceFile) ObjFileName() string {
	return src.ReplaceExt(EXT_OBJECT_FILE)
}

func (src *SourceFile) LinkedFileName() string {
	return src.ReplaceExt(EXT_OBJECT_FILE)
}

func (src *SourceFile) ReplaceExt(ext string) string {
	name := filepath.Base(src.originalName)
	return strings.TrimSuffix(name, filepath.Ext(name)) + ext
}

func (src *SourceFile) BaseName(path string) string {
	return filepath.Base(path)
}

func (src *SourceFile) ExtName(path string) string {
	return filepath.Ext(path)
}
