package sysdep

type LinkerOptions struct {
	Verbose                 bool
	NoStartFiles            bool
	NoDefaultLibs           bool
	GeneratingPIE           bool
	GeneratingSharedLibrary bool
}
