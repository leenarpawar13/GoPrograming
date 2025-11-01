package main

import "fmt"

func sum(result chan int, num1 int, num2 int) {
	myres := num1 + num2
	result <- myres

}

func main() {

	result := make(chan int)
	go sum(result, 5, 6)
	res := <-result //send data
	fmt.Println("result is", res)
}
