package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("num is-", i)
	}
	counter := 0 // infinite
	for {
		fmt.Println(" infinite loop")
		counter++
		break

	}

}
