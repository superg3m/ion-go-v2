package Codegen

import (
	"fmt"
	"ion/go/v2/AssemblyAST"
	"ion/go/v2/IR"
)

var stackOffset int

func instructionsFromIR(inst IR.Instruction) []AssemblyAST.Instruction {
	var instructions []AssemblyAST.Instruction
	switch v := inst.(type) {
	case *IR.Return:
		instructions = append(instructions, AssemblyAST.NewMoveInstruction(AssemblyAST.NewOperand(v.Value), AssemblyAST.NewRegisterOperand(AssemblyAST.EAX)))
		instructions = append(instructions, AssemblyAST.NewReturnInstruction())
	case *IR.Unary:
		destination := AssemblyAST.NewOperand(v.Destination)
		source := AssemblyAST.NewOperand(v.Source)

		instructions = append(instructions, AssemblyAST.NewMoveInstruction(source, destination))
		instructions = append(instructions, AssemblyAST.NewUnaryInstruction(v.Operator, destination))
	case *IR.Binary:
		destination := AssemblyAST.NewOperand(v.Destination)
		left := AssemblyAST.NewOperand(v.Left)
		right := AssemblyAST.NewOperand(v.Right)

		if v.Operator == "/" {
			instructions = append(instructions, AssemblyAST.NewMoveInstruction(left, AssemblyAST.NewRegisterOperand(AssemblyAST.EAX)))
			instructions = append(instructions, AssemblyAST.NewCDQInstruction())
			instructions = append(instructions, AssemblyAST.NewDivideInstruction(right))
			instructions = append(instructions, AssemblyAST.NewMoveInstruction(AssemblyAST.NewRegisterOperand(AssemblyAST.EAX), destination))

			break
		} else if v.Operator == "%" {
			instructions = append(instructions, AssemblyAST.NewMoveInstruction(left, AssemblyAST.NewRegisterOperand(AssemblyAST.EAX)))
			instructions = append(instructions, AssemblyAST.NewCDQInstruction())
			instructions = append(instructions, AssemblyAST.NewDivideInstruction(right))
			instructions = append(instructions, AssemblyAST.NewMoveInstruction(AssemblyAST.NewRegisterOperand(AssemblyAST.EDX), destination))

			break
		}

		instructions = append(instructions, AssemblyAST.NewMoveInstruction(left, destination))
		instructions = append(instructions, AssemblyAST.NewBinaryInstruction(v.Operator, destination, right))
	default:
		panic(fmt.Sprintf("Unknown instruction %T", v))
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
		case *AssemblyAST.InstructionBinary:
			v.Left = replacePseudoOperand(stackOffsetMap, v.Left)
			v.Right = replacePseudoOperand(stackOffsetMap, v.Right)
		case *AssemblyAST.InstructionDivide:
			v.Left = replacePseudoOperand(stackOffsetMap, v.Left)
		case *AssemblyAST.InstructionReturn, *AssemblyAST.InstructionCDQ:
		default:
			panic(fmt.Sprintf("Unknown instruction %T", v))
		}
	}

	return program, stackOffset
}

func ReplaceInvalidInstructions(program AssemblyAST.Program) AssemblyAST.Program {
	newInstructions := make([]AssemblyAST.Instruction, 0, len(program.FunctionDefinition.Instructions)*2)

	R10D := AssemblyAST.NewRegisterOperand(AssemblyAST.R10D)

	for _, inst := range program.FunctionDefinition.Instructions {
		switch v := inst.(type) {
		case *AssemblyAST.InstructionMove:
			if _, ok := v.Source.(*AssemblyAST.Stack); ok {
				if _, ok2 := v.Destination.(*AssemblyAST.Stack); ok2 {
					newInstructions = append(newInstructions, AssemblyAST.NewMoveInstruction(v.Source, R10D))
					v.Source = R10D
				}
			}
		case *AssemblyAST.InstructionBinary:
			if v.Operator == "*" {
				if _, ok2 := v.Left.(*AssemblyAST.Register); !ok2 {
					previousDestination := v.Left
					newInstructions = append(newInstructions, AssemblyAST.NewMoveInstruction(v.Left, R10D))
					v.Left = R10D
					newInstructions = append(newInstructions, v)
					newInstructions = append(newInstructions, AssemblyAST.NewMoveInstruction(R10D, previousDestination))
					continue
				}
			}

			if _, ok := v.Left.(*AssemblyAST.Stack); ok {
				if _, ok2 := v.Right.(*AssemblyAST.Stack); ok2 {
					previousDestination := v.Left
					newInstructions = append(newInstructions, AssemblyAST.NewMoveInstruction(v.Left, R10D))
					v.Left = R10D
					newInstructions = append(newInstructions, v)
					newInstructions = append(newInstructions, AssemblyAST.NewMoveInstruction(R10D, previousDestination))
					continue
				}
			}
		case *AssemblyAST.InstructionDivide:
			if _, ok := v.Left.(*AssemblyAST.Register); !ok {
				newInstructions = append(newInstructions, AssemblyAST.NewMoveInstruction(v.Left, R10D))
				v.Left = R10D
			}
		case *AssemblyAST.InstructionReturn, *AssemblyAST.InstructionCDQ, *AssemblyAST.InstructionStackAllocate, *AssemblyAST.InstructionUnary:
		default:
			panic(fmt.Sprintf("Unknown instruction %T", v))
		}

		newInstructions = append(newInstructions, inst)
	}

	program.FunctionDefinition.Instructions = newInstructions
	return program
}
