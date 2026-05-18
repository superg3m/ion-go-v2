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

type InstructionBinary struct {
	Operator string
	Left     Operand
	Right    Operand
}

type InstructionCDQ struct{}
type InstructionDivide struct {
	Left Operand
}

type InstructionCompare struct {
	Left  Operand
	Right Operand
}
type InstructionJump struct {
	TargetLabel string
}

type ConditionalCode string

const (
	EQUALS                 ConditionalCode = "=="
	NOT_EQUALS                             = "!="
	LESS_THAN                              = "<"
	LESS_THAN_OR_EQUALS                    = "<="
	GREATER_THAN                           = ">"
	GREATER_THAN_OR_EQUALS                 = ">="
)

func (c ConditionalCode) ToString() string {
	switch c {
	case EQUALS:
		return "e"
	case NOT_EQUALS:
		return "ne"
	case LESS_THAN:
		return "l"
	case LESS_THAN_OR_EQUALS:
		return "le"
	case GREATER_THAN:
		return "g"
	case GREATER_THAN_OR_EQUALS:
		return "ge"
	default:
		panic("unknown conditional")
	}
}

type InstructionConditionalJump struct {
	TargetLabel string
	Code        ConditionalCode
}

type InstructionSetConditionalCode struct {
	Destination Operand
	Code        ConditionalCode
}

type InstructionLabel struct {
	Identifier string
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

func NewCompareInstruction(left, right Operand) Instruction {
	return &InstructionCompare{
		Left:  left,
		Right: right,
	}
}

func NewJumpInstruction(targetLabel string) Instruction {
	return &InstructionJump{
		TargetLabel: targetLabel,
	}
}

func NewConditionalJumpInstruction(targetLabel string, code ConditionalCode) Instruction {
	return &InstructionConditionalJump{
		TargetLabel: targetLabel,
		Code:        code,
	}
}

func NewSetConditionalCodeInstruction(destination Operand, code ConditionalCode) Instruction {
	return &InstructionSetConditionalCode{
		Destination: destination,
		Code:        code,
	}
}

func NewLabelInstruction(identifier string) Instruction {
	return &InstructionLabel{
		Identifier: identifier,
	}
}

func NewUnaryInstruction(operator string, operand Operand) Instruction {
	return &InstructionUnary{
		operator,
		operand,
	}
}

func NewBinaryInstruction(operator string, left, right Operand) Instruction {
	return &InstructionBinary{
		operator,
		left,
		right,
	}
}

func NewCDQInstruction() Instruction {
	return &InstructionCDQ{}
}

func NewDivideInstruction(left Operand) Instruction {
	return &InstructionDivide{
		left,
	}
}

func (*InstructionMove) isInstruction()          {}
func (*InstructionReturn) isInstruction()        {}
func (*InstructionStackAllocate) isInstruction() {}
func (*InstructionUnary) isInstruction()         {}
func (*InstructionBinary) isInstruction()        {}
func (*InstructionCDQ) isInstruction()           {}
func (*InstructionDivide) isInstruction()        {}

func (*InstructionCompare) isInstruction()            {}
func (*InstructionJump) isInstruction()               {}
func (*InstructionConditionalJump) isInstruction()    {}
func (*InstructionSetConditionalCode) isInstruction() {}
func (*InstructionLabel) isInstruction()              {}
