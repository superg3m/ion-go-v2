package AST

import "fmt"

type Instruction interface {
	isInstruction()
}

type Operand interface {
	isOperand()
	ToString() string
}

type AssemblyProgram struct {
	FunctionDefinition *FunctionDefinition
}
type FunctionDefinition struct {
	Identifier   string
	Instructions []Instruction
}

type Immediate struct {
	Value int
}

type Register struct {
	Name string // for now this will only be %eax
}

type InstructionMove struct {
	Destination Operand
	Source      Operand
}

type InstructionReturn struct{}

func (*Immediate) isOperand() {}
func (i *Immediate) ToString() string {
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

func NewReturnInstruction() Instruction {
	return &InstructionReturn{}
}
