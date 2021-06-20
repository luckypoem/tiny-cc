package statement

import (
	"fmt"
	"tiny-cc/expression"
)

type If struct {
	Condition   expression.Expression
	Consequence Statement
	Alternate   Statement
}

func (obj If) Reduce(environment map[string]expression.Expression) (Statement, map[string]expression.Expression) {
	// 如果语句中右侧表达式可以规约
	if obj.Condition.Reducable() {
		return If{Condition: obj.Condition.Reduce(environment), Consequence: obj.Consequence, Alternate: obj.Alternate}, environment
	} else {
		switch obj.Condition.String() {
		case "true":
			return obj.Consequence, environment
		case "false":
			return obj.Alternate, environment
		default:
			panic("If Condition must be reduced to boolean")
		}
	}
}

func (obj If) String() string {
	return fmt.Sprintf("if (%s) {%s} else {%s}", obj.Condition, obj.Consequence, obj.Alternate)
}

func (obj If) Inspect() string {
	return fmt.Sprintf("<if (%s) {%s} else {%s}>", obj.Condition, obj.Consequence, obj.Alternate)
}

func (obj If) Reducable() bool {
	return true
}
