package compiler

type CompilerMode int

const (
	COMPILER_MODE_None CompilerMode = iota
	COMPILER_MODE_CheckSyntax
	COMPILER_MODE_DumpTokens
	COMPILER_MODE_DumpAST
	COMPILER_MODE_DumpStmt
	COMPILER_MODE_DumpExpr
	COMPILER_MODE_DumpSemantic
	COMPILER_MODE_DumpReference
	COMPILER_MODE_DumpIR
	COMPILER_MODE_DumpAsm
	COMPILER_MODE_PrintAsm
	COMPILER_MODE_Compile
	COMPILER_MODE_Assemble
	COMPILER_MODE_Link
)

var modes map[string]CompilerMode = map[string]CompilerMode{
	"--check-syntax":   COMPILER_MODE_CheckSyntax,
	"--dump-tokens":    COMPILER_MODE_DumpTokens,
	"--dump-ast":       COMPILER_MODE_DumpAST,
	"--dump-stmt":      COMPILER_MODE_DumpStmt,
	"--dump-expr":      COMPILER_MODE_DumpExpr,
	"--dump-semantic":  COMPILER_MODE_DumpSemantic,
	"--dump-reference": COMPILER_MODE_DumpReference,
	"--dump-ir":        COMPILER_MODE_DumpIR,
	"--dump-asm":       COMPILER_MODE_DumpAsm,
	"--print-asm":      COMPILER_MODE_PrintAsm,
	"-S":               COMPILER_MODE_Compile,
	"-c":               COMPILER_MODE_Assemble,
	"--link":           COMPILER_MODE_Link,
}

var modeOptions = []string{
	COMPILER_MODE_CheckSyntax:   "--check-syntax",
	COMPILER_MODE_DumpTokens:    "--dump-tokens",
	COMPILER_MODE_DumpAST:       "--dump-ast",
	COMPILER_MODE_DumpStmt:      "--dump-stmt",
	COMPILER_MODE_DumpExpr:      "--dump-expr",
	COMPILER_MODE_DumpSemantic:  "--dump-semantic",
	COMPILER_MODE_DumpReference: "--dump-reference",
	COMPILER_MODE_DumpIR:        "--dump-ir",
	COMPILER_MODE_DumpAsm:       "--dump-asm",
	COMPILER_MODE_PrintAsm:      "--print-asm",
	COMPILER_MODE_Compile:       "-S",
	COMPILER_MODE_Assemble:      "-c",
	COMPILER_MODE_Link:          "--link",
}

func GetCompilerMode(opt string) (CompilerMode, bool) {
	mode, ok := modes[opt]
	return mode, ok
}

func (m CompilerMode) String() string {
	return modeOptions[m]
}

func (m CompilerMode) Requires(other CompilerMode) bool {
	return m >= other
}
