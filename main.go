package main

import (
	"fmt"
	"ion/go/v2/AssemblyAST"
	"ion/go/v2/AssemblyEmitter"
	"ion/go/v2/Codegen"
	"ion/go/v2/IR"
	"ion/go/v2/Lexer"
	"ion/go/v2/Parser"
	"log"
	"os"
	"slices"
)

func main() {
	tokenStream := Lexer.GenerateTokenStream("./Test/binary.c")

	for i := 0; i < len(tokenStream); i++ {
		token := tokenStream[i]
		tokenType, tokenValue := token.Kind, token.Lexeme
		fmt.Print("Type: ", tokenType, "(", tokenValue, ") | Line:", token.Line, "\n")
	}

	program := Parser.ParseProgram(tokenStream)
	ir := IR.GenerateIntermediateRepresentation(program)
	assembly := Codegen.GenerateAssemblyProgram(ir)

	finalStackOffset := 0
	assembly, finalStackOffset = Codegen.ReplacePseudoRegisters(assembly)
	assembly.FunctionDefinition.Instructions = slices.Insert(assembly.FunctionDefinition.Instructions, 0, AssemblyAST.NewStackAllocateInstruction(finalStackOffset))
	assembly = Codegen.ReplaceInvalidInstructions(assembly)

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
