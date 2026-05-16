package IR

type Value interface {
	isValue()
}

type Constant struct {
	Value int
}

type Variable struct {
	Name string
}

func (*Constant) isValue() {}
func (*Variable) isValue() {}

func NewConstantValue(value int) Value {
	return &Constant{
		Value: value,
	}
}

func NewVariable(name string) *Variable {
	return &Variable{
		Name: name,
	}
}
