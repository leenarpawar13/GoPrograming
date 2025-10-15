package main

import "fmt"

func changeLocal(num [5]int) {
	num[0] = 55
	fmt.Println("inside local function", num)

}
func main() {
	num := [...]int{4, 5, 6, 7, 8}
	fmt.Println("before passing function", num)
	changeLocal(num) // called fun
	fmt.Println("after passing funct", num)

}
