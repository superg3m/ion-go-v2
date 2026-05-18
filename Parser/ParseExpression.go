package Parser

import (
	_ "fmt"
	"ion/go/v2/AST"
	"ion/go/v2/Token"
	"strconv"
)

// <Primary>    ::= <integer> | <float> | <boolean> | <string> | '(' <Expression> ')'
func (parser *Parser) parsePrimary() AST.Expression {
	current := parser.peekNthToken(0)
	if parser.consumeOnMatch(Token.INTEGER_LITERAL) {
		num, _ := strconv.Atoi(current.Lexeme)
		return &AST.ExpressionInteger{Value: num}
	} else if parser.consumeOnMatch(Token.BOOLEAN_LITERAL) {
		b := current.Lexeme == "true"
		return &AST.ExpressionBoolean{Value: b}
	} else if parser.consumeOnMatch(Token.FLOAT_LITERAL) {
		num, _ := strconv.ParseFloat(current.Lexeme, 32)
		return &AST.ExpressionFloat{Value: float32(num)}
	} else if parser.consumeOnMatch(Token.STRING_LITERAL) {
		return &AST.ExpressionString{Value: current.Lexeme[1 : len(current.Lexeme)-1]}
	} else if parser.consumeOnMatch(Token.IDENTIFIER) {
		// next := parser.peekNthToken(0)

		return &AST.ExpressionVariable{
			Tok:      current,
			DeclType: nil, // NOTE(Jovanni): The type checker will patch this
		}
	} else if parser.consumeOnMatch(Token.LEFT_PAREN) {
		expr := parser.parseExpression()
		if expr != nil {
			parser.expect(Token.RIGHT_PAREN)
			return &AST.ExpressionGrouping{
				Expr: expr,
			}
		}
	}

	return nil
}

// <Unary>      ::= ('+'|'-'|'!') <unary> | <Primary>
func (parser *Parser) parseUnaryExpression() AST.Expression {
	ret := &AST.ExpressionUnary{}

	if parser.consumeOnMatch(Token.NOT) || parser.consumeOnMatch(Token.BITWISE_NOT) || parser.consumeOnMatch(Token.MINUS) || parser.consumeOnMatch(Token.PLUS) {
		ret.Operator = parser.previousToken()
		ret.Operand = parser.parseUnaryExpression()

		return ret
	}

	return parser.parsePrimary()
}

// <multiplicative>     ::= <Unary> (('*'|'/'|'%') <Unary>)*
func (parser *Parser) parseMultiplicativeExpression() AST.Expression {
	expr := parser.parseUnaryExpression()

	for parser.consumeOnMatch(Token.STAR) || parser.consumeOnMatch(Token.DIVISION) || parser.consumeOnMatch(Token.MODULUS) {
		op := parser.previousToken()
		right := parser.parseUnaryExpression()
		expr = &AST.ExpressionBinary{
			Operator: op,
			Left:     expr,
			Right:    right,
		}
	}

	return expr
}

// <additive>       ::= <Factor> (('+'|'-') <Factor>)*
func (parser *Parser) parseAdditiveExpression() AST.Expression {
	expr := parser.parseMultiplicativeExpression()

	for parser.consumeOnMatch(Token.PLUS) || parser.consumeOnMatch(Token.MINUS) {
		op := parser.previousToken()
		right := parser.parseMultiplicativeExpression()
		expr = &AST.ExpressionBinary{
			Operator: op,
			Left:     expr,
			Right:    right,
		}
	}

	return expr
}

// <bitwise_and_or>       ::= <additive> (('&'|'|'|'<<'|'>>') <additive>)*
func (parser *Parser) parseBitwiseExpression() AST.Expression {
	expr := parser.parseAdditiveExpression()

	for parser.consumeOnMatch(Token.BITWISE_AND) ||
		parser.consumeOnMatch(Token.BITWISE_OR) ||
		parser.consumeOnMatch(Token.BITWISE_LS) ||
		parser.consumeOnMatch(Token.BITWISE_RS) ||
		parser.consumeOnMatch(Token.BITWISE_XOR) {
		op := parser.previousToken()
		right := parser.parseAdditiveExpression()
		expr = &AST.ExpressionBinary{
			Operator: op,
			Left:     expr,
			Right:    right,
		}
	}

	return expr
}

// <comparison> ::= <bitwise_and_or> ((<'|'<='|'>='|'>'} <bitwise_and_or>)*
func (parser *Parser) parseComparisonExpression() AST.Expression {
	expr := parser.parseBitwiseExpression()

	for parser.consumeOnMatch(Token.LESS_THAN) ||
		parser.consumeOnMatch(Token.LESS_THAN_EQUALS) ||
		parser.consumeOnMatch(Token.GREATER_THAN_EQUALS) ||
		parser.consumeOnMatch(Token.GREATER_THAN) {
		op := parser.previousToken()
		right := parser.parseBitwiseExpression()
		expr = &AST.ExpressionBinary{
			Operator: op,
			Left:     expr,
			Right:    right,
		}
	}

	return expr
}

// <equality> ::= <comparison> (('=='|'!=') <comparison>)*
func (parser *Parser) parseEqualityExpression() AST.Expression {
	expr := parser.parseComparisonExpression()

	for parser.consumeOnMatch(Token.EQUALS_EQUALS) || parser.consumeOnMatch(Token.NOT_EQUALS) {
		op := parser.previousToken()
		right := parser.parseComparisonExpression()
		expr = &AST.ExpressionBinary{
			Operator: op,
			Left:     expr,
			Right:    right,
		}
	}

	return expr
}

// <logical_and> ::= <equality> ('&&' <equality>)*
func (parser *Parser) parseLogicalAndExpression() AST.Expression {
	expr := parser.parseEqualityExpression()

	for parser.consumeOnMatch(Token.LOGICAL_AND) {
		op := parser.previousToken()
		right := parser.parseEqualityExpression()
		expr = &AST.ExpressionBinary{
			Operator: op,
			Left:     expr,
			Right:    right,
		}
	}

	return expr
}

// <logical_or> ::= <logical_and> ('||' <logical_and>)*
func (parser *Parser) parseLogicalOrExpression() AST.Expression {
	expr := parser.parseLogicalAndExpression()

	for parser.consumeOnMatch(Token.LOGICAL_OR) {
		op := parser.previousToken()
		right := parser.parseLogicalAndExpression()
		expr = &AST.ExpressionBinary{
			Operator: op,
			Left:     expr,
			Right:    right,
		}
	}

	return expr
}

// <Expression> ::= <logical_or>
func (parser *Parser) parseExpression() AST.Expression {
	// current := parser.peekNthToken(0)
	// next := parser.peekNthToken(1)
	// next2 := parser.peekNthToken(2)

	return parser.parseLogicalOrExpression()
}
