package main

import "fmt"

func main() {

	age := 25
	name := "xyz"
	height := 5.2367
	fmt.Println("age:-", age, "name:-", name, "height:-", height)

	fmt.Printf("age %d\n", age)         // %d for int
	fmt.Printf("name%s\n", name)        // % s for string
	fmt.Printf("height %.3f\n", height) // %.2f for float print 2 digit after decimal
	fmt.Printf("name %T\n", name)       // %T for check type
	fmt.Printf("*********************************\n")

	fmt.Printf("age %d, name %s, height %.3f", age, name, height)
}
