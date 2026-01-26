package compiler

import (
	"cbc/loader"
	"cbc/sysdep"
	"cbc/types"
	"fmt"
	"os"
	"strings"
)

type Options struct {
	mode           CompilerMode
	platform       sysdep.Platform
	outputFileName string
	verbose        bool
	loader         loader.LibraryLoader
	debugParser    bool
	genOptions     sysdep.CodeGeneratorOptions
	asOptions      sysdep.AssemblerOptions
	ldOptions      sysdep.LinkerOptions
	ldArgs         []LdArg
	sourceFiles    []SourceFile
}

func ParseOptions(args []string) (*Options, error) {
	opts := &Options{mode: COMPILER_MODE_None, outputFileName: "", platform: &sysdep.X86Linux{}}
	err := opts.ParserArgs(args)
	return opts, err
}

func (options *Options) Mode() CompilerMode {
	return options.mode
}

func (options *Options) IsAssembleRequired() bool {
	return options.mode.Requires(COMPILER_MODE_Assemble)
}

func (options *Options) IsLinkRequired() bool {
	return options.mode.Requires(COMPILER_MODE_Link)
}

func (options *Options) SourceFiles() []SourceFile {
	return options.sourceFiles
}

func (options *Options) GetTypeTable() types.TypeTable {
	return options.platform.GetTypeTable()
}

func (options *Options) IsGeneratingSharedLibrary() bool {
	return options.ldOptions.GeneratingSharedLibrary
}

func (options *Options) ParserArgs(args []string) error {
	gotStopsCmd := false
	for i := 0; i < len(args); i++ {
		arg := args[i]
		if gotStopsCmd {
			sourcFile := &SourceFile{originalName: arg, currentName: arg}
			options.ldArgs = append(options.ldArgs, sourcFile)
			continue
		}
		if arg == "--" {
			gotStopsCmd = true
			continue // "--" Stops command line processing
		}

		if strings.HasPrefix(arg, "-") {
			if tmp_mode, ok := GetCompilerMode(arg); ok {
				if options.mode != COMPILER_MODE_None {
					return fmt.Errorf("%s", tmp_mode.String()+" option and "+options.mode.String()+" option is exclusive")
				}
				options.mode = tmp_mode
			} else if strings.HasPrefix(arg, "-I") {
				next, err := GetOptArg(arg, args, i)
				if err != nil {
					return err
				}
				i++
				options.loader.AddLoadPath(next)
			} else if arg == "--debug-parser" {
				options.debugParser = true
			} else if strings.HasPrefix(arg, "-o") {
				next, err := NextArg(arg, args, i)
				if err != nil {
					return err
				}
				i++
				options.outputFileName = next
			} else if arg == "-fpic" || arg == "-fPIC" {
				options.genOptions.EnableGeneratePIC()
			} else if arg == "-fpie" || arg == "-fPIE" {
				options.genOptions.EnableGeneratePIE()
			} else if strings.HasPrefix(arg, "-O") {
				c := arg[2]
				// only support O0 and O1
				if len(arg) == 3 && (c == '0' || c == '1' || c == '2' || c == '3' || c == 's') {
					var level int = 0
					if c != '0' {
						level = 1
					}
					options.genOptions.SetOptimizationLevel(level)
				}
				return fmt.Errorf("unknown optimization switch: %s", arg)
			} else if arg == "-fverbose-asm" || arg == "--verbose-asm" {
				options.genOptions.GenerateVerboseAsm()
			} else if strings.HasPrefix(arg, "-Wa,") {
				for _, a := range strings.Split(arg, ",") {
					options.AddLdArg(a)
					options.asOptions.AddArg(a)
				}
			} else if arg == "-Xassembler" {
				next, err := NextArg(arg, args, i)
				if err != nil {
					return err
				}
				i++
				options.asOptions.AddArg(next)
			} else if arg == "-static" {
				options.AddLdArg(arg)
			} else if arg == "-shared" {
				options.ldOptions.GeneratingSharedLibrary = true
			} else if arg == "-pie" {
				options.ldOptions.GeneratingPIE = true
			} else if arg == "--readonly-got" {
				options.AddLdArg("-z")
				options.AddLdArg("combreloc")
				options.AddLdArg("-z")
				options.AddLdArg("now")
				options.AddLdArg("-z")
				options.AddLdArg("relro")
			} else if strings.HasPrefix(arg, "-L") {
				next, err := GetOptArg(arg, args, i)
				if err != nil {
					return err
				}
				i++
				options.AddLdArg("-L" + next)
			} else if strings.HasPrefix(arg, "-l") {
				next, err := GetOptArg(arg, args, i)
				if err != nil {
					return err
				}
				i++
				options.AddLdArg("-l" + next)
			} else if arg == "-nostartfiles" {
				options.ldOptions.NoStartFiles = true
			} else if arg == "-nodefaultlibs" {
				options.ldOptions.NoDefaultLibs = true
			} else if arg == "-nostdlib" {
				options.ldOptions.NoStartFiles = true
				options.ldOptions.NoDefaultLibs = true
			} else if strings.HasPrefix(arg, "-Wl,") {
				for _, a := range strings.Split(arg, ",") {
					options.AddLdArg(a)
				}
			} else if arg == "-Xlinker" {
				next, err := NextArg(arg, args, i)
				if err != nil {
					return err
				}
				i++
				options.AddLdArg(next)
			} else if arg == "-v" {
				options.verbose = true
				options.asOptions.Verbose = true
				options.ldOptions.Verbose = true
			} else if arg == "--version" {
				fmt.Printf("%s version %s\n", CompilerProgramName, CompilerVersion)
				os.Exit(0)
			} else if arg == "--help" {
				PrintUsage()
				os.Exit(0)
			} else {
				return fmt.Errorf("unknown option: %s", arg)
			}
		} else {
			sourcFile := &SourceFile{originalName: arg, currentName: arg}
			options.ldArgs = append(options.ldArgs, sourcFile)
		}
	}

	if options.mode == COMPILER_MODE_None {
		options.mode = COMPILER_MODE_Link
	}

	// select source file from ld args
	for _, ld_arg := range options.ldArgs {
		if ld_arg.IsSourceFile() {
			options.sourceFiles = append(options.sourceFiles, *ld_arg.(*SourceFile))
		}
	}

	if len(options.sourceFiles) == 0 {
		return fmt.Errorf("no input file")
	}

	for _, src := range options.sourceFiles {
		if !src.IsknownFileType() {
			return fmt.Errorf("unknown file type: %s", src.Path())
		}
	}

	if options.outputFileName != "" && len(options.sourceFiles) > 1 && options.IsLinkRequired() {
		return fmt.Errorf("-o option requires only 1 input (except linking)")
	}

	return nil
}

