package main

import (
	"fmt"
	"ion/go/v2/IR"
	"ion/go/v2/Lexer"
	"ion/go/v2/Parser"
)

func main() {
	tokenStream := Lexer.GenerateTokenStream("./Test/unary.c")

	for i := 0; i < len(tokenStream); i++ {
		token := tokenStream[i]
		tokenType, tokenValue := token.Kind, token.Lexeme
		fmt.Print("Type: ", tokenType, "(", tokenValue, ") | Line:", token.Line, "\n")
	}

	program := Parser.ParseProgram(tokenStream)
	ir := IR.GenerateIntermediateRepresentation(program)
	fmt.Println(ir)
	/*
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
	*/
}
