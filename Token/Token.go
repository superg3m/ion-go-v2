package Token

type TokenType string

const (
	ILLEGAL_TOKEN TokenType = "ILLEGAL_TOKEN"
	EOF                     = "EOF"

	// SYNTAX SINGLE CHARACTER
	NOT           = "NOT"           // "!"
	LEFT_PAREN    = "LEFT_PAREN"    // "("
	RIGHT_PAREN   = "RIGHT_PAREN"   // ")"
	PLUS          = "PLUS"          // "+"
	MINUS         = "MINUS"         // "-"
	STAR          = "STAR"          // "*"
	DIVISION      = "DIVISION"      // "/"
	MODULUS       = "MODULUS"       // "%"
	COMMA         = "COMMA"         // ","
	COLON         = "COLON"         // ":"
	DOT           = "DOT"           // "."
	SEMI_COLON    = "SEMI_COLON"    // ";"
	LESS_THAN     = "LESS_THAN"     // "<"
	EQUALS        = "EQUAL"         // "="
	GREATER_THAN  = "GREATER_THAN"  // ">"
	LEFT_BRACKET  = "LEFT_BRACKET"  // "["
	RIGHT_BRACKET = "RIGHT_BRACKET" // "]"
	LEFT_CURLY    = "LEFT_CURLY"    // "{"
	RIGHT_CURLY   = "RIGHT_CURLY"   // "}"
	BITWISE_NOT   = "BITWISE_NOT"   // "~"
	BITWISE_AND   = "BITWISE_AND"   // "&"
	BITWISE_OR    = "BITWISE_OR"    // "|"
	BITWISE_LS    = "BITWISE_LS"    // "<<"
	BITWISE_RS    = "BITWISE_RS"    // ">>"
	BITWISE_XOR   = "BITWISE_XOR"   // "^"

	PLUS_EQUALS        = "PLUS_EQUALS"        // "+="
	MINUS_EQUALS       = "MINUS_EQUALS"       // "-="
	STAR_EQUALS        = "STAR_EQUALS"        // "*="
	DIVIDE_EQUALS      = "DIVIDE_EQUALS"      // "/="
	BITWISE_AND_EQUALS = "BITWISE_AND_EQUALS" // "&="
	BITWISE_OR_EQUALS  = "BITWISE_OR_EQUALS"  // "|="
	BITWISE_LS_EQUALS  = "DIVIDE_EQUALS"      // "<<="
	BITWISE_RS_EQUALS  = "DIVIDE_EQUALS"      // ">>="

	// SYNTAX MULTIPLE CHARACTERS
	EQUALS_EQUALS       = "EQUALS_EQUALS"       // "=="
	NOT_EQUALS          = "NOT_EQUALS"          // "!="
	GREATER_THAN_EQUALS = "GREATER_THAN_EQUALS" // ">="
	LESS_THAN_EQUALS    = "LESS_THAN_EQUALS"    // "<="
	LOGICAL_AND         = "LOGICAL_AND"         // "&&"
	LOGICAL_OR          = "LOGICAL_OR"          // "||"
	RIGHT_ARROW         = "RIGHT_ARROW"         // "->"

	IDENTIFIER        = "IDENTIFIER"
	INTEGER_LITERAL   = "INTEGER_LITERAL"
	FLOAT_LITERAL     = "FLOAT_LITERAL"
	BOOLEAN_LITERAL   = "BOOLEAN_LITERAL"
	STRING_LITERAL    = "STRING_LITERAL"
	CHARACTER_LITERAL = "CHARACTER_LITERAL"

	// Keywords
	FN       = "FN"
	STRUCT   = "STRUCT"
	CAST     = "CAST"
	VAR      = "VAR"
	IF       = "IF"
	ELSE     = "ELSE"
	FOR      = "FOR"
	WHILE    = "WHILE"
	NULLPTR  = "NULLPTR"
	RETURN   = "RETURN"
	BREAK    = "BREAK"
	CONTINUE = "CONTINUE"
	PRINT    = "PRINT"
	PRINTLN  = "PRINTLN"
	DEFER    = "DEFER"

	// Builtin
	BUILTIN_LEN = "BUILTIN_LEN"
)

