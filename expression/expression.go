package expression

type Expression interface {
	String() string
	Inspect() string
	Reduce(environment map[string]Expression) Expression
	Reducable() bool
}
