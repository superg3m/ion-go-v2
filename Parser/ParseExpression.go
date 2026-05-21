package Parser

import (
	_ "fmt"
	"ion/go/v2/AST"
	"ion/go/v2/Token"
	"strconv"
)

func (parser *Parser) parseLValue() AST.Expression {
	current := parser.peekNthToken(0)
	if parser.consumeOnMatch(Token.IDENTIFIER) {
		return &AST.ExpressionVariable{
			Tok:      current,
			DeclType: nil, // NOTE(Jovanni): The type checker will patch this
		}
	}

	return nil
}

func (parser *Parser) parseArguments() []AST.Expression {
	var ret []AST.Expression

	parser.expect(Token.LEFT_PAREN)
	for !parser.consumeOnMatch(Token.RIGHT_PAREN) {
		expression := parser.parseExpression()

		ret = append(ret, expression)

		if parser.peekNthToken(0).Kind != Token.RIGHT_PAREN {
			parser.expect(Token.COMMA)
		}
	}

	return ret
}

// <Primary>    ::= <integer> | <float> | <boolean> | <string> | '(' <Expression> ')'
func (parser *Parser) parsePrimary() AST.Expression {
	current := parser.peekNthToken(0)
	next := parser.peekNthToken(1)
	if parser.consumeOnMatch(Token.INTEGER_LITERAL) {
		num, _ := strconv.Atoi(current.Lexeme)
		return &AST.ExpressionInteger{Value: num, Tok: current}
	} else if parser.consumeOnMatch(Token.BOOLEAN_LITERAL) {
		b := current.Lexeme == "true"
		return &AST.ExpressionBoolean{Value: b}
	} else if parser.consumeOnMatch(Token.FLOAT_LITERAL) {
		num, _ := strconv.ParseFloat(current.Lexeme, 32)
		return &AST.ExpressionFloat{Value: float32(num)}
	} else if parser.consumeOnMatch(Token.STRING_LITERAL) {
		return &AST.ExpressionString{Value: current.Lexeme[1 : len(current.Lexeme)-1]}
	} else if current.Kind == Token.IDENTIFIER && (next.Kind == Token.INCREMENT || next.Kind == Token.DECREMENT) {
		operand := parser.parseLValue()
		operator := parser.consumeNextToken()

		return &AST.ExpressionPost{
			Operator: operator,
			Operand:  operand,
		}
	} else if (current.Kind == Token.INCREMENT || current.Kind == Token.DECREMENT) && next.Kind == Token.IDENTIFIER {
		operator := parser.consumeNextToken()
		operand := parser.parseLValue()

		return &AST.ExpressionPre{
			Operator: operator,
			Operand:  operand,
		}
	} else if current.Kind == Token.IDENTIFIER {
		if next.Kind == Token.EQUALS {
			lhs := parser.parseLValue()
			parser.expect(Token.EQUALS)
			rhs := parser.parseExpression()

			return &AST.ExpressionAssignment{
				LHSIdentifierToken: current,
				LHS:                lhs,
				RHS:                rhs,
			}
		} else if Token.BinaryOperationFromCompoundAssignment(next.Lexeme) != "" {
			lhs := parser.parseLValue()
			operator := parser.consumeNextToken()
			rhs := parser.parseExpression()

			return &AST.ExpressionCompoundAssignment{
				LHSIdentifierToken: current,
				Operator:           operator,
				LHS:                lhs,
				RHS:                rhs,
			}
		} else if next.Kind == Token.LEFT_PAREN {
			parser.expect(Token.IDENTIFIER)
			arguments := parser.parseArguments()
			return &AST.ExpressionFunctionCall{
				DeclType:  nil,
				Tok:       current,
				Arguments: arguments,
			}
		}

		return parser.parseLValue()
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

// <additive>       ::= <multiplicative> (('+'|'-') <multiplicative>)*
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
