package sysdep

import (
	"cbc/models"
)

// TODO:
type IPlatform interface {
	TypeTable() *models.TypeTable
	// CodeGenerator(*CodeGeneratorOptions, *utils.ErrorHandler) CodeGenerator
	// Assembler(*utils.ErrorHandler) Assembler
	// linker(*utils.ErrorHandler) Linker
}
