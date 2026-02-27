package compiler

import (
	"cbc/models"
	"cbc/utils"
)

type IRGenerator struct {
}

func NewIRGenerator(typeTable models.TypeTable, errorHandler utils.ErrorHandler) *IRGenerator {
	return nil
}

func (this *IRGenerator) Generate(sem *models.AST) *models.IR {
	return nil
}
