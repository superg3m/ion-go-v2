package IR

import "ion/go/v2/Token"

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

type Copy struct {
	Source      Value
	Destination Value
}

type Jump struct {
	TargetLabel string
}

type ConditionalJump struct {
	TargetLabel string
	Condition   Value
	IfZero      bool
	IfNotZero   bool
}

type Label struct {
	Identifier string
}

type FunctionCall struct {
	Identifier  Token.Token
	Arguments   []Value
	Destination Value
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

func NewCopyInstruction(source Value, destination Value) Instruction {
	return &Copy{
		Source:      source,
		Destination: destination,
	}
}

func NewJumpInstruction(targetLabel string) Instruction {
	return &Jump{
		TargetLabel: targetLabel,
	}
}

func NewConditionalJumpInstruction(targetLabel string, condition Value, ifZero bool, ifNotZero bool) Instruction {
	return &ConditionalJump{
		TargetLabel: targetLabel,
		Condition:   condition,
		IfZero:      ifZero,
		IfNotZero:   ifNotZero,
	}
}

func NewLabelInstruction(identifier string) Instruction {
	return &Label{
		Identifier: identifier,
	}
}

func NewFunctionCallInstruction(identifier Token.Token, arguments []Value, destination Value) Instruction {
	return &FunctionCall{
		Identifier:  identifier,
		Arguments:   arguments,
		Destination: destination,
	}
}

func (*Return) isInstruction()          {}
func (*Unary) isInstruction()           {}
func (*Binary) isInstruction()          {}
func (*Copy) isInstruction()            {}
func (*Jump) isInstruction()            {}
func (*ConditionalJump) isInstruction() {}
func (*Label) isInstruction()           {}
func (*FunctionCall) isInstruction()    {}
