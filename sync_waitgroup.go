package main

import (
	"fmt"
	"sync"
)

func Worker(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("worker %d started\n", i)

	fmt.Printf("worker %d end\n", i)
}

func main() {
	var wg sync.WaitGroup
	//fmt.Println("grouting started")

	for i := 1; i <= 3; i++ {
		wg.Add(1) // increment waitgroup counter
		go Worker(i, &wg)

	}
	wg.Wait()
	fmt.Println("worker task complete")
}
