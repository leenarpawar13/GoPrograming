package main

import "fmt"

func modifyvalue(num *int) {
	*num = *num + 10

}
func main() {
	//var num int
	num := 2
	//var ptr *int
	ptr := &num

	fmt.Println("num value", num)
	//fmt.Println("pointer", ptr)
	fmt.Println("pinters value", *ptr)

	var pointer *int //default pointer ==  nil
	if pointer == nil {

		fmt.Println("pointer is not assigned")
	}

	value := 15
	modifyvalue(&value)
	fmt.Println("mdifyvalue", value)
}
