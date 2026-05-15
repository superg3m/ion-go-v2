package AssemblyEmitter

import "ion/go/v2/AST"

type AssemblyEmitter interface {
	EmitAssembly(program AST.AssemblyProgram)
}
