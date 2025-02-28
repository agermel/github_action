package main

import (
	"action/calc"
	"fmt"
)

func main() {
	a, b := 3724, 3721

	fmt.Println(calc.Add(a, b), calc.Minus(a, b), calc.Multiple(a, b), calc.Division(a, b))
}
