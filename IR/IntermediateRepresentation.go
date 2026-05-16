package IR

import "ion/go/v2/AST"

type Program struct {
	FunctionDefinition FunctionDefinition
}

type FunctionDefinition struct {
	identifier   string
	instructions []Instruction
}

type Instruction interface {
	isInstruction()
}

type Value interface {
	isValue()
}

type Unary struct {
	operator    string
	source      Value
	destination Value
}

type Integer struct {
	Value int
}

type Variable struct {
	identifier string
}

func (*Integer) isValue()  {}
func (*Variable) isValue() {}

type Return struct {
	Value Value
}

func (*Unary) isInstruction()  {}
func (*Return) isInstruction() {}

func EmitIntermediateRepresentation(program AST.Program) Program {
	return Program{}
}
