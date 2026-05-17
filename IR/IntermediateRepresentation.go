package IR

import (
	"fmt"
	"ion/go/v2/AST"
)

var tempVariableCounter int
var labelCounter int

func uniqueTempVariableName() string {
	tempVariableCounter += 1
	return fmt.Sprintf("temp.%d", tempVariableCounter)
}

func uniqueLabelName() string {
	labelCounter += 1
	return fmt.Sprintf("%d", labelCounter)
}

type Program struct {
	FunctionDefinition FunctionDefinition
}

type FunctionDefinition struct {
	Identifier   string
	Instructions []Instruction
}

func emitFromStatement(stmt AST.Statement, instructions []Instruction) []Instruction {
	switch v := stmt.(type) {
	case *AST.StatementReturn:
		var value Value
		instructions, value = emitFromExpression(v.Expr, instructions)
		instructions = append(instructions, NewReturnInstruction(value))
	default:
		panic(fmt.Sprintf("Unknown instruction %T", v))
	}

	return instructions
}

func emitFromExpression(expr AST.Expression, instructions []Instruction) ([]Instruction, Value) {
	switch v := expr.(type) {
	case *AST.ExpressionGrouping:
		return emitFromExpression(v.Expr, instructions)
	case *AST.ExpressionInteger:
		return instructions, NewConstantValue(v.Value)
	case *AST.ExpressionUnary:
		var source Value
		instructions, source = emitFromExpression(v.Operand, instructions)
		destination := NewVariable(uniqueTempVariableName())
		instructions = append(instructions, NewUnaryInstruction(v.Operator.Lexeme, source, destination))
		return instructions, destination
	case *AST.ExpressionBinary:
		var left, right Value
		instructions, left = emitFromExpression(v.Left, instructions)
		if v.Operator.Lexeme == "&&" {
			zeroLabel := uniqueLabelName()
			oneLabel := uniqueLabelName()
			endLabel := uniqueLabelName()

			result := NewVariable(uniqueTempVariableName())
			instructions = append(instructions, NewConditionalJumpInstruction(zeroLabel, left, true, false))
			instructions, right = emitFromExpression(v.Right, instructions)
			instructions = append(instructions, NewLabelInstruction(oneLabel))
			instructions = append(instructions, NewCopyInstruction(NewConstantValue(1), result))
			instructions = append(instructions, NewJumpInstruction(endLabel))

			instructions = append(instructions, NewLabelInstruction(zeroLabel))
			instructions = append(instructions, NewCopyInstruction(NewConstantValue(0), result))

			instructions = append(instructions, NewLabelInstruction(endLabel))

			return instructions, result
		} else if v.Operator.Lexeme == "||" {
			zeroLabel := uniqueLabelName()
			oneLabel := uniqueLabelName()
			endLabel := uniqueLabelName()

			result := NewVariable(uniqueTempVariableName())
			instructions = append(instructions, NewConditionalJumpInstruction(oneLabel, left, false, true))
			instructions, right = emitFromExpression(v.Right, instructions)
			instructions = append(instructions, NewConditionalJumpInstruction(zeroLabel, right, true, false))

			instructions = append(instructions, NewLabelInstruction(oneLabel))
			instructions = append(instructions, NewCopyInstruction(NewConstantValue(1), result))
			instructions = append(instructions, NewJumpInstruction(endLabel))

			instructions = append(instructions, NewLabelInstruction(zeroLabel))
			instructions = append(instructions, NewCopyInstruction(NewConstantValue(0), result))

			instructions = append(instructions, NewLabelInstruction(endLabel))

			return instructions, result
		}

		instructions, right = emitFromExpression(v.Right, instructions)
		destination := NewVariable(uniqueTempVariableName())
		instructions = append(instructions, NewBinaryInstruction(v.Operator.Lexeme, left, right, destination))
		return instructions, destination
	default:
		panic(fmt.Sprintf("Unknown instruction %T", v))
	}

	return instructions, nil
}

func emitFromNode(node AST.Node, instructions []Instruction) []Instruction {
	switch v := node.(type) {
	case AST.Expression:
		instructions, _ = emitFromExpression(v, instructions)
	case AST.Statement:
		return emitFromStatement(v, instructions)
	default:
		panic(fmt.Sprintf("Unknown instruction %T", v))
	}

	return instructions
}

func GenerateIntermediateRepresentation(program AST.Program) Program {
	main := FunctionDefinition{}
	main.Identifier = "main"

	decl := program.Declarations[0]
	switch v := decl.(type) {
	case *AST.DeclarationFunction:
		for _, node := range v.Block.Body {
			main.Instructions = emitFromNode(node, main.Instructions)
		}
	default:
		panic(fmt.Sprintf("Unknown instruction %T", v))
	}

	return Program{
		FunctionDefinition: main,
	}
}
