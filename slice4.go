package main

import "fmt"

func main() {
	veggies := []string{"potato", "tomatoes", "brinjal"}
	fruits := []string{"mango", "apple"}
	food := append(veggies, fruits...) //append
	fmt.Println("food", food)
}
