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

type Immediate struct {
	Value int
}

type Register struct {
	Name string // for now this will only be %eax
}

type InstructionStackAllocate struct {
	Value int
}

type InstructionMove struct {
	Destination Operand
	Source      Operand
}

type InstructionReturn struct {
	Value *Immediate
}

func (*Immediate) isOperand() {}
func (i *Immediate) ToString() string {
	return fmt.Sprintf("$%d", i.Value)
}

func (*Register) isOperand() {}
func (r *Register) ToString() string {
	return r.Name
}

type Pseudo struct {
	identifier string
}

func (*InstructionMove) isInstruction()          {}
func (*InstructionReturn) isInstruction()        {}
func (*InstructionStackAllocate) isInstruction() {}

func NewMoveInstruction(destination Operand, source Operand) Instruction {
	return &InstructionMove{
		Destination: destination,
		Source:      source,
	}
}

func NewReturnInstruction(value *Immediate) Instruction {
	return &InstructionReturn{
		Value: value,
	}
}

func NewStackAllocateInstruction(value int) Instruction {
	return &InstructionStackAllocate{
		Value: value,
	}
}
