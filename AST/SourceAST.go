package AST

import "ion/go/v2/Token"

type SourceNode interface {
	isSourceNode()
}

type Deferrable interface {
	SourceNode
	isDeferrable()
}

// ------------------------------------------------

type Expression interface {
	SourceNode
	isExpression()
}

type ExpressionInteger struct {
	Value int
}

type ExpressionFloat struct {
	Value float32
}

type ExpressionString struct {
	Value string
}

type ExpressionBoolean struct {
	Value bool
}

type ExpressionGrouping struct {
	Expr Expression
}

// ------------------------------------------------

type Statement interface {
	SourceNode
	isStatement()
}

type StatementBlock struct {
	Body []SourceNode
}

type StatementReturn struct {
	Tok  Token.Token
	Expr Expression
}

// ------------------------------------------------

type Declaration interface {
	SourceNode
	isDeclaration()
}

type DeclarationFunction struct {
	Tok   Token.Token
	Block *StatementBlock
}

type Program struct {
	Declarations []Declaration
}

func (p *Program) isNode() {}
