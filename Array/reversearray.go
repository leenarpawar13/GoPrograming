package main

import "fmt"

func main() {
	arr := [5]int{2, 3, 4, 5, 6}
	fmt.Println("print orignal array", arr)

	for i := 0; i < len(arr)/2; i++ {

		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]

	}
	fmt.Println("raverse array", arr)
}
