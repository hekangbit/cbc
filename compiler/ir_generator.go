package compiler

import (
	"cbc/models"
	"cbc/util"
)

type IRGenerator struct {
}

func NewIRGenerator(typeTable models.TypeTable, errorHandler util.ErrorHandler) *IRGenerator {
	return nil
}

func (this *IRGenerator) Generate(sem *models.AST) *models.IR {
	return nil
}
