package main

import "fmt"

func main() {

	arr := [4]int{3, 2, 7, 4}
	max, min := arr[0], arr[0]

	for _, val := range arr {
		if val > max {
			max = val
		}
		if val < min {
			min = val
		}
		fmt.Println("max value", max)
		fmt.Println("min values", min)
	}
}
