package main

type LdOption struct {
	arg string
}

func (option *LdOption) IsSourceFile() bool {
	return false
}
