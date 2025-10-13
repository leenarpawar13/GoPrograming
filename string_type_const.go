package main

import "fmt"

func main() {
	var name = "sam" // allowed
	type xyz string
	var customName xyz = "sam"
	customName = xyz(name)
	fmt.Println(customName)

}
