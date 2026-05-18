package AssemblyAST

import (
	"ion/go/v2/TS"
	"ion/go/v2/Token"
)

type Definition interface {
	isDefinition()
}

type Program struct {
	Definitions []Definition
}

type FunctionDefinition struct {
	DeclType     TS.Type
	Tok          Token.Token
	Instructions []Instruction
}

type VariableDefinition struct {
	DeclType     TS.Type
	Tok          Token.Token
	Instructions []Instruction
}

func (*FunctionDefinition) isDefinition() {}
func (*VariableDefinition) isDefinition() {}
