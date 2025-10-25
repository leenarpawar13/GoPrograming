package main

import "fmt"

type order struct {
	id     string
	amount float32
	status string
}

//method
func (o *order) chnagestatus(status string) {
	o.status = status

}

func main() {
	myorder := order{
		id:     "1",
		amount: 50.00,
		status: "recived",
	}
	myorder.chnagestatus("confirmed")
	fmt.Println("order is", myorder)
}
