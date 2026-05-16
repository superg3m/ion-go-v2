package Codegen

import (
	"ion/go/v2/AssemblyAST"
	"ion/go/v2/IR"
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

func replaceInvalidMoveInstruction(operand AssemblyAST.Operand) AssemblyAST.Operand {
	switch operand.(type) {
	case *AssemblyAST.Stack:
		operand = AssemblyAST.NewRegisterOperand(AssemblyAST.R10D)
		// Prob going to do something else here?
	}

	return operand
}

func ReplaceInvalidMoveInstructions(program AssemblyAST.Program) {
	for _, inst := range program.FunctionDefinition.Instructions {
		switch v := inst.(type) {
		case *AssemblyAST.InstructionMove:
			if _, ok := v.Source.(*AssemblyAST.Stack); ok {
				v.Destination = replaceInvalidMoveInstruction(v.Destination)
			}
		}
	}
}
