package main

import "fmt"

func main() {
	day := 30
	switch day {
	case 1:
		fmt.Println("monday")
	case 2:
		fmt.Println("tuesday")
	case 3:
		fmt.Println("wensday")
	default:
		fmt.Println("unknown day")
	}
}
