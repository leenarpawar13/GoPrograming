package main

import "fmt"

func main() {
	//if _else condition
	a := 10
	if a == 10 {
		fmt.Println("a is equla to 10")
	} else {
		fmt.Println("not equal")
	}
	// if_else_if
	b := 8
	if b >= 10 {
		fmt.Println(" b is greter than 10")
	} else if b > 5 {
		fmt.Println("b is greter than 5 but smaller than 10")

	} else {
		fmt.Println("b is smaller than 5")
	}

	c := 10
	if a > 5 && (b > 12 || c > 13) {
		fmt.Println("true...")
	} else {
		fmt.Println("false")

	}

}
