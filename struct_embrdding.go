package main

import "fmt"

type customer struct {
	name  string
	phone string
}
type order1 struct {
	id       string
	amount   float32
	status   string
	customer //embedding struct
}

//method
/*func (o *order) chnagestatus(status string) {
	o.status = status

}*/

func main() {

	myorder := order1{
		id:     "1",
		amount: 50.00,
		status: "recived",
		customer: customer{name: "leena",
			phone: "123434445"},
	}

	fmt.Println("order is", myorder)
}
