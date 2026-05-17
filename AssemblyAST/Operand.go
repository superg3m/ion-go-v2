package AssemblyAST

import (
	"fmt"
	"ion/go/v2/IR"
)

type Operand interface {
	isOperand()
	ToString() string
}

type Immediate struct {
	Value int
}

type Register string

const (
	EAX  = Register("%eax")
	R10D = "%r10d"
	R11D = "%r11d"
	EDX  = "%edx"
)

type Pseudo struct {
	Identifier string
}

type Stack struct {
	Offset int
}

func NewOperand(value IR.Value) Operand {
	switch v := value.(type) {
	case *IR.Constant:
		return &Immediate{Value: v.Value}
	case *IR.Variable:
		return &Pseudo{Identifier: v.Name}
	default:
		panic(fmt.Sprintf("Unknown instruction %T", v))
	}

	return nil
}

func NewRegisterOperand(register Register) Operand {
	ret := new(Register)
	*ret = register

	return ret
}

func NewStackOperand(offset int) Operand {
	return &Stack{Offset: offset}
}

func NewImmediateOperand(value int) Operand {
	return &Immediate{Value: value}
}

func (*Immediate) isOperand() {}
func (i *Immediate) ToString() string {
	return fmt.Sprintf("$%d", i.Value)
}

func (*Register) isOperand() {}
func (r *Register) ToString() string {
	return string(*r)
}

func (*Pseudo) isOperand() {}
func (r *Pseudo) ToString() string {
	return "PSEUDO"
}

func (*Stack) isOperand() {}
func (r *Stack) ToString() string {
	return fmt.Sprintf("-%d(%%rbp)", r.Offset)
}
