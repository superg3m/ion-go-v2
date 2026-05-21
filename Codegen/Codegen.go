package Codegen

import (
	"fmt"
	"ion/go/v2/AssemblyAST"
	"ion/go/v2/IR"
	"ion/go/v2/Symbol"
	"slices"
)

func newOperand(value IR.Value) AssemblyAST.Operand {
	switch v := value.(type) {
	case *IR.Constant:
		return &AssemblyAST.Immediate{Value: v.Value}
	case *IR.Variable:
		return &AssemblyAST.Pseudo{DeclType: v.DeclType, Tok: v.Tok}
	case *IR.ParameterVariable:
		return &AssemblyAST.Parameter{
			Tok:         v.Tok,
			DeclType:    v.DeclType,
			Register:    v.Register,
			StackOffset: v.StackOffset,
		}
	default:
		panic(fmt.Sprintf("Unknown instruction %T", v))
	}

	return nil
}

func instructionsFromIR(inst IR.Instruction) []AssemblyAST.Instruction {
	rax := AssemblyAST.NewRegisterOperand(AssemblyAST.RAX)
	zeroOperand := AssemblyAST.NewImmediateOperand(0)
	// oneOperand := AssemblyAST.NewImmediateOperand(1)

	var instructions []AssemblyAST.Instruction
	switch v := inst.(type) {

	case *IR.Return:
		instructions = append(instructions, AssemblyAST.NewMoveInstruction(newOperand(v.Value), rax))
		instructions = append(instructions, AssemblyAST.NewReturnInstruction())
	case *IR.Unary:
		destination := newOperand(v.Destination)
		source := newOperand(v.Source)
		if v.Operator == "!" {
			instructions = append(instructions, AssemblyAST.NewCompareInstruction(zeroOperand, source))
			instructions = append(instructions, AssemblyAST.NewMoveInstruction(zeroOperand, destination))
			instructions = append(instructions, AssemblyAST.NewSetConditionalCodeInstruction(destination, AssemblyAST.EQUALS))

			break
		}

		instructions = append(instructions, AssemblyAST.NewMoveInstruction(source, destination))
		instructions = append(instructions, AssemblyAST.NewUnaryInstruction(v.Operator, destination))
	case *IR.Binary:
		destination := newOperand(v.Destination)
		left := newOperand(v.Left)
		right := newOperand(v.Right)

		if v.Operator == "/" {
			instructions = append(instructions, AssemblyAST.NewMoveInstruction(left, rax))
			instructions = append(instructions, AssemblyAST.NewCDQInstruction())
			instructions = append(instructions, AssemblyAST.NewDivideInstruction(right))
			instructions = append(instructions, AssemblyAST.NewMoveInstruction(rax, destination))

			break
		} else if v.Operator == "%" {
			instructions = append(instructions, AssemblyAST.NewMoveInstruction(left, rax))
			instructions = append(instructions, AssemblyAST.NewCDQInstruction())
			instructions = append(instructions, AssemblyAST.NewDivideInstruction(right))
			instructions = append(instructions, AssemblyAST.NewMoveInstruction(AssemblyAST.NewRegisterOperand(AssemblyAST.RDX), destination))

			break
		} else if v.Operator == ">" || v.Operator == "<" || v.Operator == "<=" || v.Operator == ">=" || v.Operator == "==" || v.Operator == "!=" {
			instructions = append(instructions, AssemblyAST.NewCompareInstruction(left, right))
			instructions = append(instructions, AssemblyAST.NewMoveInstruction(zeroOperand, destination))
			instructions = append(instructions, AssemblyAST.NewSetConditionalCodeInstruction(destination, AssemblyAST.ConditionalCode(v.Operator)))

			break
		}

		instructions = append(instructions, AssemblyAST.NewMoveInstruction(left, destination))
		instructions = append(instructions, AssemblyAST.NewBinaryInstruction(v.Operator, destination, right))
	case *IR.Jump:
		instructions = append(instructions, AssemblyAST.NewJumpInstruction(v.TargetLabel))
	case *IR.Label:
		instructions = append(instructions, AssemblyAST.NewLabelInstruction(v.Identifier))
	case *IR.Copy:
		left := newOperand(v.Source)
		right := newOperand(v.Destination)
		instructions = append(instructions, AssemblyAST.NewMoveInstruction(left, right))
	case *IR.ConditionalJump:
		code := AssemblyAST.EQUALS
		if v.IfZero {
			code = AssemblyAST.EQUALS
		} else if v.IfNotZero {
			code = AssemblyAST.NOT_EQUALS
		}

		condition := newOperand(v.Condition)
		instructions = append(instructions, AssemblyAST.NewCompareInstruction(condition, zeroOperand))
		instructions = append(instructions, AssemblyAST.NewConditionalJumpInstruction(v.TargetLabel, code))
	case *IR.FunctionCall:
		// NOTE(Jovanni): negative if there are less arguments than the actual argument registers
		stackArgumentCount := max(0, len(v.Arguments)-len(AssemblyAST.ArgumentRegisters))

		var stackPadding int
		if stackArgumentCount%2 == 1 {
			stackPadding = 8
		} else {
			stackPadding = 0
		}

		if stackPadding != 0 {
			instructions = append(instructions, AssemblyAST.NewStackAllocateInstruction(stackPadding))
		}

		for i, arg := range v.Arguments {
			if i >= len(AssemblyAST.ArgumentRegisters) {
				break
			}

			source := newOperand(arg)
			if i < len(AssemblyAST.ArgumentRegisters) {
				register := AssemblyAST.NewRegisterOperand(AssemblyAST.ArgumentRegisters[i])
				instructions = append(instructions, AssemblyAST.NewMoveInstruction(source, register))
			}
		}

		if stackArgumentCount > 0 {
			for i := len(AssemblyAST.ArgumentRegisters) + (stackArgumentCount - 1); i <= len(AssemblyAST.ArgumentRegisters); i -= 1 {
				_, isVariable := v.Arguments[i].(*IR.Variable)
				_, isConstant := v.Arguments[i].(*IR.Constant)
				arg := newOperand(v.Arguments[i])
				if isVariable || isConstant {
					instructions = append(instructions, AssemblyAST.NewStackPushInstruction(arg))
				} else {
					panic(fmt.Sprintf("register not found in argument registers"))
					// emit(Mov(assembly_arg, Reg(AX)))
					//emit(Push(Reg(AX)))
				}
			}
		}

		instructions = append(instructions, AssemblyAST.NewFunctionCallInstruction(v.Identifier.Lexeme))

		bytesToRemove := 8*stackArgumentCount + stackPadding
		if bytesToRemove != 0 {
			instructions = append(instructions, AssemblyAST.NewDeallocateStackInstruction(bytesToRemove))
		}

		destination := newOperand(v.Destination)
		if destination != nil {
			instructions = append(instructions, AssemblyAST.NewMoveInstruction(rax, destination))
		}
	default:
		panic(fmt.Sprintf("Unknown instruction %T", v))
	}

	return instructions
}

