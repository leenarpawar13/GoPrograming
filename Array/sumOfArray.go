package main

import "fmt"

func main() {
	arr := [3]int{10, 20, 30}
	sum := 0
	for _, val := range arr {
		sum += val
	}
	fmt.Println("sum of array elemnets", sum)
}
