package AST

import "ion/go/v2/Token"

type Node interface {
	isNode()
}

type Deferrable interface {
	Node
	isDeferrable()
}

// ------------------------------------------------

type Expression interface {
	Node
	isExpression()
}

type ExpressionBoolean struct {
	Value bool
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

// ------------------------------------------------

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

// ------------------------------------------------

type Declaration interface {
	Node
	isDeclaration()
}

type DeclarationFunction struct {
	Tok   Token.Token
	Block *StatementBlock
}

func (*DeclarationFunction) isNode()        {}
func (*DeclarationFunction) isDeclaration() {}

type Program struct {
	Declarations []Declaration
}

func (p *Program) isNode() {}