func (options *Options) AddLdArg(arg string) {
	opt := &LdOption{arg: arg}
	options.ldArgs = append(options.ldArgs, opt)
}

func (options *Options) AsmFileNameOf(src SourceFile) string {
	if options.outputFileName != "" && options.mode == COMPILER_MODE_Compile {
		return options.outputFileName
	}
	return src.AsmFileName()
}

func (options *Options) ObjFileNameOf(src SourceFile) string {
	if options.outputFileName != "" && options.mode == COMPILER_MODE_Assemble {
		return options.outputFileName
	}
	return src.ObjFileName()
}

func NextArg(opt string, args []string, iter int) (string, error) {
	if iter >= len(args)-1 {
		return "", fmt.Errorf("missing argument for: %s", opt)
	}
	return args[iter+1], nil
}

func GetOptArg(opt string, args []string, iter int) (string, error) {
	path := opt[2:]
	if len(path) > 0 {
		return path, nil
	}
	return NextArg(opt, args, iter)
}

func PrintUsage() {
	fmt.Println("Usage: cbc [options] file...")
	fmt.Println("Global Options:")
	fmt.Println("  --check-syntax   Checks syntax and quit.")
	fmt.Println("  --dump-tokens    Dumps tokens and quit.")
	// --dump-stmt is a hidden option.
	// --dump-expr is a hidden option.
	fmt.Println("  --dump-ast       Dumps AST and quit.")
	fmt.Println("  --dump-semantic  Dumps AST after semantic checks and quit.")
	// --dump-reference is a hidden option.
	fmt.Println("  --dump-ir        Dumps IR and quit.")
	fmt.Println("  --dump-asm       Dumps AssemblyCode and quit.")
	fmt.Println("  --print-asm      Prints assembly code and quit.")
	fmt.Println("  -S               Generates an assembly file and quit.")
	fmt.Println("  -c               Generates an object file and quit.")
	fmt.Println("  -o PATH          Places output in file PATH.")
	fmt.Println("  -v               Turn on verbose mode.")
	fmt.Println("  --version        Shows compiler version and quit.")
	fmt.Println("  --help           Prints this message and quit.")
	fmt.Println("")
	fmt.Println("Optimization Options:")
	fmt.Println("  -O               Enables optimization.")
	fmt.Println("  -O1, -O2, -O3    Equivalent to -O.")
	fmt.Println("  -Os              Equivalent to -O.")
	fmt.Println("  -O0              Disables optimization (default).")
	fmt.Println("")
	fmt.Println("Parser Options:")
	fmt.Println("  -I PATH          Adds PATH as import file directory.")
	fmt.Println("  --debug-parser   Dumps parsing process.")
	fmt.Println("")
	fmt.Println("Code Generator Options:")
	fmt.Println("  -O               Enables optimization.")
	fmt.Println("  -O1, -O2, -O3    Equivalent to -O.")
	fmt.Println("  -Os              Equivalent to -O.")
	fmt.Println("  -O0              Disables optimization (default).")
	fmt.Println("  -fPIC            Generates PIC assembly.")
	fmt.Println("  -fpic            Equivalent to -fPIC.")
	fmt.Println("  -fPIE            Generates PIE assembly.")
	fmt.Println("  -fpie            Equivalent to -fPIE.")
	fmt.Println("  -fverbose-asm    Generate assembly with verbose comments.")
	fmt.Println("")
	fmt.Println("Assembler Options:")
	fmt.Println("  -Wa,OPT          Passes OPT to the assembler (as).")
	fmt.Println("  -Xassembler OPT  Passes OPT to the assembler (as).")
	fmt.Println("")
	fmt.Println("Linker Options:")
	fmt.Println("  -l LIB           Links the library LIB.")
	fmt.Println("  -L PATH          Adds PATH as library directory.")
	fmt.Println("  -shared          Generates shared library rather than executable.")
	fmt.Println("  -static          Linkes only with static libraries.")
	fmt.Println("  -pie             Generates PIE.")
	fmt.Println("  --readonly-got   Generates read-only GOT (ld -z combreloc -z now -z relro).")
	fmt.Println("  -nostartfiles    Do not link startup files.")
	fmt.Println("  -nodefaultlibs   Do not link default libraries.")
	fmt.Println("  -nostdlib        Enables -nostartfiles and -nodefaultlibs.")
	fmt.Println("  -Wl,OPT          Passes OPT to the linker (ld).")
	fmt.Println("  -Xlinker OPT     Passes OPT to the linker (ld).")
}
