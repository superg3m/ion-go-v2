package AssemblyEmitter

import (
	"ion/go/v2/AssemblyAST"
)

type AssemblyEmitter interface {
	EmitAssembly(program AssemblyAST.Program)
}
