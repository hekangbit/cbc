package sysdep

import (
	"cbc/models"
)

// TODO:
type X86Linux struct {
}

var _ IPlatform = &X86Linux{}

func NewX86Linux() *X86Linux {
	return &X86Linux{}
}

func (this *X86Linux) TypeTable() *models.TypeTable {
	return models.ILP64()
}
