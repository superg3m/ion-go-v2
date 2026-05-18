package AST

import (
	"ion/go/v2/TS"
	"ion/go/v2/Token"
)

type Expression interface {
	Node
	isExpression()
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
	Identifier Token.Token
	DeclType   TS.Type
}

type ExpressionAssignment struct {
	Tok Token.Token
	LHS Expression
	RHS Expression
}

func (*ExpressionBoolean) isNode()       {}
func (*ExpressionBoolean) isExpression() {}

func (*ExpressionInteger) isNode()       {}
func (*ExpressionInteger) isExpression() {}

func (*ExpressionFloat) isNode()       {}
func (*ExpressionFloat) isExpression() {}

func (*ExpressionString) isNode()       {}
func (*ExpressionString) isExpression() {}

func (*ExpressionGrouping) isNode()       {}
func (*ExpressionGrouping) isExpression() {}

func (*ExpressionUnary) isNode()       {}
func (*ExpressionUnary) isExpression() {}

func (*ExpressionBinary) isNode()       {}
func (*ExpressionBinary) isExpression() {}

func (*ExpressionVariable) isNode()       {}
func (*ExpressionVariable) isExpression() {}

func (*ExpressionAssignment) isNode()       {}
func (*ExpressionAssignment) isExpression() {}
