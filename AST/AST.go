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

func (*ExpressionBoolean) isSourceNode() {}
func (*ExpressionBoolean) isExpression() {}

func (*ExpressionInteger) isSourceNode() {}
func (*ExpressionInteger) isExpression() {}

func (*ExpressionFloat) isSourceNode() {}
func (*ExpressionFloat) isExpression() {}

func (*ExpressionString) isSourceNode() {}
func (*ExpressionString) isExpression() {}

func (*ExpressionGrouping) isSourceNode() {}
func (*ExpressionGrouping) isExpression() {}

func (*ExpressionUnary) isSourceNode() {}
func (*ExpressionUnary) isExpression() {}

func (*ExpressionBinary) isSourceNode() {}
func (*ExpressionBinary) isExpression() {}

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

func (*StatementBlock) isSourceNode() {}
func (*StatementBlock) isStatement()  {}

func (*StatementReturn) isSourceNode() {}
func (*StatementReturn) isStatement()  {}

// ------------------------------------------------

type Declaration interface {
	SourceNode
	isDeclaration()
}

type DeclarationFunction struct {
	Tok   Token.Token
	Block *StatementBlock
}

func (*DeclarationFunction) isSourceNode()  {}
func (*DeclarationFunction) isDeclaration() {}

type Program struct {
	Declarations []Declaration
}

func (p *Program) isNode() {}
