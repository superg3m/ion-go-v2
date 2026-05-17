package IR

import (
	"fmt"
	"ion/go/v2/AST"
)

var counter int

func uniqueTempVariableName() string {
	counter += 1
	return fmt.Sprintf("temp.%d", counter)
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
