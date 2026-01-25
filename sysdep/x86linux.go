package sysdep

import "cbc/types"

type X86Linux struct {
}

func (platform *X86Linux) GetTypeTable() types.TypeTable {
	return types.TypeTable{}
}
