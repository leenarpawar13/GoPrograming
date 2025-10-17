package main

import "fmt"

func main() {

	fruitarray := []string{"aaple", "mango", "orange", "pine apple", "chikoo"}
	fruitslice := fruitarray[1:4]
	fmt.Println("fruitslice is", fruitslice)
	fmt.Printf("length %d and capcity %d ", len(fruitslice), cap(fruitslice))
	fruitslice = fruitslice[:cap(fruitslice)] // re slicing
	fmt.Println("afret re slicing length", len(fruitslice), "capcity", cap(fruitslice))
}
