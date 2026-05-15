package AST

type Instruction interface {
	isInstruction()
}

type Operand interface {
	isOperand()
}

type AssemblyProgram struct {
	FunctionDefinition FunctionDefinition
}
type FunctionDefinition struct {
	Identifier   string
	Instructions []Instruction
}

type Immediate struct {
	Immediate int
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
func (*Register) isOperand()  {}

func (*InstructionMove) isInstruction() {}
func (*InstructionRet) isInstruction()  {}
