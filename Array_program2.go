package main

import "fmt"

func main() {

	a := []string{"BadExpr", "BadExpr", "BadExpr", "BadExpr"}
	b := a         // copy of a assing to b
	b[1] = "pawar" // changes 1st index value
	fmt.Println(" a is :=", a)
	fmt.Println("b is", b)
}
