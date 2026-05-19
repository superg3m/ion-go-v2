package IR

import (
	"fmt"
	"ion/go/v2/AST"
	"ion/go/v2/TS"
	"ion/go/v2/Token"
)

var tempVariableCounter int
var labelCounter int

func uniqueTempVariableToken(token Token.Token) Token.Token {
	tempVariableCounter += 1
	token.Lexeme = fmt.Sprintf("temp.%d", tempVariableCounter)
	return token
}

func uniqueLabelName() string {
	labelCounter += 1
	return fmt.Sprintf("%d", labelCounter)
}

type Definition interface {
	isDefinition()
}

type Program struct {
	Definitions []Definition
}

type FunctionDefinition struct {
	DeclType     TS.Type
	Tok          Token.Token
	Instructions []Instruction
}

type VariableDefinition struct {
	DeclType     TS.Type
	Tok          Token.Token
	Instructions []Instruction
}

func (*FunctionDefinition) isDefinition() {}
func (*VariableDefinition) isDefinition() {}

func emitFromStatement(stmt AST.Statement, instructions []Instruction) []Instruction {
	switch v := stmt.(type) {
	case *AST.StatementReturn:
		var value Value
		instructions, value = emitFromExpression(v.Expr, instructions)
		instructions = append(instructions, NewReturnInstruction(value))
	case *AST.StatementExpression:
		instructions, _ = emitFromExpression(v.Expr, instructions)
	case *AST.StatementIfElse:
		elseLabel := uniqueLabelName()
		endLabel := uniqueLabelName()

		var condition Value
		instructions, condition = emitFromExpression(v.Condition, instructions)
		instructions = append(instructions, NewConditionalJumpInstruction(elseLabel, condition, true, false))

		instructions = emitFromStatement(v.ThenBlock, instructions)
		instructions = append(instructions, NewJumpInstruction(elseLabel))

		instructions = append(instructions, NewLabelInstruction(elseLabel))
		if v.ElseBlock != nil {
			instructions = emitFromStatement(v.ElseBlock, instructions)
		}

		instructions = append(instructions, NewLabelInstruction(endLabel))
	case *AST.StatementBlock:
		for _, node := range v.Body {
			instructions = append(instructions, emitFromNode(node)...)
		}
	case *AST.StatementCompoundAssignment:
		var left, right Value
		instructions, left = emitFromExpression(v.LHS, instructions)
		instructions, right = emitFromExpression(v.RHS, instructions)
		destination := NewVariable(v.LHSIdentifierToken, left.GetDeclType())
		instructions = append(instructions, NewBinaryInstruction(Token.BinaryOperationFromCompoundAssignment(v.Operator.Lexeme), left, right, destination))
	default:
		panic(fmt.Sprintf("Unknown instruction %T", v))
	}

	return instructions
}

