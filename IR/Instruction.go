package IR

type Instruction interface {
	isInstruction()
}

type Return struct {
	Value Value
}

type Unary struct {
	Operator    string
	Source      Value
	Destination *Variable
}

type Binary struct {
	Operator    string
	Left        Value
	Right       Value
	Destination *Variable
}

func NewReturnInstruction(value Value) Instruction {
	return &Return{
		Value: value,
	}
}

func NewUnaryInstruction(operator string, source Value, destination *Variable) Instruction {
	return &Unary{
		Operator:    operator,
		Source:      source,
		Destination: destination,
	}
}

func NewBinaryInstruction(operator string, left Value, right Value, destination *Variable) Instruction {
	return &Binary{
		Operator:    operator,
		Left:        left,
		Right:       right,
		Destination: destination,
	}
}

func (*Return) isInstruction() {}
func (*Unary) isInstruction()  {}
func (*Binary) isInstruction() {}
