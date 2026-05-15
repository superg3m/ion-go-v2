package Parser

import (
	"ion/go/v2/AST"
	"ion/go/v2/Token"
)

func (parser *Parser) parseFunctionDeclaration() AST.Declaration {
	parser.expect(Token.IDENTIFIER) // return_type
	ident := parser.expect(Token.IDENTIFIER)
	parser.expect(Token.LEFT_PAREN)
	parser.expect(Token.RIGHT_PAREN)
	block := parser.parseStatementBlock().(*AST.StatementBlock)

	return &AST.DeclarationFunction{
		Tok:   ident,
		Block: block,
	}
}

func (parser *Parser) parseDeclaration() AST.Declaration {
	current := parser.peekNthToken(0)

	if current.Kind == Token.IDENTIFIER {
		return parser.parseFunctionDeclaration()
	}

	return nil
}
