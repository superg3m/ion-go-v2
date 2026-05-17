package AssemblyEmitter

import (
	"fmt"
	"ion/go/v2/AssemblyAST"
)

type ATTx64Emitter struct{}

func EmitFunctionPrologue() []string {
	return []string{
		"\tpushq %rbp",
		"\tmovq %rsp, %rbp",
	}
}

func EmitFunctionEpilogue() []string {
	return []string{
		"\tmovq %rbp, %rsp",
		"\tpopq %rbp",
		"\tret",
	}
}

func (e *ATTx64Emitter) EmitReturnInstruction() {}

func (e *ATTx64Emitter) EmitInstruction(inst AssemblyAST.Instruction) []string {
	var instructions []string
	switch v := inst.(type) {
	case *AssemblyAST.InstructionMove:
		instructions = append(instructions, fmt.Sprintf("\tmovl %s, %s", v.Source.ToString(), v.Destination.ToString()))
	case *AssemblyAST.InstructionReturn:
		instructions = EmitFunctionEpilogue()
	case *AssemblyAST.InstructionStackAllocate:
		instructions = []string{fmt.Sprintf("\tsubq $%d, %%rsp", v.AllocationSize)}
	case *AssemblyAST.InstructionUnary:
		switch v.Operator {
		case "-":
			instructions = []string{fmt.Sprintf("\tnegl %s", v.Operand.ToString())}
		case "~":
			instructions = []string{fmt.Sprintf("\tnotl %s", v.Operand.ToString())}
		}
	case *AssemblyAST.InstructionBinary:
		switch v.Operator {
		case "+":
			instructions = []string{fmt.Sprintf("\taddl %s, %s", v.Right.ToString(), v.Left.ToString())}
		case "-":
			instructions = []string{fmt.Sprintf("\tsubl %s, %s", v.Right.ToString(), v.Left.ToString())}
		case "*":
			instructions = []string{fmt.Sprintf("\timull %s, %s", v.Right.ToString(), v.Left.ToString())}
		default:
			panic(fmt.Sprintf("Unknown operator %s", v.Operator))
		}
	case *AssemblyAST.InstructionCDQ:
		instructions = []string{fmt.Sprintf("\tcdq")}
	case *AssemblyAST.InstructionDivide:
		instructions = []string{fmt.Sprintf("\tidiv %s", v.Left.ToString())}
	case *AssemblyAST.InstructionCompare:
		instructions = []string{fmt.Sprintf("\tcmpl %s, %s", v.C1.ToString(), v.C2.ToString())}
	case *AssemblyAST.InstructionJump:
		instructions = []string{fmt.Sprintf("\tjmp .L%s", v.TargetLabel)}
	case *AssemblyAST.InstructionConditionalJump:
		instructions = []string{fmt.Sprintf("\tj%s .L%s", v.Code.ToString(), v.TargetLabel)}
	case *AssemblyAST.InstructionLabel:
		instructions = []string{fmt.Sprintf(".L%s:", v.Identifier)}
	default:
		panic(fmt.Sprintf("Unknown instruction %T", inst))
	}

	return instructions
}

func (e *ATTx64Emitter) EmitFunctionDefinition(functionDefinition *AssemblyAST.FunctionDefinition) []string {
	instructions := []string{
		fmt.Sprintf(".global %s", functionDefinition.Identifier),
		fmt.Sprintf("%s:", functionDefinition.Identifier),
	}

	instructions = append(instructions, EmitFunctionPrologue()...)

	for _, inst := range functionDefinition.Instructions {
		instructions = append(instructions, e.EmitInstruction(inst)...)
	}

	return instructions
}

func (e *ATTx64Emitter) EmitAssemblyProgram(program AssemblyAST.Program) []string {
	instructions := []string{
		".text",
	}
	instructions = append(instructions, e.EmitFunctionDefinition(program.FunctionDefinition)...)

	return instructions
}
