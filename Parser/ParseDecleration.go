package Parser

import (
	"ion/go/v2/AST"
	"ion/go/v2/TS"
	"ion/go/v2/Token"
)

func (parser *Parser) parseFunctionDeclaration() AST.Declaration {
	returnType := parser.parseType()
	ident := parser.expect(Token.IDENTIFIER)
	params := parser.parseParameters()
	block := parser.parseStatementBlock().(*AST.StatementBlock)

	declType := TS.NewTypeFunction(returnType, params)
	return &AST.DeclarationFunction{
		DeclType: declType,
		Tok:      ident,
		Block:    block,
	}
}

func (parser *Parser) parseVariableDeclaration() AST.Declaration {
	returnType := parser.parseType()
	ident := parser.expect(Token.IDENTIFIER)

	var rhs AST.Expression
	if parser.peekNthToken(0).Kind == Token.EQUALS {
		parser.expect(Token.EQUALS)
		rhs = parser.parseExpression()
	}
	parser.expect(Token.SEMI_COLON)

	return &AST.DeclarationVariable{
		DeclType: returnType,
		Tok:      ident,
		RHS:      rhs,
	}
}

func (parser *Parser) fakeParseDeclaration() bool {
	current := parser.peekNthToken(0)
	next := parser.peekNthToken(1)
	next2 := parser.peekNthToken(2)

	if current.Kind == Token.IDENTIFIER && next.Kind == Token.IDENTIFIER && (next2.Kind == Token.EQUALS || next2.Kind == Token.SEMI_COLON) {
		return true
	} else if current.Kind == Token.IDENTIFIER && next.Kind == Token.IDENTIFIER && next2.Kind == Token.LEFT_PAREN {
		return true
	}

	return false
}

func (parser *Parser) parseDeclaration() AST.Declaration {
	current := parser.peekNthToken(0)
	next := parser.peekNthToken(1)
	next2 := parser.peekNthToken(2)

	if current.Kind == Token.IDENTIFIER && next.Kind == Token.IDENTIFIER && (next2.Kind == Token.EQUALS || next2.Kind == Token.SEMI_COLON) {
		return parser.parseVariableDeclaration()
	} else if current.Kind == Token.IDENTIFIER && next.Kind == Token.IDENTIFIER && next2.Kind == Token.LEFT_PAREN {
		return parser.parseFunctionDeclaration()
	}

	return nil
}
