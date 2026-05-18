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

func (parser *Parser) parseAssignmentStatement() AST.Statement {
	tok := parser.peekNthToken(0)
	lhs := parser.parseLValue()
	parser.expect(Token.EQUALS)
	rhs := parser.parseExpression()
	parser.expect(Token.SEMI_COLON)
	/*
		if !parser.ctx.ParsingForIncrement {
			parser.expect(Token.SEMI_COLON)
		}
	*/

	return &AST.StatementExpression{
		Tok: tok,
		Expr: &AST.ExpressionAssignment{
			LHSIdentifierToken: tok,
			LHS:                lhs,
			RHS:                rhs,
		},
	}
}

func (parser *Parser) parseCompoundAssignmentStatement(compoundAssignmentToken Token.TokenType) AST.Statement {
	tok := parser.peekNthToken(0)

	lhs := parser.parseExpression()
	operator := parser.expect(compoundAssignmentToken)
	rhs := parser.parseExpression()
	parser.expect(Token.SEMI_COLON)

	return &AST.StatementCompoundAssignment{
		Operator:           operator,
		LHSIdentifierToken: tok,
		LHS:                lhs,
		RHS:                rhs,
	}
}

func (parser *Parser) parseStatement() AST.Statement {
	current := parser.peekNthToken(0)
	next := parser.peekNthToken(1)

	if current.Kind == Token.LEFT_CURLY {
		return parser.parseStatementBlock()
	} else if current.Kind == Token.RETURN {
		return parser.parseStatementReturn()
	} else if current.Kind == Token.IDENTIFIER && next.Kind == Token.EQUALS {
		return parser.parseAssignmentStatement()
	} else if current.Kind == Token.IDENTIFIER && Token.BinaryOperationFromCompoundAssignment(next.Lexeme) != "" {
		return parser.parseCompoundAssignmentStatement(next.Kind)
	}

	parser.reportError("INVALID STATEMENT")
	return nil
}
