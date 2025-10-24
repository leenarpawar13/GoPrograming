package main

import "fmt"

func sum(nums ...int) int { // declared varidic fun
	total := 0
	for _, nums := range nums {
		total = total + nums
	}
	return total
}

func main() {
	//result := sum(3, 5, 7, 54) //same type value allow
	nums := []int{3, 4, 5, 6} // using slice
	result := sum(nums...)

	fmt.Println(" rsult is:=", result)

}
