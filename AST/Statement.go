package AST

import "ion/go/v2/Token"

type Statement interface {
	Node
	isStatement()
}

type StatementBlock struct {
	Body []Node
}

type StatementReturn struct {
	Tok  Token.Token
	Expr Expression
}

type StatementCompoundAssignment struct {
	Operator Token.Token // operator
	LHS      Expression
	RHS      Expression
}

type StatementExpression struct {
	Tok  Token.Token
	Expr Expression
}

type StatementNull struct {
	Expr Expression
}

func (*StatementBlock) isNode()      {}
func (*StatementBlock) isStatement() {}

func (*StatementReturn) isNode()      {}
func (*StatementReturn) isStatement() {}

func (*StatementCompoundAssignment) isNode()      {}
func (*StatementCompoundAssignment) isStatement() {}

func (*StatementExpression) isNode()      {}
func (*StatementExpression) isStatement() {}

func (*StatementNull) isNode()      {}
func (*StatementNull) isStatement() {}
