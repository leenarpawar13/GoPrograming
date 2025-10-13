package main

import (
	"fmt"
	"math"
)

func main() {
	name, count := "asd", 23 //short hand  declaration
	fmt.Println("name is =", name, "scount is =", count)

	a, b := 10, 20 // declared variable
	fmt.Println("a is", a, "b is", b)
	b, c := 30, 40 // b is declred but c is new
	fmt.Println("b is", b, "c is", c)
	b, c = 60, 70 //chnage value

	fmt.Println(" changed b is", b, "c is", c)

	x, y := 123.56, 567.23
	z := math.Min(x, y) // math package used
	fmt.Println("minimum value is", z)
	p := math.Max(x, y)

	fmt.Println("max value is", p)

}
