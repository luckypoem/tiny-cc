package statement

import "tiny-cc/expression"

type Statement interface {
	String() string
	Inspect() string
	Reduce(environment map[string]expression.Expression) (Statement, map[string]expression.Expression)
	Reducable() bool
}
