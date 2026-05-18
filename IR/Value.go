package IR

import (
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

func (*Constant) isValue() {}
func (*Variable) isValue() {}

func (c *Constant) GetDeclType() TS.Type {
	return TS.NewTypeInteger(true, 4)
}

func (v *Variable) GetDeclType() TS.Type {
	return v.DeclType
}

func (c *Constant) GetToken() Token.Token {
	return c.Token
}

func (v *Variable) GetToken() Token.Token {
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
