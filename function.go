package main

import "fmt"

func simplefunc() {
	fmt.Println("its simple function....")

}
func add(a, b int) int {
	return a + b

}

func multiply(a, b int) (result int) {
	result = a * b
	return
}
func minus(a, b, c int) int {
	return a - b - c
}

func main() {
	fmt.Println("function learning")
	simplefunc()
	ans := add(2, 4)
	fmt.Println("addition is:-", ans)
	ans1 := multiply(3, 4)
	fmt.Println("multiply is:=", ans1)
	ans2 := minus(8, 2, 3)
	fmt.Println("ans =", ans2)
}
