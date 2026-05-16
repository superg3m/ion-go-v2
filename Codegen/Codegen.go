package Codegen

import (
	"ion/go/v2/AST"
	"ion/go/v2/IR"
)
import "ion/go/v2/AssemblyAST"

func instructionsFromStatement(stmt AST.Statement) []AssemblyAST.Instruction {
	var instructions []AssemblyAST.Instruction
	switch v := stmt.(type) {
	case *AST.StatementReturn:
		destination := &AssemblyAST.Register{Name: "%eax"}
		source := &AssemblyAST.IntegerConstant{Value: v.Expr.(*AST.ExpressionInteger).Value}
		instructions = append(instructions, AssemblyAST.NewMoveInstruction(destination, source))
		instructions = append(instructions, AssemblyAST.NewReturnInstruction(source))
	}

	return instructions
}

func instructionsFromExpression(expr AST.Expression) []AssemblyAST.Instruction {
	/*
		var instructions []AST.Instruction
		switch v := expr.(type) {
		case *AST.ExpressionInteger:
			inst := &AST.Immediate{
			}
			instructions = append(instructions, inst)
		}

		return instructions
	*/

	return nil
}

func instructionsFromDeclaration(decl AST.Declaration) []AssemblyAST.Instruction {
	var instructions []AssemblyAST.Instruction
	switch v := decl.(type) {
	case *AST.DeclarationFunction:
		for _, node := range v.Block.Body {
			instructions = append(instructions, instructionsFromNode(node)...)
		}
	}

	return instructions
}

func instructionsFromIR(inst IR.Instruction) []AssemblyAST.Instruction {
	var instructions []AssemblyAST.Instruction
	switch v := inst.(type) {
	case *IR.Return:
		instructions = append(instructions, AssemblyAST.NewMoveInstruction(v.Value, AssemblyAST.NewRegisterOperand(AssemblyAST.AX)))
		instructions = append(instructions, AssemblyAST.NewUnaryInstruction(v.Operator, v.Destination))
	case *IR.Unary:
		instructions = append(instructions, AssemblyAST.NewMoveInstruction(v.Destination, v.Source))
		instructions = append(instructions, AssemblyAST.NewUnaryInstruction(v.Operator, v.Destination))
	}

	return instructions
}

func GenerateAssemblyProgram(program IR.Program) AssemblyAST.Program {
	main := &AssemblyAST.FunctionDefinition{}
	main.Identifier = "main"

	for _, inst := range program.FunctionDefinition.Instructions {
		main.Instructions = append(main.Instructions, instructionsFromNode(node)...)
	}

	return AssemblyAST.Program{
		FunctionDefinition: main,
	}
}
