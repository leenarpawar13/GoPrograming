package main

import "fmt"

func main() {
	pls := [][]string{{"c", "c++"},
		{"java", "javascript"},
		{"go"},
	}
	for _, v1 := range pls {
		for _, v2 := range v1 {
			fmt.Printf("%s", v2)
		}
	}
	fmt.Printf("\n")
}
