package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6} // declred slices
	fmt.Println("number is", numbers)
	numbers = append(numbers, 7, 8, 9, 10, 11, 12, 13, 14)
	fmt.Println("append numbers", numbers)
	fmt.Println("length", len(numbers))
	fmt.Println("length", cap(numbers))

	num := make([]int, 3, 5) // p and caproper way to declared slice using make func make func allow to set len

	num = append(num, 22, 23, 24)

	fmt.Println("append numbers", num)
	fmt.Println("length", len(num))
	fmt.Println("length", cap(num))

}