func getNext16ByteAligned(totalAllocationSize int) int {
	alignment := 16

	// NOTE(Jovanni): Because alignment is a power of two I can do this
	if (totalAllocationSize & (alignment - 1)) != 0 {
		totalAllocationSize += alignment - (totalAllocationSize & (alignment - 1))
	}

	return totalAllocationSize
}

func assemblyDefinitionFromIRDefinition(definition IR.Definition, globalSymbolTable *Symbol.SymbolTable) AssemblyAST.Definition {
	var instructions []AssemblyAST.Instruction

	switch v := definition.(type) {
	case *IR.FunctionDefinition:
		for _, inst := range v.Instructions {
			instructions = append(instructions, instructionsFromIR(inst)...)
		}

		table := Symbol.CreateSymbolTable(globalSymbolTable)
		instructions = ReplacePseudoRegisters(instructions, &table)
		table.StackOffset = getNext16ByteAligned(table.StackOffset)
		instructions = slices.Insert(instructions, 0, AssemblyAST.NewStackAllocateInstruction(table.StackOffset))
		return &AssemblyAST.FunctionDefinition{
			DeclType:     v.DeclType,
			Tok:          v.Tok,
			Instructions: instructions,
		}
	case *IR.VariableDefinition:
		// TODO(Jovanni): THIS IS NOT RIGHT BECAUSE IT NEEDS TO BE A GLOBAL ADDRESS NOT STACK
		for _, inst := range v.Instructions {
			instructions = append(instructions, instructionsFromIR(inst)...)
		}

		instructions = ReplacePseudoRegisters(instructions, globalSymbolTable)
		return &AssemblyAST.VariableDefinition{
			DeclType:     v.DeclType,
			Tok:          v.Tok,
			Instructions: instructions,
		}
	}

	return nil
}

func GenerateAssemblyProgram(program IR.Program, globalSymbolTable *Symbol.SymbolTable) AssemblyAST.Program {
	var definitions []AssemblyAST.Definition

	for _, def := range program.Definitions {
		definitions = append(definitions, assemblyDefinitionFromIRDefinition(def, globalSymbolTable))
	}

	return AssemblyAST.Program{
		Definitions: definitions,
	}
}

func replacePseudoOperand(table *Symbol.SymbolTable, operand AssemblyAST.Operand) AssemblyAST.Operand {
	switch v := operand.(type) {
	case *AssemblyAST.Pseudo:
		if table.Has(v.Tok) {
			offset := table.GetOffset(v.Tok)
			operand = AssemblyAST.NewStackOperand(offset)
		} else {
			table.StackOffset += v.DeclType.Size()
			table.Set(v.Tok, Symbol.CreateSymbol(v.Tok, v.DeclType), table.StackOffset)
			operand = AssemblyAST.NewStackOperand(table.StackOffset)
		}
	}

	return operand
}

