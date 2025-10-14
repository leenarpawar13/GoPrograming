package main

import "fmt"

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("denominator must be 0")
	}
	return a / b, nil
}
func main() {
	fmt.Println("error handling concept")
	ans, _ := divide(10, 0) // blank indentifier
	/*ans, err := (divide(10, 0))
	if err != nil {
		fmt.Println("error handling")
	}*/
	fmt.Println("division is:=", ans)
}
