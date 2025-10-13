package main

import "fmt"

func main() {
	var price, quntity int // initial value
	fmt.Println("price is:- ", price, "quntity is", quntity)

	price = 3450 // multiple variable declaretion
	quntity = 6
	fmt.Println("price is:- ", price, "quntity is", quntity)

	var weight, height = 50, 5 //multiple variable with type inferance
	fmt.Println("weight is", weight, "height is", height)

	var (
		name   = "abc"
		age    = 23
		rollno int
	) //diff way to declared variable
	fmt.Println("name is =", name, "age is=", age, "roll no =", rollno)
}
