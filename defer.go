package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("1 st program")
	data := add(3, 2)
	fmt.Println("addition is:=", data)
	defer fmt.Println("2nd program") // using defer keyword excecute this line befor closing main fun
	fmt.Println("3rd program")
}
