package Codegen

import "ion/go/v2/AST"

func instructionsFromStatement(stmt AST.Statement) []AST.Instruction {
	var instructions []AST.Instruction
	switch v := stmt.(type) {
	case *AST.StatementReturn:
		destination := &AST.Register{Name: "%eax"}
		source := &AST.Immediate{Immediate: v.Expr.(*AST.ExpressionInteger).Value}
		instructions = append(instructions, &AST.InstructionMove{destination, source})
		instructions = append(instructions, &AST.InstructionRet{})
	}

	return instructions
}

func instructionsFromExpression(expr AST.Expression) []AST.Instruction {
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

func instructionsFromDeclaration(decl AST.Declaration) []AST.Instruction {
	var instructions []AST.Instruction
	switch v := decl.(type) {
	case *AST.DeclarationFunction:
		for _, node := range v.Block.Body {
			instructions = append(instructions, instructionsFromNode(node)...)
		}
	}

	return instructions
}

func instructionsFromNode(node AST.SourceNode) []AST.Instruction {
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

func GenerateAssemblyProgram(program AST.Program) AST.AssemblyProgram {
	var main AST.FunctionDefinition
	main.Identifier = "main"
	function := program.Declarations[0].(*AST.DeclarationFunction)

	for _, node := range function.Block.Body {
		main.Instructions = append(main.Instructions, instructionsFromNode(node)...)
	}

	return AST.AssemblyProgram{
		FunctionDefinition: main,
	}
}
