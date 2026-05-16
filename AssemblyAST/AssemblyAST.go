package AssemblyAST

type Program struct {
	FunctionDefinition *FunctionDefinition
}
type FunctionDefinition struct {
	Identifier   string
	Instructions []Instruction
}
