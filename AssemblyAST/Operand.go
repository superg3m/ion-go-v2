package AssemblyAST

import "fmt"

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
	identifier string
}

func NewImmediateOperand(immediate int) Operand {
	return &Immediate{Value: immediate}
}

func NewRegisterOperand(register Register) Operand {
	ret := new(Register)
	*ret = register

	return ret
}

func NewPseudoOperand(identifier string) Operand {
	return &Pseudo{identifier}
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
