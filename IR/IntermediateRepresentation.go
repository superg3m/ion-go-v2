package IR

import (
	"fmt"
	"ion/go/v2/AST"
)

var counter int

func uniqueTempVariableName() string {
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

func (*Constant) isValue() {}
func (*Variable) isValue() {}

func (*Unary) isInstruction()  {}
func (*Return) isInstruction() {}

func emitValueFromStatement(stmt AST.Statement, instructions []Instruction) Value {
	switch v := stmt.(type) {
	case *AST.StatementReturn:
		return NewConstantValue(v.Expr.(*AST.ExpressionInteger).Value)
	}

	return nil
}

func emitValueFromExpression(expr AST.Expression, instructions []Instruction) Value {
	switch v := expr.(type) {
	case *AST.ExpressionUnary:
		source := emitValueFromNode(v.Operand, instructions)
		destination := NewVariable(uniqueTempVariableName())
		instructions = append(instructions, NewUnaryInstruction(v.Operator.Lexeme, source, destination))
		return destination
	}

	return nil
}

func emitValueFromNode(node AST.Node, instructions []Instruction) Value {
	switch v := node.(type) {
	case AST.Expression:
		return emitValueFromNode(v, instructions)
	}

	return nil
}

func EmitIntermediateRepresentation(program AST.Program) Program {
	main := FunctionDefinition{}
	main.identifier = "main"

	decl := program.Declarations[0]
	switch v := decl.(type) {
	case *AST.DeclarationFunction:
		for _, node := range v.Block.Body {
			emitValueFromNode(node, main.instructions)
		}
	}

	return Program{
		FunctionDefinition: main,
	}
}
