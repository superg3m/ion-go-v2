package main

import (
	"fmt"
	"ion/go/v2/Codegen"
	"ion/go/v2/Lexer"
	"ion/go/v2/Parser"
)

func main() {
	tokenStream := Lexer.GenerateTokenStream("./Test/return.c")

	for i := 0; i < len(tokenStream); i++ {
		token := tokenStream[i]
		tokenType, tokenValue := token.Kind, token.Lexeme
		fmt.Print("Type: ", tokenType, "(", tokenValue, ") | Line:", token.Line, "\n")
	}

	program := Parser.ParseProgram(tokenStream)
	assembly := Codegen.GenerateAssemblyProgram(program)
	fmt.Println(assembly)
}
