package AST

import (
	"ion/go/v2/TS"
	"ion/go/v2/Token"
)

type Expression interface {
	Node
	isExpression()
	GetDeclType() TS.Type
}

type ExpressionBoolean struct {
	Value bool
}

type ExpressionInteger struct {
	Tok   Token.Token
	Value int
}

type ExpressionFloat struct {
	Value float32
}

type ExpressionString struct {
	Value string
}

type ExpressionGrouping struct {
	Expr Expression
}

type ExpressionUnary struct {
	Operator Token.Token
	Operand  Expression
}

type ExpressionBinary struct {
	Operator Token.Token
	Left     Expression
	Right    Expression
}

type ExpressionVariable struct {
	Tok      Token.Token
	DeclType TS.Type
}

type ExpressionAssignment struct {
	LHSIdentifierToken Token.Token
	LHS                Expression
	RHS                Expression
}

func (*ExpressionBoolean) isNode()       {}
func (*ExpressionBoolean) isExpression() {}
func (*ExpressionBoolean) GetDeclType() TS.Type {
	return TS.NewTypeBool()
}

func (*ExpressionInteger) isNode()       {}
func (*ExpressionInteger) isExpression() {}
func (*ExpressionInteger) GetDeclType() TS.Type {
	return TS.NewTypeInteger(true, 4)
}

func (*ExpressionFloat) isNode()       {}
func (*ExpressionFloat) isExpression() {}
func (*ExpressionFloat) GetDeclType() TS.Type {
	return TS.NewTypeFloat(4)
}

func (*ExpressionString) isNode()       {}
func (*ExpressionString) isExpression() {}
func (*ExpressionString) GetDeclType() TS.Type {
	return TS.NewTypeString()
}

func (*ExpressionGrouping) isNode()       {}
func (*ExpressionGrouping) isExpression() {}
func (g *ExpressionGrouping) GetDeclType() TS.Type {
	return g.Expr.GetDeclType()
}

func (*ExpressionUnary) isNode()       {}
func (*ExpressionUnary) isExpression() {}
func (u *ExpressionUnary) GetDeclType() TS.Type {
	return u.Operand.GetDeclType()
}

func (*ExpressionBinary) isNode()       {}
func (*ExpressionBinary) isExpression() {}
func (b *ExpressionBinary) GetDeclType() TS.Type {
	return b.Left.GetDeclType()
}

func (*ExpressionVariable) isNode()       {}
func (*ExpressionVariable) isExpression() {}
func (v *ExpressionVariable) GetDeclType() TS.Type {
	return v.DeclType
}

func (*ExpressionAssignment) isNode()       {}
func (*ExpressionAssignment) isExpression() {}
func (a *ExpressionAssignment) GetDeclType() TS.Type {
	return a.LHS.GetDeclType()
}
