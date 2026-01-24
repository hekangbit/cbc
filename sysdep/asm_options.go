package sysdep

type AssemblerOptions struct {
	Verbose bool
	args    []string
}

func (opts *AssemblerOptions) AddArg(arg string) {
	opts.args = append(opts.args, arg)
}
