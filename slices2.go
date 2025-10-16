package main

import "fmt"

func main() {
	a := []int{12, 13, 14, 15, 16, 17}
	var b []int = a[2:5]
	fmt.Println(b)
	for i := range b {
		b[i]++
	}
	fmt.Println("b", b)
}
