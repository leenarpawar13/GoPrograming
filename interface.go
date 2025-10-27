package main

import "fmt"

// craete inteface
type paymenter interface {
	pay(amount float32)
}

type payment struct {
	getway stipe
}

func (p payment) makePayment(amount float32) {
	//razorpayPaymentGW := razorpay{}
	//razorpayPaymentGW.pay(amount)
	//stipePaymentGW := stipe{}
	p.getway.pay(amount)
}

type razorpay struct {
}

func (r razorpay) pay(amount float32) {
	fmt.Println("making payment using razorpay", amount)
}

type stipe struct {
}

func (s stipe) pay(amount float32) {
	fmt.Println("making payment using stipe", amount)
}

func main() {
	stipePaymentGW := stipe{}
	Newpayment := payment{
		getway: stipePaymentGW,
	}
	Newpayment.makePayment(100)

}