type Token struct {
	Kind   TokenType
	Lexeme string
	Line   int
}

func CreateToken(kind TokenType, lexeme string, line int) Token {
	return Token{kind, lexeme, line}
}

func IsCompoundAssignment(tokenType TokenType) bool {
	switch tokenType {
	case PLUS_EQUALS:
		return true
	case MINUS_EQUALS:
		return true
	case STAR_EQUALS:
		return true
	case DIVIDE_EQUALS:
		return true
	}

	return false
}

func BinaryOperationFromCompoundAssignment(compound string) string {
	switch compound {
	case "+=":
		return "+"
	case "-=":
		return "-"
	case "*=":
		return "*"
	case "/=":
		return "/"
	case "&=":
		return "&"
	case "|=":
		return "|"
	case "<<=":
		return "<<"
	case ">>=":
		return ">>"
	}

	panic("BinaryOperationFromCompoundAssignment: invalid compound")
	return ""
}

func GetKeywordToken(input string) (TokenType, bool) {
	var m = map[string]TokenType{
		"fn":       FN,
		"struct":   STRUCT,
		"cast":     CAST,
		"var":      VAR,
		"if":       IF,
		"else":     ELSE,
		"for":      FOR,
		"while":    WHILE,
		"nullptr":  NULLPTR,
		"return":   RETURN,
		"break":    BREAK,
		"continue": CONTINUE,
		"print":    PRINT,
		"println":  PRINTLN,
		"defer":    DEFER,
		"true":     BOOLEAN_LITERAL,
		"false":    BOOLEAN_LITERAL,
	}

	token, ok := m[input]

	return token, ok
}

func GetBuiltinToken(input string) (TokenType, bool) {
	var m = map[string]TokenType{
		"len": BUILTIN_LEN,
	}

	token, ok := m[input]

	return token, ok
}

func GetSyntaxToken(input string) (TokenType, bool) {
	var m = map[string]TokenType{
		"=":   EQUALS,
		"%":   MODULUS,
		"+":   PLUS,
		"-":   MINUS,
		"/":   DIVISION,
		"*":   STAR,
		"<":   LESS_THAN,
		">":   GREATER_THAN,
		"!":   NOT,
		"(":   LEFT_PAREN,
		")":   RIGHT_PAREN,
		",":   COMMA,
		":":   COLON,
		".":   DOT,
		";":   SEMI_COLON,
		"[":   LEFT_BRACKET,
		"]":   RIGHT_BRACKET,
		"{":   LEFT_CURLY,
		"}":   RIGHT_CURLY,
		"~":   BITWISE_NOT,
		"&":   BITWISE_AND,
		"|":   BITWISE_OR,
		"<<":  BITWISE_LS,
		">>":  BITWISE_RS,
		"^":   BITWISE_XOR,
		"==":  EQUALS_EQUALS,
		"!=":  NOT_EQUALS,
		">=":  GREATER_THAN_EQUALS,
		"<=":  LESS_THAN_EQUALS,
		"&&":  LOGICAL_AND,
		"||":  LOGICAL_OR,
		"->":  RIGHT_ARROW,
		"+=":  PLUS_EQUALS,
		"-=":  MINUS_EQUALS,
		"*=":  STAR_EQUALS,
		"/=":  DIVIDE_EQUALS,
		"&=":  BITWISE_AND_EQUALS,
		"|=":  BITWISE_OR_EQUALS,
		"<<=": BITWISE_LS_EQUALS,
		">>=": BITWISE_RS_EQUALS,
	}

	token, ok := m[input]

	return token, ok
}
