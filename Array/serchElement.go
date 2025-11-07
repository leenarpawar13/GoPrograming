package main

import "fmt"

func main() {
	arr := [5]int{2, 3, 4, 5, 6}
	var search int
	fmt.Print("enter element")
	fmt.Scan(&search)

	found := false
	for i, val := range arr {
		if val == search {
			fmt.Printf("elements %d found at index%d\n", search, i)
			found = true
			break
		}

	}
	if !found {
		fmt.Println("not found")
	}
}
