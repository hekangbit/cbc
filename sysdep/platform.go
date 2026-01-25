package sysdep

import "cbc/types"

type Platform interface {
	GetTypeTable() types.TypeTable
}
