package compiler

type LdOption struct {
	arg string
}

func (option *LdOption) IsSourceFile() bool {
	return false
}
