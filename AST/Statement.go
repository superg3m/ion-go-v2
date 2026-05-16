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

func (*StatementBlock) isNode()      {}
func (*StatementBlock) isStatement() {}

func (*StatementReturn) isNode()      {}
func (*StatementReturn) isStatement() {}
