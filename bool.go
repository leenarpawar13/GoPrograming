package main

import "fmt"

func main() {
	a := true
	b := false
	fmt.Println(" a is ", a, "b is ", b)
	c := a && b // true when both  of operand true athorwise false
	fmt.Println("c is :=", c)
	d := a || b // or operator ,false when both  of operand false athorwise true
	fmt.Println(" d is :=", d)

}
