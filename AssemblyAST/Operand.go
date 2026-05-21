package AssemblyAST

import (
	"fmt"
	"ion/go/v2/TS"
	"ion/go/v2/Token"
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
	INVALID Register = iota
	RAX     Register = iota
	RBX
	RCX
	RDX
	RSI
	RDI
	RBP
	RSP
	R8
	R9
	R10
	R11
	R12
	R13
	R14
	R15
)

var ArgumentRegisters = []Register{RDI, RSI, RDX, RCX, R8, R9}

type RegisterNames struct {
	X64 string
	X32 string
	X16 string
	X8  string
}

var RegisterLookup = map[Register]RegisterNames{
	RAX: {"%rax", "%eax", "%ax", "%al"},
	RBX: {"%rbx", "%ebx", "%bx", "%bl"},
	RCX: {"%rcx", "%ecx", "%cx", "%cl"},
	RDX: {"%rdx", "%edx", "%dx", "%dl"},
	RSI: {"%rsi", "%esi", "%si", "%sil"},
	RDI: {"%rdi", "%edi", "%di", "%dil"},
	RBP: {"%rbp", "%ebp", "%bp", "%bpl"},
	RSP: {"%rsp", "%esp", "%sp", "%spl"},

	R8:  {"%r8", "%r8d", "%r8w", "%r8b"},
	R9:  {"%r9", "%r9d", "%r9w", "%r9b"},
	R10: {"%r10", "%r10d", "%r10w", "%r10b"},
	R11: {"%r11", "%r11d", "%r11w", "%r11b"},
	R12: {"%r12", "%r12d", "%r12w", "%r12b"},
	R13: {"%r13", "%r13d", "%r13w", "%r13b"},
	R14: {"%r14", "%r14d", "%r14w", "%r14b"},
	R15: {"%r15", "%r15d", "%r15w", "%r15b"},
}

type Pseudo struct {
	Tok      Token.Token
	DeclType TS.Type
}

type Stack struct {
	Offset int
}

type Parameter struct {
	Tok         Token.Token
	DeclType    TS.Type
	Register    Register
	StackOffset int
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
	return r.X32BitName()
}

func (r *Register) X64BitName() string {
	return RegisterLookup[*r].X64
}

func (r *Register) X32BitName() string {
	return RegisterLookup[*r].X32
}

func (r *Register) X16BitName() string {
	return RegisterLookup[*r].X16
}

func (r *Register) X8BitName() string {
	return RegisterLookup[*r].X8
}

func (*Pseudo) isOperand() {}
func (r *Pseudo) ToString() string {
	return "PSEUDO"
}

func (*Stack) isOperand() {}
func (r *Stack) ToString() string {
	return fmt.Sprintf("-%d(%%rbp)", r.Offset)
}

func (*Parameter) isOperand() {}
func (p *Parameter) ToString() string {
	if p.Register == INVALID {
		return fmt.Sprintf("-%d(%%rbp)", p.StackOffset)
	}

	return p.Register.ToString()
}
