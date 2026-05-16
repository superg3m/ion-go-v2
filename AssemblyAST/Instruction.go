package AssemblyAST

type Instruction interface {
	isInstruction()
}

type InstructionMove struct {
	Destination Operand
	Source      Operand
}

type InstructionReturn struct {
	Value *Immediate
}

type InstructionStackAllocate struct {
	Value int
}

type InstructionUnary struct {
	operator string
	operand  Operand
}

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

func NewUnaryInstruction(operator string, operand Operand) Instruction {
	return &InstructionUnary{
		operator,
		operand,
	}
}

func (*InstructionMove) isInstruction()          {}
func (*InstructionReturn) isInstruction()        {}
func (*InstructionStackAllocate) isInstruction() {}
func (*InstructionUnary) isInstruction()         {}
