package Codegen

import "ion/go/v2/AST"
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

func instructionsFromNode(node AST.SourceNode) []AssemblyAST.Instruction {
	switch v := node.(type) {
	case AST.Statement:
		return instructionsFromStatement(v)
	case AST.Expression:
		return instructionsFromExpression(v)
	case AST.Declaration:
		return instructionsFromDeclaration(v)
	}

	return nil
}

func GenerateAssemblyProgram(program AST.Program) AssemblyAST.Program {
	main := &AssemblyAST.FunctionDefinition{}
	main.Identifier = "main"
	function := program.Declarations[0].(*AST.DeclarationFunction)

	for _, node := range function.Block.Body {
		main.Instructions = append(main.Instructions, instructionsFromNode(node)...)
	}

	return AssemblyAST.Program{
		FunctionDefinition: main,
	}
}
