package expression

import "fmt"

type Boolean struct {
	Value bool
}

func (obj Boolean) String() string {
	return fmt.Sprintf("%t", obj.Value)
}

func (obj Boolean) Inspect() string {
	return fmt.Sprintf("<%t>", obj.Value)
}

func (obj Boolean) Reduce(environment map[string]Expression) Expression {
	return obj
}

func (obj Boolean) Reducable() bool {
	return false
}
