package main

import (
	"fmt"
	"ion/go/v2/AssemblyEmitter"
	"ion/go/v2/Codegen"
	"ion/go/v2/Lexer"
	"ion/go/v2/Parser"
	"log"
	"os"
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

	emitter := &AssemblyEmitter.ATTx64Emitter{}
	instructions := emitter.EmitAssemblyProgram(assembly)

	f, err := os.Create("Test/output.s")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	for _, line := range instructions {
		_, err := f.WriteString(line + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
}
