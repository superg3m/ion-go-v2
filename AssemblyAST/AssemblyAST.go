package AssemblyAST

import "fmt"

type Instruction interface {
	isInstruction()
}

type Operand interface {
	isOperand()
	ToString() string
}

type Program struct {
	FunctionDefinition *FunctionDefinition
}
type FunctionDefinition struct {
	Identifier   string
	Instructions []Instruction
}

type IntegerConstant struct {
	Value int
}

type Register struct {
	Name string // for now this will only be %eax
}

type InstructionMove struct {
	Destination Operand
	Source      Operand
}

type InstructionReturn struct {
	Value *IntegerConstant
}

func (*IntegerConstant) isOperand() {}
func (i *IntegerConstant) ToString() string {
	return fmt.Sprintf("$%d", i.Value)
}

func (*Register) isOperand() {}
func (r *Register) ToString() string {
	return r.Name
}

func (*InstructionMove) isInstruction()   {}
func (*InstructionReturn) isInstruction() {}

func NewMoveInstruction(destination Operand, source Operand) Instruction {
	return &InstructionMove{
		Destination: destination,
		Source:      source,
	}
}

func NewReturnInstruction(value *IntegerConstant) Instruction {
	return &InstructionReturn{
		Value: value,
	}
}
