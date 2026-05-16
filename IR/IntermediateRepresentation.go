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
	identifier   string
	instructions []Instruction
}

type Instruction interface {
	isInstruction()
}

type Value interface {
	isValue()
}

type Unary struct {
	operator    string
	source      Value
	destination Value
}

func NewUnaryInstruction(operator string, source Value, destination Value) Instruction {
	return &Unary{
		operator:    operator,
		source:      source,
		destination: destination,
	}
}

type Constant struct {
	Value int
}

type Variable struct {
	Name string
}

func NewConstantValue(value int) Value {
	return &Constant{
		Value: value,
	}
}

func NewVariable(name string) Value {
	return &Variable{
		Name: name,
	}
}

type Return struct {
	Value Value
}

func NewReturnInstruction(value Value) Instruction {
	return &Return{
		Value: value,
	}
}

func (*Constant) isValue() {}
func (*Variable) isValue() {}

func (*Unary) isInstruction()  {}
func (*Return) isInstruction() {}

func emitFromStatement(stmt AST.Statement, instructions []Instruction) []Instruction {
	switch v := stmt.(type) {
	case *AST.StatementReturn:
		var value Value
		instructions, value = emitFromExpression(v.Expr, instructions)
		instructions = append(instructions, NewReturnInstruction(value))
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
	}

	return instructions, nil
}

func emitFromNode(node AST.Node, instructions []Instruction) []Instruction {
	switch v := node.(type) {
	case AST.Expression:
		instructions, _ = emitFromExpression(v, instructions)
	case AST.Statement:
		return emitFromStatement(v, instructions)
	}

	return instructions
}

func GenerateIntermediateRepresentation(program AST.Program) Program {
	main := FunctionDefinition{}
	main.identifier = "main"

	decl := program.Declarations[0]
	switch v := decl.(type) {
	case *AST.DeclarationFunction:
		for _, node := range v.Block.Body {
			main.instructions = emitFromNode(node, main.instructions)
		}
	}

	return Program{
		FunctionDefinition: main,
	}
}
