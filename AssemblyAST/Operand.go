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

type Register int

const (
	AX = Register(iota)
	R10
)

type Pseudo struct {
	Identifier string
}

func NewOperand(value IR.Value) Operand {
	switch v := value.(type) {
	case *IR.Constant:
		return &Immediate{Value: v.Value}
	case *IR.Variable:
		return &Pseudo{Identifier: v.Name}
	}

	panic("Not a Operand")
	return nil
}

func NewRegisterOperand(register Register) Operand {
	ret := new(Register)
	*ret = register

	return ret
}

func (*Immediate) isOperand() {}
func (i *Immediate) ToString() string {
	return fmt.Sprintf("$%d", i.Value)
}

func (*Register) isOperand() {}
func (r *Register) ToString() string {
	return "r.Name"
}

func (*Pseudo) isOperand() {}
func (r *Pseudo) ToString() string {
	return "r.Name"
}
