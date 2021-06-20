package main

import (
	"fmt"
	"tiny-cc/expression"
	"tiny-cc/statement"
)

type Machine struct {
	Statement   statement.Statement
	environment map[string]expression.Expression
}

func (machine *Machine) Step() {
	machine.Statement, machine.environment = machine.Statement.Reduce(machine.environment)
}

func (machine *Machine) Run() {
	for machine.Statement.Reducable() {
		fmt.Println(machine.Statement, machine.environment)
		machine.Step()
	}
	fmt.Println(machine.Statement, machine.environment)
}

func main() {
	var state statement.Statement
	state = statement.Sequence{
		First:  statement.Assign{Name: "x", Exp: expression.Boolean{Value: false}},
		Second: statement.Assign{Name: "x", Exp: expression.Add{Left: expression.Variable{Name: "x"}, Right: expression.Number{Value: 3}}},
	}
	env := make(map[string]expression.Expression)
	m := Machine{Statement: state, environment: env}
	m.Run()
	return
}
