package models

type Op int

const (
	OpADD Op = iota
	OpSUB
	OpMUL
	OpSDIV
	OpUDIV
	OpSMOD
	OpUMOD
	OpBITAND
	OpBITOR
	OpBITXOR
	OpBITLSHIFT
	OpBITRSHIFT
	OpARITHRSHIFT

	OpEQ
	OpNEQ
	OpSGT
	OpSGTEQ
	OpSLT
	OpSLTEQ
	OpUGT
	OpUGTEQ
	OpULT
	OpULTEQ

	OpUMINUS
	OpBITNOT
	OpNOT

	OpSCAST
	OpUCAST
)

func InternBinary(op string, isSigned bool) Op {
	switch op {
	case "+":
		return OpADD
	case "-":
		return OpSUB
	case "*":
		return OpMUL
	case "/":
		if isSigned {
			return OpSDIV
		}
		return OpUDIV
	case "%":
		if isSigned {
			return OpSMOD
		}
		return OpUMOD
	case "&":
		return OpBITAND
	case "|":
		return OpBITOR
	case "^":
		return OpBITXOR
	case "<<":
		return OpBITLSHIFT
	case ">>":
		if isSigned {
			return OpARITHRSHIFT
		}
		return OpBITRSHIFT
	case "==":
		return OpEQ
	case "!=":
		return OpNEQ
	case "<":
		if isSigned {
			return OpSLT
		}
		return OpULT
	case "<=":
		if isSigned {
			return OpSLTEQ
		}
		return OpULTEQ
	case ">":
		if isSigned {
			return OpSGT
		}
		return OpUGT
	case ">=":
		if isSigned {
			return OpSGTEQ
		}
		return OpUGTEQ
	default:
		panic("unknown binary op: " + op)
	}
}

func InternUnary(op string) Op {
	switch op {
	case "-":
		return OpUMINUS
	case "~":
		return OpBITNOT
	case "!":
		return OpNOT
	case "+":
		panic("unary+ should not be in IR")
	default:
		panic("unknown unary op: " + op)
	}
}
