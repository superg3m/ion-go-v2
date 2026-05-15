package AssemblyEmitter

import (
	"fmt"
	"ion/go/v2/AST"
)

type ATTx64Emitter struct {
}

func (e *ATTx64Emitter) EmitReturnInstruction() {}

func (e *ATTx64Emitter) EmitInstruction(inst AST.Instruction) []string {
	var instructions []string
	switch v := inst.(type) {
	case *AST.InstructionMove:
		instructions = append(instructions, fmt.Sprintf("\tmovl %s, %s", v.Source.ToString(), v.Destination.ToString()))
	case *AST.InstructionReturn:
		instructions = append(instructions, "\tret")
	}

	return instructions
}

func (e *ATTx64Emitter) EmitFunctionDefinition(functionDefinition *AST.FunctionDefinition) []string {
	instructions := []string{
		".text",
		fmt.Sprintf(".global %s", functionDefinition.Identifier),
		fmt.Sprintf(".%s:", functionDefinition.Identifier),
	}

	for _, inst := range functionDefinition.Instructions {
		instructions = append(instructions, e.EmitInstruction(inst)...)
	}

	return instructions
}

func (e *ATTx64Emitter) EmitAssembly(program AST.AssemblyProgram) []string {
	var instructions []string
	instructions = append(instructions, e.EmitFunctionDefinition(program.FunctionDefinition)...)

	return instructions
}
