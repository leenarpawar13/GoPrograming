package main

import "fmt"

func main() {
	a := [...]float64{12.34, 34.4, 45.56, 56.6}
	for i := 0; i < len(a); i++ {
		fmt.Printf("%d th element of a %.2f\n", i, a[i])
	}
	fmt.Println("*********************************************\n")
	sum := float64(0)
	for i, v := range a {
		fmt.Printf("%d th element of a %.2f\n", i, a[i])
		sum += v
	}
	fmt.Println("\n sum of all element", sum)
}
