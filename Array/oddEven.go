package main

import "fmt"

func main() {

	arr := [6]int{1, 2, 3, 4, 5, 6}
	even, odd := 0, 0

	for _, val := range arr {
		if val%2 == 0 {
			fmt.Println("even", val)
			even++
		} else {
			fmt.Println("odd", val)
			odd++
		}
	}
	fmt.Println("even count", even)
	fmt.Println("odd count", odd)

}
