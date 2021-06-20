package expression

import (
	"fmt"
	"strconv"
)

type LessThan struct {
	Left  Expression
	Right Expression
}

func (obj LessThan) String() string {
	return fmt.Sprintf("%s < %s", obj.Left, obj.Right)
}

func (obj LessThan) Inspect() string {
	return fmt.Sprintf("<%s < %s>", obj.Left, obj.Right)
}

func (obj LessThan) Reduce(environment map[string]Expression) Expression {
	if obj.Left.Reducable() {
		return LessThan{Left: obj.Left.Reduce(environment), Right: obj.Right}
	} else if obj.Right.Reducable() {
		return LessThan{Left: obj.Left, Right: obj.Right.Reduce(environment)}
	}

	li, _ := strconv.ParseInt(obj.Left.String(), 10, 64)
	ri, _ := strconv.ParseInt(obj.Right.String(), 10, 64)
	return Boolean{Value: li < ri}
}

func (obj LessThan) Reducable() bool {
	return true
}
