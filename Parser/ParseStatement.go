package Parser

import (
	"ion/go/v2/AST"
	"ion/go/v2/Token"
)

func (parser *Parser) parseStatementReturn() AST.Statement {
	tok := parser.expect(Token.RETURN)
	expr := parser.parseExpression()
	parser.expect(Token.SEMI_COLON)

	return &AST.StatementReturn{
		Tok:  tok,
		Expr: expr,
	}
}

func (parser *Parser) parseStatementBlock() AST.Statement {
	var body []AST.Node
	parser.expect(Token.LEFT_CURLY)
	for !parser.consumeOnMatch(Token.RIGHT_CURLY) {
		if decl := parser.parseDeclaration(); decl != nil {
			body = append(body, decl)
			continue
		}

		if stmt := parser.parseStatement(); stmt != nil {
			body = append(body, stmt)
			continue
		}
	}

	return &AST.StatementBlock{
		Body: body,
	}
}

func (parser *Parser) parseStatement() AST.Statement {
	current := parser.peekNthToken(0)

	if current.Kind == Token.LEFT_CURLY {
		return parser.parseStatementBlock()
	} else if current.Kind == Token.RETURN {
		return parser.parseStatementReturn()
	}

	panic("INVALID STATEMENT!")
	return nil
}
