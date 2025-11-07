package main

import "fmt"

func main() {
	arr := [5]int{10, 23, 32, 34, 35}
	fmt.Println("array elements")
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}
