package expression

import (
	"fmt"
	"strconv"
)

type Multiply struct {
	Left  Expression
	Right Expression
}

func (mul Multiply) String() string {
	return fmt.Sprintf("%s * %s", mul.Left.String(), mul.Right.String())
}

func (mul Multiply) Inspect() string {
	return fmt.Sprintf("<%s * %s>", mul.Left.String(), mul.Right.String())
}

func (mul Multiply) Reduce(environment map[string]Expression) Expression {
	if mul.Left.Reducable() {
		return Multiply{Left: mul.Left.Reduce(environment), Right: mul.Right}
	} else if mul.Right.Reducable() {
		return Multiply{Left: mul.Left, Right: mul.Right.Reduce(environment)}
	}

	li, _ := strconv.ParseInt(mul.Left.String(), 10, 64)
	ri, _ := strconv.ParseInt(mul.Right.String(), 10, 64)
	return Number{Value: li * ri}
}

func (mul Multiply) Reducable() bool {
	return true
}
