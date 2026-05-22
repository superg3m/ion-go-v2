package IR

import (
	"ion/go/v2/AssemblyAST"
	"ion/go/v2/TS"
	"ion/go/v2/Token"
)

type Value interface {
	isValue()
	GetToken() Token.Token
	GetDeclType() TS.Type
}

type Constant struct {
	Value int
	Token Token.Token
}

type Variable struct {
	Tok      Token.Token
	DeclType TS.Type
}

type ParameterVariable struct {
	Tok         Token.Token
	Register    *AssemblyAST.Register // either this or the StackOffset is used
	StackOffset int                   // either this or the Register is used
	DeclType    TS.Type
}

func (*Constant) isValue()          {}
func (*Variable) isValue()          {}
func (*ParameterVariable) isValue() {}

func (c *Constant) GetDeclType() TS.Type {
	return TS.NewTypeInteger(true, 4)
}

func (v *Variable) GetDeclType() TS.Type {
	return v.DeclType
}

func (v *ParameterVariable) GetDeclType() TS.Type {
	return v.DeclType
}

func (c *Constant) GetToken() Token.Token {
	return c.Token
}

func (v *Variable) GetToken() Token.Token {
	return v.Tok
}

func (v *ParameterVariable) GetToken() Token.Token {
	return v.Tok
}

func NewConstantValue(value int, token Token.Token) Value {
	return &Constant{
		Token: token,
		Value: value,
	}
}

func NewVariable(token Token.Token, declType TS.Type) *Variable {
	return &Variable{
		Tok:      token,
		DeclType: declType,
	}
}
