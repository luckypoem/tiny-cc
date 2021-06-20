package statement

import "tiny-cc/expression"

type DoNothing struct {
}

func (obj DoNothing) Reduce(environment map[string]expression.Expression) (Statement, map[string]expression.Expression) {
	panic("Can't reduce DoNothing")
}

func (obj DoNothing) String() string {
	return "DoNothing"
}

func (obj DoNothing) Inspect() string {
	return "<DoNothing>"
}

func (obj DoNothing) Reducable() bool {
	return false
}