func emitFromDeclaration(decl AST.Declaration, instructions []Instruction) []Instruction {
	switch v := decl.(type) {
	case *AST.DeclarationVariable:
		value := NewConstantValue(0, v.Tok)
		if v.RHS != nil {
			instructions, value = emitFromExpression(v.RHS, instructions)
		}

		destination := NewVariable(v.Tok, v.DeclType)
		instructions = append(instructions, NewCopyInstruction(value, destination))
	case *AST.DeclarationFunction:
		for _, node := range v.Block.Body {
			instructions = append(instructions, emitFromNode(node)...)
		}
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
		return instructions, NewConstantValue(v.Value, v.Tok)
	case *AST.ExpressionVariable:
		return instructions, NewVariable(v.Tok, v.DeclType)
	case *AST.ExpressionUnary:
		var source Value
		instructions, source = emitFromExpression(v.Operand, instructions)
		destination := NewVariable(uniqueTempVariableToken(v.Operator), source.GetDeclType())
		instructions = append(instructions, NewUnaryInstruction(v.Operator.Lexeme, source, destination))
		return instructions, destination
	case *AST.ExpressionAssignment:
		var source Value
		instructions, source = emitFromExpression(v.RHS, instructions)
		destination := NewVariable(v.LHSIdentifierToken, source.GetDeclType())
		instructions = append(instructions, NewCopyInstruction(source, destination))
		return instructions, destination
	case *AST.ExpressionBinary:
		var left, right Value
		instructions, left = emitFromExpression(v.Left, instructions)
		if v.Operator.Lexeme == "&&" {
			zeroLabel := uniqueLabelName()
			oneLabel := uniqueLabelName()
			endLabel := uniqueLabelName()

			result := NewVariable(uniqueTempVariableToken(v.Operator), left.GetDeclType())
			instructions = append(instructions, NewConditionalJumpInstruction(zeroLabel, left, true, false))
			instructions, right = emitFromExpression(v.Right, instructions)
			instructions = append(instructions, NewLabelInstruction(oneLabel))
			instructions = append(instructions, NewConditionalJumpInstruction(zeroLabel, right, true, false))
			instructions = append(instructions, NewCopyInstruction(NewConstantValue(1, v.Operator), result))
			instructions = append(instructions, NewJumpInstruction(endLabel))

			instructions = append(instructions, NewLabelInstruction(zeroLabel))
			instructions = append(instructions, NewCopyInstruction(NewConstantValue(0, v.Operator), result))

			instructions = append(instructions, NewLabelInstruction(endLabel))

			return instructions, result
		} else if v.Operator.Lexeme == "||" {
			zeroLabel := uniqueLabelName()
			oneLabel := uniqueLabelName()
			endLabel := uniqueLabelName()

			result := NewVariable(uniqueTempVariableToken(v.Operator), left.GetDeclType())
			instructions = append(instructions, NewConditionalJumpInstruction(oneLabel, left, false, true))
			instructions, right = emitFromExpression(v.Right, instructions)
			instructions = append(instructions, NewConditionalJumpInstruction(zeroLabel, right, true, false))

			instructions = append(instructions, NewLabelInstruction(oneLabel))
			instructions = append(instructions, NewCopyInstruction(NewConstantValue(1, v.Operator), result))
			instructions = append(instructions, NewJumpInstruction(endLabel))

			instructions = append(instructions, NewLabelInstruction(zeroLabel))
			instructions = append(instructions, NewCopyInstruction(NewConstantValue(0, v.Operator), result))

			instructions = append(instructions, NewLabelInstruction(endLabel))

			return instructions, result
		}

		instructions, right = emitFromExpression(v.Right, instructions)
		destination := NewVariable(uniqueTempVariableToken(v.Operator), left.GetDeclType())
		instructions = append(instructions, NewBinaryInstruction(v.Operator.Lexeme, left, right, destination))
		return instructions, destination
	default:
		panic(fmt.Sprintf("Unknown instruction %T", v))
	}

	return instructions, nil
}

func emitFromNode(node AST.Node) []Instruction {
	var instructions []Instruction
	switch v := node.(type) {
	case AST.Expression:
		instructions, _ = emitFromExpression(v, instructions)
	case AST.Statement:
		instructions = emitFromStatement(v, instructions)
	case AST.Declaration:
		instructions = emitFromDeclaration(v, instructions)
	default:
		panic(fmt.Sprintf("Unknown instruction %T", v))
	}

	return instructions
}

func emitDefinitionFromDeclaration(decl AST.Declaration) Definition {
	var instructions []Instruction

	switch v := decl.(type) {
	case *AST.DeclarationVariable:
		var value Value
		instructions, value = emitFromExpression(v.RHS, instructions)
		destination := NewVariable(v.Tok, v.DeclType)
		instructions = append(instructions, NewCopyInstruction(value, destination))
		return &VariableDefinition{
			DeclType:     v.DeclType,
			Tok:          v.Tok,
			Instructions: instructions,
		}
	case *AST.DeclarationFunction:
		instructions = append(instructions, emitFromStatement(v.Block, instructions)...)

		return &FunctionDefinition{
			DeclType:     v.DeclType,
			Tok:          v.Tok,
			Instructions: instructions,
		}
	default:
		panic(fmt.Sprintf("Unknown instruction %T", v))
	}

	return nil
}

func GenerateIntermediateRepresentation(program AST.Program) Program {
	var definitions []Definition

	for _, decl := range program.Declarations {
		definitions = append(definitions, emitDefinitionFromDeclaration(decl))
	}

	return Program{
		Definitions: definitions,
	}
}
