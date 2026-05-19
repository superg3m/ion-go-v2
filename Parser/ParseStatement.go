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

	if parser.consumeOnMatch(Token.LEFT_CURLY) {
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
	} else {
		if decl := parser.parseDeclaration(); decl != nil {
			body = append(body, decl)
		} else if stmt := parser.parseStatement(); stmt != nil {
			body = append(body, stmt)
		}
	}

	return &AST.StatementBlock{
		Body: body,
	}
}

func (parser *Parser) parseStatementExpression() AST.Statement {
	tok := parser.peekNthToken(0)
	expr := parser.parseExpression()
	if !parser.ctx.ParsingForIncrement {
		parser.expect(Token.SEMI_COLON)
	}

	return &AST.StatementExpression{
		Tok:  tok,
		Expr: expr,
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

func (parser *Parser) parseForStatement() AST.Statement {
	parser.expect(Token.FOR)
	parser.expect(Token.LEFT_PAREN)
	initializer := parser.parseVariableDeclaration()
	condition := parser.parseExpression()
	parser.expect(Token.SEMI_COLON)
	parser.ctx.ParsingForIncrement = true
	increment := parser.parseStatementExpression()
	parser.ctx.ParsingForIncrement = false
	parser.expect(Token.RIGHT_PAREN)
	block := parser.parseStatementBlock()

	return &AST.StatementFor{
		Initializer: initializer.(*AST.DeclarationVariable),
		Condition:   condition,
		Increment:   increment.(*AST.StatementExpression),
		Block:       block.(*AST.StatementBlock),
	}
}

func (parser *Parser) parseWhileStatement() AST.Statement {
	parser.expect(Token.WHILE)
	parser.expect(Token.LEFT_PAREN)
	condition := parser.parseExpression()
	parser.expect(Token.RIGHT_PAREN)
	block := parser.parseStatementBlock()

	return &AST.StatementWhile{
		Condition: condition,
		Block:     block.(*AST.StatementBlock),
	}
}

func (parser *Parser) parseIfElseStatement() AST.Statement {
	parser.expect(Token.IF)
	parser.expect(Token.LEFT_PAREN)
	condition := parser.parseExpression()
	parser.expect(Token.RIGHT_PAREN)
	ifBlock := parser.parseStatementBlock()

	var elseBlock *AST.StatementBlock = nil
	if parser.consumeOnMatch(Token.ELSE) {
		elseBlock = parser.parseStatementBlock().(*AST.StatementBlock)
	}

	return &AST.StatementIfElse{
		Condition: condition,
		ThenBlock: ifBlock.(*AST.StatementBlock),
		ElseBlock: elseBlock,
	}
}

func (parser *Parser) parseStatement() AST.Statement {
	current := parser.peekNthToken(0)
	next := parser.peekNthToken(1)

	if current.Kind == Token.LEFT_CURLY {
		return parser.parseStatementBlock()
	} else if current.Kind == Token.RETURN {
		return parser.parseStatementReturn()
	} else if current.Kind == Token.IF {
		return parser.parseIfElseStatement()
	} else if current.Kind == Token.FOR {
		return parser.parseForStatement()
	} else if current.Kind == Token.WHILE {
		return parser.parseWhileStatement()
	} else if parser.consumeOnMatch(Token.BREAK) {
		parser.expect(Token.SEMI_COLON)
		return &AST.StatementBreak{}
	} else if parser.consumeOnMatch(Token.CONTINUE) {
		parser.expect(Token.SEMI_COLON)
		return &AST.StatementContinue{}
	} else if current.Kind == Token.IDENTIFIER && next.Kind == Token.EQUALS {
		return parser.parseStatementExpression()
	} else if current.Kind == Token.IDENTIFIER && Token.BinaryOperationFromCompoundAssignment(next.Lexeme) != "" {
		return parser.parseCompoundAssignmentStatement(next.Kind)
	}

	parser.reportError("INVALID STATEMENT")
	return nil
}
