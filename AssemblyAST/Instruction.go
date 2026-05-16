package AssemblyAST

type Instruction interface {
	isInstruction()
}

type InstructionMove struct {
	Source      Operand
	Destination Operand
}

type InstructionReturn struct{}

type InstructionStackAllocate struct {
	AllocationSize int
}

type InstructionUnary struct {
	Operator string
	Operand  Operand
}

func NewMoveInstruction(source Operand, destination Operand) Instruction {
	return &InstructionMove{
		Destination: destination,
		Source:      source,
	}
}

func NewReturnInstruction() Instruction {
	return &InstructionReturn{}
}

func NewStackAllocateInstruction(allocationSize int) Instruction {
	return &InstructionStackAllocate{
		AllocationSize: allocationSize,
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
