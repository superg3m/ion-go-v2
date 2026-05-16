package IR

type Instruction interface {
	isInstruction()
}

type Unary struct {
	Operator    string
	Source      Value
	Destination Value
}

type Return struct {
	Value Value
}

func NewUnaryInstruction(operator string, source Value, destination Value) Instruction {
	return &Unary{
		Operator:    operator,
		Source:      source,
		Destination: destination,
	}
}

func NewReturnInstruction(value Value) Instruction {
	return &Return{
		Value: value,
	}
}

func (*Unary) isInstruction()  {}
func (*Return) isInstruction() {}
