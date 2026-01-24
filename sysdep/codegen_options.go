package sysdep

type CodeGeneratorOptions struct {
	OptimizeLevel int
	GeneratePIC   bool
	GeneratePIE   bool
	VerboseAsm    bool
}

func (opts *CodeGeneratorOptions) SetOptimizationLevel(level int) {
	opts.OptimizeLevel = level
}

func (opts *CodeGeneratorOptions) GenerateVerboseAsm() {
	opts.VerboseAsm = true
}

func (opts *CodeGeneratorOptions) EnableGeneratePIC() {
	opts.GeneratePIC = true
}

func (opts *CodeGeneratorOptions) IsPICRequired() bool {
	return opts.GeneratePIC
}

func (opts *CodeGeneratorOptions) EnableGeneratePIE() {
	opts.GeneratePIE = true
}

func (opts *CodeGeneratorOptions) IsPIERequired() bool {
	return opts.GeneratePIE
}
