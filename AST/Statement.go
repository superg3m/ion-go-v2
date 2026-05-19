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

type StatementExpression struct {
	Tok  Token.Token
	Expr Expression
}

type StatementIfElse struct {
	Condition Expression
	ThenBlock *StatementBlock
	ElseBlock *StatementBlock
}

type StatementFor struct {
	StartLoopLabel string
	EndLoopLabel   string

	Condition   Expression
	Initializer *DeclarationVariable
	Increment   *StatementExpression
	Block       *StatementBlock
}

type StatementWhile struct {
	StartLoopLabel string
	EndLoopLabel   string

	Condition Expression
	Block     *StatementBlock
}

type StatementContinue struct {
	StartLoopLabel string
}

type StatementBreak struct {
	EndLoopLabel string
}

type StatementNull struct {
	Expr Expression
}

func (*StatementBlock) isNode()      {}
func (*StatementBlock) isStatement() {}

func (*StatementReturn) isNode()      {}
func (*StatementReturn) isStatement() {}

func (*StatementExpression) isNode()      {}
func (*StatementExpression) isStatement() {}

func (*StatementNull) isNode()      {}
func (*StatementNull) isStatement() {}

func (*StatementIfElse) isNode()      {}
func (*StatementIfElse) isStatement() {}

func (*StatementFor) isNode()      {}
func (*StatementFor) isStatement() {}

func (*StatementWhile) isNode()      {}
func (*StatementWhile) isStatement() {}

func (*StatementContinue) isNode()      {}
func (*StatementContinue) isStatement() {}

func (*StatementBreak) isNode()      {}
func (*StatementBreak) isStatement() {}
