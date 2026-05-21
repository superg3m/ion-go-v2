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
	case *AssemblyAST.InstructionDeallocateStack:
		instructions = []string{fmt.Sprintf("\taddq $%d, %%rsp", v.Value)}
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
		case "&":
			instructions = []string{fmt.Sprintf("\tandl %s, %s", v.Right.ToString(), v.Left.ToString())}
		case "|":
			instructions = []string{fmt.Sprintf("\torl %s, %s", v.Right.ToString(), v.Left.ToString())}
		case "<<":
			instructions = []string{fmt.Sprintf("\tsall %s, %s", v.Right.ToString(), v.Left.ToString())}
		case ">>":
			instructions = []string{fmt.Sprintf("\tsarl %s, %s", v.Right.ToString(), v.Left.ToString())}
		case "^":
			instructions = []string{fmt.Sprintf("\txorl %s, %s", v.Right.ToString(), v.Left.ToString())}
		default:
			panic(fmt.Sprintf("Unknown operator %s", v.Operator))
		}
	case *AssemblyAST.InstructionCDQ:
		instructions = []string{fmt.Sprintf("\tcdq")}
	case *AssemblyAST.InstructionDivide:
		instructions = []string{fmt.Sprintf("\tidiv %s", v.Left.ToString())}
	case *AssemblyAST.InstructionCompare:
		instructions = []string{fmt.Sprintf("\tcmpl %s, %s", v.Right.ToString(), v.Left.ToString())}
	case *AssemblyAST.InstructionJump:
		instructions = []string{fmt.Sprintf("\tjmp .L%s", v.TargetLabel)}
	case *AssemblyAST.InstructionConditionalJump:
		instructions = []string{fmt.Sprintf("\tj%s .L%s", v.Code.ToString(), v.TargetLabel)}
	case *AssemblyAST.InstructionSetConditionalCode:
		switch t := v.Destination.(type) {
		case *AssemblyAST.Stack:
			instructions = []string{
				fmt.Sprintf("\tset%s %s", v.Code.ToString(), v.Destination.ToString()),
			}
		case *AssemblyAST.Register:
			instructions = []string{
				fmt.Sprintf("\tset%s %s", v.Code.ToString(), t.X8BitName()),
			}
		}
	case *AssemblyAST.InstructionStackPush:
		switch t := v.Source.(type) {
		case *AssemblyAST.Register:
			instructions = []string{fmt.Sprintf("\tpushq %s", t.X64BitName())}
		case *AssemblyAST.Parameter:
			if t.Register == AssemblyAST.INVALID {
				instructions = []string{fmt.Sprintf("\tpushq %s", t.Register.X64BitName())}
			} else {
				instructions = []string{fmt.Sprintf("\tpushq %s", t.Register.ToString())}
			}
		default:
			instructions = []string{fmt.Sprintf("\tpushq %s", t.ToString())}
		}

	case *AssemblyAST.InstructionFunctionCall:
		instructions = []string{fmt.Sprintf("\tcall _%s", v.Identifier)}
	case *AssemblyAST.InstructionLabel:
		instructions = []string{fmt.Sprintf(".L%s:", v.Identifier)}
	case *AssemblyAST.InstructionStackPop:
		switch t := v.Destination.(type) {
		case *AssemblyAST.Register:
			instructions = []string{fmt.Sprintf("\tpopq %s", t.X64BitName())}
		case *AssemblyAST.Parameter:
			if t.Register == AssemblyAST.INVALID {
				instructions = []string{fmt.Sprintf("\tpopq %s", t.Register.X64BitName())}
			} else {
				instructions = []string{fmt.Sprintf("\tpopq %s", t.Register.ToString())}
			}
		default:
			instructions = []string{fmt.Sprintf("\tpopq %s", t.ToString())}
		}
	default:
		panic(fmt.Sprintf("Unknown instruction %T", inst))
	}

	return instructions
}

func (e *ATTx64Emitter) EmitFunctionDefinition(functionDefinition *AssemblyAST.FunctionDefinition) []string {
	instructions := []string{
		fmt.Sprintf(".global _%s", functionDefinition.Tok.Lexeme),
		fmt.Sprintf("_%s:", functionDefinition.Tok.Lexeme),
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
	for _, def := range program.Definitions {
		switch v := def.(type) {
		case *AssemblyAST.FunctionDefinition:
			instructions = append(instructions, e.EmitFunctionDefinition(v)...)
		case *AssemblyAST.VariableDefinition:
			panic("NOT IMPLEMENTED")
			// instructions = append(instructions, e.EmitFunctionDefinition(v)...)
		}
	}

	return instructions
}
