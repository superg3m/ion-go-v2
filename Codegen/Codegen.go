package Codegen

import (
	"ion/go/v2/IR"
)
import "ion/go/v2/AssemblyAST"

func instructionsFromIR(inst IR.Instruction) []AssemblyAST.Instruction {
	var instructions []AssemblyAST.Instruction
	switch v := inst.(type) {
	case *IR.Return:
		instructions = append(instructions, AssemblyAST.NewMoveInstruction(AssemblyAST.NewRegisterOperand(AssemblyAST.AX), AssemblyAST.NewOperand(v.Value)))
		instructions = append(instructions, AssemblyAST.NewReturnInstruction())
	case *IR.Unary:
		destination := AssemblyAST.NewOperand(v.Destination)
		source := AssemblyAST.NewOperand(v.Source)

		instructions = append(instructions, AssemblyAST.NewMoveInstruction(destination, source))
		instructions = append(instructions, AssemblyAST.NewUnaryInstruction(v.Operator, destination))
	}

	return instructions
}

func GenerateAssemblyProgram(program IR.Program) AssemblyAST.Program {
	main := &AssemblyAST.FunctionDefinition{}
	main.Identifier = "main"

	for _, inst := range program.FunctionDefinition.Instructions {
		main.Instructions = append(main.Instructions, instructionsFromIR(inst)...)
	}

	return AssemblyAST.Program{
		FunctionDefinition: main,
	}
}
