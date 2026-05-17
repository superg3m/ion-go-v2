package Codegen

import (
	"ion/go/v2/AssemblyAST"
	"ion/go/v2/IR"
	"slices"
)

var stackOffset int

func instructionsFromIR(inst IR.Instruction) []AssemblyAST.Instruction {
	var instructions []AssemblyAST.Instruction
	switch v := inst.(type) {
	case *IR.Return:
		instructions = append(instructions, AssemblyAST.NewMoveInstruction(AssemblyAST.NewRegisterOperand(AssemblyAST.EAX), AssemblyAST.NewOperand(v.Value)))
		instructions = append(instructions, AssemblyAST.NewReturnInstruction())
	case *IR.Unary:
		destination := AssemblyAST.NewOperand(v.Destination)
		source := AssemblyAST.NewOperand(v.Source)

		instructions = append(instructions, AssemblyAST.NewMoveInstruction(source, destination))
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

func replacePseudoOperand(stackOffsetMap map[string]int, operand AssemblyAST.Operand) AssemblyAST.Operand {
	switch v := operand.(type) {
	case *AssemblyAST.Pseudo:
		if offset, ok := stackOffsetMap[v.Identifier]; ok {
			operand = AssemblyAST.NewStackOperand(offset)
		} else {
			stackOffset += 4
			stackOffsetMap[v.Identifier] = stackOffset
			operand = AssemblyAST.NewStackOperand(stackOffset)
		}
	}

	return operand
}

func ReplacePseudoRegisters(program AssemblyAST.Program) (AssemblyAST.Program, int) {
	stackOffsetMap := make(map[string]int)

	for _, inst := range program.FunctionDefinition.Instructions {
		switch v := inst.(type) {
		case *AssemblyAST.InstructionMove:
			v.Source = replacePseudoOperand(stackOffsetMap, v.Source)
			v.Destination = replacePseudoOperand(stackOffsetMap, v.Destination)
		case *AssemblyAST.InstructionUnary:
			v.Operand = replacePseudoOperand(stackOffsetMap, v.Operand)
		}
	}

	return program, stackOffset
}

func ReplaceInvalidMoveInstructions(program AssemblyAST.Program) AssemblyAST.Program {
	for i, inst := range program.FunctionDefinition.Instructions {
		switch v := inst.(type) {
		case *AssemblyAST.InstructionMove:
			if _, ok := v.Source.(*AssemblyAST.Stack); ok {
				if _, ok := v.Destination.(*AssemblyAST.Stack); ok {
					previousDestination := v.Destination
					v.Destination = AssemblyAST.NewRegisterOperand(AssemblyAST.R10D)
					program.FunctionDefinition.Instructions = slices.Insert(program.FunctionDefinition.Instructions, i+1, AssemblyAST.NewMoveInstruction(AssemblyAST.NewRegisterOperand(AssemblyAST.R10D), previousDestination))
				}
			}
		}
	}

	return program
}
