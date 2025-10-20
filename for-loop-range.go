package main

import "fmt"

func main() {
	num := []int{12, 13, 14, 15, 15} // range keyword for int
	for index, value := range num {

		fmt.Printf("index is %d ,value is %d\n", index, value)
	}
	data := "hello world" // range keyworld for string
	for index, value := range data {
		fmt.Printf("index is %d ,value is %c\n", index, value)
	}

}
