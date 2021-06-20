package expression

import (
	"fmt"
	"log"
	"strconv"
)

type Add struct {
	Left  Expression
	Right Expression
}

func (add Add) String() string {
	return fmt.Sprintf("%s + %s", add.Left.String(), add.Right.String())
}

func (add Add) Inspect() string {
	return fmt.Sprintf("<%s + %s>", add.Left.String(), add.Right.String())
}

func (add Add) Reduce(environment map[string]Expression) Expression {
	if add.Left.Reducable() {
		return Add{Left: add.Left.Reduce(environment), Right: add.Right}
	} else if add.Right.Reducable() {
		return Add{Left: add.Left, Right: add.Right.Reduce(environment)}
	}

	li, err := strconv.ParseInt(add.Left.String(), 10, 64)
	if err != nil {
		log.Fatalf("add left must be number but have %s", add.Left.String())
	}
	ri, err := strconv.ParseInt(add.Right.String(), 10, 64)
	if err != nil {
		log.Fatalf("add right must be number but have %s", add.Right.String())
	}
	return Number{Value: li + ri}
}

func (add Add) Reducable() bool {
	return true
}
