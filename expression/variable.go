package expression

import (
	"fmt"
)

type Variable struct {
	Name string
}

func (obj Variable) Reduce(environment map[string]Expression) Expression {
	return environment[obj.Name]
}

func (obj Variable) String() string {
	return fmt.Sprintf("%s", obj.Name)
}

func (obj Variable) Inspect() string {
	return fmt.Sprintf("<%s>", obj.Name)
}

func (obj Variable) Reducable() bool {
	return true
}
