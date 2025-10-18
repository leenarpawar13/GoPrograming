package main

import "fmt"

func main() {
	month := "may"

	switch month {
	case "january", "feb", "march": // switch stament multiple value
		fmt.Println("winter")

	case "april", "may", "june":
		fmt.Println("Spring")
	default:
		fmt.Println("other season")
	}

	temp := 10
	switch {
	case temp < 0:
		fmt.Println("freezing")
	case temp > 0 && temp < 10:

		fmt.Println("cold")

	case temp >= 10 && temp < 20:
		fmt.Println("cool")
	default:
		fmt.Println("hot")

	}

}
