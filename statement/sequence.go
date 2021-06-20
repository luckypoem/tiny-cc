package statement

import (
	"fmt"
	"tiny-cc/expression"
)

type Sequence struct {
	First  Statement
	Second Statement
}

func (obj Sequence) Reduce(environment map[string]expression.Expression) (Statement, map[string]expression.Expression) {
	// 如果语句中右侧表达式可以规约
	switch obj.First.(type) {
	case DoNothing:
		return obj.Second, environment
	default:
		reducedFirst, newEnv := obj.First.Reduce(environment)
		return Sequence{First: reducedFirst, Second: obj.Second}, newEnv
	}
}

func (obj Sequence) String() string {
	return fmt.Sprintf("%s; %s", obj.First, obj.Second)
}

func (obj Sequence) Inspect() string {
	return fmt.Sprintf("<%s; %s>", obj.First, obj.Second)
}

func (obj Sequence) Reducable() bool {
	return true
}
