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

type InstructionRet struct{}

func (*Immediate) isOperand() {}
func (i *Immediate) ToString() string {
	return fmt.Sprintf("$%d", i.Value)
}
func (*Register) isOperand() {}
func (r *Register) ToString() string {
	return r.Name
}

func (*InstructionMove) isInstruction() {}
func (*InstructionRet) isInstruction()  {}
