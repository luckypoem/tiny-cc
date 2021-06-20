package statement

import (
	"fmt"
	"tiny-cc/expression"
)

type While struct {
	Condition expression.Expression
	Body      Statement
}

func (obj While) Reduce(environment map[string]expression.Expression) (Statement, map[string]expression.Expression) {
	return If{Condition: obj.Condition, Consequence: Sequence{First: obj.Body, Second: obj}, Alternate: DoNothing{}}, environment
}

func (obj While) String() string {
	return fmt.Sprintf("while(%s) { %s }", obj.Condition, obj.Body)
}

func (obj While) Inspect() string {
	return fmt.Sprintf("<while(%s) { %s }>", obj.Condition, obj.Body)
}

func (obj While) Reducable() bool {
	return true
}
