package sysdep

import (
	"cbc/models"
)

type X86Linux struct {
}

func (platform *X86Linux) GetTypeTable() models.TypeTable {
	return models.TypeTable{}
}
