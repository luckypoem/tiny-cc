package statement

import (
	"fmt"
	"tiny-cc/expression"
)

type Assign struct {
	Name string
	Exp  expression.Expression
}

func CopyMap(m map[string]expression.Expression) map[string]expression.Expression {
	maps2 := make(map[string]expression.Expression)
	for k2, v2 := range m {
		maps2[k2] = v2
	}
	return maps2
}

func (obj Assign) Reduce(environment map[string]expression.Expression) (Statement, map[string]expression.Expression) {
	// 如果语句中右侧表达式可以规约
	if obj.Exp.Reducable() {
		return Assign{Name: obj.Name, Exp: obj.Exp.Reduce(environment)}, environment
	} else {
		// 如果语句中右侧表达式规约完毕, 则更新环境
		envCopy := CopyMap(environment)
		envCopy[obj.Name] = obj.Exp
		// 规约成donothing的同时, 返回新环境
		return DoNothing{}, envCopy
	}
}

func (obj Assign) String() string {
	return fmt.Sprintf("%s = %s", obj.Name, obj.Exp)
}

func (obj Assign) Inspect() string {
	return fmt.Sprintf("<%s = %s>", obj.Name, obj.Exp)
}

func (obj Assign) Reducable() bool {
	return true
}
