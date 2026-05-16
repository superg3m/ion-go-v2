package AssemblyEmitter

import (
	"fmt"
	"ion/go/v2/AssemblyAST"
)

type ATTx64Emitter struct{}

func EmitFunctionPrologue() []string {
	return []string{
		"\tpushq %rbp",
		"\tmovq %rsp, %rbp",
		"\tsubq $8, %rsp",
	}
}

func EmitFunctionEpilogue() []string {
	return []string{
		"\tmovq %rbp, %rsp",
		"\tpopq %rbp",
		"\tret",
	}
}

func (e *ATTx64Emitter) EmitReturnInstruction() {}

func (e *ATTx64Emitter) EmitInstruction(inst AssemblyAST.Instruction) []string {
	var instructions []string
	switch v := inst.(type) {
	case *AssemblyAST.InstructionMove:
		instructions = append(instructions, fmt.Sprintf("\tmovl %s, %s", v.Source.ToString(), v.Destination.ToString()))
	case *AssemblyAST.InstructionReturn:
		instructions = EmitFunctionEpilogue()
	}

	return instructions
}

func (e *ATTx64Emitter) EmitFunctionDefinition(functionDefinition *AssemblyAST.FunctionDefinition) []string {
	instructions := []string{
		fmt.Sprintf(".global %s", functionDefinition.Identifier),
		fmt.Sprintf("%s:", functionDefinition.Identifier),
	}

	instructions = append(instructions, EmitFunctionPrologue()...)

	for _, inst := range functionDefinition.Instructions {
		instructions = append(instructions, e.EmitInstruction(inst)...)
	}

	// instructions = append(instructions, EmitFunctionEpilogue()...)

	return instructions
}

func (e *ATTx64Emitter) EmitAssemblyProgram(program AssemblyAST.Program) []string {
	instructions := []string{
		".text",
	}
	instructions = append(instructions, e.EmitFunctionDefinition(program.FunctionDefinition)...)

	return instructions
}
