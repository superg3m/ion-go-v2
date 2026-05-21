package AST

import (
	"ion/go/v2/TS"
	"ion/go/v2/Token"
)

type Declaration interface {
	Node
	isDeclaration()
}

type DeclarationFunctionPrototype struct {
	DeclType TS.Type
	Tok      Token.Token
}

type DeclarationFunction struct {
	DeclType TS.Type
	Tok      Token.Token
	Block    *StatementBlock
}

type DeclarationVariable struct {
	DeclType TS.Type
	Tok      Token.Token
	RHS      Expression
}

func (*DeclarationFunctionPrototype) isNode()        {}
func (*DeclarationFunctionPrototype) isDeclaration() {}

func (*DeclarationFunction) isNode()        {}
func (*DeclarationFunction) isDeclaration() {}

func (*DeclarationVariable) isNode()        {}
func (*DeclarationVariable) isDeclaration() {}