func ReplacePseudoRegisters(instructions []AssemblyAST.Instruction, table *Symbol.SymbolTable) []AssemblyAST.Instruction {
	for _, inst := range instructions {
		switch v := inst.(type) {
		case *AssemblyAST.InstructionMove:
			v.Source = replacePseudoOperand(table, v.Source)
			v.Destination = replacePseudoOperand(table, v.Destination)
		case *AssemblyAST.InstructionUnary:
			v.Operand = replacePseudoOperand(table, v.Operand)
		case *AssemblyAST.InstructionBinary:
			v.Left = replacePseudoOperand(table, v.Left)
			v.Right = replacePseudoOperand(table, v.Right)
		case *AssemblyAST.InstructionDivide:
			v.Left = replacePseudoOperand(table, v.Left)
		case *AssemblyAST.InstructionSetConditionalCode:
			v.Destination = replacePseudoOperand(table, v.Destination)
		case *AssemblyAST.InstructionCompare:
			v.Left = replacePseudoOperand(table, v.Left)
			v.Right = replacePseudoOperand(table, v.Right)
		case *AssemblyAST.InstructionReturn, *AssemblyAST.InstructionCDQ,
			*AssemblyAST.InstructionConditionalJump, *AssemblyAST.InstructionLabel,
			*AssemblyAST.InstructionJump, *AssemblyAST.InstructionFunctionCall:
		default:
			panic(fmt.Sprintf("Unknown instruction %T", v))
		}
	}

	return instructions
}

func ReplaceInvalidInstructions(program AssemblyAST.Program) AssemblyAST.Program {
	for _, def := range program.Definitions {
		switch v := def.(type) {
		case *AssemblyAST.FunctionDefinition:
			newInstructions := make([]AssemblyAST.Instruction, 0, len(v.Instructions)*2)
			R10 := AssemblyAST.NewRegisterOperand(AssemblyAST.R10)
			for _, inst := range v.Instructions {
				switch v := inst.(type) {
				case *AssemblyAST.InstructionMove:
					if _, ok := v.Source.(*AssemblyAST.Stack); ok {
						if _, ok2 := v.Destination.(*AssemblyAST.Stack); ok2 {
							newInstructions = append(newInstructions, AssemblyAST.NewMoveInstruction(v.Source, R10))
							v.Source = R10
						}
					}
				case *AssemblyAST.InstructionBinary:
					if v.Operator == "*" {
						if _, ok2 := v.Left.(*AssemblyAST.Register); !ok2 {
							previousDestination := v.Left
							newInstructions = append(newInstructions, AssemblyAST.NewMoveInstruction(v.Left, R10))
							v.Left = R10
							newInstructions = append(newInstructions, v)
							newInstructions = append(newInstructions, AssemblyAST.NewMoveInstruction(R10, previousDestination))
							continue
						}
					}

					if _, ok := v.Left.(*AssemblyAST.Stack); ok {
						if _, ok2 := v.Right.(*AssemblyAST.Stack); ok2 {
							previousDestination := v.Left
							newInstructions = append(newInstructions, AssemblyAST.NewMoveInstruction(v.Left, R10))
							v.Left = R10
							newInstructions = append(newInstructions, v)
							newInstructions = append(newInstructions, AssemblyAST.NewMoveInstruction(R10, previousDestination))
							continue
						}
					}
				case *AssemblyAST.InstructionDivide:
					if _, ok := v.Left.(*AssemblyAST.Register); !ok {
						newInstructions = append(newInstructions, AssemblyAST.NewMoveInstruction(v.Left, R10))
						v.Left = R10
					}
				case *AssemblyAST.InstructionCompare:
					if _, ok := v.Left.(*AssemblyAST.Register); !ok {
						newInstructions = append(newInstructions, AssemblyAST.NewMoveInstruction(v.Left, R10))
						v.Left = R10
					}
				case *AssemblyAST.InstructionReturn, *AssemblyAST.InstructionCDQ,
					*AssemblyAST.InstructionStackAllocate, *AssemblyAST.InstructionUnary,
					*AssemblyAST.InstructionConditionalJump, *AssemblyAST.InstructionLabel,
					*AssemblyAST.InstructionJump, *AssemblyAST.InstructionSetConditionalCode,
					*AssemblyAST.InstructionFunctionCall:
				default:
					panic(fmt.Sprintf("Unknown instruction %T", v))
				}

				newInstructions = append(newInstructions, inst)
			}

			v.Instructions = newInstructions
		}
	}

	return program
}
