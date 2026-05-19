package Unique

import (
	"fmt"
	"ion/go/v2/Token"
)

var tempVariableCounter int
var labelCounter int

func TempVariableToken(token Token.Token) Token.Token {
	tempVariableCounter += 1
	token.Lexeme = fmt.Sprintf("temp.%d", tempVariableCounter)
	return token
}

func LabelName() string {
	labelCounter += 1
	return fmt.Sprintf("%d", labelCounter)
}
