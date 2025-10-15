package main

import "fmt"

func printarray(a [3][2]string) {
	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Printf("%s", v2)

		}
		fmt.Println("\n")
	}
}
func main() {
	a := [3][2]string{

		{"aaaa", "bbbb"},
		{"cccc", "ddddd"},
		{"eee", "fff"},
	}
	printarray(a)

	var b [3][2]string
	b[0][0] = "apple"
	b[0][1] = "samsung"
	b[1][0] = "microsoft"
	b[1][1] = "google"
	b[2][0] = "AT&t"
	b[2][1] = "asd"

	fmt.Printf("\n")
	printarray(b)
}
