package expression

import "fmt"

type Number struct {
	Value int64
}

func (number Number) String() string {
	return fmt.Sprintf("%d", number.Value)
}

func (number Number) Inspect() string {
	return fmt.Sprintf("<%d>", number.Value)
}

func (number Number) Reduce(environment map[string]Expression) Expression {
	panic("don't reduce number")
}

func (number Number) Reducable() bool {
	return false
}
