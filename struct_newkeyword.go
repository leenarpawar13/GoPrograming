package main

import "fmt"

type order struct {
	id     string
	amount float32
	status string
}

//new keyword linke constructor
func Newmyoder(id string, amount float32, status string) *order {
	myorder := order{
		id:     id,
		amount: amount,
		status: status,
	}
	return &myorder
}

func main() {

	myorder := Newmyoder("1", 23.43, "recived")
	fmt.Println("order is", myorder)
}
