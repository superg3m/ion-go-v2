package AST

import "ion/go/v2/Token"

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
